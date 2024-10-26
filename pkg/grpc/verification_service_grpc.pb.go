// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: verification_service.proto

package grpc

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

const (
	VerificationService_ChangeState_FullMethodName                   = "/centray.VerificationService/ChangeState"
	VerificationService_VerificationStateServer_FullMethodName       = "/centray.VerificationService/VerificationStateServer"
	VerificationService_VerificationDataBidirectional_FullMethodName = "/centray.VerificationService/VerificationDataBidirectional"
)

// VerificationServiceClient is the client API for VerificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationServiceClient interface {
	ChangeState(ctx context.Context, in *ChangeStateRequest, opts ...grpc.CallOption) (*ChangeStateResponse, error)
	VerificationStateServer(ctx context.Context, in *VerificationStateServerRequest, opts ...grpc.CallOption) (VerificationService_VerificationStateServerClient, error)
	VerificationDataBidirectional(ctx context.Context, opts ...grpc.CallOption) (VerificationService_VerificationDataBidirectionalClient, error)
}

type verificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationServiceClient(cc grpc.ClientConnInterface) VerificationServiceClient {
	return &verificationServiceClient{cc}
}

func (c *verificationServiceClient) ChangeState(ctx context.Context, in *ChangeStateRequest, opts ...grpc.CallOption) (*ChangeStateResponse, error) {
	out := new(ChangeStateResponse)
	err := c.cc.Invoke(ctx, VerificationService_ChangeState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationServiceClient) VerificationStateServer(ctx context.Context, in *VerificationStateServerRequest, opts ...grpc.CallOption) (VerificationService_VerificationStateServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &VerificationService_ServiceDesc.Streams[0], VerificationService_VerificationStateServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &verificationServiceVerificationStateServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VerificationService_VerificationStateServerClient interface {
	Recv() (*VerificationStateServerResponse, error)
	grpc.ClientStream
}

type verificationServiceVerificationStateServerClient struct {
	grpc.ClientStream
}

func (x *verificationServiceVerificationStateServerClient) Recv() (*VerificationStateServerResponse, error) {
	m := new(VerificationStateServerResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *verificationServiceClient) VerificationDataBidirectional(ctx context.Context, opts ...grpc.CallOption) (VerificationService_VerificationDataBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &VerificationService_ServiceDesc.Streams[1], VerificationService_VerificationDataBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &verificationServiceVerificationDataBidirectionalClient{stream}
	return x, nil
}

type VerificationService_VerificationDataBidirectionalClient interface {
	Send(*VerificationDataBidirectionalRequest) error
	Recv() (*VerificationDataBidirectionalResponse, error)
	grpc.ClientStream
}

type verificationServiceVerificationDataBidirectionalClient struct {
	grpc.ClientStream
}

func (x *verificationServiceVerificationDataBidirectionalClient) Send(m *VerificationDataBidirectionalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *verificationServiceVerificationDataBidirectionalClient) Recv() (*VerificationDataBidirectionalResponse, error) {
	m := new(VerificationDataBidirectionalResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VerificationServiceServer is the server API for VerificationService service.
// All implementations must embed UnimplementedVerificationServiceServer
// for forward compatibility
type VerificationServiceServer interface {
	ChangeState(context.Context, *ChangeStateRequest) (*ChangeStateResponse, error)
	VerificationStateServer(*VerificationStateServerRequest, VerificationService_VerificationStateServerServer) error
	VerificationDataBidirectional(VerificationService_VerificationDataBidirectionalServer) error
	mustEmbedUnimplementedVerificationServiceServer()
}

// UnimplementedVerificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVerificationServiceServer struct {
}

func (UnimplementedVerificationServiceServer) ChangeState(context.Context, *ChangeStateRequest) (*ChangeStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeState not implemented")
}
func (UnimplementedVerificationServiceServer) VerificationStateServer(*VerificationStateServerRequest, VerificationService_VerificationStateServerServer) error {
	return status.Errorf(codes.Unimplemented, "method VerificationStateServer not implemented")
}
func (UnimplementedVerificationServiceServer) VerificationDataBidirectional(VerificationService_VerificationDataBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method VerificationDataBidirectional not implemented")
}
func (UnimplementedVerificationServiceServer) mustEmbedUnimplementedVerificationServiceServer() {}

// UnsafeVerificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationServiceServer will
// result in compilation errors.
type UnsafeVerificationServiceServer interface {
	mustEmbedUnimplementedVerificationServiceServer()
}

func RegisterVerificationServiceServer(s grpc.ServiceRegistrar, srv VerificationServiceServer) {
	s.RegisterService(&VerificationService_ServiceDesc, srv)
}

func _VerificationService_ChangeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationServiceServer).ChangeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VerificationService_ChangeState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationServiceServer).ChangeState(ctx, req.(*ChangeStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationService_VerificationStateServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(VerificationStateServerRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VerificationServiceServer).VerificationStateServer(m, &verificationServiceVerificationStateServerServer{stream})
}

type VerificationService_VerificationStateServerServer interface {
	Send(*VerificationStateServerResponse) error
	grpc.ServerStream
}

type verificationServiceVerificationStateServerServer struct {
	grpc.ServerStream
}

func (x *verificationServiceVerificationStateServerServer) Send(m *VerificationStateServerResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VerificationService_VerificationDataBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VerificationServiceServer).VerificationDataBidirectional(&verificationServiceVerificationDataBidirectionalServer{stream})
}

type VerificationService_VerificationDataBidirectionalServer interface {
	Send(*VerificationDataBidirectionalResponse) error
	Recv() (*VerificationDataBidirectionalRequest, error)
	grpc.ServerStream
}

type verificationServiceVerificationDataBidirectionalServer struct {
	grpc.ServerStream
}

func (x *verificationServiceVerificationDataBidirectionalServer) Send(m *VerificationDataBidirectionalResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *verificationServiceVerificationDataBidirectionalServer) Recv() (*VerificationDataBidirectionalRequest, error) {
	m := new(VerificationDataBidirectionalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VerificationService_ServiceDesc is the grpc.ServiceDesc for VerificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "centray.VerificationService",
	HandlerType: (*VerificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChangeState",
			Handler:    _VerificationService_ChangeState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "VerificationStateServer",
			Handler:       _VerificationService_VerificationStateServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "VerificationDataBidirectional",
			Handler:       _VerificationService_VerificationDataBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "verification_service.proto",
}