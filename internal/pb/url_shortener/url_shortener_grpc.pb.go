// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc2
// source: API/url_shortener.proto

package urlshortener

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UrlShortener_ShortenLink_FullMethodName = "/url_shortener.UrlShortener/ShortenLink"
	UrlShortener_GetLink_FullMethodName     = "/url_shortener.UrlShortener/GetLink"
)

// UrlShortenerClient is the client API for UrlShortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerClient interface {
	ShortenLink(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error)
	GetLink(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
}

type urlShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerClient(cc grpc.ClientConnInterface) UrlShortenerClient {
	return &urlShortenerClient{cc}
}

func (c *urlShortenerClient) ShortenLink(ctx context.Context, in *ShortenRequest, opts ...grpc.CallOption) (*ShortenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ShortenResponse)
	err := c.cc.Invoke(ctx, UrlShortener_ShortenLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerClient) GetLink(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, UrlShortener_GetLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServer is the server API for UrlShortener service.
// All implementations must embed UnimplementedUrlShortenerServer
// for forward compatibility.
type UrlShortenerServer interface {
	ShortenLink(context.Context, *ShortenRequest) (*ShortenResponse, error)
	GetLink(context.Context, *ResolveRequest) (*ResolveResponse, error)
	mustEmbedUnimplementedUrlShortenerServer()
}

// UnimplementedUrlShortenerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUrlShortenerServer struct{}

func (UnimplementedUrlShortenerServer) ShortenLink(context.Context, *ShortenRequest) (*ShortenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenLink not implemented")
}
func (UnimplementedUrlShortenerServer) GetLink(context.Context, *ResolveRequest) (*ResolveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (UnimplementedUrlShortenerServer) mustEmbedUnimplementedUrlShortenerServer() {}
func (UnimplementedUrlShortenerServer) testEmbeddedByValue()                      {}

// UnsafeUrlShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServer will
// result in compilation errors.
type UnsafeUrlShortenerServer interface {
	mustEmbedUnimplementedUrlShortenerServer()
}

func RegisterUrlShortenerServer(s grpc.ServiceRegistrar, srv UrlShortenerServer) {
	// If the following call pancis, it indicates UnimplementedUrlShortenerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UrlShortener_ServiceDesc, srv)
}

func _UrlShortener_ShortenLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).ShortenLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortener_ShortenLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).ShortenLink(ctx, req.(*ShortenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortener_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortener_GetLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).GetLink(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortener_ServiceDesc is the grpc.ServiceDesc for UrlShortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "url_shortener.UrlShortener",
	HandlerType: (*UrlShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenLink",
			Handler:    _UrlShortener_ShortenLink_Handler,
		},
		{
			MethodName: "GetLink",
			Handler:    _UrlShortener_GetLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "API/url_shortener.proto",
}
