package github

// <!-- START clutchdoc -->
// description: GitHub client that combines the REST/GraphQL APIs and raw git capabilities into a single interface.
// <!-- END clutchdoc -->

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bradleyfalzon/ghinstallation"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	gittransport "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
	"github.com/golang/protobuf/ptypes/any"
	githubv3 "github.com/google/go-github/v37/github"
	"github.com/shurcooL/githubv4"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	githubv1 "github.com/lyft/clutch/backend/api/config/service/github/v1"
	sourcecontrolv1 "github.com/lyft/clutch/backend/api/sourcecontrol/v1"
	"github.com/lyft/clutch/backend/service"
	"github.com/lyft/clutch/backend/service/authn"
)

const Name = "clutch.service.github"
const CurrentUser = ""

type FileMap map[string]io.ReadCloser

type StatsRoundTripper struct {
	Wrapped http.RoundTripper
	scope   tally.Scope
}

func (st *StatsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := st.Wrapped.RoundTrip(req)

	if hdr := resp.Header.Get("X-RateLimit-Remaining"); hdr != "" {
		if v, err := strconv.Atoi(hdr); err == nil {
			st.scope.Gauge("rate_limit_remaining").Update(float64(v))
		}
	}
	return resp, err
}

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	config := &githubv1.Config{}
	if err := cfg.UnmarshalTo(config); err != nil {
		return nil, err
	}
	return newService(config, scope, logger)
}

// Remote ref points to a git reference using a combination of the repository and the reference itself.
type RemoteRef struct {
	// Organization or user that owns the repository.
	RepoOwner string
	// Name of the repository.
	RepoName string
	// SHA, branch name, or tag.
	Ref string
}

// Repository contains information about a requested repository.
type Repository struct {
	Name          string
	Owner         string
	DefaultBranch string
}

// File contains information about a requested file, including its content.
type File struct {
	Path             string
	Contents         io.ReadCloser
	SHA              string
	LastModifiedTime time.Time
	LastModifiedSHA  string
}

// Client allows various interactions with remote repositories on GitHub.
type Client interface {
	GetFile(ctx context.Context, ref *RemoteRef, path string) (*File, error)
	CreateBranch(ctx context.Context, req *CreateBranchRequest) error
	CreatePullRequest(ctx context.Context, ref *RemoteRef, base, title, body string) (*PullRequestInfo, error)
	CreateRepository(ctx context.Context, req *sourcecontrolv1.CreateRepositoryRequest) (*sourcecontrolv1.CreateRepositoryResponse, error)
	CreateIssueComment(ctx context.Context, ref *RemoteRef, number int, body string) error
	CompareCommits(ctx context.Context, ref *RemoteRef, compareSHA string) (*githubv3.CommitsComparison, error)
	GetCommit(ctx context.Context, ref *RemoteRef) (*Commit, error)
	GetRepository(ctx context.Context, ref *RemoteRef) (*Repository, error)
	GetOrganization(ctx context.Context, organization string) (*githubv3.Organization, error)
	ListOrganizations(ctx context.Context, user string) ([]*githubv3.Organization, error)
	ListPullRequestsWithCommit(ctx context.Context, ref *RemoteRef, sha string, opts *githubv3.PullRequestListOptions) ([]*PullRequestInfo, error)
	GetOrgMembership(ctx context.Context, user, org string) (*githubv3.Membership, error)
	GetUser(ctx context.Context, username string) (*githubv3.User, error)
}

// This func can be used to create comments for PRs or Issues
func (s *svc) CreateIssueComment(ctx context.Context, ref *RemoteRef, number int, body string) error {
	com := &githubv3.IssueComment{
		Body: strPtr(body),
	}
	_, _, err := s.rest.Issues.CreateComment(ctx, ref.RepoOwner, ref.RepoName, number, com)
	return err
}

type PullRequestInfo struct {
	Number     int
	HTMLURL    string
	BranchName string
}

type svc struct {
	scope  tally.Scope
	logger *zap.Logger

	graphQL v4client
	rest    v3client

	appTransport        *ghinstallation.Transport
	personalAccessToken string
}

func (s *svc) basicAuth(ctx context.Context) *gittransport.BasicAuth {
	ret := &gittransport.BasicAuth{
		Username: "token",
	}

	if s.appTransport != nil {
		password, err := s.appTransport.Token(ctx)
		ret.Password = password

		if err != nil {
			s.logger.Error("could not refresh token from transport", zap.Error(err))
		}
	} else {
		ret.Password = s.personalAccessToken
	}

	return ret
}

func (s *svc) GetOrganization(ctx context.Context, organization string) (*githubv3.Organization, error) {
	org, _, err := s.rest.Organizations.Get(ctx, organization)
	if err != nil {
		return nil, err
	}
	return org, nil
}

// ListOrganizations returns all organizations for a specified user.
// To list organizations for the currently authenticated user set user to "".
func (s *svc) ListOrganizations(ctx context.Context, user string) ([]*githubv3.Organization, error) {
	organizations, _, err := s.rest.Organizations.List(ctx, user, &githubv3.ListOptions{})
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

// GetOrgMembership returns a specified users membership within a specified organization.
// To list organizations for the currently authenticated user set user to "".
func (s *svc) GetOrgMembership(ctx context.Context, user, org string) (*githubv3.Membership, error) {
	membership, response, err := s.rest.Organizations.GetOrgMembership(ctx, user, org)
	if err != nil {
		// A user might be part of an org but not have permissions to get memerbship information if auth is behind SSO.
		// In this case we return a default Membership.
		if response.StatusCode == 403 {
			return &githubv3.Membership{}, nil
		}
		return nil, err
	}

	return membership, nil
}

// GetUser returns information about the specified user.
// To list organizations for the currently authenticated user set user to "".
func (s *svc) GetUser(ctx context.Context, username string) (*githubv3.User, error) {
	user, _, err := s.rest.Users.Get(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *svc) CreateRepository(ctx context.Context, req *sourcecontrolv1.CreateRepositoryRequest) (*sourcecontrolv1.CreateRepositoryResponse, error) {
	// Validate that we received GitHub Options.
	_, ok := req.Options.(*sourcecontrolv1.CreateRepositoryRequest_GithubOptions)
	if !ok {
		return nil, status.New(codes.InvalidArgument, "GitHub options were not provided to GitHub service").Err()
	}

	opts := req.GetGithubOptions()
	currentUser, _, err := s.rest.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	var org string
	if org = req.Owner; currentUser.GetLogin() == req.Owner {
		// If the specified owner is the same as the current user the GitHub API expects an empty string.
		org = ""
	}

	repo := &githubv3.Repository{
		Name:        strPtr(req.Name),
		Description: strPtr(req.Description),
		Private:     boolPtr(opts.Parameters.Visibility.String() == sourcecontrolv1.Visibility_PRIVATE.String()),
		AutoInit:    boolPtr(opts.AutoInit),
	}
	newRepo, _, err := s.rest.Repositories.Create(ctx, org, repo)
	if err != nil {
		return nil, err
	}

	resp := &sourcecontrolv1.CreateRepositoryResponse{
		Url: *newRepo.HTMLURL,
	}
	return resp, nil
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func (s *svc) CreatePullRequest(ctx context.Context, ref *RemoteRef, base, title, body string) (*PullRequestInfo, error) {
	req := &githubv3.NewPullRequest{
		Title:               strPtr(title),
		Head:                strPtr(ref.Ref),
		Base:                strPtr(base),
		Body:                strPtr(body),
		MaintainerCanModify: boolPtr(true),
	}
	pr, _, err := s.rest.PullRequests.Create(ctx, ref.RepoOwner, ref.RepoName, req)
	if err != nil {
		return nil, err
	}

	return &PullRequestInfo{
		Number: pr.GetNumber(),
		// There are many possible URLs to return, but the HTML one is most human friendly
		HTMLURL: pr.GetHTMLURL(),
	}, nil
}

func (s *svc) ListPullRequestsWithCommit(ctx context.Context, ref *RemoteRef, sha string, opts *githubv3.PullRequestListOptions) ([]*PullRequestInfo, error) {
	respPRs, _, err := s.rest.PullRequests.ListPullRequestsWithCommit(ctx, ref.RepoOwner, ref.RepoName, sha, opts)
	if err != nil {
		return nil, err
	}

	prInfos := make([]*PullRequestInfo, len(respPRs))
	for i, pr := range respPRs {
		prInfos[i] = &PullRequestInfo{
			Number:     pr.GetNumber(),
			HTMLURL:    pr.GetHTMLURL(),
			BranchName: pr.GetHead().GetRef(),
		}
	}

	return prInfos, nil
}

type CreateBranchRequest struct {
	// The base for the new branch.
	Ref *RemoteRef

	// The name of the new branch.
	BranchName string

	// Files and their content. Files will be clobbered with new content or created if they don't already exist.
	Files FileMap

	// The commit message for files added.
	CommitMessage string

	// Fetch only ReferenceName specified branch.
	SingleBranch bool
}

func commitOptionsFromClaims(ctx context.Context, commitTime time.Time) *git.CommitOptions {
	ret := &git.CommitOptions{Author: &object.Signature{When: commitTime}, All: true}

	subject := "Anonymous User" // Used if auth is disabled or it's the actual anonymous user.
	if claims, err := authn.ClaimsFromContext(ctx); err == nil && claims.Subject != authn.AnonymousSubject {
		subject = claims.Subject
	}
	ret.Author.Name = fmt.Sprintf("%s via Clutch", subject)

	// If it looks like an email, make it the email, otherwise the email will be left blank. This could be enhanced
	// via a default email from config if needed for other use cases.
	email := ""
	if strings.Contains(subject, "@") {
		email = subject
	}
	ret.Author.Email = email

	return ret
}

// Creates a new branch with a commit containing files and pushes it to the remote.
func (s *svc) CreateBranch(ctx context.Context, req *CreateBranchRequest) error {
	cloneOpts := &git.CloneOptions{
		SingleBranch:  req.SingleBranch,
		Depth:         1,
		URL:           fmt.Sprintf("https://github.com/%s/%s", req.Ref.RepoOwner, req.Ref.RepoName),
		ReferenceName: plumbing.NewBranchReferenceName(req.Ref.Ref),
		Auth:          s.basicAuth(ctx),
	}

	repo, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), cloneOpts)
	if err != nil {
		return err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	checkoutOpts := &git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(req.BranchName),
		Create: true,
	}
	if err := wt.Checkout(checkoutOpts); err != nil {
		return err
	}

	for filename, contents := range req.Files {
		fh, err := wt.Filesystem.Create(filename)
		if err != nil {
			return err
		}
		if _, err := io.Copy(fh, contents); err != nil {
			return err
		}
	}

	opts := commitOptionsFromClaims(ctx, time.Now())
	if _, err := wt.Commit(req.CommitMessage, opts); err != nil {
		return err
	}

	pushOpts := &git.PushOptions{Auth: s.basicAuth(ctx)}
	if err := repo.PushContext(ctx, pushOpts); err != nil {
		return err
	}

	return nil
}

func newService(config *githubv1.Config, scope tally.Scope, logger *zap.Logger) (Client, error) {
	ret := &svc{
		scope:  scope,
		logger: logger,
	}

	var httpClient *http.Client

	switch auth := config.GetAuth().(type) {
	case *githubv1.Config_AccessToken:
		token := config.GetAccessToken()
		ret.personalAccessToken = token

		tokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(context.Background(), tokenSource)
	case *githubv1.Config_AppConfig:
		config := config.GetAppConfig()
		tr := http.DefaultTransport
		var pem []byte
		switch config.GetPem().(type) {
		case *githubv1.AppConfig_Base64Pem:
			p, err := base64.StdEncoding.DecodeString(config.GetBase64Pem())
			if err != nil {
				return nil, err
			}
			pem = p
		case *githubv1.AppConfig_KeyPem:
			pem = []byte(config.GetKeyPem())
		}
		itr, err := ghinstallation.New(tr, config.AppId, config.InstallationId, pem)
		if err != nil {
			return nil, err
		}
		ret.appTransport = itr

		httpClient = &http.Client{Transport: itr}
	default:
		return nil, fmt.Errorf("did not recognize auth config type '%T'", auth)
	}
	httpClient.Transport = &StatsRoundTripper{Wrapped: httpClient.Transport, scope: scope}

	restClient := githubv3.NewClient(httpClient)
	ret.rest = v3client{
		Repositories:  restClient.Repositories,
		PullRequests:  restClient.PullRequests,
		Issues:        restClient.Issues,
		Users:         restClient.Users,
		Organizations: restClient.Organizations,
	}

	ret.graphQL = githubv4.NewClient(httpClient)

	return ret, nil
}

func (s *svc) GetFile(ctx context.Context, ref *RemoteRef, path string) (*File, error) {
	q := &getFileQuery{}
	params := map[string]interface{}{
		"owner":   githubv4.String(ref.RepoOwner),
		"name":    githubv4.String(ref.RepoName),
		"path":    githubv4.String(path),
		"ref":     githubv4.String(ref.Ref),
		"refPath": githubv4.String(fmt.Sprintf("%s:%s", ref.Ref, path)),
	}

	err := s.graphQL.Query(ctx, q, params)
	if err != nil {
		return nil, err
	}

	switch {
	case q.Repository.Ref.Commit.ID == nil:
		return nil, errors.New("ref not found")
	case q.Repository.Object.Blob.ID == nil:
		return nil, errors.New("object not found")
	case bool(q.Repository.Object.Blob.IsTruncated):
		return nil, errors.New("object was too large and was truncated by the API")
	case bool(q.Repository.Object.Blob.IsBinary):
		return nil, errors.New("object is a binary object and cannot be retrieved directly via the API")
	}

	f := &File{
		Path:     path,
		Contents: ioutil.NopCloser(strings.NewReader(string(q.Repository.Object.Blob.Text))),
		SHA:      string(q.Repository.Object.Blob.OID),
	}
	if len(q.Repository.Ref.Commit.History.Nodes) > 0 {
		f.LastModifiedTime = q.Repository.Ref.Commit.History.Nodes[0].CommittedDate.Time
		f.LastModifiedSHA = string(q.Repository.Ref.Commit.History.Nodes[0].OID)
	}

	return f, nil
}

/*
 * Rather than calling GetCommit() multiple times, we can use CompareCommits to get a range of commits
 */
func (s *svc) CompareCommits(ctx context.Context, ref *RemoteRef, compareSHA string) (*githubv3.CommitsComparison, error) {
	comp, _, err := s.rest.Repositories.CompareCommits(ctx, ref.RepoOwner, ref.RepoName, compareSHA, ref.Ref)
	if err != nil {
		return nil, fmt.Errorf("Could not get comparison for %s and %s. %+v", ref.Ref, compareSHA, err)
	}

	return comp, nil
}

type Commit struct {
	Files     []*githubv3.CommitFile
	Message   string
	Author    *githubv3.User
	ParentRef string
}

func (s *svc) GetCommit(ctx context.Context, ref *RemoteRef) (*Commit, error) {
	commit, _, err := s.rest.Repositories.GetCommit(ctx, ref.RepoOwner, ref.RepoName, ref.Ref)
	if err != nil {
		return nil, err
	}

	// Currently we are using the Author (Github) rather than commit Author (Git)
	retCommit := &Commit{
		Files:   commit.Files,
		Message: commit.GetCommit().GetMessage(),
		Author:  commit.GetAuthor(),
	}

	if commit.Parents != nil && len(commit.Parents) > 0 {
		retCommit.ParentRef = commit.Parents[0].GetSHA()
	}

	return retCommit, nil
}

func (s *svc) GetRepository(ctx context.Context, repo *RemoteRef) (*Repository, error) {
	q := &getRepositoryQuery{}
	params := map[string]interface{}{
		"owner": githubv4.String(repo.RepoOwner),
		"name":  githubv4.String(repo.RepoName),
	}

	err := s.graphQL.Query(ctx, q, params)
	if err != nil {
		return nil, err
	}

	r := &Repository{
		Name:          repo.RepoName,
		Owner:         repo.RepoOwner,
		DefaultBranch: q.Repository.DefaultBranchRef.Name,
	}

	return r, nil
}
