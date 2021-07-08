// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: user/cgi.proto

package cgi

import (
	user "github.com/iegad/mmo/ds/user"
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

type AddEntryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntryList []*user.Entry `protobuf:"bytes,1,rep,name=EntryList,proto3" json:"EntryList,omitempty"`
}

func (x *AddEntryReq) Reset() {
	*x = AddEntryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryReq) ProtoMessage() {}

func (x *AddEntryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryReq.ProtoReflect.Descriptor instead.
func (*AddEntryReq) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{0}
}

func (x *AddEntryReq) GetEntryList() []*user.Entry {
	if x != nil {
		return x.EntryList
	}
	return nil
}

type AddEntryRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32         `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	EntryList []*user.Entry `protobuf:"bytes,2,rep,name=EntryList,proto3" json:"EntryList,omitempty"`
}

func (x *AddEntryRsp) Reset() {
	*x = AddEntryRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEntryRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEntryRsp) ProtoMessage() {}

func (x *AddEntryRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEntryRsp.ProtoReflect.Descriptor instead.
func (*AddEntryRsp) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{1}
}

func (x *AddEntryRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *AddEntryRsp) GetEntryList() []*user.Entry {
	if x != nil {
		return x.EntryList
	}
	return nil
}

type ModEntryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entry *user.Entry `protobuf:"bytes,1,opt,name=Entry,proto3" json:"Entry,omitempty"`
}

func (x *ModEntryReq) Reset() {
	*x = ModEntryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModEntryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModEntryReq) ProtoMessage() {}

func (x *ModEntryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModEntryReq.ProtoReflect.Descriptor instead.
func (*ModEntryReq) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{2}
}

func (x *ModEntryReq) GetEntry() *user.Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type ModEntryRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32       `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Entry *user.Entry `protobuf:"bytes,2,opt,name=Entry,proto3" json:"Entry,omitempty"`
}

func (x *ModEntryRsp) Reset() {
	*x = ModEntryRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModEntryRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModEntryRsp) ProtoMessage() {}

func (x *ModEntryRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModEntryRsp.ProtoReflect.Descriptor instead.
func (*ModEntryRsp) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{3}
}

func (x *ModEntryRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ModEntryRsp) GetEntry() *user.Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type RmvEntryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *RmvEntryReq) Reset() {
	*x = RmvEntryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RmvEntryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RmvEntryReq) ProtoMessage() {}

func (x *RmvEntryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RmvEntryReq.ProtoReflect.Descriptor instead.
func (*RmvEntryReq) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{4}
}

func (x *RmvEntryReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type RmvEntryRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int32       `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Entry *user.Entry `protobuf:"bytes,2,opt,name=Entry,proto3" json:"Entry,omitempty"`
}

func (x *RmvEntryRsp) Reset() {
	*x = RmvEntryRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RmvEntryRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RmvEntryRsp) ProtoMessage() {}

func (x *RmvEntryRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RmvEntryRsp.ProtoReflect.Descriptor instead.
func (*RmvEntryRsp) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{5}
}

func (x *RmvEntryRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RmvEntryRsp) GetEntry() *user.Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

type GetEntryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64  `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Gender int32  `protobuf:"varint,2,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Key    string `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Offset int64  `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int64  `protobuf:"varint,5,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetEntryReq) Reset() {
	*x = GetEntryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryReq) ProtoMessage() {}

func (x *GetEntryReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryReq.ProtoReflect.Descriptor instead.
func (*GetEntryReq) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{6}
}

func (x *GetEntryReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *GetEntryReq) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *GetEntryReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetEntryReq) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetEntryReq) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetEntryRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      int32         `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Total     int64         `protobuf:"varint,2,opt,name=Total,proto3" json:"Total,omitempty"`
	EntryList []*user.Entry `protobuf:"bytes,3,rep,name=EntryList,proto3" json:"EntryList,omitempty"`
}

func (x *GetEntryRsp) Reset() {
	*x = GetEntryRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_cgi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntryRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntryRsp) ProtoMessage() {}

func (x *GetEntryRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_cgi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntryRsp.ProtoReflect.Descriptor instead.
func (*GetEntryRsp) Descriptor() ([]byte, []int) {
	return file_user_cgi_proto_rawDescGZIP(), []int{7}
}

func (x *GetEntryRsp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetEntryRsp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetEntryRsp) GetEntryList() []*user.Entry {
	if x != nil {
		return x.EntryList
	}
	return nil
}

var File_user_cgi_proto protoreflect.FileDescriptor

var file_user_cgi_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x63, 0x67, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a,
	0x0b, 0x41, 0x64, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x09,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4c, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x09, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x44, 0x0a, 0x0b, 0x4d, 0x6f, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x25, 0x0a,
	0x0b, 0x52, 0x6d, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x22, 0x44, 0x0a, 0x0b, 0x52, 0x6d, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x7d, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x4f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x29, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x1a, 0x5a,
	0x18, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x65, 0x67, 0x61,
	0x64, 0x2f, 0x6d, 0x6d, 0x6f, 0x2f, 0x63, 0x67, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_cgi_proto_rawDescOnce sync.Once
	file_user_cgi_proto_rawDescData = file_user_cgi_proto_rawDesc
)

func file_user_cgi_proto_rawDescGZIP() []byte {
	file_user_cgi_proto_rawDescOnce.Do(func() {
		file_user_cgi_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_cgi_proto_rawDescData)
	})
	return file_user_cgi_proto_rawDescData
}

var file_user_cgi_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_cgi_proto_goTypes = []interface{}{
	(*AddEntryReq)(nil), // 0: user.AddEntryReq
	(*AddEntryRsp)(nil), // 1: user.AddEntryRsp
	(*ModEntryReq)(nil), // 2: user.ModEntryReq
	(*ModEntryRsp)(nil), // 3: user.ModEntryRsp
	(*RmvEntryReq)(nil), // 4: user.RmvEntryReq
	(*RmvEntryRsp)(nil), // 5: user.RmvEntryRsp
	(*GetEntryReq)(nil), // 6: user.GetEntryReq
	(*GetEntryRsp)(nil), // 7: user.GetEntryRsp
	(*user.Entry)(nil),  // 8: user.Entry
}
var file_user_cgi_proto_depIdxs = []int32{
	8, // 0: user.AddEntryReq.EntryList:type_name -> user.Entry
	8, // 1: user.AddEntryRsp.EntryList:type_name -> user.Entry
	8, // 2: user.ModEntryReq.Entry:type_name -> user.Entry
	8, // 3: user.ModEntryRsp.Entry:type_name -> user.Entry
	8, // 4: user.RmvEntryRsp.Entry:type_name -> user.Entry
	8, // 5: user.GetEntryRsp.EntryList:type_name -> user.Entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_user_cgi_proto_init() }
func file_user_cgi_proto_init() {
	if File_user_cgi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_cgi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntryReq); i {
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
		file_user_cgi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEntryRsp); i {
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
		file_user_cgi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModEntryReq); i {
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
		file_user_cgi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModEntryRsp); i {
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
		file_user_cgi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RmvEntryReq); i {
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
		file_user_cgi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RmvEntryRsp); i {
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
		file_user_cgi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryReq); i {
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
		file_user_cgi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntryRsp); i {
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
			RawDescriptor: file_user_cgi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_cgi_proto_goTypes,
		DependencyIndexes: file_user_cgi_proto_depIdxs,
		MessageInfos:      file_user_cgi_proto_msgTypes,
	}.Build()
	File_user_cgi_proto = out.File
	file_user_cgi_proto_rawDesc = nil
	file_user_cgi_proto_goTypes = nil
	file_user_cgi_proto_depIdxs = nil
}
