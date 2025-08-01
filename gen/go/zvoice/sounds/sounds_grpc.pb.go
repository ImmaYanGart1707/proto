// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: zvoice/sounds/sounds.proto

package sounds

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
	SoundsService_GetSound_FullMethodName = "/sounds.SoundsService/GetSound"
)

// SoundsServiceClient is the client API for SoundsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoundsServiceClient interface {
	GetSound(ctx context.Context, in *GetSoundRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetSoundResponse], error)
}

type soundsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSoundsServiceClient(cc grpc.ClientConnInterface) SoundsServiceClient {
	return &soundsServiceClient{cc}
}

func (c *soundsServiceClient) GetSound(ctx context.Context, in *GetSoundRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetSoundResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SoundsService_ServiceDesc.Streams[0], SoundsService_GetSound_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetSoundRequest, GetSoundResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SoundsService_GetSoundClient = grpc.ServerStreamingClient[GetSoundResponse]

// SoundsServiceServer is the server API for SoundsService service.
// All implementations must embed UnimplementedSoundsServiceServer
// for forward compatibility.
type SoundsServiceServer interface {
	GetSound(*GetSoundRequest, grpc.ServerStreamingServer[GetSoundResponse]) error
	mustEmbedUnimplementedSoundsServiceServer()
}

// UnimplementedSoundsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSoundsServiceServer struct{}

func (UnimplementedSoundsServiceServer) GetSound(*GetSoundRequest, grpc.ServerStreamingServer[GetSoundResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetSound not implemented")
}
func (UnimplementedSoundsServiceServer) mustEmbedUnimplementedSoundsServiceServer() {}
func (UnimplementedSoundsServiceServer) testEmbeddedByValue()                       {}

// UnsafeSoundsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoundsServiceServer will
// result in compilation errors.
type UnsafeSoundsServiceServer interface {
	mustEmbedUnimplementedSoundsServiceServer()
}

func RegisterSoundsServiceServer(s grpc.ServiceRegistrar, srv SoundsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSoundsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SoundsService_ServiceDesc, srv)
}

func _SoundsService_GetSound_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetSoundRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SoundsServiceServer).GetSound(m, &grpc.GenericServerStream[GetSoundRequest, GetSoundResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SoundsService_GetSoundServer = grpc.ServerStreamingServer[GetSoundResponse]

// SoundsService_ServiceDesc is the grpc.ServiceDesc for SoundsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SoundsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sounds.SoundsService",
	HandlerType: (*SoundsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSound",
			Handler:       _SoundsService_GetSound_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zvoice/sounds/sounds.proto",
}
