// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: compilations.proto

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

// MovieCompilationsClient is the client API for MovieCompilations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MovieCompilationsClient interface {
	GetAllMovies(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetAllSeries(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetMainCompilations(ctx context.Context, in *GetMainCompilationsOptions, opts ...grpc.CallOption) (*MovieCompilationsArr, error)
	GetByGenre(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetByCountry(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetByMovie(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetByPerson(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetTopByYear(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
	GetTop(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error)
}

type movieCompilationsClient struct {
	cc grpc.ClientConnInterface
}

func NewMovieCompilationsClient(cc grpc.ClientConnInterface) MovieCompilationsClient {
	return &movieCompilationsClient{cc}
}

func (c *movieCompilationsClient) GetAllMovies(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetAllMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetAllSeries(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetAllSeries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetMainCompilations(ctx context.Context, in *GetMainCompilationsOptions, opts ...grpc.CallOption) (*MovieCompilationsArr, error) {
	out := new(MovieCompilationsArr)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetMainCompilations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetByGenre(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetByGenre", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetByCountry(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetByCountry", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetByMovie(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetByMovie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetByPerson(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetByPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetTopByYear(ctx context.Context, in *GetByIDOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetTopByYear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieCompilationsClient) GetTop(ctx context.Context, in *GetCompilationOptions, opts ...grpc.CallOption) (*MovieCompilation, error) {
	out := new(MovieCompilation)
	err := c.cc.Invoke(ctx, "/MovieCompilations/GetTop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieCompilationsServer is the server API for MovieCompilations service.
// All implementations must embed UnimplementedMovieCompilationsServer
// for forward compatibility
type MovieCompilationsServer interface {
	GetAllMovies(context.Context, *GetCompilationOptions) (*MovieCompilation, error)
	GetAllSeries(context.Context, *GetCompilationOptions) (*MovieCompilation, error)
	GetMainCompilations(context.Context, *GetMainCompilationsOptions) (*MovieCompilationsArr, error)
	GetByGenre(context.Context, *GetByIDOptions) (*MovieCompilation, error)
	GetByCountry(context.Context, *GetByIDOptions) (*MovieCompilation, error)
	GetByMovie(context.Context, *GetByIDOptions) (*MovieCompilation, error)
	GetByPerson(context.Context, *GetByIDOptions) (*MovieCompilation, error)
	GetTopByYear(context.Context, *GetByIDOptions) (*MovieCompilation, error)
	GetTop(context.Context, *GetCompilationOptions) (*MovieCompilation, error)
	mustEmbedUnimplementedMovieCompilationsServer()
}

// UnimplementedMovieCompilationsServer must be embedded to have forward compatible implementations.
type UnimplementedMovieCompilationsServer struct {
}

func (UnimplementedMovieCompilationsServer) GetAllMovies(context.Context, *GetCompilationOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMovies not implemented")
}
func (UnimplementedMovieCompilationsServer) GetAllSeries(context.Context, *GetCompilationOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSeries not implemented")
}
func (UnimplementedMovieCompilationsServer) GetMainCompilations(context.Context, *GetMainCompilationsOptions) (*MovieCompilationsArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMainCompilations not implemented")
}
func (UnimplementedMovieCompilationsServer) GetByGenre(context.Context, *GetByIDOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByGenre not implemented")
}
func (UnimplementedMovieCompilationsServer) GetByCountry(context.Context, *GetByIDOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCountry not implemented")
}
func (UnimplementedMovieCompilationsServer) GetByMovie(context.Context, *GetByIDOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByMovie not implemented")
}
func (UnimplementedMovieCompilationsServer) GetByPerson(context.Context, *GetByIDOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPerson not implemented")
}
func (UnimplementedMovieCompilationsServer) GetTopByYear(context.Context, *GetByIDOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopByYear not implemented")
}
func (UnimplementedMovieCompilationsServer) GetTop(context.Context, *GetCompilationOptions) (*MovieCompilation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTop not implemented")
}
func (UnimplementedMovieCompilationsServer) mustEmbedUnimplementedMovieCompilationsServer() {}

// UnsafeMovieCompilationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MovieCompilationsServer will
// result in compilation errors.
type UnsafeMovieCompilationsServer interface {
	mustEmbedUnimplementedMovieCompilationsServer()
}

func RegisterMovieCompilationsServer(s grpc.ServiceRegistrar, srv MovieCompilationsServer) {
	s.RegisterService(&MovieCompilations_ServiceDesc, srv)
}

func _MovieCompilations_GetAllMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompilationOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetAllMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetAllMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetAllMovies(ctx, req.(*GetCompilationOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetAllSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompilationOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetAllSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetAllSeries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetAllSeries(ctx, req.(*GetCompilationOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetMainCompilations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMainCompilationsOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetMainCompilations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetMainCompilations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetMainCompilations(ctx, req.(*GetMainCompilationsOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetByGenre_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetByGenre(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetByGenre",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetByGenre(ctx, req.(*GetByIDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetByCountry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetByCountry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetByCountry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetByCountry(ctx, req.(*GetByIDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetByMovie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetByMovie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetByMovie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetByMovie(ctx, req.(*GetByIDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetByPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetByPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetByPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetByPerson(ctx, req.(*GetByIDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetTopByYear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIDOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetTopByYear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetTopByYear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetTopByYear(ctx, req.(*GetByIDOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieCompilations_GetTop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompilationOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieCompilationsServer).GetTop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MovieCompilations/GetTop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieCompilationsServer).GetTop(ctx, req.(*GetCompilationOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// MovieCompilations_ServiceDesc is the grpc.ServiceDesc for MovieCompilations service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MovieCompilations_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MovieCompilations",
	HandlerType: (*MovieCompilationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMovies",
			Handler:    _MovieCompilations_GetAllMovies_Handler,
		},
		{
			MethodName: "GetAllSeries",
			Handler:    _MovieCompilations_GetAllSeries_Handler,
		},
		{
			MethodName: "GetMainCompilations",
			Handler:    _MovieCompilations_GetMainCompilations_Handler,
		},
		{
			MethodName: "GetByGenre",
			Handler:    _MovieCompilations_GetByGenre_Handler,
		},
		{
			MethodName: "GetByCountry",
			Handler:    _MovieCompilations_GetByCountry_Handler,
		},
		{
			MethodName: "GetByMovie",
			Handler:    _MovieCompilations_GetByMovie_Handler,
		},
		{
			MethodName: "GetByPerson",
			Handler:    _MovieCompilations_GetByPerson_Handler,
		},
		{
			MethodName: "GetTopByYear",
			Handler:    _MovieCompilations_GetTopByYear_Handler,
		},
		{
			MethodName: "GetTop",
			Handler:    _MovieCompilations_GetTop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compilations.proto",
}
