// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: movie.proto

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
	GetMovie(ctx context.Context, in *GetMovieOptions, opts ...grpc.CallOption) (*GetMovieResponse, error)
}

type moviesClient struct {
	cc grpc.ClientConnInterface
}

func NewMoviesClient(cc grpc.ClientConnInterface) MoviesClient {
	return &moviesClient{cc}
}

func (c *moviesClient) GetMovie(ctx context.Context, in *GetMovieOptions, opts ...grpc.CallOption) (*GetMovieResponse, error) {
	out := new(GetMovieResponse)
	err := c.cc.Invoke(ctx, "/Movies/GetMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoviesServer is the server API for Movies service.
// All implementations must embed UnimplementedMoviesServer
// for forward compatibility
type MoviesServer interface {
	GetMovie(context.Context, *GetMovieOptions) (*GetMovieResponse, error)
	mustEmbedUnimplementedMoviesServer()
}

// UnimplementedMoviesServer must be embedded to have forward compatible implementations.
type UnimplementedMoviesServer struct {
}

func (UnimplementedMoviesServer) GetMovie(context.Context, *GetMovieOptions) (*GetMovieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
}
func (UnimplementedMoviesServer) mustEmbedUnimplementedMoviesServer() {}

// UnsafeMoviesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoviesServer will
// result in compilation errors.
type UnsafeMoviesServer interface {
	mustEmbedUnimplementedMoviesServer()
}

func RegisterMoviesServer(s grpc.ServiceRegistrar, srv MoviesServer) {
	s.RegisterService(&Movies_ServiceDesc, srv)
}

func _Movies_GetMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoviesServer).GetMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Movies/GetMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoviesServer).GetMovie(ctx, req.(*GetMovieOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// Movies_ServiceDesc is the grpc.ServiceDesc for Movies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Movies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Movies",
	HandlerType: (*MoviesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovie",
			Handler:    _Movies_GetMovie_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}
