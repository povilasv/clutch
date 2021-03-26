// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: k8s/v1/status.proto

package k8sv1

import (
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

// Kubernetes protos are generated using non-standard mechanisms, and the Go objects panic when serializing
// with the v2 proto APIs. These are just repetitions of the structs from K8s metav1 package so they are compliant
// with proto v2 serialization.
// https://github.com/kubernetes/apimachinery/blob/44b9a379dc1834b9cee463d1da4dc3e9d7302b35/pkg/apis/meta/v1/types.go#L620-L650
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Reason  string         `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty"`
	Code    int32          `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Details *StatusDetails `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_v1_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_v1_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_k8s_v1_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetDetails() *StatusDetails {
	if x != nil {
		return x.Details
	}
	return nil
}

type StatusDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Group  string         `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	Kind   string         `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Uid    string         `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
	Causes []*StatusCause `protobuf:"bytes,5,rep,name=causes,proto3" json:"causes,omitempty"`
}

func (x *StatusDetails) Reset() {
	*x = StatusDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_v1_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusDetails) ProtoMessage() {}

func (x *StatusDetails) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_v1_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusDetails.ProtoReflect.Descriptor instead.
func (*StatusDetails) Descriptor() ([]byte, []int) {
	return file_k8s_v1_status_proto_rawDescGZIP(), []int{1}
}

func (x *StatusDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StatusDetails) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *StatusDetails) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *StatusDetails) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *StatusDetails) GetCauses() []*StatusCause {
	if x != nil {
		return x.Causes
	}
	return nil
}

type StatusCause struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Field   string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *StatusCause) Reset() {
	*x = StatusCause{}
	if protoimpl.UnsafeEnabled {
		mi := &file_k8s_v1_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusCause) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusCause) ProtoMessage() {}

func (x *StatusCause) ProtoReflect() protoreflect.Message {
	mi := &file_k8s_v1_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusCause.ProtoReflect.Descriptor instead.
func (*StatusCause) Descriptor() ([]byte, []int) {
	return file_k8s_v1_status_proto_rawDescGZIP(), []int{2}
}

func (x *StatusCause) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *StatusCause) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StatusCause) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_k8s_v1_status_proto protoreflect.FileDescriptor

var file_k8s_v1_status_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x76, 0x31, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x36, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x63, 0x61, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x61,
	0x75, 0x73, 0x65, 0x52, 0x06, 0x63, 0x61, 0x75, 0x73, 0x65, 0x73, 0x22, 0x51, 0x0a, 0x0b, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66,
	0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38, 0x73, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_k8s_v1_status_proto_rawDescOnce sync.Once
	file_k8s_v1_status_proto_rawDescData = file_k8s_v1_status_proto_rawDesc
)

func file_k8s_v1_status_proto_rawDescGZIP() []byte {
	file_k8s_v1_status_proto_rawDescOnce.Do(func() {
		file_k8s_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_k8s_v1_status_proto_rawDescData)
	})
	return file_k8s_v1_status_proto_rawDescData
}

var file_k8s_v1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_k8s_v1_status_proto_goTypes = []interface{}{
	(*Status)(nil),        // 0: clutch.k8s.v1.Status
	(*StatusDetails)(nil), // 1: clutch.k8s.v1.StatusDetails
	(*StatusCause)(nil),   // 2: clutch.k8s.v1.StatusCause
}
var file_k8s_v1_status_proto_depIdxs = []int32{
	1, // 0: clutch.k8s.v1.Status.details:type_name -> clutch.k8s.v1.StatusDetails
	2, // 1: clutch.k8s.v1.StatusDetails.causes:type_name -> clutch.k8s.v1.StatusCause
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_k8s_v1_status_proto_init() }
func file_k8s_v1_status_proto_init() {
	if File_k8s_v1_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_k8s_v1_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_k8s_v1_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusDetails); i {
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
		file_k8s_v1_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusCause); i {
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
			RawDescriptor: file_k8s_v1_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_k8s_v1_status_proto_goTypes,
		DependencyIndexes: file_k8s_v1_status_proto_depIdxs,
		MessageInfos:      file_k8s_v1_status_proto_msgTypes,
	}.Build()
	File_k8s_v1_status_proto = out.File
	file_k8s_v1_status_proto_rawDesc = nil
	file_k8s_v1_status_proto_goTypes = nil
	file_k8s_v1_status_proto_depIdxs = nil
}
