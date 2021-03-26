// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: envoytriage/v1/output.proto

package envoytriagev1

import (
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

type HostStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Healthy bool   `protobuf:"varint,2,opt,name=healthy,proto3" json:"healthy,omitempty"`
}

func (x *HostStatus) Reset() {
	*x = HostStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostStatus) ProtoMessage() {}

func (x *HostStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostStatus.ProtoReflect.Descriptor instead.
func (*HostStatus) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{0}
}

func (x *HostStatus) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *HostStatus) GetHealthy() bool {
	if x != nil {
		return x.Healthy
	}
	return false
}

type ClusterStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	HostStatuses []*HostStatus `protobuf:"bytes,2,rep,name=host_statuses,json=hostStatuses,proto3" json:"host_statuses,omitempty"`
}

func (x *ClusterStatus) Reset() {
	*x = ClusterStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterStatus) ProtoMessage() {}

func (x *ClusterStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterStatus.ProtoReflect.Descriptor instead.
func (*ClusterStatus) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClusterStatus) GetHostStatuses() []*HostStatus {
	if x != nil {
		return x.HostStatuses
	}
	return nil
}

type Clusters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterStatuses []*ClusterStatus `protobuf:"bytes,1,rep,name=cluster_statuses,json=clusterStatuses,proto3" json:"cluster_statuses,omitempty"`
}

func (x *Clusters) Reset() {
	*x = Clusters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Clusters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Clusters) ProtoMessage() {}

func (x *Clusters) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Clusters.ProtoReflect.Descriptor instead.
func (*Clusters) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{2}
}

func (x *Clusters) GetClusterStatuses() []*ClusterStatus {
	if x != nil {
		return x.ClusterStatuses
	}
	return nil
}

type ConfigDump struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *structpb.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigDump) Reset() {
	*x = ConfigDump{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDump) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDump) ProtoMessage() {}

func (x *ConfigDump) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDump.ProtoReflect.Descriptor instead.
func (*ConfigDump) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{3}
}

func (x *ConfigDump) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type ListenerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LocalAddress string `protobuf:"bytes,2,opt,name=local_address,json=localAddress,proto3" json:"local_address,omitempty"`
}

func (x *ListenerStatus) Reset() {
	*x = ListenerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenerStatus) ProtoMessage() {}

func (x *ListenerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenerStatus.ProtoReflect.Descriptor instead.
func (*ListenerStatus) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{4}
}

func (x *ListenerStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListenerStatus) GetLocalAddress() string {
	if x != nil {
		return x.LocalAddress
	}
	return ""
}

type Listeners struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerStatuses []*ListenerStatus `protobuf:"bytes,1,rep,name=listener_statuses,json=listenerStatuses,proto3" json:"listener_statuses,omitempty"`
}

func (x *Listeners) Reset() {
	*x = Listeners{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listeners) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listeners) ProtoMessage() {}

func (x *Listeners) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listeners.ProtoReflect.Descriptor instead.
func (*Listeners) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{5}
}

func (x *Listeners) GetListenerStatuses() []*ListenerStatus {
	if x != nil {
		return x.ListenerStatuses
	}
	return nil
}

type Runtime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Runtime_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Runtime) Reset() {
	*x = Runtime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runtime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runtime) ProtoMessage() {}

func (x *Runtime) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runtime.ProtoReflect.Descriptor instead.
func (*Runtime) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{6}
}

func (x *Runtime) GetEntries() []*Runtime_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *structpb.Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{7}
}

func (x *ServerInfo) GetValue() *structpb.Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Counters and gauges are returned here.
	Stats []*Stats_Stat `protobuf:"bytes,1,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{8}
}

func (x *Stats) GetStats() []*Stats_Stat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type Runtime_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Types that are assignable to Type:
	//	*Runtime_Entry_Value
	Type isRuntime_Entry_Type `protobuf_oneof:"type"`
}

func (x *Runtime_Entry) Reset() {
	*x = Runtime_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Runtime_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Runtime_Entry) ProtoMessage() {}

func (x *Runtime_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Runtime_Entry.ProtoReflect.Descriptor instead.
func (*Runtime_Entry) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Runtime_Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *Runtime_Entry) GetType() isRuntime_Entry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Runtime_Entry) GetValue() string {
	if x, ok := x.GetType().(*Runtime_Entry_Value); ok {
		return x.Value
	}
	return ""
}

type isRuntime_Entry_Type interface {
	isRuntime_Entry_Type()
}

type Runtime_Entry_Value struct {
	Value string `protobuf:"bytes,2,opt,name=value,proto3,oneof"`
}

func (*Runtime_Entry_Value) isRuntime_Entry_Type() {}

type Stats_Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Stats_Stat) Reset() {
	*x = Stats_Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoytriage_v1_output_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats_Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats_Stat) ProtoMessage() {}

func (x *Stats_Stat) ProtoReflect() protoreflect.Message {
	mi := &file_envoytriage_v1_output_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats_Stat.ProtoReflect.Descriptor instead.
func (*Stats_Stat) Descriptor() ([]byte, []int) {
	return file_envoytriage_v1_output_proto_rawDescGZIP(), []int{8, 0}
}

func (x *Stats_Stat) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Stats_Stat) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_envoytriage_v1_output_proto protoreflect.FileDescriptor

var file_envoytriage_v1_output_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x79, 0x22, 0x6b, 0x0a, 0x0d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x68, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74,
	0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x0c, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x22, 0x5b, 0x0a, 0x08, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4f, 0x0a,
	0x10, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x3a,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x75, 0x6d, 0x70, 0x12, 0x2c, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x0e, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5f, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x07, 0x52, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x1a, 0x39, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x0a,
	0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x70, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x37, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x2e, 0x0a, 0x04, 0x53,
	0x74, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x74, 0x72, 0x69, 0x61, 0x67, 0x65, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoytriage_v1_output_proto_rawDescOnce sync.Once
	file_envoytriage_v1_output_proto_rawDescData = file_envoytriage_v1_output_proto_rawDesc
)

func file_envoytriage_v1_output_proto_rawDescGZIP() []byte {
	file_envoytriage_v1_output_proto_rawDescOnce.Do(func() {
		file_envoytriage_v1_output_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoytriage_v1_output_proto_rawDescData)
	})
	return file_envoytriage_v1_output_proto_rawDescData
}

var file_envoytriage_v1_output_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_envoytriage_v1_output_proto_goTypes = []interface{}{
	(*HostStatus)(nil),     // 0: clutch.envoytriage.v1.HostStatus
	(*ClusterStatus)(nil),  // 1: clutch.envoytriage.v1.ClusterStatus
	(*Clusters)(nil),       // 2: clutch.envoytriage.v1.Clusters
	(*ConfigDump)(nil),     // 3: clutch.envoytriage.v1.ConfigDump
	(*ListenerStatus)(nil), // 4: clutch.envoytriage.v1.ListenerStatus
	(*Listeners)(nil),      // 5: clutch.envoytriage.v1.Listeners
	(*Runtime)(nil),        // 6: clutch.envoytriage.v1.Runtime
	(*ServerInfo)(nil),     // 7: clutch.envoytriage.v1.ServerInfo
	(*Stats)(nil),          // 8: clutch.envoytriage.v1.Stats
	(*Runtime_Entry)(nil),  // 9: clutch.envoytriage.v1.Runtime.Entry
	(*Stats_Stat)(nil),     // 10: clutch.envoytriage.v1.Stats.Stat
	(*structpb.Value)(nil), // 11: google.protobuf.Value
}
var file_envoytriage_v1_output_proto_depIdxs = []int32{
	0,  // 0: clutch.envoytriage.v1.ClusterStatus.host_statuses:type_name -> clutch.envoytriage.v1.HostStatus
	1,  // 1: clutch.envoytriage.v1.Clusters.cluster_statuses:type_name -> clutch.envoytriage.v1.ClusterStatus
	11, // 2: clutch.envoytriage.v1.ConfigDump.value:type_name -> google.protobuf.Value
	4,  // 3: clutch.envoytriage.v1.Listeners.listener_statuses:type_name -> clutch.envoytriage.v1.ListenerStatus
	9,  // 4: clutch.envoytriage.v1.Runtime.entries:type_name -> clutch.envoytriage.v1.Runtime.Entry
	11, // 5: clutch.envoytriage.v1.ServerInfo.value:type_name -> google.protobuf.Value
	10, // 6: clutch.envoytriage.v1.Stats.stats:type_name -> clutch.envoytriage.v1.Stats.Stat
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_envoytriage_v1_output_proto_init() }
func file_envoytriage_v1_output_proto_init() {
	if File_envoytriage_v1_output_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoytriage_v1_output_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostStatus); i {
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
		file_envoytriage_v1_output_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterStatus); i {
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
		file_envoytriage_v1_output_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Clusters); i {
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
		file_envoytriage_v1_output_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDump); i {
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
		file_envoytriage_v1_output_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenerStatus); i {
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
		file_envoytriage_v1_output_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listeners); i {
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
		file_envoytriage_v1_output_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runtime); i {
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
		file_envoytriage_v1_output_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerInfo); i {
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
		file_envoytriage_v1_output_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_envoytriage_v1_output_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Runtime_Entry); i {
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
		file_envoytriage_v1_output_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats_Stat); i {
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
	file_envoytriage_v1_output_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Runtime_Entry_Value)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoytriage_v1_output_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoytriage_v1_output_proto_goTypes,
		DependencyIndexes: file_envoytriage_v1_output_proto_depIdxs,
		MessageInfos:      file_envoytriage_v1_output_proto_msgTypes,
	}.Build()
	File_envoytriage_v1_output_proto = out.File
	file_envoytriage_v1_output_proto_rawDesc = nil
	file_envoytriage_v1_output_proto_goTypes = nil
	file_envoytriage_v1_output_proto_depIdxs = nil
}
