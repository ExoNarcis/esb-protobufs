// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/mindbox.proto

package mindbox

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Mindbox service

type MindboxService interface {
	Ping(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*PingResponse, error)
	UserInformation(ctx context.Context, in *ParamsUserInformation, opts ...client.CallOption) (*ResponseUserInformation, error)
	OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, opts ...client.CallOption) (*ResponseOrdersHistory, error)
}

type mindboxService struct {
	c    client.Client
	name string
}

func NewMindboxService(name string, c client.Client) MindboxService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mindbox"
	}
	return &mindboxService{
		c:    c,
		name: name,
	}
}

func (c *mindboxService) Ping(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*PingResponse, error) {
	req := c.c.NewRequest(c.name, "Mindbox.Ping", in)
	out := new(PingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mindboxService) UserInformation(ctx context.Context, in *ParamsUserInformation, opts ...client.CallOption) (*ResponseUserInformation, error) {
	req := c.c.NewRequest(c.name, "Mindbox.UserInformation", in)
	out := new(ResponseUserInformation)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mindboxService) OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, opts ...client.CallOption) (*ResponseOrdersHistory, error) {
	req := c.c.NewRequest(c.name, "Mindbox.OrdersHistory", in)
	out := new(ResponseOrdersHistory)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Mindbox service

type MindboxHandler interface {
	Ping(context.Context, *empty.Empty, *PingResponse) error
	UserInformation(context.Context, *ParamsUserInformation, *ResponseUserInformation) error
	OrdersHistory(context.Context, *ParamsOrdersHistory, *ResponseOrdersHistory) error
}

func RegisterMindboxHandler(s server.Server, hdlr MindboxHandler, opts ...server.HandlerOption) error {
	type mindbox interface {
		Ping(ctx context.Context, in *empty.Empty, out *PingResponse) error
		UserInformation(ctx context.Context, in *ParamsUserInformation, out *ResponseUserInformation) error
		OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, out *ResponseOrdersHistory) error
	}
	type Mindbox struct {
		mindbox
	}
	h := &mindboxHandler{hdlr}
	return s.Handle(s.NewHandler(&Mindbox{h}, opts...))
}

type mindboxHandler struct {
	MindboxHandler
}

func (h *mindboxHandler) Ping(ctx context.Context, in *empty.Empty, out *PingResponse) error {
	return h.MindboxHandler.Ping(ctx, in, out)
}

func (h *mindboxHandler) UserInformation(ctx context.Context, in *ParamsUserInformation, out *ResponseUserInformation) error {
	return h.MindboxHandler.UserInformation(ctx, in, out)
}

func (h *mindboxHandler) OrdersHistory(ctx context.Context, in *ParamsOrdersHistory, out *ResponseOrdersHistory) error {
	return h.MindboxHandler.OrdersHistory(ctx, in, out)
}

// Client API for MobileAuthentication service

type MobileAuthenticationService interface {
	Creation(ctx context.Context, in *ParamsCreation, opts ...client.CallOption) (*ResponseCreation, error)
	Authorization(ctx context.Context, in *ParamsAuthorization, opts ...client.CallOption) (*ResponseAuthorization, error)
	Registration(ctx context.Context, in *ParamsRegistration, opts ...client.CallOption) (*ResponseRegistration, error)
	Code(ctx context.Context, in *ParamsCode, opts ...client.CallOption) (*ResponseCode, error)
	CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...client.CallOption) (*ResponseCheckCode, error)
}

type mobileAuthenticationService struct {
	c    client.Client
	name string
}

func NewMobileAuthenticationService(name string, c client.Client) MobileAuthenticationService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mindbox"
	}
	return &mobileAuthenticationService{
		c:    c,
		name: name,
	}
}

func (c *mobileAuthenticationService) Creation(ctx context.Context, in *ParamsCreation, opts ...client.CallOption) (*ResponseCreation, error) {
	req := c.c.NewRequest(c.name, "MobileAuthentication.Creation", in)
	out := new(ResponseCreation)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileAuthenticationService) Authorization(ctx context.Context, in *ParamsAuthorization, opts ...client.CallOption) (*ResponseAuthorization, error) {
	req := c.c.NewRequest(c.name, "MobileAuthentication.Authorization", in)
	out := new(ResponseAuthorization)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileAuthenticationService) Registration(ctx context.Context, in *ParamsRegistration, opts ...client.CallOption) (*ResponseRegistration, error) {
	req := c.c.NewRequest(c.name, "MobileAuthentication.Registration", in)
	out := new(ResponseRegistration)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileAuthenticationService) Code(ctx context.Context, in *ParamsCode, opts ...client.CallOption) (*ResponseCode, error) {
	req := c.c.NewRequest(c.name, "MobileAuthentication.Code", in)
	out := new(ResponseCode)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mobileAuthenticationService) CheckCode(ctx context.Context, in *ParamsCheckCode, opts ...client.CallOption) (*ResponseCheckCode, error) {
	req := c.c.NewRequest(c.name, "MobileAuthentication.CheckCode", in)
	out := new(ResponseCheckCode)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MobileAuthentication service

type MobileAuthenticationHandler interface {
	Creation(context.Context, *ParamsCreation, *ResponseCreation) error
	Authorization(context.Context, *ParamsAuthorization, *ResponseAuthorization) error
	Registration(context.Context, *ParamsRegistration, *ResponseRegistration) error
	Code(context.Context, *ParamsCode, *ResponseCode) error
	CheckCode(context.Context, *ParamsCheckCode, *ResponseCheckCode) error
}

func RegisterMobileAuthenticationHandler(s server.Server, hdlr MobileAuthenticationHandler, opts ...server.HandlerOption) error {
	type mobileAuthentication interface {
		Creation(ctx context.Context, in *ParamsCreation, out *ResponseCreation) error
		Authorization(ctx context.Context, in *ParamsAuthorization, out *ResponseAuthorization) error
		Registration(ctx context.Context, in *ParamsRegistration, out *ResponseRegistration) error
		Code(ctx context.Context, in *ParamsCode, out *ResponseCode) error
		CheckCode(ctx context.Context, in *ParamsCheckCode, out *ResponseCheckCode) error
	}
	type MobileAuthentication struct {
		mobileAuthentication
	}
	h := &mobileAuthenticationHandler{hdlr}
	return s.Handle(s.NewHandler(&MobileAuthentication{h}, opts...))
}

type mobileAuthenticationHandler struct {
	MobileAuthenticationHandler
}

func (h *mobileAuthenticationHandler) Creation(ctx context.Context, in *ParamsCreation, out *ResponseCreation) error {
	return h.MobileAuthenticationHandler.Creation(ctx, in, out)
}

func (h *mobileAuthenticationHandler) Authorization(ctx context.Context, in *ParamsAuthorization, out *ResponseAuthorization) error {
	return h.MobileAuthenticationHandler.Authorization(ctx, in, out)
}

func (h *mobileAuthenticationHandler) Registration(ctx context.Context, in *ParamsRegistration, out *ResponseRegistration) error {
	return h.MobileAuthenticationHandler.Registration(ctx, in, out)
}

func (h *mobileAuthenticationHandler) Code(ctx context.Context, in *ParamsCode, out *ResponseCode) error {
	return h.MobileAuthenticationHandler.Code(ctx, in, out)
}

func (h *mobileAuthenticationHandler) CheckCode(ctx context.Context, in *ParamsCheckCode, out *ResponseCheckCode) error {
	return h.MobileAuthenticationHandler.CheckCode(ctx, in, out)
}
