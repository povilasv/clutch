// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: bot/slackbot/v1/slackbot.proto

package slackbotv1

import (
	_ "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Information on the bot user
type Bot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// bot id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// if the bot is deleted
	Deleted bool `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// bot name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// timestamp of when the bot app was last updated
	Updated int64 `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty"`
	// unique identifier of the installed Slack application
	AppId string `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// image urls of the bot's display picture
	Icons map[string]string `protobuf:"bytes,6,rep,name=icons,proto3" json:"icons,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// unique identifier of the workspace where the event occurred
	TeamId string `protobuf:"bytes,7,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
}

func (x *Bot) Reset() {
	*x = Bot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bot) ProtoMessage() {}

func (x *Bot) ProtoReflect() protoreflect.Message {
	mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bot.ProtoReflect.Descriptor instead.
func (*Bot) Descriptor() ([]byte, []int) {
	return file_bot_slackbot_v1_slackbot_proto_rawDescGZIP(), []int{0}
}

func (x *Bot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bot) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *Bot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bot) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

func (x *Bot) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Bot) GetIcons() map[string]string {
	if x != nil {
		return x.Icons
	}
	return nil
}

func (x *Bot) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

// The structure of events vary among types. Full list of event types: https://api.slack.com/events.
// We will be receiving app_mention events (events that mention the bot) or message.im events (a message posted in a DM
// with the bot).
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type of event
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// the user id of the user who messaged the bot
	User  string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	BotId string `protobuf:"bytes,3,opt,name=bot_id,json=botId,proto3" json:"bot_id,omitempty"`
	// metadata on the bot
	BotProfile *Bot `protobuf:"bytes,4,opt,name=bot_profile,json=botProfile,proto3" json:"bot_profile,omitempty"`
	// the message text
	Text string `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	// timestamp associated with the object the event is describing
	Ts string `protobuf:"bytes,6,opt,name=ts,proto3" json:"ts,omitempty"`
	// the channel id of the channel where the event happened
	Channel string `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	// for DM messages, this will be "im"
	ChannelType string `protobuf:"bytes,8,opt,name=channel_type,json=channelType,proto3" json:"channel_type,omitempty"`
	// timestamp associated with the streamed event
	EventTs string `protobuf:"bytes,9,opt,name=event_ts,json=eventTs,proto3" json:"event_ts,omitempty"`
	// unclear what this field is but it's sent as part of the request from the Events API. Seems like it's a mistake
	// and it hasn't been addressed https://github.com/slackapi/python-slack-sdk/issues/736, so we have to support it for
	// now.
	ClientMsgId string `protobuf:"bytes,10,opt,name=client_msg_id,json=clientMsgId,proto3" json:"client_msg_id,omitempty"`
	// unique identifier of the workspace where the event occurred
	// identical to the team field value sent in the outer layer of the request
	Team string `protobuf:"bytes,11,opt,name=team,proto3" json:"team,omitempty"`
	// received when a user interacts with a Block Kit component, schema can vary.
	// https://api.slack.com/reference/block-kit/interactive-components
	Blocks *structpb.Value `protobuf:"bytes,12,opt,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_bot_slackbot_v1_slackbot_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Event) GetBotId() string {
	if x != nil {
		return x.BotId
	}
	return ""
}

func (x *Event) GetBotProfile() *Bot {
	if x != nil {
		return x.BotProfile
	}
	return nil
}

func (x *Event) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Event) GetTs() string {
	if x != nil {
		return x.Ts
	}
	return ""
}

func (x *Event) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *Event) GetChannelType() string {
	if x != nil {
		return x.ChannelType
	}
	return ""
}

func (x *Event) GetEventTs() string {
	if x != nil {
		return x.EventTs
	}
	return ""
}

func (x *Event) GetClientMsgId() string {
	if x != nil {
		return x.ClientMsgId
	}
	return ""
}

func (x *Event) GetTeam() string {
	if x != nil {
		return x.Team
	}
	return ""
}

func (x *Event) GetBlocks() *structpb.Value {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// For more details:
// https://api.slack.com/apis/connections/events-api#the-events-api__receiving-events
// https://api.slack.com/enterprise/apps/reference#event
type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// verification token to validate the event originated from Slack
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// if where the event occured is an external shared channel
	IsExtSharedChannel bool `protobuf:"varint,2,opt,name=is_ext_shared_channel,json=isExtSharedChannel,proto3" json:"is_ext_shared_channel,omitempty"`
	// unique identifier of the workspace where the event occurred
	TeamId string `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	// unique identifier of the installed Slack application
	ApiAppId string `protobuf:"bytes,4,opt,name=api_app_id,json=apiAppId,proto3" json:"api_app_id,omitempty"`
	// TODO: (sperry) if we expand the types of events we suscribe to and their fields vary alot, maybe we should use
	// google.protobuf.Value.
	Event *Event `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	// indicates which kind of event this is
	Type string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
	// identifier for this specific event, globally unique across all workspaces
	EventId string `protobuf:"bytes,7,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// epoch timestamp in seconds indicating when this event was dispatched
	EventTime int64 `protobuf:"varint,8,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// an identifier for this specific event, can be used with list of authorizations to obtain a full list of
	// installations of your app that this event is visible to
	EventContext string `protobuf:"bytes,9,opt,name=event_context,json=eventContext,proto3" json:"event_context,omitempty"`
	// describes the installation of the app that the event is visible to
	// https://api.slack.com/apis/connections/events-api#authorizations
	Authorizations *structpb.Value `protobuf:"bytes,10,opt,name=authorizations,proto3" json:"authorizations,omitempty"`
	// randomly generated string used as part of the URL verification handshake,
	// https://api.slack.com/apis/connections/events-api#the-events-api__subscribing-to-event-types__events-api-request-urls__request-url-configuration--verification__url-verification-handshake
	Challenge string `protobuf:"bytes,11,opt,name=challenge,proto3" json:"challenge,omitempty"`
	// sent in the request if we receive more than 30,000 events in 60 minutes
	// https://api.slack.com/apis/connections/events-api#the-events-api__responding-to-events__rate-limiting
	MinuteRateLimited string `protobuf:"bytes,12,opt,name=minute_rate_limited,json=minuteRateLimited,proto3" json:"minute_rate_limited,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_bot_slackbot_v1_slackbot_proto_rawDescGZIP(), []int{2}
}

func (x *EventRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *EventRequest) GetIsExtSharedChannel() bool {
	if x != nil {
		return x.IsExtSharedChannel
	}
	return false
}

func (x *EventRequest) GetTeamId() string {
	if x != nil {
		return x.TeamId
	}
	return ""
}

func (x *EventRequest) GetApiAppId() string {
	if x != nil {
		return x.ApiAppId
	}
	return ""
}

func (x *EventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *EventRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *EventRequest) GetEventTime() int64 {
	if x != nil {
		return x.EventTime
	}
	return 0
}

func (x *EventRequest) GetEventContext() string {
	if x != nil {
		return x.EventContext
	}
	return ""
}

func (x *EventRequest) GetAuthorizations() *structpb.Value {
	if x != nil {
		return x.Authorizations
	}
	return nil
}

func (x *EventRequest) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *EventRequest) GetMinuteRateLimited() string {
	if x != nil {
		return x.MinuteRateLimited
	}
	return ""
}

// Respond back to the Slack Events API with the challenge or a 2xx,
// https://api.slack.com/apis/connections/events-api#the-events-api__responding-to-events
type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the request will include the challenge and we respond back with the same challenge to complete the URL verification
	// handshake
	Challenge string `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_slackbot_v1_slackbot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_bot_slackbot_v1_slackbot_proto_rawDescGZIP(), []int{3}
}

func (x *EventResponse) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

var File_bot_slackbot_v1_slackbot_proto protoreflect.FileDescriptor

var file_bot_slackbot_v1_slackbot_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x62, 0x6f, 0x74, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x16, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x73, 0x6c, 0x61,
	0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85,
	0x02, 0x0a, 0x03, 0x42, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x05, 0x69, 0x63, 0x6f, 0x6e, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f,
	0x74, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x74, 0x2e, 0x49, 0x63, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x69, 0x63,
	0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x1a, 0x38, 0x0a, 0x0a,
	0x49, 0x63, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xe8, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x62, 0x6f, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x74, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x0b, 0x62, 0x6f, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f,
	0x74, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x74, 0x52, 0x0a, 0x62, 0x6f, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x73, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x61, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x61,
	0x6d, 0x12, 0x2e, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x22, 0xca, 0x03, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x31, 0x0a, 0x15, 0x69, 0x73, 0x5f, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x69, 0x73, 0x45, 0x78, 0x74, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65,
	0x61, 0x6d, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x61, 0x70, 0x69, 0x5f, 0x61, 0x70, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x70, 0x69, 0x41, 0x70, 0x70,
	0x49, 0x64, 0x12, 0x33, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x73,
	0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6d, 0x69, 0x6e, 0x75,
	0x74, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x3a, 0x04, 0xb8, 0xe1, 0x1c, 0x01, 0x22, 0x33,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x3a, 0x04, 0xb8,
	0xe1, 0x1c, 0x01, 0x32, 0x8c, 0x01, 0x0a, 0x0b, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x42, 0x6f, 0x74,
	0x41, 0x50, 0x49, 0x12, 0x7d, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f, 0x74, 0x2e, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62,
	0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x62, 0x6f, 0x74, 0x2e,
	0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x22, 0x16, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x73, 0x6c, 0x61, 0x63, 0x6b,
	0x62, 0x6f, 0x74, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02,
	0x08, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x6f, 0x74, 0x2f, 0x73, 0x6c, 0x61,
	0x63, 0x6b, 0x62, 0x6f, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f,
	0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bot_slackbot_v1_slackbot_proto_rawDescOnce sync.Once
	file_bot_slackbot_v1_slackbot_proto_rawDescData = file_bot_slackbot_v1_slackbot_proto_rawDesc
)

func file_bot_slackbot_v1_slackbot_proto_rawDescGZIP() []byte {
	file_bot_slackbot_v1_slackbot_proto_rawDescOnce.Do(func() {
		file_bot_slackbot_v1_slackbot_proto_rawDescData = protoimpl.X.CompressGZIP(file_bot_slackbot_v1_slackbot_proto_rawDescData)
	})
	return file_bot_slackbot_v1_slackbot_proto_rawDescData
}

var file_bot_slackbot_v1_slackbot_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bot_slackbot_v1_slackbot_proto_goTypes = []interface{}{
	(*Bot)(nil),            // 0: clutch.bot.slackbot.v1.Bot
	(*Event)(nil),          // 1: clutch.bot.slackbot.v1.Event
	(*EventRequest)(nil),   // 2: clutch.bot.slackbot.v1.EventRequest
	(*EventResponse)(nil),  // 3: clutch.bot.slackbot.v1.EventResponse
	nil,                    // 4: clutch.bot.slackbot.v1.Bot.IconsEntry
	(*structpb.Value)(nil), // 5: google.protobuf.Value
}
var file_bot_slackbot_v1_slackbot_proto_depIdxs = []int32{
	4, // 0: clutch.bot.slackbot.v1.Bot.icons:type_name -> clutch.bot.slackbot.v1.Bot.IconsEntry
	0, // 1: clutch.bot.slackbot.v1.Event.bot_profile:type_name -> clutch.bot.slackbot.v1.Bot
	5, // 2: clutch.bot.slackbot.v1.Event.blocks:type_name -> google.protobuf.Value
	1, // 3: clutch.bot.slackbot.v1.EventRequest.event:type_name -> clutch.bot.slackbot.v1.Event
	5, // 4: clutch.bot.slackbot.v1.EventRequest.authorizations:type_name -> google.protobuf.Value
	2, // 5: clutch.bot.slackbot.v1.SlackBotAPI.Event:input_type -> clutch.bot.slackbot.v1.EventRequest
	3, // 6: clutch.bot.slackbot.v1.SlackBotAPI.Event:output_type -> clutch.bot.slackbot.v1.EventResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_bot_slackbot_v1_slackbot_proto_init() }
func file_bot_slackbot_v1_slackbot_proto_init() {
	if File_bot_slackbot_v1_slackbot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bot_slackbot_v1_slackbot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_slackbot_v1_slackbot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_slackbot_v1_slackbot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bot_slackbot_v1_slackbot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bot_slackbot_v1_slackbot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bot_slackbot_v1_slackbot_proto_goTypes,
		DependencyIndexes: file_bot_slackbot_v1_slackbot_proto_depIdxs,
		MessageInfos:      file_bot_slackbot_v1_slackbot_proto_msgTypes,
	}.Build()
	File_bot_slackbot_v1_slackbot_proto = out.File
	file_bot_slackbot_v1_slackbot_proto_rawDesc = nil
	file_bot_slackbot_v1_slackbot_proto_goTypes = nil
	file_bot_slackbot_v1_slackbot_proto_depIdxs = nil
}
