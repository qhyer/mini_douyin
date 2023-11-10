// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: seq.proto

package v1

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
	Seq_GetID_FullMethodName        = "/api.seq.v1.Seq/GetID"
	Seq_UpdateMaxSeq_FullMethodName = "/api.seq.v1.Seq/UpdateMaxSeq"
)

// SeqClient is the client API for Seq service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SeqClient interface {
	GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error)
	UpdateMaxSeq(ctx context.Context, in *UpdateMaxSeqRequest, opts ...grpc.CallOption) (*UpdateMaxSeqResponse, error)
}

type seqClient struct {
	cc grpc.ClientConnInterface
}

func NewSeqClient(cc grpc.ClientConnInterface) SeqClient {
	return &seqClient{cc}
}

func (c *seqClient) GetID(ctx context.Context, in *GetIDRequest, opts ...grpc.CallOption) (*GetIDResponse, error) {
	out := new(GetIDResponse)
	err := c.cc.Invoke(ctx, Seq_GetID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *seqClient) UpdateMaxSeq(ctx context.Context, in *UpdateMaxSeqRequest, opts ...grpc.CallOption) (*UpdateMaxSeqResponse, error) {
	out := new(UpdateMaxSeqResponse)
	err := c.cc.Invoke(ctx, Seq_UpdateMaxSeq_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SeqServer is the server API for Seq service.
// All implementations must embed UnimplementedSeqServer
// for forward compatibility
type SeqServer interface {
	GetID(context.Context, *GetIDRequest) (*GetIDResponse, error)
	UpdateMaxSeq(context.Context, *UpdateMaxSeqRequest) (*UpdateMaxSeqResponse, error)
	mustEmbedUnimplementedSeqServer()
}

// UnimplementedSeqServer must be embedded to have forward compatible implementations.
type UnimplementedSeqServer struct {
}

func (UnimplementedSeqServer) GetID(context.Context, *GetIDRequest) (*GetIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetID not implemented")
}
func (UnimplementedSeqServer) UpdateMaxSeq(context.Context, *UpdateMaxSeqRequest) (*UpdateMaxSeqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaxSeq not implemented")
}
func (UnimplementedSeqServer) mustEmbedUnimplementedSeqServer() {}

// UnsafeSeqServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SeqServer will
// result in compilation errors.
type UnsafeSeqServer interface {
	mustEmbedUnimplementedSeqServer()
}

func RegisterSeqServer(s grpc.ServiceRegistrar, srv SeqServer) {
	s.RegisterService(&Seq_ServiceDesc, srv)
}

func _Seq_GetID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).GetID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_GetID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).GetID(ctx, req.(*GetIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Seq_UpdateMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMaxSeqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SeqServer).UpdateMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Seq_UpdateMaxSeq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SeqServer).UpdateMaxSeq(ctx, req.(*UpdateMaxSeqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Seq_ServiceDesc is the grpc.ServiceDesc for Seq service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Seq_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.seq.v1.Seq",
	HandlerType: (*SeqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetID",
			Handler:    _Seq_GetID_Handler,
		},
		{
			MethodName: "UpdateMaxSeq",
			Handler:    _Seq_UpdateMaxSeq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seq.proto",
}
