// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: chaos/experimentation/v1/create_experiment_data.proto

package experimentationv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The data used as an input for experiment creation.
type CreateExperimentData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique identifier of experiment run that's created as part of the experiment creation process.
	// A random run identifier is generated and assigned to the experiment if it's not provided by a caller.
	// The identifier is supposed to be user-readable and URL renderable - for this reason, allowed characters
	// are limited to English characters, digits and the following special characters: "-", ".", "_" and "~".
	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// The experiment configuration specific to the type of experiment.
	Config *anypb.Any `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The time when the experiment should start. If not provided, defaults to 'now'. It cannot be in the past.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// The time when the experiment should end, If not provided, the experiment runs until it's manually stopped.
	// If provided, it has to be after `start_time`.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *CreateExperimentData) Reset() {
	*x = CreateExperimentData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaos_experimentation_v1_create_experiment_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateExperimentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateExperimentData) ProtoMessage() {}

func (x *CreateExperimentData) ProtoReflect() protoreflect.Message {
	mi := &file_chaos_experimentation_v1_create_experiment_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateExperimentData.ProtoReflect.Descriptor instead.
func (*CreateExperimentData) Descriptor() ([]byte, []int) {
	return file_chaos_experimentation_v1_create_experiment_data_proto_rawDescGZIP(), []int{0}
}

func (x *CreateExperimentData) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *CreateExperimentData) GetConfig() *anypb.Any {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *CreateExperimentData) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *CreateExperimentData) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_chaos_experimentation_v1_create_experiment_data_proto protoreflect.FileDescriptor

var file_chaos_experimentation_v1_create_experiment_data_proto_rawDesc = []byte{
	0x0a, 0x35, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x18, 0x64, 0x32,
	0x12, 0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x2e, 0x5f, 0x7e,
	0x5d, 0x2a, 0x24, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73,
	0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaos_experimentation_v1_create_experiment_data_proto_rawDescOnce sync.Once
	file_chaos_experimentation_v1_create_experiment_data_proto_rawDescData = file_chaos_experimentation_v1_create_experiment_data_proto_rawDesc
)

func file_chaos_experimentation_v1_create_experiment_data_proto_rawDescGZIP() []byte {
	file_chaos_experimentation_v1_create_experiment_data_proto_rawDescOnce.Do(func() {
		file_chaos_experimentation_v1_create_experiment_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaos_experimentation_v1_create_experiment_data_proto_rawDescData)
	})
	return file_chaos_experimentation_v1_create_experiment_data_proto_rawDescData
}

var file_chaos_experimentation_v1_create_experiment_data_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chaos_experimentation_v1_create_experiment_data_proto_goTypes = []interface{}{
	(*CreateExperimentData)(nil),  // 0: clutch.chaos.experimentation.v1.CreateExperimentData
	(*anypb.Any)(nil),             // 1: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_chaos_experimentation_v1_create_experiment_data_proto_depIdxs = []int32{
	1, // 0: clutch.chaos.experimentation.v1.CreateExperimentData.config:type_name -> google.protobuf.Any
	2, // 1: clutch.chaos.experimentation.v1.CreateExperimentData.start_time:type_name -> google.protobuf.Timestamp
	2, // 2: clutch.chaos.experimentation.v1.CreateExperimentData.end_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chaos_experimentation_v1_create_experiment_data_proto_init() }
func file_chaos_experimentation_v1_create_experiment_data_proto_init() {
	if File_chaos_experimentation_v1_create_experiment_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaos_experimentation_v1_create_experiment_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateExperimentData); i {
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
			RawDescriptor: file_chaos_experimentation_v1_create_experiment_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaos_experimentation_v1_create_experiment_data_proto_goTypes,
		DependencyIndexes: file_chaos_experimentation_v1_create_experiment_data_proto_depIdxs,
		MessageInfos:      file_chaos_experimentation_v1_create_experiment_data_proto_msgTypes,
	}.Build()
	File_chaos_experimentation_v1_create_experiment_data_proto = out.File
	file_chaos_experimentation_v1_create_experiment_data_proto_rawDesc = nil
	file_chaos_experimentation_v1_create_experiment_data_proto_goTypes = nil
	file_chaos_experimentation_v1_create_experiment_data_proto_depIdxs = nil
}
