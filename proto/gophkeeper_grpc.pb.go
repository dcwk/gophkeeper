// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: proto/gophkeeper.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Gophkeeper_Auth_FullMethodName               = "/gophkeeper.Gophkeeper/Auth"
	Gophkeeper_Register_FullMethodName           = "/gophkeeper.Gophkeeper/Register"
	Gophkeeper_GetUserSecretsList_FullMethodName = "/gophkeeper.Gophkeeper/GetUserSecretsList"
	Gophkeeper_SaveAuthPair_FullMethodName       = "/gophkeeper.Gophkeeper/SaveAuthPair"
	Gophkeeper_SavePayCard_FullMethodName        = "/gophkeeper.Gophkeeper/SavePayCard"
	Gophkeeper_SaveBinaryData_FullMethodName     = "/gophkeeper.Gophkeeper/SaveBinaryData"
	Gophkeeper_SaveTextData_FullMethodName       = "/gophkeeper.Gophkeeper/SaveTextData"
	Gophkeeper_DeleteItem_FullMethodName         = "/gophkeeper.Gophkeeper/DeleteItem"
	Gophkeeper_GetAuthPair_FullMethodName        = "/gophkeeper.Gophkeeper/GetAuthPair"
	Gophkeeper_GetPayCard_FullMethodName         = "/gophkeeper.Gophkeeper/GetPayCard"
	Gophkeeper_GetBinaryData_FullMethodName      = "/gophkeeper.Gophkeeper/GetBinaryData"
	Gophkeeper_GetTextData_FullMethodName        = "/gophkeeper.Gophkeeper/GetTextData"
)

// GophkeeperClient is the client API for Gophkeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophkeeperClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUserSecretsList(ctx context.Context, in *GetUserSecretsRequest, opts ...grpc.CallOption) (*GetUserSecretsResponse, error)
	SaveAuthPair(ctx context.Context, in *SaveAuthPairRequest, opts ...grpc.CallOption) (*SaveAuthPairResponse, error)
	SavePayCard(ctx context.Context, in *SavePayCardRequest, opts ...grpc.CallOption) (*SavePayCardResponse, error)
	SaveBinaryData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, SaveBinaryDataResponse], error)
	SaveTextData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, SaveTextDataResponse], error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error)
	GetAuthPair(ctx context.Context, in *GetAuthPairRequest, opts ...grpc.CallOption) (*GetAuthPairResponse, error)
	GetPayCard(ctx context.Context, in *GetPayCardRequest, opts ...grpc.CallOption) (*GetPayCardResponse, error)
	GetBinaryData(ctx context.Context, in *GetBinaryDataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error)
	GetTextData(ctx context.Context, in *GetTextDataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error)
}

type gophkeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophkeeperClient(cc grpc.ClientConnInterface) GophkeeperClient {
	return &gophkeeperClient{cc}
}

func (c *gophkeeperClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_Auth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetUserSecretsList(ctx context.Context, in *GetUserSecretsRequest, opts ...grpc.CallOption) (*GetUserSecretsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserSecretsResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_GetUserSecretsList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveAuthPair(ctx context.Context, in *SaveAuthPairRequest, opts ...grpc.CallOption) (*SaveAuthPairResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveAuthPairResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_SaveAuthPair_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SavePayCard(ctx context.Context, in *SavePayCardRequest, opts ...grpc.CallOption) (*SavePayCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SavePayCardResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_SavePayCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) SaveBinaryData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, SaveBinaryDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Gophkeeper_ServiceDesc.Streams[0], Gophkeeper_SaveBinaryData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, SaveBinaryDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_SaveBinaryDataClient = grpc.ClientStreamingClient[wrapperspb.StringValue, SaveBinaryDataResponse]

func (c *gophkeeperClient) SaveTextData(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[wrapperspb.StringValue, SaveTextDataResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Gophkeeper_ServiceDesc.Streams[1], Gophkeeper_SaveTextData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[wrapperspb.StringValue, SaveTextDataResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_SaveTextDataClient = grpc.ClientStreamingClient[wrapperspb.StringValue, SaveTextDataResponse]

func (c *gophkeeperClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*DeleteItemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteItemResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_DeleteItem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetAuthPair(ctx context.Context, in *GetAuthPairRequest, opts ...grpc.CallOption) (*GetAuthPairResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAuthPairResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_GetAuthPair_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetPayCard(ctx context.Context, in *GetPayCardRequest, opts ...grpc.CallOption) (*GetPayCardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPayCardResponse)
	err := c.cc.Invoke(ctx, Gophkeeper_GetPayCard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophkeeperClient) GetBinaryData(ctx context.Context, in *GetBinaryDataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Gophkeeper_ServiceDesc.Streams[2], Gophkeeper_GetBinaryData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetBinaryDataRequest, wrapperspb.StringValue]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_GetBinaryDataClient = grpc.ServerStreamingClient[wrapperspb.StringValue]

func (c *gophkeeperClient) GetTextData(ctx context.Context, in *GetTextDataRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[wrapperspb.StringValue], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Gophkeeper_ServiceDesc.Streams[3], Gophkeeper_GetTextData_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetTextDataRequest, wrapperspb.StringValue]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_GetTextDataClient = grpc.ServerStreamingClient[wrapperspb.StringValue]

// GophkeeperServer is the server API for Gophkeeper service.
// All implementations must embed UnimplementedGophkeeperServer
// for forward compatibility.
type GophkeeperServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUserSecretsList(context.Context, *GetUserSecretsRequest) (*GetUserSecretsResponse, error)
	SaveAuthPair(context.Context, *SaveAuthPairRequest) (*SaveAuthPairResponse, error)
	SavePayCard(context.Context, *SavePayCardRequest) (*SavePayCardResponse, error)
	SaveBinaryData(grpc.ClientStreamingServer[wrapperspb.StringValue, SaveBinaryDataResponse]) error
	SaveTextData(grpc.ClientStreamingServer[wrapperspb.StringValue, SaveTextDataResponse]) error
	DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error)
	GetAuthPair(context.Context, *GetAuthPairRequest) (*GetAuthPairResponse, error)
	GetPayCard(context.Context, *GetPayCardRequest) (*GetPayCardResponse, error)
	GetBinaryData(*GetBinaryDataRequest, grpc.ServerStreamingServer[wrapperspb.StringValue]) error
	GetTextData(*GetTextDataRequest, grpc.ServerStreamingServer[wrapperspb.StringValue]) error
	mustEmbedUnimplementedGophkeeperServer()
}

// UnimplementedGophkeeperServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGophkeeperServer struct{}

func (UnimplementedGophkeeperServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedGophkeeperServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophkeeperServer) GetUserSecretsList(context.Context, *GetUserSecretsRequest) (*GetUserSecretsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserSecretsList not implemented")
}
func (UnimplementedGophkeeperServer) SaveAuthPair(context.Context, *SaveAuthPairRequest) (*SaveAuthPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveAuthPair not implemented")
}
func (UnimplementedGophkeeperServer) SavePayCard(context.Context, *SavePayCardRequest) (*SavePayCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePayCard not implemented")
}
func (UnimplementedGophkeeperServer) SaveBinaryData(grpc.ClientStreamingServer[wrapperspb.StringValue, SaveBinaryDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SaveBinaryData not implemented")
}
func (UnimplementedGophkeeperServer) SaveTextData(grpc.ClientStreamingServer[wrapperspb.StringValue, SaveTextDataResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SaveTextData not implemented")
}
func (UnimplementedGophkeeperServer) DeleteItem(context.Context, *DeleteItemRequest) (*DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedGophkeeperServer) GetAuthPair(context.Context, *GetAuthPairRequest) (*GetAuthPairResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthPair not implemented")
}
func (UnimplementedGophkeeperServer) GetPayCard(context.Context, *GetPayCardRequest) (*GetPayCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayCard not implemented")
}
func (UnimplementedGophkeeperServer) GetBinaryData(*GetBinaryDataRequest, grpc.ServerStreamingServer[wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method GetBinaryData not implemented")
}
func (UnimplementedGophkeeperServer) GetTextData(*GetTextDataRequest, grpc.ServerStreamingServer[wrapperspb.StringValue]) error {
	return status.Errorf(codes.Unimplemented, "method GetTextData not implemented")
}
func (UnimplementedGophkeeperServer) mustEmbedUnimplementedGophkeeperServer() {}
func (UnimplementedGophkeeperServer) testEmbeddedByValue()                    {}

// UnsafeGophkeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophkeeperServer will
// result in compilation errors.
type UnsafeGophkeeperServer interface {
	mustEmbedUnimplementedGophkeeperServer()
}

func RegisterGophkeeperServer(s grpc.ServiceRegistrar, srv GophkeeperServer) {
	// If the following call pancis, it indicates UnimplementedGophkeeperServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Gophkeeper_ServiceDesc, srv)
}

func _Gophkeeper_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_Auth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetUserSecretsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetUserSecretsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_GetUserSecretsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetUserSecretsList(ctx, req.(*GetUserSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveAuthPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveAuthPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SaveAuthPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_SaveAuthPair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SaveAuthPair(ctx, req.(*SaveAuthPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SavePayCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePayCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).SavePayCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_SavePayCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).SavePayCard(ctx, req.(*SavePayCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_SaveBinaryData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GophkeeperServer).SaveBinaryData(&grpc.GenericServerStream[wrapperspb.StringValue, SaveBinaryDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_SaveBinaryDataServer = grpc.ClientStreamingServer[wrapperspb.StringValue, SaveBinaryDataResponse]

func _Gophkeeper_SaveTextData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GophkeeperServer).SaveTextData(&grpc.GenericServerStream[wrapperspb.StringValue, SaveTextDataResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_SaveTextDataServer = grpc.ClientStreamingServer[wrapperspb.StringValue, SaveTextDataResponse]

func _Gophkeeper_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_DeleteItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetAuthPair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthPairRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetAuthPair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_GetAuthPair_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetAuthPair(ctx, req.(*GetAuthPairRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetPayCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPayCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophkeeperServer).GetPayCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gophkeeper_GetPayCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophkeeperServer).GetPayCard(ctx, req.(*GetPayCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gophkeeper_GetBinaryData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetBinaryDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GophkeeperServer).GetBinaryData(m, &grpc.GenericServerStream[GetBinaryDataRequest, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_GetBinaryDataServer = grpc.ServerStreamingServer[wrapperspb.StringValue]

func _Gophkeeper_GetTextData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTextDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GophkeeperServer).GetTextData(m, &grpc.GenericServerStream[GetTextDataRequest, wrapperspb.StringValue]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Gophkeeper_GetTextDataServer = grpc.ServerStreamingServer[wrapperspb.StringValue]

// Gophkeeper_ServiceDesc is the grpc.ServiceDesc for Gophkeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gophkeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.Gophkeeper",
	HandlerType: (*GophkeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _Gophkeeper_Auth_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Gophkeeper_Register_Handler,
		},
		{
			MethodName: "GetUserSecretsList",
			Handler:    _Gophkeeper_GetUserSecretsList_Handler,
		},
		{
			MethodName: "SaveAuthPair",
			Handler:    _Gophkeeper_SaveAuthPair_Handler,
		},
		{
			MethodName: "SavePayCard",
			Handler:    _Gophkeeper_SavePayCard_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _Gophkeeper_DeleteItem_Handler,
		},
		{
			MethodName: "GetAuthPair",
			Handler:    _Gophkeeper_GetAuthPair_Handler,
		},
		{
			MethodName: "GetPayCard",
			Handler:    _Gophkeeper_GetPayCard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveBinaryData",
			Handler:       _Gophkeeper_SaveBinaryData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SaveTextData",
			Handler:       _Gophkeeper_SaveTextData_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetBinaryData",
			Handler:       _Gophkeeper_GetBinaryData_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTextData",
			Handler:       _Gophkeeper_GetTextData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/gophkeeper.proto",
}
