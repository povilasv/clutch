// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: authz/v1/authz.proto

package authzv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/lyft/clutch/backend/api/api/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Decision int32

const (
	Decision_UNSPECIFIED Decision = 0
	Decision_DENY        Decision = 1
	Decision_ALLOW       Decision = 2
)

// Enum value maps for Decision.
var (
	Decision_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DENY",
		2: "ALLOW",
	}
	Decision_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DENY":        1,
		"ALLOW":       2,
	}
)

func (x Decision) Enum() *Decision {
	p := new(Decision)
	*p = x
	return p
}

func (x Decision) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Decision) Descriptor() protoreflect.EnumDescriptor {
	return file_authz_v1_authz_proto_enumTypes[0].Descriptor()
}

func (Decision) Type() protoreflect.EnumType {
	return &file_authz_v1_authz_proto_enumTypes[0]
}

func (x Decision) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Decision.Descriptor instead.
func (Decision) EnumDescriptor() ([]byte, []int) {
	return file_authz_v1_authz_proto_rawDescGZIP(), []int{0}
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Groups []string `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_v1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_authz_v1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_authz_v1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *Subject) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Subject) GetGroups() []string {
	if x != nil {
		return x.Groups
	}
	return nil
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject    *Subject      `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Method     string        `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	ActionType v1.ActionType `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3,enum=clutch.api.v1.ActionType" json:"action_type,omitempty"`
	Resource   string        `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_v1_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_authz_v1_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_authz_v1_authz_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRequest) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *CheckRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CheckRequest) GetActionType() v1.ActionType {
	if x != nil {
		return x.ActionType
	}
	return v1.ActionType(0)
}

func (x *CheckRequest) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Decision Decision `protobuf:"varint,1,opt,name=decision,proto3,enum=clutch.authz.v1.Decision" json:"decision,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authz_v1_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_authz_v1_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_authz_v1_authz_proto_rawDescGZIP(), []int{2}
}

func (x *CheckResponse) GetDecision() Decision {
	if x != nil {
		return x.Decision
	}
	return Decision_UNSPECIFIED
}

var File_authz_v1_authz_proto protoreflect.FileDescriptor

var file_authz_v1_authz_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a,
	0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x0b, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x46, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x30, 0x0a, 0x08, 0x44,
	0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x4e, 0x59,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x32, 0x74, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x41, 0x50, 0x49, 0x12, 0x68, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c,
	0x02, 0x08, 0x02, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_authz_v1_authz_proto_rawDescOnce sync.Once
	file_authz_v1_authz_proto_rawDescData = file_authz_v1_authz_proto_rawDesc
)

func file_authz_v1_authz_proto_rawDescGZIP() []byte {
	file_authz_v1_authz_proto_rawDescOnce.Do(func() {
		file_authz_v1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authz_v1_authz_proto_rawDescData)
	})
	return file_authz_v1_authz_proto_rawDescData
}

var file_authz_v1_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authz_v1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_authz_v1_authz_proto_goTypes = []interface{}{
	(Decision)(0),         // 0: clutch.authz.v1.Decision
	(*Subject)(nil),       // 1: clutch.authz.v1.Subject
	(*CheckRequest)(nil),  // 2: clutch.authz.v1.CheckRequest
	(*CheckResponse)(nil), // 3: clutch.authz.v1.CheckResponse
	(v1.ActionType)(0),    // 4: clutch.api.v1.ActionType
}
var file_authz_v1_authz_proto_depIdxs = []int32{
	1, // 0: clutch.authz.v1.CheckRequest.subject:type_name -> clutch.authz.v1.Subject
	4, // 1: clutch.authz.v1.CheckRequest.action_type:type_name -> clutch.api.v1.ActionType
	0, // 2: clutch.authz.v1.CheckResponse.decision:type_name -> clutch.authz.v1.Decision
	2, // 3: clutch.authz.v1.AuthzAPI.Check:input_type -> clutch.authz.v1.CheckRequest
	3, // 4: clutch.authz.v1.AuthzAPI.Check:output_type -> clutch.authz.v1.CheckResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_authz_v1_authz_proto_init() }
func file_authz_v1_authz_proto_init() {
	if File_authz_v1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authz_v1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_authz_v1_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_authz_v1_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
			RawDescriptor: file_authz_v1_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_authz_v1_authz_proto_goTypes,
		DependencyIndexes: file_authz_v1_authz_proto_depIdxs,
		EnumInfos:         file_authz_v1_authz_proto_enumTypes,
		MessageInfos:      file_authz_v1_authz_proto_msgTypes,
	}.Build()
	File_authz_v1_authz_proto = out.File
	file_authz_v1_authz_proto_rawDesc = nil
	file_authz_v1_authz_proto_goTypes = nil
	file_authz_v1_authz_proto_depIdxs = nil
}
