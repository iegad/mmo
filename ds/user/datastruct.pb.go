// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.17.3
// source: user/datastruct.proto

package user

import (
	ds "github.com/iegad/mmo/ds"
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

type ArchiveLogType int32

const (
	ArchiveLogType_ArchiveLogType_Invalid  ArchiveLogType = 0
	ArchiveLogType_ArchiveLogType_Basic    ArchiveLogType = 1 // 条目信息修改
	ArchiveLogType_ArchiveLogType_Personal ArchiveLogType = 2 // 个人信息修改
)

// Enum value maps for ArchiveLogType.
var (
	ArchiveLogType_name = map[int32]string{
		0: "ArchiveLogType_Invalid",
		1: "ArchiveLogType_Basic",
		2: "ArchiveLogType_Personal",
	}
	ArchiveLogType_value = map[string]int32{
		"ArchiveLogType_Invalid":  0,
		"ArchiveLogType_Basic":    1,
		"ArchiveLogType_Personal": 2,
	}
)

func (x ArchiveLogType) Enum() *ArchiveLogType {
	p := new(ArchiveLogType)
	*p = x
	return p
}

func (x ArchiveLogType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArchiveLogType) Descriptor() protoreflect.EnumDescriptor {
	return file_user_datastruct_proto_enumTypes[0].Descriptor()
}

func (ArchiveLogType) Type() protoreflect.EnumType {
	return &file_user_datastruct_proto_enumTypes[0]
}

func (x ArchiveLogType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchiveLogType.Descriptor instead.
func (ArchiveLogType) EnumDescriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{0}
}

// 用户条目
//  存放于ES中, 用于快速检索用户
//  @相关操作: 增,删,改,查,计数
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gender   int32  `protobuf:"varint,1,opt,name=Gender,proto3" json:"Gender,omitempty"`    // 性别
	UserID   int64  `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`    // 用户ID
	Email    string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`       // 邮箱
	Nickname string `protobuf:"bytes,4,opt,name=Nickname,proto3" json:"Nickname,omitempty"` // 昵称
	PhoneNum string `protobuf:"bytes,5,opt,name=PhoneNum,proto3" json:"PhoneNum,omitempty"` // 手机号
	Avator   string `protobuf:"bytes,6,opt,name=Avator,proto3" json:"Avator,omitempty"`     // 头像
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Entry) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Entry) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Entry) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Entry) GetPhoneNum() string {
	if x != nil {
		return x.PhoneNum
	}
	return ""
}

func (x *Entry) GetAvator() string {
	if x != nil {
		return x.Avator
	}
	return ""
}

// 用户基础信息
//  存放于MYSQL
//  @相关操作: 增,删,改,查,计数
type Basic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VerCode    uint32 `protobuf:"varint,1,opt,name=VerCode,proto3" json:"VerCode,omitempty"`       // 版本号
	CreateTime int64  `protobuf:"varint,2,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"` // 创建时间
	LastUpdate int64  `protobuf:"varint,3,opt,name=LastUpdate,proto3" json:"LastUpdate,omitempty"` // 最后更新时间
	Entry      *Entry `protobuf:"bytes,4,opt,name=Entry,proto3" json:"Entry,omitempty"`            // 条目信息
}

func (x *Basic) Reset() {
	*x = Basic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Basic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Basic) ProtoMessage() {}

func (x *Basic) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Basic.ProtoReflect.Descriptor instead.
func (*Basic) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{1}
}

func (x *Basic) GetVerCode() uint32 {
	if x != nil {
		return x.VerCode
	}
	return 0
}

func (x *Basic) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Basic) GetLastUpdate() int64 {
	if x != nil {
		return x.LastUpdate
	}
	return 0
}

func (x *Basic) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// 用户个人信息
//  存放于MySQL
type Personal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nationality int32  `protobuf:"varint,1,opt,name=Nationality,proto3" json:"Nationality,omitempty"` // 国籍
	Gender      int32  `protobuf:"varint,2,opt,name=Gender,proto3" json:"Gender,omitempty"`           // 性别
	VerCode     uint32 `protobuf:"varint,3,opt,name=VerCode,proto3" json:"VerCode,omitempty"`         // 版本号
	UserID      int64  `protobuf:"varint,4,opt,name=UserID,proto3" json:"UserID,omitempty"`           // 用户ID
	Birth       int64  `protobuf:"varint,5,opt,name=Birth,proto3" json:"Birth,omitempty"`             // 出生
	Name        string `protobuf:"bytes,6,opt,name=Name,proto3" json:"Name,omitempty"`                // 姓名
	ID          string `protobuf:"bytes,7,opt,name=ID,proto3" json:"ID,omitempty"`                    // 身份证号
}

func (x *Personal) Reset() {
	*x = Personal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Personal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Personal) ProtoMessage() {}

func (x *Personal) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Personal.ProtoReflect.Descriptor instead.
func (*Personal) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{2}
}

func (x *Personal) GetNationality() int32 {
	if x != nil {
		return x.Nationality
	}
	return 0
}

func (x *Personal) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Personal) GetVerCode() uint32 {
	if x != nil {
		return x.VerCode
	}
	return 0
}

func (x *Personal) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Personal) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

func (x *Personal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Personal) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// 用户会话
//  存放于Redis中
//  @相关操作: SET, DEL, GET
type Session struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OS         ds.OSType `protobuf:"varint,1,opt,name=OS,proto3,enum=common.OSType" json:"OS,omitempty"` // 操作系统
	ActiveTime int64     `protobuf:"varint,2,opt,name=ActiveTime,proto3" json:"ActiveTime,omitempty"`    // 会话激活时间
	Token      []byte    `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`               // 令牌
	DeviceCode []byte    `protobuf:"bytes,4,opt,name=DeviceCode,proto3" json:"DeviceCode,omitempty"`     // 设备编码: md5(deviceType:operatingSystem:deviceUniqueIdentifier:macAddr)
	MountAddr  string    `protobuf:"bytes,5,opt,name=MountAddr,proto3" json:"MountAddr,omitempty"`       // 登录网关(挂载点)地址
	Entry      *Entry    `protobuf:"bytes,6,opt,name=Entry,proto3" json:"Entry,omitempty"`               // 用户信息
}

func (x *Session) Reset() {
	*x = Session{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Session) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Session) ProtoMessage() {}

func (x *Session) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Session.ProtoReflect.Descriptor instead.
func (*Session) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{3}
}

func (x *Session) GetOS() ds.OSType {
	if x != nil {
		return x.OS
	}
	return ds.OSType(0)
}

func (x *Session) GetActiveTime() int64 {
	if x != nil {
		return x.ActiveTime
	}
	return 0
}

func (x *Session) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *Session) GetDeviceCode() []byte {
	if x != nil {
		return x.DeviceCode
	}
	return nil
}

func (x *Session) GetMountAddr() string {
	if x != nil {
		return x.MountAddr
	}
	return ""
}

func (x *Session) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// 用户关系
type Relation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateTime int64  `protobuf:"varint,1,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"` // 关系确定时间
	Entry      *Entry `protobuf:"bytes,2,opt,name=Entry,proto3" json:"Entry,omitempty"`            // 用户条目
}

func (x *Relation) Reset() {
	*x = Relation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Relation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Relation) ProtoMessage() {}

func (x *Relation) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Relation.ProtoReflect.Descriptor instead.
func (*Relation) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{4}
}

func (x *Relation) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Relation) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// 用户关系信息
//  存放于Mongo
type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         int64       `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`                // 用户ID
	ContactList    []*Relation `protobuf:"bytes,2,rep,name=ContactList,proto3" json:"ContactList,omitempty"`       // 联系人
	BlackList      []*Relation `protobuf:"bytes,3,rep,name=BlackList,proto3" json:"BlackList,omitempty"`           // 黑名单
	FansList       []*Relation `protobuf:"bytes,4,rep,name=FansList,proto3" json:"FansList,omitempty"`             // 粉丝
	SubscriberList []*Relation `protobuf:"bytes,5,rep,name=SubscriberList,proto3" json:"SubscriberList,omitempty"` // 订阅
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{5}
}

func (x *Contact) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Contact) GetContactList() []*Relation {
	if x != nil {
		return x.ContactList
	}
	return nil
}

func (x *Contact) GetBlackList() []*Relation {
	if x != nil {
		return x.BlackList
	}
	return nil
}

func (x *Contact) GetFansList() []*Relation {
	if x != nil {
		return x.FansList
	}
	return nil
}

func (x *Contact) GetSubscriberList() []*Relation {
	if x != nil {
		return x.SubscriberList
	}
	return nil
}

// 登录日志
//  存放于ES
type LoginLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OSType                 ds.OSType `protobuf:"varint,1,opt,name=OSType,proto3,enum=common.OSType" json:"OSType,omitempty"`             // 系统类型
	DeviceType             int32     `protobuf:"varint,2,opt,name=DeviceType,proto3" json:"DeviceType,omitempty"`                        // 设备类型
	UserID                 int64     `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`                                // 用户ID
	LoginTime              int64     `protobuf:"varint,4,opt,name=LoginTime,proto3" json:"LoginTime,omitempty"`                          // 登录时间
	Token                  []byte    `protobuf:"bytes,5,opt,name=Token,proto3" json:"Token,omitempty"`                                   // 令牌
	MacAddr                []byte    `protobuf:"bytes,6,opt,name=MacAddr,proto3" json:"MacAddr,omitempty"`                               // MAC地址
	OperatingSystem        string    `protobuf:"bytes,7,opt,name=OperatingSystem,proto3" json:"OperatingSystem,omitempty"`               // 操作系统
	DeviceUniqueIdentifier string    `protobuf:"bytes,8,opt,name=DeviceUniqueIdentifier,proto3" json:"DeviceUniqueIdentifier,omitempty"` // 唯一码
	MountAddr              string    `protobuf:"bytes,9,opt,name=MountAddr,proto3" json:"MountAddr,omitempty"`                           // 挂载点
	Entry                  *Entry    `protobuf:"bytes,10,opt,name=Entry,proto3" json:"Entry,omitempty"`                                  // 用户条目信息
}

func (x *LoginLog) Reset() {
	*x = LoginLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginLog) ProtoMessage() {}

func (x *LoginLog) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginLog.ProtoReflect.Descriptor instead.
func (*LoginLog) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{6}
}

func (x *LoginLog) GetOSType() ds.OSType {
	if x != nil {
		return x.OSType
	}
	return ds.OSType(0)
}

func (x *LoginLog) GetDeviceType() int32 {
	if x != nil {
		return x.DeviceType
	}
	return 0
}

func (x *LoginLog) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LoginLog) GetLoginTime() int64 {
	if x != nil {
		return x.LoginTime
	}
	return 0
}

func (x *LoginLog) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *LoginLog) GetMacAddr() []byte {
	if x != nil {
		return x.MacAddr
	}
	return nil
}

func (x *LoginLog) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

func (x *LoginLog) GetDeviceUniqueIdentifier() string {
	if x != nil {
		return x.DeviceUniqueIdentifier
	}
	return ""
}

func (x *LoginLog) GetMountAddr() string {
	if x != nil {
		return x.MountAddr
	}
	return ""
}

func (x *LoginLog) GetEntry() *Entry {
	if x != nil {
		return x.Entry
	}
	return nil
}

// 存档日志
//  存放于ES, 当数据修改后需要存档原始部分
type ArchiveLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ArchiveLogType `protobuf:"varint,1,opt,name=Type,proto3,enum=user.ArchiveLogType" json:"Type,omitempty"` // 存档类型
	VerCode     uint32         `protobuf:"varint,2,opt,name=VerCode,proto3" json:"VerCode,omitempty"`                    // 版本号
	UserID      int64          `protobuf:"varint,3,opt,name=UserID,proto3" json:"UserID,omitempty"`                      // 用户ID
	ArchiveTime int64          `protobuf:"varint,4,opt,name=ArchiveTime,proto3" json:"ArchiveTime,omitempty"`            // 存档时间
	Raw         []byte         `protobuf:"bytes,5,opt,name=Raw,proto3" json:"Raw,omitempty"`                             // 原始数据
	Changed     []byte         `protobuf:"bytes,6,opt,name=Changed,proto3" json:"Changed,omitempty"`                     // 修改后的数据
}

func (x *ArchiveLog) Reset() {
	*x = ArchiveLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_datastruct_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArchiveLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArchiveLog) ProtoMessage() {}

func (x *ArchiveLog) ProtoReflect() protoreflect.Message {
	mi := &file_user_datastruct_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArchiveLog.ProtoReflect.Descriptor instead.
func (*ArchiveLog) Descriptor() ([]byte, []int) {
	return file_user_datastruct_proto_rawDescGZIP(), []int{7}
}

func (x *ArchiveLog) GetType() ArchiveLogType {
	if x != nil {
		return x.Type
	}
	return ArchiveLogType_ArchiveLogType_Invalid
}

func (x *ArchiveLog) GetVerCode() uint32 {
	if x != nil {
		return x.VerCode
	}
	return 0
}

func (x *ArchiveLog) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *ArchiveLog) GetArchiveTime() int64 {
	if x != nil {
		return x.ArchiveTime
	}
	return 0
}

func (x *ArchiveLog) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *ArchiveLog) GetChanged() []byte {
	if x != nil {
		return x.Changed
	}
	return nil
}

var File_user_datastruct_proto protoreflect.FileDescriptor

var file_user_datastruct_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x01, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x05,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x08, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x12,
	0x20, 0x0a, 0x0b, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0xc0, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1e, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x52, 0x02, 0x4f,
	0x53, 0x12, 0x1e, 0x0a, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x75, 0x6e, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xe5, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a,
	0x09, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x09, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x46,
	0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x46,
	0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22,
	0xdb, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x26, 0x0a, 0x06,
	0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x53, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x4f, 0x53,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x4d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x36, 0x0a, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09,
	0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x21, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xb6, 0x01,
	0x0a, 0x0a, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x28, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x41,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x52, 0x61,
	0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x52, 0x61, 0x77, 0x12, 0x18, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2a, 0x63, 0x0a, 0x0e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x65, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x72, 0x63, 0x68,
	0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4c,
	0x6f, 0x67, 0x54, 0x79, 0x70, 0x65, 0x5f, 0x42, 0x61, 0x73, 0x69, 0x63, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x5f, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x10, 0x02, 0x42, 0x25, 0x5a, 0x1c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x65, 0x67, 0x61, 0x64, 0x2f,
	0x6d, 0x6d, 0x6f, 0x2f, 0x64, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0xaa, 0x02, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_datastruct_proto_rawDescOnce sync.Once
	file_user_datastruct_proto_rawDescData = file_user_datastruct_proto_rawDesc
)

func file_user_datastruct_proto_rawDescGZIP() []byte {
	file_user_datastruct_proto_rawDescOnce.Do(func() {
		file_user_datastruct_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_datastruct_proto_rawDescData)
	})
	return file_user_datastruct_proto_rawDescData
}

var file_user_datastruct_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_datastruct_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_user_datastruct_proto_goTypes = []interface{}{
	(ArchiveLogType)(0), // 0: user.ArchiveLogType
	(*Entry)(nil),       // 1: user.Entry
	(*Basic)(nil),       // 2: user.Basic
	(*Personal)(nil),    // 3: user.Personal
	(*Session)(nil),     // 4: user.Session
	(*Relation)(nil),    // 5: user.Relation
	(*Contact)(nil),     // 6: user.Contact
	(*LoginLog)(nil),    // 7: user.LoginLog
	(*ArchiveLog)(nil),  // 8: user.ArchiveLog
	(ds.OSType)(0),      // 9: common.OSType
}
var file_user_datastruct_proto_depIdxs = []int32{
	1,  // 0: user.Basic.Entry:type_name -> user.Entry
	9,  // 1: user.Session.OS:type_name -> common.OSType
	1,  // 2: user.Session.Entry:type_name -> user.Entry
	1,  // 3: user.Relation.Entry:type_name -> user.Entry
	5,  // 4: user.Contact.ContactList:type_name -> user.Relation
	5,  // 5: user.Contact.BlackList:type_name -> user.Relation
	5,  // 6: user.Contact.FansList:type_name -> user.Relation
	5,  // 7: user.Contact.SubscriberList:type_name -> user.Relation
	9,  // 8: user.LoginLog.OSType:type_name -> common.OSType
	1,  // 9: user.LoginLog.Entry:type_name -> user.Entry
	0,  // 10: user.ArchiveLog.Type:type_name -> user.ArchiveLogType
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_user_datastruct_proto_init() }
func file_user_datastruct_proto_init() {
	if File_user_datastruct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_datastruct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_user_datastruct_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Basic); i {
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
		file_user_datastruct_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Personal); i {
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
		file_user_datastruct_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Session); i {
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
		file_user_datastruct_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Relation); i {
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
		file_user_datastruct_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_user_datastruct_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginLog); i {
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
		file_user_datastruct_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArchiveLog); i {
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
			RawDescriptor: file_user_datastruct_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_datastruct_proto_goTypes,
		DependencyIndexes: file_user_datastruct_proto_depIdxs,
		EnumInfos:         file_user_datastruct_proto_enumTypes,
		MessageInfos:      file_user_datastruct_proto_msgTypes,
	}.Build()
	File_user_datastruct_proto = out.File
	file_user_datastruct_proto_rawDesc = nil
	file_user_datastruct_proto_goTypes = nil
	file_user_datastruct_proto_depIdxs = nil
}
