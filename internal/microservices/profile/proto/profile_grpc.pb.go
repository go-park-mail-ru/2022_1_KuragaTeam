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

// ProfileServer is the server API for Profile service.
// All implementations should embed UnimplementedProfileServer
// for forward compatibility
type ProfileServer interface {
	GetUserProfile(context.Context, *UserID) (*ProfileData, error)
	EditProfile(context.Context, *EditProfileData) (*Empty, error)
	EditAvatar(context.Context, *EditAvatarData) (*Empty, error)
	UploadAvatar(context.Context, *UploadInputFile) (*FileName, error)
	GetAvatar(context.Context, *UserID) (*FileName, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile.proto",
}
