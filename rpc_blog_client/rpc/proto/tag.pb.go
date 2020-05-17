// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.2.0
// source: rpc/proto/tag.proto

package out

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Tag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedBy  string `protobuf:"bytes,3,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	ModifiedBy string `protobuf:"bytes,4,opt,name=modifiedBy,proto3" json:"modifiedBy,omitempty"`
	State      int32  `protobuf:"varint,5,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Tag) Reset() {
	*x = Tag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tag) ProtoMessage() {}

func (x *Tag) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tag.ProtoReflect.Descriptor instead.
func (*Tag) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{0}
}

func (x *Tag) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tag) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Tag) GetModifiedBy() string {
	if x != nil {
		return x.ModifiedBy
	}
	return ""
}

func (x *Tag) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

type GetTagsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32             `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32             `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Maps   map[string]string `protobuf:"bytes,3,rep,name=maps,proto3" json:"maps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetTagsRequest) Reset() {
	*x = GetTagsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagsRequest) ProtoMessage() {}

func (x *GetTagsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagsRequest.ProtoReflect.Descriptor instead.
func (*GetTagsRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{1}
}

func (x *GetTagsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetTagsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetTagsRequest) GetMaps() map[string]string {
	if x != nil {
		return x.Maps
	}
	return nil
}

type ExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exist bool  `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	Code  int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *ExistResponse) Reset() {
	*x = ExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistResponse) ProtoMessage() {}

func (x *ExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistResponse.ProtoReflect.Descriptor instead.
func (*ExistResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{2}
}

func (x *ExistResponse) GetExist() bool {
	if x != nil {
		return x.Exist
	}
	return false
}

func (x *ExistResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetTagResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag  *Tag  `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Code int32 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetTagResponse) Reset() {
	*x = GetTagResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagResponse) ProtoMessage() {}

func (x *GetTagResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagResponse.ProtoReflect.Descriptor instead.
func (*GetTagResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{3}
}

func (x *GetTagResponse) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

func (x *GetTagResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetTagsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GetTagsResponse_Data `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Code int32                 `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *GetTagsResponse) Reset() {
	*x = GetTagsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagsResponse) ProtoMessage() {}

func (x *GetTagsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagsResponse.ProtoReflect.Descriptor instead.
func (*GetTagsResponse) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{4}
}

func (x *GetTagsResponse) GetData() *GetTagsResponse_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetTagsResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{5}
}

func (x *Id) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{6}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TagRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag *Tag `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *TagRequest) Reset() {
	*x = TagRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagRequest) ProtoMessage() {}

func (x *TagRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagRequest.ProtoReflect.Descriptor instead.
func (*TagRequest) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{7}
}

func (x *TagRequest) GetTag() *Tag {
	if x != nil {
		return x.Tag
	}
	return nil
}

type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{8}
}

func (x *Code) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type GetTagsResponse_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags  []*Tag `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	Total int32  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GetTagsResponse_Data) Reset() {
	*x = GetTagsResponse_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_proto_tag_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTagsResponse_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTagsResponse_Data) ProtoMessage() {}

func (x *GetTagsResponse_Data) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_proto_tag_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTagsResponse_Data.ProtoReflect.Descriptor instead.
func (*GetTagsResponse_Data) Descriptor() ([]byte, []int) {
	return file_rpc_proto_tag_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetTagsResponse_Data) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetTagsResponse_Data) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_rpc_proto_tag_proto protoreflect.FileDescriptor

var file_rpc_proto_tag_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6f, 0x75, 0x74, 0x22, 0x7d, 0x0a, 0x03, 0x54, 0x61,
	0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x42, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x42, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x31, 0x0a, 0x04, 0x6d, 0x61,
	0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61,
	0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x61, 0x70, 0x73, 0x1a, 0x37, 0x0a,
	0x09, 0x4d, 0x61, 0x70, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x0d, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x78, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x22, 0x40, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54,
	0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x3a, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x22, 0x1a, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xb6,
	0x02, 0x0a, 0x0a, 0x54, 0x61, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a,
	0x0c, 0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49, 0x64, 0x12, 0x07, 0x2e,
	0x6f, 0x75, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0e,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09,
	0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x12, 0x2e, 0x6f, 0x75, 0x74, 0x2e,
	0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x28, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x12, 0x07, 0x2e, 0x6f, 0x75, 0x74, 0x2e,
	0x49, 0x64, 0x1a, 0x13, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6f, 0x75, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x61, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x1f, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67, 0x12, 0x08, 0x2e, 0x6f, 0x75,
	0x74, 0x2e, 0x54, 0x61, 0x67, 0x1a, 0x09, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x00, 0x12, 0x20, 0x0a, 0x07, 0x45, 0x64, 0x69, 0x74, 0x54, 0x61, 0x67, 0x12, 0x08, 0x2e,
	0x6f, 0x75, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x1a, 0x09, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61,
	0x67, 0x12, 0x07, 0x2e, 0x6f, 0x75, 0x74, 0x2e, 0x49, 0x64, 0x1a, 0x09, 0x2e, 0x6f, 0x75, 0x74,
	0x2e, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_proto_tag_proto_rawDescOnce sync.Once
	file_rpc_proto_tag_proto_rawDescData = file_rpc_proto_tag_proto_rawDesc
)

func file_rpc_proto_tag_proto_rawDescGZIP() []byte {
	file_rpc_proto_tag_proto_rawDescOnce.Do(func() {
		file_rpc_proto_tag_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_proto_tag_proto_rawDescData)
	})
	return file_rpc_proto_tag_proto_rawDescData
}

var file_rpc_proto_tag_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_proto_tag_proto_goTypes = []interface{}{
	(*Tag)(nil),                  // 0: out.Tag
	(*GetTagsRequest)(nil),       // 1: out.GetTagsRequest
	(*ExistResponse)(nil),        // 2: out.ExistResponse
	(*GetTagResponse)(nil),       // 3: out.GetTagResponse
	(*GetTagsResponse)(nil),      // 4: out.GetTagsResponse
	(*Id)(nil),                   // 5: out.Id
	(*Name)(nil),                 // 6: out.Name
	(*TagRequest)(nil),           // 7: out.TagRequest
	(*Code)(nil),                 // 8: out.Code
	nil,                          // 9: out.GetTagsRequest.MapsEntry
	(*GetTagsResponse_Data)(nil), // 10: out.GetTagsResponse.Data
}
var file_rpc_proto_tag_proto_depIdxs = []int32{
	9,  // 0: out.GetTagsRequest.maps:type_name -> out.GetTagsRequest.MapsEntry
	0,  // 1: out.GetTagResponse.tag:type_name -> out.Tag
	10, // 2: out.GetTagsResponse.data:type_name -> out.GetTagsResponse.Data
	0,  // 3: out.TagRequest.tag:type_name -> out.Tag
	0,  // 4: out.GetTagsResponse.Data.tags:type_name -> out.Tag
	5,  // 5: out.TagService.ExistTagById:input_type -> out.Id
	6,  // 6: out.TagService.ExistTagByName:input_type -> out.Name
	5,  // 7: out.TagService.GetTag:input_type -> out.Id
	1,  // 8: out.TagService.GetTags:input_type -> out.GetTagsRequest
	0,  // 9: out.TagService.AddTag:input_type -> out.Tag
	0,  // 10: out.TagService.EditTag:input_type -> out.Tag
	5,  // 11: out.TagService.DeleteTag:input_type -> out.Id
	2,  // 12: out.TagService.ExistTagById:output_type -> out.ExistResponse
	2,  // 13: out.TagService.ExistTagByName:output_type -> out.ExistResponse
	3,  // 14: out.TagService.GetTag:output_type -> out.GetTagResponse
	4,  // 15: out.TagService.GetTags:output_type -> out.GetTagsResponse
	8,  // 16: out.TagService.AddTag:output_type -> out.Code
	8,  // 17: out.TagService.EditTag:output_type -> out.Code
	8,  // 18: out.TagService.DeleteTag:output_type -> out.Code
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_proto_tag_proto_init() }
func file_rpc_proto_tag_proto_init() {
	if File_rpc_proto_tag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_proto_tag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tag); i {
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
		file_rpc_proto_tag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagsRequest); i {
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
		file_rpc_proto_tag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExistResponse); i {
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
		file_rpc_proto_tag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagResponse); i {
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
		file_rpc_proto_tag_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagsResponse); i {
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
		file_rpc_proto_tag_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_rpc_proto_tag_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_rpc_proto_tag_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagRequest); i {
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
		file_rpc_proto_tag_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
		file_rpc_proto_tag_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTagsResponse_Data); i {
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
			RawDescriptor: file_rpc_proto_tag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_proto_tag_proto_goTypes,
		DependencyIndexes: file_rpc_proto_tag_proto_depIdxs,
		MessageInfos:      file_rpc_proto_tag_proto_msgTypes,
	}.Build()
	File_rpc_proto_tag_proto = out.File
	file_rpc_proto_tag_proto_rawDesc = nil
	file_rpc_proto_tag_proto_goTypes = nil
	file_rpc_proto_tag_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	ExistTagById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ExistResponse, error)
	ExistTagByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*ExistResponse, error)
	GetTag(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetTagResponse, error)
	GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsResponse, error)
	AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Code, error)
	EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Code, error)
	DeleteTag(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Code, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) ExistTagById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*ExistResponse, error) {
	out := new(ExistResponse)
	err := c.cc.Invoke(ctx, "/out.TagService/ExistTagById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) ExistTagByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*ExistResponse, error) {
	out := new(ExistResponse)
	err := c.cc.Invoke(ctx, "/out.TagService/ExistTagByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTag(ctx context.Context, in *Id, opts ...grpc.CallOption) (*GetTagResponse, error) {
	out := new(GetTagResponse)
	err := c.cc.Invoke(ctx, "/out.TagService/GetTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) GetTags(ctx context.Context, in *GetTagsRequest, opts ...grpc.CallOption) (*GetTagsResponse, error) {
	out := new(GetTagsResponse)
	err := c.cc.Invoke(ctx, "/out.TagService/GetTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) AddTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/out.TagService/AddTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) EditTag(ctx context.Context, in *Tag, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/out.TagService/EditTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tagServiceClient) DeleteTag(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Code, error) {
	out := new(Code)
	err := c.cc.Invoke(ctx, "/out.TagService/DeleteTag", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	ExistTagById(context.Context, *Id) (*ExistResponse, error)
	ExistTagByName(context.Context, *Name) (*ExistResponse, error)
	GetTag(context.Context, *Id) (*GetTagResponse, error)
	GetTags(context.Context, *GetTagsRequest) (*GetTagsResponse, error)
	AddTag(context.Context, *Tag) (*Code, error)
	EditTag(context.Context, *Tag) (*Code, error)
	DeleteTag(context.Context, *Id) (*Code, error)
}

// UnimplementedTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (*UnimplementedTagServiceServer) ExistTagById(context.Context, *Id) (*ExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTagById not implemented")
}
func (*UnimplementedTagServiceServer) ExistTagByName(context.Context, *Name) (*ExistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExistTagByName not implemented")
}
func (*UnimplementedTagServiceServer) GetTag(context.Context, *Id) (*GetTagResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTag not implemented")
}
func (*UnimplementedTagServiceServer) GetTags(context.Context, *GetTagsRequest) (*GetTagsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTags not implemented")
}
func (*UnimplementedTagServiceServer) AddTag(context.Context, *Tag) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTag not implemented")
}
func (*UnimplementedTagServiceServer) EditTag(context.Context, *Tag) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditTag not implemented")
}
func (*UnimplementedTagServiceServer) DeleteTag(context.Context, *Id) (*Code, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTag not implemented")
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_ExistTagById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ExistTagById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/ExistTagById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ExistTagById(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_ExistTagByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).ExistTagByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/ExistTagByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).ExistTagByName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/GetTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTag(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_GetTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/GetTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTags(ctx, req.(*GetTagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_AddTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).AddTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/AddTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).AddTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_EditTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Tag)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).EditTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/EditTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).EditTag(ctx, req.(*Tag))
	}
	return interceptor(ctx, in, info, handler)
}

func _TagService_DeleteTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).DeleteTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/out.TagService/DeleteTag",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).DeleteTag(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "out.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExistTagById",
			Handler:    _TagService_ExistTagById_Handler,
		},
		{
			MethodName: "ExistTagByName",
			Handler:    _TagService_ExistTagByName_Handler,
		},
		{
			MethodName: "GetTag",
			Handler:    _TagService_GetTag_Handler,
		},
		{
			MethodName: "GetTags",
			Handler:    _TagService_GetTags_Handler,
		},
		{
			MethodName: "AddTag",
			Handler:    _TagService_AddTag_Handler,
		},
		{
			MethodName: "EditTag",
			Handler:    _TagService_EditTag_Handler,
		},
		{
			MethodName: "DeleteTag",
			Handler:    _TagService_DeleteTag_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/proto/tag.proto",
}
