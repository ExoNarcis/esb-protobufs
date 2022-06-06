// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: zone_group_restriction.proto

package logistics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ZoneGroupRestrictionServiceClient is the client API for ZoneGroupRestrictionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZoneGroupRestrictionServiceClient interface {
	List(ctx context.Context, in *ListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error)
	Create(ctx context.Context, in *CreateListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error)
	Update(ctx context.Context, in *UpdateListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error)
	Get(ctx context.Context, in *ZoneGroupRestrictionId, opts ...grpc.CallOption) (*ZoneGroupRestriction, error)
	Delete(ctx context.Context, in *ZoneGroupRestrictionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type zoneGroupRestrictionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZoneGroupRestrictionServiceClient(cc grpc.ClientConnInterface) ZoneGroupRestrictionServiceClient {
	return &zoneGroupRestrictionServiceClient{cc}
}

func (c *zoneGroupRestrictionServiceClient) List(ctx context.Context, in *ListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error) {
	out := new(ListZoneGroupRestrictionResponse)
	err := c.cc.Invoke(ctx, "/logistics.ZoneGroupRestrictionService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupRestrictionServiceClient) Create(ctx context.Context, in *CreateListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error) {
	out := new(ListZoneGroupRestrictionResponse)
	err := c.cc.Invoke(ctx, "/logistics.ZoneGroupRestrictionService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupRestrictionServiceClient) Update(ctx context.Context, in *UpdateListZoneGroupRestrictionRequest, opts ...grpc.CallOption) (*ListZoneGroupRestrictionResponse, error) {
	out := new(ListZoneGroupRestrictionResponse)
	err := c.cc.Invoke(ctx, "/logistics.ZoneGroupRestrictionService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupRestrictionServiceClient) Get(ctx context.Context, in *ZoneGroupRestrictionId, opts ...grpc.CallOption) (*ZoneGroupRestriction, error) {
	out := new(ZoneGroupRestriction)
	err := c.cc.Invoke(ctx, "/logistics.ZoneGroupRestrictionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneGroupRestrictionServiceClient) Delete(ctx context.Context, in *ZoneGroupRestrictionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.ZoneGroupRestrictionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneGroupRestrictionServiceServer is the server API for ZoneGroupRestrictionService service.
// All implementations should embed UnimplementedZoneGroupRestrictionServiceServer
// for forward compatibility
type ZoneGroupRestrictionServiceServer interface {
	List(context.Context, *ListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error)
	Create(context.Context, *CreateListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error)
	Update(context.Context, *UpdateListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error)
	Get(context.Context, *ZoneGroupRestrictionId) (*ZoneGroupRestriction, error)
	Delete(context.Context, *ZoneGroupRestrictionId) (*emptypb.Empty, error)
}

// UnimplementedZoneGroupRestrictionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedZoneGroupRestrictionServiceServer struct {
}

func (UnimplementedZoneGroupRestrictionServiceServer) List(context.Context, *ListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedZoneGroupRestrictionServiceServer) Create(context.Context, *CreateListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedZoneGroupRestrictionServiceServer) Update(context.Context, *UpdateListZoneGroupRestrictionRequest) (*ListZoneGroupRestrictionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedZoneGroupRestrictionServiceServer) Get(context.Context, *ZoneGroupRestrictionId) (*ZoneGroupRestriction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedZoneGroupRestrictionServiceServer) Delete(context.Context, *ZoneGroupRestrictionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeZoneGroupRestrictionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZoneGroupRestrictionServiceServer will
// result in compilation errors.
type UnsafeZoneGroupRestrictionServiceServer interface {
	mustEmbedUnimplementedZoneGroupRestrictionServiceServer()
}

func RegisterZoneGroupRestrictionServiceServer(s grpc.ServiceRegistrar, srv ZoneGroupRestrictionServiceServer) {
	s.RegisterService(&ZoneGroupRestrictionService_ServiceDesc, srv)
}

func _ZoneGroupRestrictionService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZoneGroupRestrictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupRestrictionServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.ZoneGroupRestrictionService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupRestrictionServiceServer).List(ctx, req.(*ListZoneGroupRestrictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupRestrictionService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateListZoneGroupRestrictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupRestrictionServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.ZoneGroupRestrictionService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupRestrictionServiceServer).Create(ctx, req.(*CreateListZoneGroupRestrictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupRestrictionService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateListZoneGroupRestrictionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupRestrictionServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.ZoneGroupRestrictionService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupRestrictionServiceServer).Update(ctx, req.(*UpdateListZoneGroupRestrictionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupRestrictionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneGroupRestrictionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupRestrictionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.ZoneGroupRestrictionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupRestrictionServiceServer).Get(ctx, req.(*ZoneGroupRestrictionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneGroupRestrictionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ZoneGroupRestrictionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneGroupRestrictionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.ZoneGroupRestrictionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneGroupRestrictionServiceServer).Delete(ctx, req.(*ZoneGroupRestrictionId))
	}
	return interceptor(ctx, in, info, handler)
}

// ZoneGroupRestrictionService_ServiceDesc is the grpc.ServiceDesc for ZoneGroupRestrictionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZoneGroupRestrictionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.ZoneGroupRestrictionService",
	HandlerType: (*ZoneGroupRestrictionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ZoneGroupRestrictionService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ZoneGroupRestrictionService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ZoneGroupRestrictionService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ZoneGroupRestrictionService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ZoneGroupRestrictionService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zone_group_restriction.proto",
}