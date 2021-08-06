// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/github/v1/github.proto

package githubv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on AppConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AppConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetAppId()) < 1 {
		return AppConfigValidationError{
			field:  "AppId",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetInstallationId()) < 1 {
		return AppConfigValidationError{
			field:  "InstallationId",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetPem()) < 1 {
		return AppConfigValidationError{
			field:  "Pem",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// AppConfigValidationError is the validation error returned by
// AppConfig.Validate if the designated constraints aren't met.
type AppConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AppConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AppConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AppConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AppConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AppConfigValidationError) ErrorName() string { return "AppConfigValidationError" }

// Error satisfies the builtin error interface
func (e AppConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAppConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AppConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AppConfigValidationError{}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Auth.(type) {

	case *Config_AccessToken:

		if len(m.GetAccessToken()) < 1 {
			return ConfigValidationError{
				field:  "AccessToken",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *Config_AppConfig:

		if v, ok := interface{}(m.GetAppConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  "AppConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ConfigValidationError{
			field:  "Auth",
			reason: "value is required",
		}

	}

	return nil
}

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}
