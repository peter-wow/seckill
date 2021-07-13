// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.2
// source: v1/goods.proto

package v1

import (
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

type CreateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGoodsRequest) Reset() {
	*x = CreateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsRequest) ProtoMessage() {}

func (x *CreateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{0}
}

type CreateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateGoodsReply) Reset() {
	*x = CreateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodsReply) ProtoMessage() {}

func (x *CreateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodsReply.ProtoReflect.Descriptor instead.
func (*CreateGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{1}
}

type UpdateGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGoodsRequest) Reset() {
	*x = UpdateGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsRequest) ProtoMessage() {}

func (x *UpdateGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{2}
}

type UpdateGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateGoodsReply) Reset() {
	*x = UpdateGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodsReply) ProtoMessage() {}

func (x *UpdateGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodsReply.ProtoReflect.Descriptor instead.
func (*UpdateGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{3}
}

type DeleteGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodsRequest) Reset() {
	*x = DeleteGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsRequest) ProtoMessage() {}

func (x *DeleteGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{4}
}

type DeleteGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodsReply) Reset() {
	*x = DeleteGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodsReply) ProtoMessage() {}

func (x *DeleteGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodsReply.ProtoReflect.Descriptor instead.
func (*DeleteGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{5}
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{6}
}

func (x *GetGoodsRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Intro string `protobuf:"bytes,3,opt,name=intro,proto3" json:"intro,omitempty"`
}

func (x *GetGoodsReply) Reset() {
	*x = GetGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsReply) ProtoMessage() {}

func (x *GetGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsReply.ProtoReflect.Descriptor instead.
func (*GetGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{7}
}

func (x *GetGoodsReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetGoodsReply) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetGoodsReply) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

type ListGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGoodsRequest) Reset() {
	*x = ListGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsRequest) ProtoMessage() {}

func (x *ListGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsRequest.ProtoReflect.Descriptor instead.
func (*ListGoodsRequest) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{8}
}

type ListGoodsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGoodsReply) Reset() {
	*x = ListGoodsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_goods_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGoodsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGoodsReply) ProtoMessage() {}

func (x *ListGoodsReply) ProtoReflect() protoreflect.Message {
	mi := &file_v1_goods_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGoodsReply.ProtoReflect.Descriptor instead.
func (*ListGoodsReply) Descriptor() ([]byte, []int) {
	return file_v1_goods_proto_rawDescGZIP(), []int{9}
}

var File_v1_goods_proto protoreflect.FileDescriptor

var file_v1_goods_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x12,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x74,
	0x72, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xf2, 0x03, 0x0a, 0x05, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x5f, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x5f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f,
	0x6f, 0x64, 0x73, 0x12, 0x28, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x6b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x59, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x12,
	0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x4d, 0x0a,
	0x14, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x4a, 0x57, 0x2f, 0x73, 0x65, 0x63,
	0x6b, 0x69, 0x6c, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_goods_proto_rawDescOnce sync.Once
	file_v1_goods_proto_rawDescData = file_v1_goods_proto_rawDesc
)

func file_v1_goods_proto_rawDescGZIP() []byte {
	file_v1_goods_proto_rawDescOnce.Do(func() {
		file_v1_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_goods_proto_rawDescData)
	})
	return file_v1_goods_proto_rawDescData
}

var file_v1_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v1_goods_proto_goTypes = []interface{}{
	(*CreateGoodsRequest)(nil), // 0: api.goods.service.v1.CreateGoodsRequest
	(*CreateGoodsReply)(nil),   // 1: api.goods.service.v1.CreateGoodsReply
	(*UpdateGoodsRequest)(nil), // 2: api.goods.service.v1.UpdateGoodsRequest
	(*UpdateGoodsReply)(nil),   // 3: api.goods.service.v1.UpdateGoodsReply
	(*DeleteGoodsRequest)(nil), // 4: api.goods.service.v1.DeleteGoodsRequest
	(*DeleteGoodsReply)(nil),   // 5: api.goods.service.v1.DeleteGoodsReply
	(*GetGoodsRequest)(nil),    // 6: api.goods.service.v1.GetGoodsRequest
	(*GetGoodsReply)(nil),      // 7: api.goods.service.v1.GetGoodsReply
	(*ListGoodsRequest)(nil),   // 8: api.goods.service.v1.ListGoodsRequest
	(*ListGoodsReply)(nil),     // 9: api.goods.service.v1.ListGoodsReply
}
var file_v1_goods_proto_depIdxs = []int32{
	0, // 0: api.goods.service.v1.Goods.CreateGoods:input_type -> api.goods.service.v1.CreateGoodsRequest
	2, // 1: api.goods.service.v1.Goods.UpdateGoods:input_type -> api.goods.service.v1.UpdateGoodsRequest
	4, // 2: api.goods.service.v1.Goods.DeleteGoods:input_type -> api.goods.service.v1.DeleteGoodsRequest
	6, // 3: api.goods.service.v1.Goods.GetGoods:input_type -> api.goods.service.v1.GetGoodsRequest
	8, // 4: api.goods.service.v1.Goods.ListGoods:input_type -> api.goods.service.v1.ListGoodsRequest
	1, // 5: api.goods.service.v1.Goods.CreateGoods:output_type -> api.goods.service.v1.CreateGoodsReply
	3, // 6: api.goods.service.v1.Goods.UpdateGoods:output_type -> api.goods.service.v1.UpdateGoodsReply
	5, // 7: api.goods.service.v1.Goods.DeleteGoods:output_type -> api.goods.service.v1.DeleteGoodsReply
	7, // 8: api.goods.service.v1.Goods.GetGoods:output_type -> api.goods.service.v1.GetGoodsReply
	9, // 9: api.goods.service.v1.Goods.ListGoods:output_type -> api.goods.service.v1.ListGoodsReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_goods_proto_init() }
func file_v1_goods_proto_init() {
	if File_v1_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodsReply); i {
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
		file_v1_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodsReply); i {
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
		file_v1_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodsReply); i {
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
		file_v1_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsReply); i {
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
		file_v1_goods_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsRequest); i {
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
		file_v1_goods_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGoodsReply); i {
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
			RawDescriptor: file_v1_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_goods_proto_goTypes,
		DependencyIndexes: file_v1_goods_proto_depIdxs,
		MessageInfos:      file_v1_goods_proto_msgTypes,
	}.Build()
	File_v1_goods_proto = out.File
	file_v1_goods_proto_rawDesc = nil
	file_v1_goods_proto_goTypes = nil
	file_v1_goods_proto_depIdxs = nil
}
