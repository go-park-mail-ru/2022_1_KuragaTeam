// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: microservices/movie/proto/movie.proto

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

// MoviesClient is the client API for Movies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoviesClient interface {
	GetByID(ctx context.Context, in *GetMovieOptions, opts ...grpc.CallOption) (*Movie, error)
	GetRandom(ctx context.Context, in *GetRandomOptions, opts ...grpc.CallOption) (*MoviesArr, error)
	GetMainMovie(ctx context.Context, in *GetMainMovieOptions, opts ...grpc.CallOption) (*MainMovie, error)
	AddMovieRating(ctx context.Context, in *AddRatingOptions, opts ...grpc.CallOption) (*NewMovieRating, error)
}

type moviesClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesClient(cc grpc.ClientConnInterface) MoviesClient {
	return &moviesClient{cc}
}

func (c *moviesClient) GetByID(ctx context.Context, in *GetMovieOptions, opts ...grpc.CallOption) (*Movie, error) {
	out := new(Movie)
	err := c.cc.Invoke(ctx, "/proto.Movies/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) GetRandom(ctx context.Context, in *GetRandomOptions, opts ...grpc.CallOption) (*MoviesArr, error) {
	out := new(MoviesArr)
	err := c.cc.Invoke(ctx, "/proto.Movies/GetRandom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) GetMainMovie(ctx context.Context, in *GetMainMovieOptions, opts ...grpc.CallOption) (*MainMovie, error) {
	out := new(MainMovie)
	err := c.cc.Invoke(ctx, "/proto.Movies/GetMainMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moviesClient) AddMovieRating(ctx context.Context, in *AddRatingOptions, opts ...grpc.CallOption) (*NewMovieRating, error) {
	out := new(NewMovieRating)
	err := c.cc.Invoke(ctx, "/proto.Movies/AddMovieRating", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServer is the server API for Movies service.
// All implementations should embed UnimplementedMoviesServer
// for forward compatibility
type MoviesServer interface {
	GetByID(context.Context, *GetMovieOptions) (*Movie, error)
	GetRandom(context.Context, *GetRandomOptions) (*MoviesArr, error)
	GetMainMovie(context.Context, *GetMainMovieOptions) (*MainMovie, error)
	AddMovieRating(context.Context, *AddRatingOptions) (*NewMovieRating, error)
}

// UnimplementedMoviesServer should be embedded to have forward compatible implementations.
type UnimplementedMoviesServer struct {
}

func (UnimplementedMoviesServer) GetByID(context.Context, *GetMovieOptions) (*Movie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedMoviesServer) GetRandom(context.Context, *GetRandomOptions) (*MoviesArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRandom not implemented")
}
func (UnimplementedMoviesServer) GetMainMovie(context.Context, *GetMainMovieOptions) (*MainMovie, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainMovie not implemented")
}
func (UnimplementedMoviesServer) AddMovieRating(context.Context, *AddRatingOptions) (*NewMovieRating, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMovieRating not implemented")
}

// UnsafeMoviesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesServer will
// result in compilation errors.
type UnsafeMoviesServer interface {
	mustEmbedUnimplementedMoviesServer()
}

func RegisterMoviesServer(s grpc.ServiceRegistrar, srv MoviesServer) {
	s.RegisterService(&Movies_ServiceDesc, srv)
}

func _Movies_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Movies/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetByID(ctx, req.(*GetMovieOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_GetRandom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRandomOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetRandom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Movies/GetRandom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetRandom(ctx, req.(*GetRandomOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_GetMainMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMainMovieOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetMainMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Movies/GetMainMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetMainMovie(ctx, req.(*GetMainMovieOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _Movies_AddMovieRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRatingOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).AddMovieRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Movies/AddMovieRating",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).AddMovieRating(ctx, req.(*AddRatingOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// Movies_ServiceDesc is the grpc.ServiceDesc for Movies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Movies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Movies",
	HandlerType: (*MoviesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByID",
			Handler:    _Movies_GetByID_Handler,
		},
		{
			MethodName: "GetRandom",
			Handler:    _Movies_GetRandom_Handler,
		},
		{
			MethodName: "GetMainMovie",
			Handler:    _Movies_GetMainMovie_Handler,
		},
		{
			MethodName: "AddMovieRating",
			Handler:    _Movies_AddMovieRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "microservices/movie/proto/movie.proto",
}
