// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: resolver/v1/schema.proto

package resolverv1

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type StringField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Placeholder  string `protobuf:"bytes,1,opt,name=placeholder,proto3" json:"placeholder,omitempty"`
	DefaultValue string `protobuf:"bytes,2,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
}

func (x *StringField) Reset() {
	*x = StringField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringField) ProtoMessage() {}

func (x *StringField) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringField.ProtoReflect.Descriptor instead.
func (*StringField) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{0}
}

func (x *StringField) GetPlaceholder() string {
	if x != nil {
		return x.Placeholder
	}
	return ""
}

func (x *StringField) GetDefaultValue() string {
	if x != nil {
		return x.DefaultValue
	}
	return ""
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Types that are assignable to Value:
	//	*Option_StringValue
	Value isOption_Value `protobuf_oneof:"value"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (m *Option) GetValue() isOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Option) GetStringValue() string {
	if x, ok := x.GetValue().(*Option_StringValue); ok {
		return x.StringValue
	}
	return ""
}

type isOption_Value interface {
	isOption_Value()
}

type Option_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*Option_StringValue) isOption_Value() {}

type OptionField struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncludeAllOption      bool      `protobuf:"varint,1,opt,name=include_all_option,json=includeAllOption,proto3" json:"include_all_option,omitempty"`
	IncludeDynamicOptions []string  `protobuf:"bytes,2,rep,name=include_dynamic_options,json=includeDynamicOptions,proto3" json:"include_dynamic_options,omitempty"`
	Options               []*Option `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (x *OptionField) Reset() {
	*x = OptionField{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OptionField) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OptionField) ProtoMessage() {}

func (x *OptionField) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OptionField.ProtoReflect.Descriptor instead.
func (*OptionField) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{2}
}

func (x *OptionField) GetIncludeAllOption() bool {
	if x != nil {
		return x.IncludeAllOption
	}
	return false
}

func (x *OptionField) GetIncludeDynamicOptions() []string {
	if x != nil {
		return x.IncludeDynamicOptions
	}
	return nil
}

func (x *OptionField) GetOptions() []*Option {
	if x != nil {
		return x.Options
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Metadata *FieldMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{3}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetMetadata() *FieldMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type FieldMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Required    bool   `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	// Types that are assignable to Type:
	//	*FieldMetadata_StringField
	//	*FieldMetadata_OptionField
	Type isFieldMetadata_Type `protobuf_oneof:"type"`
}

func (x *FieldMetadata) Reset() {
	*x = FieldMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMetadata) ProtoMessage() {}

func (x *FieldMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMetadata.ProtoReflect.Descriptor instead.
func (*FieldMetadata) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{4}
}

func (x *FieldMetadata) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *FieldMetadata) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (m *FieldMetadata) GetType() isFieldMetadata_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *FieldMetadata) GetStringField() *StringField {
	if x, ok := x.GetType().(*FieldMetadata_StringField); ok {
		return x.StringField
	}
	return nil
}

func (x *FieldMetadata) GetOptionField() *OptionField {
	if x, ok := x.GetType().(*FieldMetadata_OptionField); ok {
		return x.OptionField
	}
	return nil
}

type isFieldMetadata_Type interface {
	isFieldMetadata_Type()
}

type FieldMetadata_StringField struct {
	StringField *StringField `protobuf:"bytes,3,opt,name=string_field,json=stringField,proto3,oneof"`
}

type FieldMetadata_OptionField struct {
	OptionField *OptionField `protobuf:"bytes,4,opt,name=option_field,json=optionField,proto3,oneof"`
}

func (*FieldMetadata_StringField) isFieldMetadata_Type() {}

func (*FieldMetadata_OptionField) isFieldMetadata_Type() {}

type SearchMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// This field is controlled by the backend resolver
	// Acting as a feature flag to the frontend which enables the autocomplete components
	AutocompleteEnabled bool `protobuf:"varint,2,opt,name=autocomplete_enabled,json=autocompleteEnabled,proto3" json:"autocomplete_enabled,omitempty"`
}

func (x *SearchMetadata) Reset() {
	*x = SearchMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchMetadata) ProtoMessage() {}

func (x *SearchMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchMetadata.ProtoReflect.Descriptor instead.
func (*SearchMetadata) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{5}
}

func (x *SearchMetadata) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SearchMetadata) GetAutocompleteEnabled() bool {
	if x != nil {
		return x.AutocompleteEnabled
	}
	return false
}

type SchemaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DisplayName string `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Searchable has been replaced by SearchMetadata.enabled and will be deprecated soon
	Searchable bool            `protobuf:"varint,2,opt,name=searchable,proto3" json:"searchable,omitempty"`
	Search     *SearchMetadata `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *SchemaMetadata) Reset() {
	*x = SchemaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaMetadata) ProtoMessage() {}

func (x *SchemaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaMetadata.ProtoReflect.Descriptor instead.
func (*SchemaMetadata) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{6}
}

func (x *SchemaMetadata) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *SchemaMetadata) GetSearchable() bool {
	if x != nil {
		return x.Searchable
	}
	return false
}

func (x *SchemaMetadata) GetSearch() *SearchMetadata {
	if x != nil {
		return x.Search
	}
	return nil
}

type Schema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type URL of the object the schema was produced from, which becomes the 'have' type URL when submitting a
	// filled-in schema.
	TypeUrl  string          `protobuf:"bytes,1,opt,name=type_url,json=typeUrl,proto3" json:"type_url,omitempty"`
	Metadata *SchemaMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Fields   []*Field        `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty"`
	// If the schema is broken, e.g. a required option field is missing options, an error will be returned here to be
	// displayed when that schema is selected.
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Schema) Reset() {
	*x = Schema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resolver_v1_schema_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Schema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Schema) ProtoMessage() {}

func (x *Schema) ProtoReflect() protoreflect.Message {
	mi := &file_resolver_v1_schema_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Schema.ProtoReflect.Descriptor instead.
func (*Schema) Descriptor() ([]byte, []int) {
	return file_resolver_v1_schema_proto_rawDescGZIP(), []int{7}
}

func (x *Schema) GetTypeUrl() string {
	if x != nil {
		return x.TypeUrl
	}
	return ""
}

func (x *Schema) GetMetadata() *SchemaMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Schema) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Schema) GetError() *status.Status {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_resolver_v1_schema_proto protoreflect.FileDescriptor

var file_resolver_v1_schema_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6c, 0x75, 0x74,
	0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x59, 0x0a,
	0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa9, 0x01, 0x0a, 0x0b, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41, 0x6c, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x5f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34,
	0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0x5a, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0xe2, 0x01, 0x0a, 0x0d, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x44, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x48, 0x00,
	0x52, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x06, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5d, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x31, 0x0a, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x13, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x79, 0x70, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x3e, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x06,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75,
	0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resolver_v1_schema_proto_rawDescOnce sync.Once
	file_resolver_v1_schema_proto_rawDescData = file_resolver_v1_schema_proto_rawDesc
)

func file_resolver_v1_schema_proto_rawDescGZIP() []byte {
	file_resolver_v1_schema_proto_rawDescOnce.Do(func() {
		file_resolver_v1_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_resolver_v1_schema_proto_rawDescData)
	})
	return file_resolver_v1_schema_proto_rawDescData
}

var file_resolver_v1_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_resolver_v1_schema_proto_goTypes = []interface{}{
	(*StringField)(nil),    // 0: clutch.resolver.v1.StringField
	(*Option)(nil),         // 1: clutch.resolver.v1.Option
	(*OptionField)(nil),    // 2: clutch.resolver.v1.OptionField
	(*Field)(nil),          // 3: clutch.resolver.v1.Field
	(*FieldMetadata)(nil),  // 4: clutch.resolver.v1.FieldMetadata
	(*SearchMetadata)(nil), // 5: clutch.resolver.v1.SearchMetadata
	(*SchemaMetadata)(nil), // 6: clutch.resolver.v1.SchemaMetadata
	(*Schema)(nil),         // 7: clutch.resolver.v1.Schema
	(*status.Status)(nil),  // 8: google.rpc.Status
}
var file_resolver_v1_schema_proto_depIdxs = []int32{
	1, // 0: clutch.resolver.v1.OptionField.options:type_name -> clutch.resolver.v1.Option
	4, // 1: clutch.resolver.v1.Field.metadata:type_name -> clutch.resolver.v1.FieldMetadata
	0, // 2: clutch.resolver.v1.FieldMetadata.string_field:type_name -> clutch.resolver.v1.StringField
	2, // 3: clutch.resolver.v1.FieldMetadata.option_field:type_name -> clutch.resolver.v1.OptionField
	5, // 4: clutch.resolver.v1.SchemaMetadata.search:type_name -> clutch.resolver.v1.SearchMetadata
	6, // 5: clutch.resolver.v1.Schema.metadata:type_name -> clutch.resolver.v1.SchemaMetadata
	3, // 6: clutch.resolver.v1.Schema.fields:type_name -> clutch.resolver.v1.Field
	8, // 7: clutch.resolver.v1.Schema.error:type_name -> google.rpc.Status
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_resolver_v1_schema_proto_init() }
func file_resolver_v1_schema_proto_init() {
	if File_resolver_v1_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resolver_v1_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringField); i {
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
		file_resolver_v1_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_resolver_v1_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OptionField); i {
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
		file_resolver_v1_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_resolver_v1_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMetadata); i {
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
		file_resolver_v1_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchMetadata); i {
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
		file_resolver_v1_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaMetadata); i {
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
		file_resolver_v1_schema_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Schema); i {
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
	file_resolver_v1_schema_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Option_StringValue)(nil),
	}
	file_resolver_v1_schema_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*FieldMetadata_StringField)(nil),
		(*FieldMetadata_OptionField)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resolver_v1_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resolver_v1_schema_proto_goTypes,
		DependencyIndexes: file_resolver_v1_schema_proto_depIdxs,
		MessageInfos:      file_resolver_v1_schema_proto_msgTypes,
	}.Build()
	File_resolver_v1_schema_proto = out.File
	file_resolver_v1_schema_proto_rawDesc = nil
	file_resolver_v1_schema_proto_goTypes = nil
	file_resolver_v1_schema_proto_depIdxs = nil
}
