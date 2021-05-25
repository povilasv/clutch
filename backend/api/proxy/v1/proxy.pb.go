// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: proxy/v1/proxy.proto

package proxyv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RequestProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of a service that is configured
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// The HTTP method that will be used for the request
	HttpMethod string `protobuf:"bytes,2,opt,name=http_method,json=httpMethod,proto3" json:"http_method,omitempty"`
	// The URI path to call
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// The request body
	Request *structpb.Value `protobuf:"bytes,4,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *RequestProxyRequest) Reset() {
	*x = RequestProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_v1_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestProxyRequest) ProtoMessage() {}

func (x *RequestProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestProxyRequest.ProtoReflect.Descriptor instead.
func (*RequestProxyRequest) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *RequestProxyRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RequestProxyRequest) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

func (x *RequestProxyRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RequestProxyRequest) GetRequest() *structpb.Value {
	if x != nil {
		return x.Request
	}
	return nil
}

type RequestProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The http response code from the service response
	HttpStatus int32 `protobuf:"varint,1,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	// All of the headers from teh service response
	Headers map[string]string `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The full response body
	Response *structpb.Value `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RequestProxyResponse) Reset() {
	*x = RequestProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_v1_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestProxyResponse) ProtoMessage() {}

func (x *RequestProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_v1_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestProxyResponse.ProtoReflect.Descriptor instead.
func (*RequestProxyResponse) Descriptor() ([]byte, []int) {
	return file_proxy_v1_proxy_proto_rawDescGZIP(), []int{1}
}

func (x *RequestProxyResponse) GetHttpStatus() int32 {
	if x != nil {
		return x.HttpStatus
	}
	return 0
}

func (x *RequestProxyResponse) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *RequestProxyResponse) GetResponse() *structpb.Value {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_proxy_v1_proxy_proto protoreflect.FileDescriptor

var file_proxy_v1_proxy_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x28, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xf5, 0x01, 0x0a, 0x14, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x32, 0x8b, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x41, 0x50, 0x49, 0x12,
	0x7f, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12,
	0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x01,
	0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proxy_v1_proxy_proto_rawDescOnce sync.Once
	file_proxy_v1_proxy_proto_rawDescData = file_proxy_v1_proxy_proto_rawDesc
)

func file_proxy_v1_proxy_proto_rawDescGZIP() []byte {
	file_proxy_v1_proxy_proto_rawDescOnce.Do(func() {
		file_proxy_v1_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_v1_proxy_proto_rawDescData)
	})
	return file_proxy_v1_proxy_proto_rawDescData
}

var file_proxy_v1_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_v1_proxy_proto_goTypes = []interface{}{
	(*RequestProxyRequest)(nil),  // 0: clutch.proxy.v1.RequestProxyRequest
	(*RequestProxyResponse)(nil), // 1: clutch.proxy.v1.RequestProxyResponse
	nil,                          // 2: clutch.proxy.v1.RequestProxyResponse.HeadersEntry
	(*structpb.Value)(nil),       // 3: google.protobuf.Value
}
var file_proxy_v1_proxy_proto_depIdxs = []int32{
	3, // 0: clutch.proxy.v1.RequestProxyRequest.request:type_name -> google.protobuf.Value
	2, // 1: clutch.proxy.v1.RequestProxyResponse.headers:type_name -> clutch.proxy.v1.RequestProxyResponse.HeadersEntry
	3, // 2: clutch.proxy.v1.RequestProxyResponse.response:type_name -> google.protobuf.Value
	0, // 3: clutch.proxy.v1.ProxyAPI.RequestProxy:input_type -> clutch.proxy.v1.RequestProxyRequest
	1, // 4: clutch.proxy.v1.ProxyAPI.RequestProxy:output_type -> clutch.proxy.v1.RequestProxyResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proxy_v1_proxy_proto_init() }
func file_proxy_v1_proxy_proto_init() {
	if File_proxy_v1_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_v1_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestProxyRequest); i {
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
		file_proxy_v1_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestProxyResponse); i {
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
			RawDescriptor: file_proxy_v1_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proxy_v1_proxy_proto_goTypes,
		DependencyIndexes: file_proxy_v1_proxy_proto_depIdxs,
		MessageInfos:      file_proxy_v1_proxy_proto_msgTypes,
	}.Build()
	File_proxy_v1_proxy_proto = out.File
	file_proxy_v1_proxy_proto_rawDesc = nil
	file_proxy_v1_proxy_proto_goTypes = nil
	file_proxy_v1_proxy_proto_depIdxs = nil
}
