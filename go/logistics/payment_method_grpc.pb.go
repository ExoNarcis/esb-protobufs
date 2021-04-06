// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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
const _ = grpc.SupportPackageIsVersion7

// PaymentMethodsClient is the client API for PaymentMethods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaymentMethodsClient interface {
	Create(ctx context.Context, in *PaymentMethod, opts ...grpc.CallOption) (*PaymentMethodId, error)
	Get(ctx context.Context, in *PaymentMethodId, opts ...grpc.CallOption) (*PaymentMethod, error)
	List(ctx context.Context, in *ListPaymentMethodsRequest, opts ...grpc.CallOption) (*ListPaymentMethodsResponse, error)
	Update(ctx context.Context, in *PaymentMethod, opts ...grpc.CallOption) (*PaymentMethod, error)
	Delete(ctx context.Context, in *PaymentMethodId, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type paymentMethodsClient struct {
	cc grpc.ClientConnInterface
}

func NewPaymentMethodsClient(cc grpc.ClientConnInterface) PaymentMethodsClient {
	return &paymentMethodsClient{cc}
}

func (c *paymentMethodsClient) Create(ctx context.Context, in *PaymentMethod, opts ...grpc.CallOption) (*PaymentMethodId, error) {
	out := new(PaymentMethodId)
	err := c.cc.Invoke(ctx, "/logistics.PaymentMethods/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentMethodsClient) Get(ctx context.Context, in *PaymentMethodId, opts ...grpc.CallOption) (*PaymentMethod, error) {
	out := new(PaymentMethod)
	err := c.cc.Invoke(ctx, "/logistics.PaymentMethods/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentMethodsClient) List(ctx context.Context, in *ListPaymentMethodsRequest, opts ...grpc.CallOption) (*ListPaymentMethodsResponse, error) {
	out := new(ListPaymentMethodsResponse)
	err := c.cc.Invoke(ctx, "/logistics.PaymentMethods/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentMethodsClient) Update(ctx context.Context, in *PaymentMethod, opts ...grpc.CallOption) (*PaymentMethod, error) {
	out := new(PaymentMethod)
	err := c.cc.Invoke(ctx, "/logistics.PaymentMethods/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paymentMethodsClient) Delete(ctx context.Context, in *PaymentMethodId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/logistics.PaymentMethods/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaymentMethodsServer is the server API for PaymentMethods service.
// All implementations must embed UnimplementedPaymentMethodsServer
// for forward compatibility
type PaymentMethodsServer interface {
	Create(context.Context, *PaymentMethod) (*PaymentMethodId, error)
	Get(context.Context, *PaymentMethodId) (*PaymentMethod, error)
	List(context.Context, *ListPaymentMethodsRequest) (*ListPaymentMethodsResponse, error)
	Update(context.Context, *PaymentMethod) (*PaymentMethod, error)
	Delete(context.Context, *PaymentMethodId) (*emptypb.Empty, error)
	mustEmbedUnimplementedPaymentMethodsServer()
}

// UnimplementedPaymentMethodsServer must be embedded to have forward compatible implementations.
type UnimplementedPaymentMethodsServer struct {
}

func (UnimplementedPaymentMethodsServer) Create(context.Context, *PaymentMethod) (*PaymentMethodId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPaymentMethodsServer) Get(context.Context, *PaymentMethodId) (*PaymentMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPaymentMethodsServer) List(context.Context, *ListPaymentMethodsRequest) (*ListPaymentMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPaymentMethodsServer) Update(context.Context, *PaymentMethod) (*PaymentMethod, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPaymentMethodsServer) Delete(context.Context, *PaymentMethodId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPaymentMethodsServer) mustEmbedUnimplementedPaymentMethodsServer() {}

// UnsafePaymentMethodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaymentMethodsServer will
// result in compilation errors.
type UnsafePaymentMethodsServer interface {
	mustEmbedUnimplementedPaymentMethodsServer()
}

func RegisterPaymentMethodsServer(s grpc.ServiceRegistrar, srv PaymentMethodsServer) {
	s.RegisterService(&_PaymentMethods_serviceDesc, srv)
}

func _PaymentMethods_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMethod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.PaymentMethods/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).Create(ctx, req.(*PaymentMethod))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentMethods_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMethodId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.PaymentMethods/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).Get(ctx, req.(*PaymentMethodId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentMethods_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPaymentMethodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.PaymentMethods/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).List(ctx, req.(*ListPaymentMethodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentMethods_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMethod)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.PaymentMethods/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).Update(ctx, req.(*PaymentMethod))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaymentMethods_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMethodId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaymentMethodsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logistics.PaymentMethods/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaymentMethodsServer).Delete(ctx, req.(*PaymentMethodId))
	}
	return interceptor(ctx, in, info, handler)
}

var _PaymentMethods_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logistics.PaymentMethods",
	HandlerType: (*PaymentMethodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PaymentMethods_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PaymentMethods_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PaymentMethods_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PaymentMethods_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PaymentMethods_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/logistics/payment_method.proto",
}
