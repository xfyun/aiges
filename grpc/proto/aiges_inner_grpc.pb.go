// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: grpc/proto/aiges_inner.proto

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

// WrapperServiceClient is the client API for WrapperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WrapperServiceClient interface {
	WrapperInit(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Ret, error)
	WrapperOnceExec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	TestStream(ctx context.Context, opts ...grpc.CallOption) (WrapperService_TestStreamClient, error)
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	Communicate(ctx context.Context, opts ...grpc.CallOption) (WrapperService_CommunicateClient, error)
}

type wrapperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWrapperServiceClient(cc grpc.ClientConnInterface) WrapperServiceClient {
	return &wrapperServiceClient{cc}
}

func (c *wrapperServiceClient) WrapperInit(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Ret, error) {
	out := new(Ret)
	err := c.cc.Invoke(ctx, "/aiges.WrapperService/wrapperInit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrapperServiceClient) WrapperOnceExec(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/aiges.WrapperService/wrapperOnceExec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wrapperServiceClient) TestStream(ctx context.Context, opts ...grpc.CallOption) (WrapperService_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &WrapperService_ServiceDesc.Streams[0], "/aiges.WrapperService/testStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &wrapperServiceTestStreamClient{stream}
	return x, nil
}

type WrapperService_TestStreamClient interface {
	Send(*StreamRequest) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type wrapperServiceTestStreamClient struct {
	grpc.ClientStream
}

func (x *wrapperServiceTestStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wrapperServiceTestStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *wrapperServiceClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (WrapperService_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &WrapperService_ServiceDesc.Streams[1], "/aiges.WrapperService/communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &wrapperServiceCommunicateClient{stream}
	return x, nil
}

type WrapperService_CommunicateClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type wrapperServiceCommunicateClient struct {
	grpc.ClientStream
}

func (x *wrapperServiceCommunicateClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *wrapperServiceCommunicateClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WrapperServiceServer is the server API for WrapperService service.
// All implementations should embed UnimplementedWrapperServiceServer
// for forward compatibility
type WrapperServiceServer interface {
	WrapperInit(context.Context, *InitRequest) (*Ret, error)
	WrapperOnceExec(context.Context, *Request) (*Response, error)
	TestStream(WrapperService_TestStreamServer) error
	//
	// Accepts a stream of RouteNotes sent while a route is being traversed,
	// while receiving other RouteNotes (e.g. from other users).
	Communicate(WrapperService_CommunicateServer) error
}

// UnimplementedWrapperServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWrapperServiceServer struct {
}

func (UnimplementedWrapperServiceServer) WrapperInit(context.Context, *InitRequest) (*Ret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WrapperInit not implemented")
}
func (UnimplementedWrapperServiceServer) WrapperOnceExec(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WrapperOnceExec not implemented")
}
func (UnimplementedWrapperServiceServer) TestStream(WrapperService_TestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}
func (UnimplementedWrapperServiceServer) Communicate(WrapperService_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}

// UnsafeWrapperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WrapperServiceServer will
// result in compilation errors.
type UnsafeWrapperServiceServer interface {
	mustEmbedUnimplementedWrapperServiceServer()
}

func RegisterWrapperServiceServer(s grpc.ServiceRegistrar, srv WrapperServiceServer) {
	s.RegisterService(&WrapperService_ServiceDesc, srv)
}

func _WrapperService_WrapperInit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrapperServiceServer).WrapperInit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aiges.WrapperService/wrapperInit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrapperServiceServer).WrapperInit(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WrapperService_WrapperOnceExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WrapperServiceServer).WrapperOnceExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/aiges.WrapperService/wrapperOnceExec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WrapperServiceServer).WrapperOnceExec(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _WrapperService_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WrapperServiceServer).TestStream(&wrapperServiceTestStreamServer{stream})
}

type WrapperService_TestStreamServer interface {
	Send(*Response) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type wrapperServiceTestStreamServer struct {
	grpc.ServerStream
}

func (x *wrapperServiceTestStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wrapperServiceTestStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _WrapperService_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WrapperServiceServer).Communicate(&wrapperServiceCommunicateServer{stream})
}

type WrapperService_CommunicateServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type wrapperServiceCommunicateServer struct {
	grpc.ServerStream
}

func (x *wrapperServiceCommunicateServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *wrapperServiceCommunicateServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WrapperService_ServiceDesc is the grpc.ServiceDesc for WrapperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WrapperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aiges.WrapperService",
	HandlerType: (*WrapperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "wrapperInit",
			Handler:    _WrapperService_WrapperInit_Handler,
		},
		{
			MethodName: "wrapperOnceExec",
			Handler:    _WrapperService_WrapperOnceExec_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "testStream",
			Handler:       _WrapperService_TestStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "communicate",
			Handler:       _WrapperService_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/proto/aiges_inner.proto",
}
