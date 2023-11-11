// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: publish.proto

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
	Publish_PublishVideo_FullMethodName                    = "/api.publish.v1.Publish/PublishVideo"
	Publish_GetUserPublishedVideoList_FullMethodName       = "/api.publish.v1.Publish/GetUserPublishedVideoList"
	Publish_GetPublishedVideoByLatestTime_FullMethodName   = "/api.publish.v1.Publish/GetPublishedVideoByLatestTime"
	Publish_GetVideoInfoByVideoIds_FullMethodName          = "/api.publish.v1.Publish/GetVideoInfoByVideoIds"
	Publish_MGetVideoInfoByVideoIds_FullMethodName         = "/api.publish.v1.Publish/MGetVideoInfoByVideoIds"
	Publish_CountUserPublishedVideoByUserId_FullMethodName = "/api.publish.v1.Publish/CountUserPublishedVideoByUserId"
)

// PublishClient is the client API for Publish service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublishClient interface {
	PublishVideo(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error)
	GetUserPublishedVideoList(ctx context.Context, in *GetUserPublishedVideoListRequest, opts ...grpc.CallOption) (*GetUserPublishedVideoListResponse, error)
	GetPublishedVideoByLatestTime(ctx context.Context, in *GetPublishedVideoByLatestTimeRequest, opts ...grpc.CallOption) (*GetPublishedVideoByLatestTimeResponse, error)
	GetVideoInfoByVideoIds(ctx context.Context, in *GetVideoInfoRequest, opts ...grpc.CallOption) (*GetVideoInfoResponse, error)
	MGetVideoInfoByVideoIds(ctx context.Context, in *MGetVideoInfoRequest, opts ...grpc.CallOption) (*MGetVideoInfoResponse, error)
	CountUserPublishedVideoByUserId(ctx context.Context, in *CountUserPublishedVideoByUserIdRequest, opts ...grpc.CallOption) (*CountUserPublishedVideoByUserIdResponse, error)
}

type publishClient struct {
	cc grpc.ClientConnInterface
}

func NewPublishClient(cc grpc.ClientConnInterface) PublishClient {
	return &publishClient{cc}
}

func (c *publishClient) PublishVideo(ctx context.Context, in *PublishActionRequest, opts ...grpc.CallOption) (*PublishActionResponse, error) {
	out := new(PublishActionResponse)
	err := c.cc.Invoke(ctx, Publish_PublishVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetUserPublishedVideoList(ctx context.Context, in *GetUserPublishedVideoListRequest, opts ...grpc.CallOption) (*GetUserPublishedVideoListResponse, error) {
	out := new(GetUserPublishedVideoListResponse)
	err := c.cc.Invoke(ctx, Publish_GetUserPublishedVideoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetPublishedVideoByLatestTime(ctx context.Context, in *GetPublishedVideoByLatestTimeRequest, opts ...grpc.CallOption) (*GetPublishedVideoByLatestTimeResponse, error) {
	out := new(GetPublishedVideoByLatestTimeResponse)
	err := c.cc.Invoke(ctx, Publish_GetPublishedVideoByLatestTime_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) GetVideoInfoByVideoIds(ctx context.Context, in *GetVideoInfoRequest, opts ...grpc.CallOption) (*GetVideoInfoResponse, error) {
	out := new(GetVideoInfoResponse)
	err := c.cc.Invoke(ctx, Publish_GetVideoInfoByVideoIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) MGetVideoInfoByVideoIds(ctx context.Context, in *MGetVideoInfoRequest, opts ...grpc.CallOption) (*MGetVideoInfoResponse, error) {
	out := new(MGetVideoInfoResponse)
	err := c.cc.Invoke(ctx, Publish_MGetVideoInfoByVideoIds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishClient) CountUserPublishedVideoByUserId(ctx context.Context, in *CountUserPublishedVideoByUserIdRequest, opts ...grpc.CallOption) (*CountUserPublishedVideoByUserIdResponse, error) {
	out := new(CountUserPublishedVideoByUserIdResponse)
	err := c.cc.Invoke(ctx, Publish_CountUserPublishedVideoByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublishServer is the server API for Publish service.
// All implementations must embed UnimplementedPublishServer
// for forward compatibility
type PublishServer interface {
	PublishVideo(context.Context, *PublishActionRequest) (*PublishActionResponse, error)
	GetUserPublishedVideoList(context.Context, *GetUserPublishedVideoListRequest) (*GetUserPublishedVideoListResponse, error)
	GetPublishedVideoByLatestTime(context.Context, *GetPublishedVideoByLatestTimeRequest) (*GetPublishedVideoByLatestTimeResponse, error)
	GetVideoInfoByVideoIds(context.Context, *GetVideoInfoRequest) (*GetVideoInfoResponse, error)
	MGetVideoInfoByVideoIds(context.Context, *MGetVideoInfoRequest) (*MGetVideoInfoResponse, error)
	CountUserPublishedVideoByUserId(context.Context, *CountUserPublishedVideoByUserIdRequest) (*CountUserPublishedVideoByUserIdResponse, error)
	mustEmbedUnimplementedPublishServer()
}

// UnimplementedPublishServer must be embedded to have forward compatible implementations.
type UnimplementedPublishServer struct {
}

func (UnimplementedPublishServer) PublishVideo(context.Context, *PublishActionRequest) (*PublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVideo not implemented")
}
func (UnimplementedPublishServer) GetUserPublishedVideoList(context.Context, *GetUserPublishedVideoListRequest) (*GetUserPublishedVideoListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPublishedVideoList not implemented")
}
func (UnimplementedPublishServer) GetPublishedVideoByLatestTime(context.Context, *GetPublishedVideoByLatestTimeRequest) (*GetPublishedVideoByLatestTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishedVideoByLatestTime not implemented")
}
func (UnimplementedPublishServer) GetVideoInfoByVideoIds(context.Context, *GetVideoInfoRequest) (*GetVideoInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoInfoByVideoIds not implemented")
}
func (UnimplementedPublishServer) MGetVideoInfoByVideoIds(context.Context, *MGetVideoInfoRequest) (*MGetVideoInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MGetVideoInfoByVideoIds not implemented")
}
func (UnimplementedPublishServer) CountUserPublishedVideoByUserId(context.Context, *CountUserPublishedVideoByUserIdRequest) (*CountUserPublishedVideoByUserIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountUserPublishedVideoByUserId not implemented")
}
func (UnimplementedPublishServer) mustEmbedUnimplementedPublishServer() {}

// UnsafePublishServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublishServer will
// result in compilation errors.
type UnsafePublishServer interface {
	mustEmbedUnimplementedPublishServer()
}

func RegisterPublishServer(s grpc.ServiceRegistrar, srv PublishServer) {
	s.RegisterService(&Publish_ServiceDesc, srv)
}

func _Publish_PublishVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).PublishVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_PublishVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).PublishVideo(ctx, req.(*PublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetUserPublishedVideoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPublishedVideoListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetUserPublishedVideoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_GetUserPublishedVideoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetUserPublishedVideoList(ctx, req.(*GetUserPublishedVideoListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetPublishedVideoByLatestTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishedVideoByLatestTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetPublishedVideoByLatestTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_GetPublishedVideoByLatestTime_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetPublishedVideoByLatestTime(ctx, req.(*GetPublishedVideoByLatestTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_GetVideoInfoByVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).GetVideoInfoByVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_GetVideoInfoByVideoIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).GetVideoInfoByVideoIds(ctx, req.(*GetVideoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_MGetVideoInfoByVideoIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MGetVideoInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).MGetVideoInfoByVideoIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_MGetVideoInfoByVideoIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).MGetVideoInfoByVideoIds(ctx, req.(*MGetVideoInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publish_CountUserPublishedVideoByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountUserPublishedVideoByUserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublishServer).CountUserPublishedVideoByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Publish_CountUserPublishedVideoByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublishServer).CountUserPublishedVideoByUserId(ctx, req.(*CountUserPublishedVideoByUserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Publish_ServiceDesc is the grpc.ServiceDesc for Publish service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Publish_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.publish.v1.Publish",
	HandlerType: (*PublishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishVideo",
			Handler:    _Publish_PublishVideo_Handler,
		},
		{
			MethodName: "GetUserPublishedVideoList",
			Handler:    _Publish_GetUserPublishedVideoList_Handler,
		},
		{
			MethodName: "GetPublishedVideoByLatestTime",
			Handler:    _Publish_GetPublishedVideoByLatestTime_Handler,
		},
		{
			MethodName: "GetVideoInfoByVideoIds",
			Handler:    _Publish_GetVideoInfoByVideoIds_Handler,
		},
		{
			MethodName: "MGetVideoInfoByVideoIds",
			Handler:    _Publish_MGetVideoInfoByVideoIds_Handler,
		},
		{
			MethodName: "CountUserPublishedVideoByUserId",
			Handler:    _Publish_CountUserPublishedVideoByUserId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "publish.proto",
}
