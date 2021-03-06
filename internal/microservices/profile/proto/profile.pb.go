// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: profile.proto

//protoc --go_out=. --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. *.proto

package proto

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

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Avatar string `protobuf:"bytes,3,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Date   string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ProfileData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ProfileData) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type EditProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *EditProfileData) Reset() {
	*x = EditProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProfileData) ProtoMessage() {}

func (x *EditProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProfileData.ProtoReflect.Descriptor instead.
func (*EditProfileData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *EditProfileData) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EditProfileData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EditProfileData) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type EditAvatarData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Avatar string `protobuf:"bytes,2,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
}

func (x *EditAvatarData) Reset() {
	*x = EditAvatarData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditAvatarData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditAvatarData) ProtoMessage() {}

func (x *EditAvatarData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditAvatarData.ProtoReflect.Descriptor instead.
func (*EditAvatarData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

func (x *EditAvatarData) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *EditAvatarData) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type UploadInputFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	File        []byte `protobuf:"bytes,2,opt,name=File,proto3" json:"File,omitempty"`
	Size        int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	ContentType string `protobuf:"bytes,4,opt,name=ContentType,proto3" json:"ContentType,omitempty"`
}

func (x *UploadInputFile) Reset() {
	*x = UploadInputFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInputFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInputFile) ProtoMessage() {}

func (x *UploadInputFile) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInputFile.ProtoReflect.Descriptor instead.
func (*UploadInputFile) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *UploadInputFile) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UploadInputFile) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *UploadInputFile) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadInputFile) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

type FileName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FileName) Reset() {
	*x = FileName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileName) ProtoMessage() {}

func (x *FileName) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileName.ProtoReflect.Descriptor instead.
func (*FileName) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *FileName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{5}
}

func (x *UserID) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type LikeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MovieID int64 `protobuf:"varint,2,opt,name=movieID,proto3" json:"movieID,omitempty"`
}

func (x *LikeData) Reset() {
	*x = LikeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeData) ProtoMessage() {}

func (x *LikeData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeData.ProtoReflect.Descriptor instead.
func (*LikeData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{6}
}

func (x *LikeData) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LikeData) GetMovieID() int64 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

type MovieRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	MovieID int64 `protobuf:"varint,2,opt,name=movieID,proto3" json:"movieID,omitempty"`
}

func (x *MovieRating) Reset() {
	*x = MovieRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieRating) ProtoMessage() {}

func (x *MovieRating) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieRating.ProtoReflect.Descriptor instead.
func (*MovieRating) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{7}
}

func (x *MovieRating) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *MovieRating) GetMovieID() int64 {
	if x != nil {
		return x.MovieID
	}
	return 0
}

type Rating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rating int32 `protobuf:"varint,1,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *Rating) Reset() {
	*x = Rating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rating) ProtoMessage() {}

func (x *Rating) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rating.ProtoReflect.Descriptor instead.
func (*Rating) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{8}
}

func (x *Rating) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type Favorites struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *Favorites) Reset() {
	*x = Favorites{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorites) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorites) ProtoMessage() {}

func (x *Favorites) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorites.ProtoReflect.Descriptor instead.
func (*Favorites) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{9}
}

func (x *Favorites) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{10}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type CheckTokenData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Id    int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CheckTokenData) Reset() {
	*x = CheckTokenData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTokenData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTokenData) ProtoMessage() {}

func (x *CheckTokenData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTokenData.ProtoReflect.Descriptor instead.
func (*CheckTokenData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{11}
}

func (x *CheckTokenData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CheckTokenData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SubscribeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Amount float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SubscribeData) Reset() {
	*x = SubscribeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeData) ProtoMessage() {}

func (x *SubscribeData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeData.ProtoReflect.Descriptor instead.
func (*SubscribeData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{12}
}

func (x *SubscribeData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SubscribeData) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subscription bool `protobuf:"varint,1,opt,name=subscription,proto3" json:"subscription,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{13}
}

func (x *Subscription) GetSubscription() bool {
	if x != nil {
		return x.Subscription
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_profile_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{14}
}

var File_profile_proto protoreflect.FileDescriptor

var file_profile_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x63, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0x51, 0x0a,
	0x0f, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x38, 0x0a, 0x0e, 0x45, 0x64, 0x69, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x6b, 0x0a, 0x0f, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x3c, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44, 0x22,
	0x3f, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x44,
	0x22, 0x20, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x1b, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x36,
	0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x32, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0xd5, 0x06, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x39,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0b, 0x45, 0x64, 0x69,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x45, 0x64, 0x69, 0x74, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x64, 0x69,
	0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x12,
	0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x0a, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x1a, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x12, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x0a,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0e, 0x49, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData = file_profile_proto_rawDesc
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_profile_proto_rawDescData)
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_profile_proto_goTypes = []interface{}{
	(*ProfileData)(nil),     // 0: profile.ProfileData
	(*EditProfileData)(nil), // 1: profile.EditProfileData
	(*EditAvatarData)(nil),  // 2: profile.EditAvatarData
	(*UploadInputFile)(nil), // 3: profile.UploadInputFile
	(*FileName)(nil),        // 4: profile.FileName
	(*UserID)(nil),          // 5: profile.UserID
	(*LikeData)(nil),        // 6: profile.LikeData
	(*MovieRating)(nil),     // 7: profile.MovieRating
	(*Rating)(nil),          // 8: profile.Rating
	(*Favorites)(nil),       // 9: profile.Favorites
	(*Token)(nil),           // 10: profile.Token
	(*CheckTokenData)(nil),  // 11: profile.CheckTokenData
	(*SubscribeData)(nil),   // 12: profile.SubscribeData
	(*Subscription)(nil),    // 13: profile.Subscription
	(*Empty)(nil),           // 14: profile.Empty
}
var file_profile_proto_depIdxs = []int32{
	5,  // 0: profile.Profile.GetUserProfile:input_type -> profile.UserID
	1,  // 1: profile.Profile.EditProfile:input_type -> profile.EditProfileData
	2,  // 2: profile.Profile.EditAvatar:input_type -> profile.EditAvatarData
	3,  // 3: profile.Profile.UploadAvatar:input_type -> profile.UploadInputFile
	5,  // 4: profile.Profile.GetAvatar:input_type -> profile.UserID
	6,  // 5: profile.Profile.AddLike:input_type -> profile.LikeData
	6,  // 6: profile.Profile.RemoveLike:input_type -> profile.LikeData
	5,  // 7: profile.Profile.GetFavorites:input_type -> profile.UserID
	7,  // 8: profile.Profile.GetMovieRating:input_type -> profile.MovieRating
	5,  // 9: profile.Profile.GetPaymentsToken:input_type -> profile.UserID
	11, // 10: profile.Profile.CheckPaymentsToken:input_type -> profile.CheckTokenData
	10, // 11: profile.Profile.CheckToken:input_type -> profile.Token
	11, // 12: profile.Profile.CreatePayment:input_type -> profile.CheckTokenData
	12, // 13: profile.Profile.CreateSubscribe:input_type -> profile.SubscribeData
	5,  // 14: profile.Profile.IsSubscription:input_type -> profile.UserID
	0,  // 15: profile.Profile.GetUserProfile:output_type -> profile.ProfileData
	14, // 16: profile.Profile.EditProfile:output_type -> profile.Empty
	14, // 17: profile.Profile.EditAvatar:output_type -> profile.Empty
	4,  // 18: profile.Profile.UploadAvatar:output_type -> profile.FileName
	4,  // 19: profile.Profile.GetAvatar:output_type -> profile.FileName
	14, // 20: profile.Profile.AddLike:output_type -> profile.Empty
	14, // 21: profile.Profile.RemoveLike:output_type -> profile.Empty
	9,  // 22: profile.Profile.GetFavorites:output_type -> profile.Favorites
	8,  // 23: profile.Profile.GetMovieRating:output_type -> profile.Rating
	10, // 24: profile.Profile.GetPaymentsToken:output_type -> profile.Token
	14, // 25: profile.Profile.CheckPaymentsToken:output_type -> profile.Empty
	14, // 26: profile.Profile.CheckToken:output_type -> profile.Empty
	14, // 27: profile.Profile.CreatePayment:output_type -> profile.Empty
	14, // 28: profile.Profile.CreateSubscribe:output_type -> profile.Empty
	14, // 29: profile.Profile.IsSubscription:output_type -> profile.Empty
	15, // [15:30] is the sub-list for method output_type
	0,  // [0:15] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditProfileData); i {
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
		file_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EditAvatarData); i {
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
		file_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInputFile); i {
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
		file_profile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileName); i {
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
		file_profile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_profile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeData); i {
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
		file_profile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieRating); i {
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
		file_profile_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rating); i {
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
		file_profile_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Favorites); i {
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
		file_profile_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_profile_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTokenData); i {
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
		file_profile_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeData); i {
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
		file_profile_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subscription); i {
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
		file_profile_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_profile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_rawDesc = nil
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
