// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: profile.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProfileClient is the client API for Profile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileClient interface {
	GetUserProfile(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*ProfileData, error)
	EditProfile(ctx context.Context, in *EditProfileData, opts ...grpc.CallOption) (*Empty, error)
	EditAvatar(ctx context.Context, in *EditAvatarData, opts ...grpc.CallOption) (*Empty, error)
	UploadAvatar(ctx context.Context, in *UploadInputFile, opts ...grpc.CallOption) (*FileName, error)
	GetAvatar(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FileName, error)
	AddLike(ctx context.Context, in *LikeData, opts ...grpc.CallOption) (*Empty, error)
	RemoveLike(ctx context.Context, in *LikeData, opts ...grpc.CallOption) (*Empty, error)
	GetFavorites(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Favorites, error)
	GetPaymentsToken(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Token, error)
	CheckPaymentsToken(ctx context.Context, in *CheckTokenData, opts ...grpc.CallOption) (*Empty, error)
	CheckToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error)
	CreatePayment(ctx context.Context, in *CheckTokenData, opts ...grpc.CallOption) (*Empty, error)
	CreateSubscribe(ctx context.Context, in *SubscribeData, opts ...grpc.CallOption) (*Empty, error)
	IsSubscription(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Empty, error)
}

type profileClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileClient(cc grpc.ClientConnInterface) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) GetUserProfile(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*ProfileData, error) {
	out := new(ProfileData)
	err := c.cc.Invoke(ctx, "/profile.Profile/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditProfile(ctx context.Context, in *EditProfileData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/EditProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) EditAvatar(ctx context.Context, in *EditAvatarData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/EditAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) UploadAvatar(ctx context.Context, in *UploadInputFile, opts ...grpc.CallOption) (*FileName, error) {
	out := new(FileName)
	err := c.cc.Invoke(ctx, "/profile.Profile/UploadAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetAvatar(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*FileName, error) {
	out := new(FileName)
	err := c.cc.Invoke(ctx, "/profile.Profile/GetAvatar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) AddLike(ctx context.Context, in *LikeData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/AddLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) RemoveLike(ctx context.Context, in *LikeData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/RemoveLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetFavorites(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Favorites, error) {
	out := new(Favorites)
	err := c.cc.Invoke(ctx, "/profile.Profile/GetFavorites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) GetPaymentsToken(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/profile.Profile/GetPaymentsToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CheckPaymentsToken(ctx context.Context, in *CheckTokenData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/CheckPaymentsToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CheckToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/CheckToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CreatePayment(ctx context.Context, in *CheckTokenData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/CreatePayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) CreateSubscribe(ctx context.Context, in *SubscribeData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/CreateSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) IsSubscription(ctx context.Context, in *UserID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/profile.Profile/IsSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServer is the server API for Profile service.
// All implementations should embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	GetUserProfile(context.Context, *UserID) (*ProfileData, error)
	EditProfile(context.Context, *EditProfileData) (*Empty, error)
	EditAvatar(context.Context, *EditAvatarData) (*Empty, error)
	UploadAvatar(context.Context, *UploadInputFile) (*FileName, error)
	GetAvatar(context.Context, *UserID) (*FileName, error)
	AddLike(context.Context, *LikeData) (*Empty, error)
	RemoveLike(context.Context, *LikeData) (*Empty, error)
	GetFavorites(context.Context, *UserID) (*Favorites, error)
	GetPaymentsToken(context.Context, *UserID) (*Token, error)
	CheckPaymentsToken(context.Context, *CheckTokenData) (*Empty, error)
	CheckToken(context.Context, *Token) (*Empty, error)
	CreatePayment(context.Context, *CheckTokenData) (*Empty, error)
	CreateSubscribe(context.Context, *SubscribeData) (*Empty, error)
	IsSubscription(context.Context, *UserID) (*Empty, error)
}

// UnimplementedProfileServer should be embedded to have forward compatible implementations.
type UnimplementedProfileServer struct {
}

func (UnimplementedProfileServer) GetUserProfile(context.Context, *UserID) (*ProfileData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedProfileServer) EditProfile(context.Context, *EditProfileData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProfile not implemented")
}
func (UnimplementedProfileServer) EditAvatar(context.Context, *EditAvatarData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditAvatar not implemented")
}
func (UnimplementedProfileServer) UploadAvatar(context.Context, *UploadInputFile) (*FileName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadAvatar not implemented")
}
func (UnimplementedProfileServer) GetAvatar(context.Context, *UserID) (*FileName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvatar not implemented")
}
func (UnimplementedProfileServer) AddLike(context.Context, *LikeData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLike not implemented")
}
func (UnimplementedProfileServer) RemoveLike(context.Context, *LikeData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLike not implemented")
}
func (UnimplementedProfileServer) GetFavorites(context.Context, *UserID) (*Favorites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavorites not implemented")
}
func (UnimplementedProfileServer) GetPaymentsToken(context.Context, *UserID) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentsToken not implemented")
}
func (UnimplementedProfileServer) CheckPaymentsToken(context.Context, *CheckTokenData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPaymentsToken not implemented")
}
func (UnimplementedProfileServer) CheckToken(context.Context, *Token) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedProfileServer) CreatePayment(context.Context, *CheckTokenData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePayment not implemented")
}
func (UnimplementedProfileServer) CreateSubscribe(context.Context, *SubscribeData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscribe not implemented")
}
func (UnimplementedProfileServer) IsSubscription(context.Context, *UserID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsSubscription not implemented")
}

// UnsafeProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServer will
// result in compilation errors.
type UnsafeProfileServer interface {
	mustEmbedUnimplementedProfileServer()
}

func RegisterProfileServer(s grpc.ServiceRegistrar, srv ProfileServer) {
	s.RegisterService(&Profile_ServiceDesc, srv)
}

func _Profile_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetUserProfile(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProfileData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/EditProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditProfile(ctx, req.(*EditProfileData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_EditAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditAvatarData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).EditAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/EditAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).EditAvatar(ctx, req.(*EditAvatarData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_UploadAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadInputFile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).UploadAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/UploadAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).UploadAvatar(ctx, req.(*UploadInputFile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetAvatar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetAvatar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetAvatar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetAvatar(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_AddLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).AddLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/AddLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).AddLike(ctx, req.(*LikeData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_RemoveLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).RemoveLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/RemoveLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).RemoveLike(ctx, req.(*LikeData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetFavorites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetFavorites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetFavorites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetFavorites(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_GetPaymentsToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).GetPaymentsToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/GetPaymentsToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).GetPaymentsToken(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CheckPaymentsToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CheckPaymentsToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/CheckPaymentsToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CheckPaymentsToken(ctx, req.(*CheckTokenData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CheckToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CreatePayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CreatePayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/CreatePayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CreatePayment(ctx, req.(*CheckTokenData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_CreateSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).CreateSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/CreateSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).CreateSubscribe(ctx, req.(*SubscribeData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_IsSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).IsSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.Profile/IsSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).IsSubscription(ctx, req.(*UserID))
	}
	return interceptor(ctx, in, info, handler)
}

// Profile_ServiceDesc is the grpc.ServiceDesc for Profile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Profile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProfile",
			Handler:    _Profile_GetUserProfile_Handler,
		},
		{
			MethodName: "EditProfile",
			Handler:    _Profile_EditProfile_Handler,
		},
		{
			MethodName: "EditAvatar",
			Handler:    _Profile_EditAvatar_Handler,
		},
		{
			MethodName: "UploadAvatar",
			Handler:    _Profile_UploadAvatar_Handler,
		},
		{
			MethodName: "GetAvatar",
			Handler:    _Profile_GetAvatar_Handler,
		},
		{
			MethodName: "AddLike",
			Handler:    _Profile_AddLike_Handler,
		},
		{
			MethodName: "RemoveLike",
			Handler:    _Profile_RemoveLike_Handler,
		},
		{
			MethodName: "GetFavorites",
			Handler:    _Profile_GetFavorites_Handler,
		},
		{
			MethodName: "GetPaymentsToken",
			Handler:    _Profile_GetPaymentsToken_Handler,
		},
		{
			MethodName: "CheckPaymentsToken",
			Handler:    _Profile_CheckPaymentsToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Profile_CheckToken_Handler,
		},
		{
			MethodName: "CreatePayment",
			Handler:    _Profile_CreatePayment_Handler,
		},
		{
			MethodName: "CreateSubscribe",
			Handler:    _Profile_CreateSubscribe_Handler,
		},
		{
			MethodName: "IsSubscription",
			Handler:    _Profile_IsSubscription_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
