// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: users.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	users.proto

It has these top-level messages:
	User
	RegisterRequest
	LoginRequest
	LoginResponse
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SUsers service

type SUsersService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*User, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type sUsersService struct {
	c           client.Client
	serviceName string
}

func SUsersServiceClient(serviceName string, c client.Client) SUsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "protocol"
	}
	return &sUsersService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *sUsersService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.serviceName, "SUsers.Register", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sUsersService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.serviceName, "SUsers.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SUsers service

type SUsersHandler interface {
	Register(context.Context, *RegisterRequest, *User) error
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterSUsersHandler(s server.Server, hdlr SUsersHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SUsers{hdlr}, opts...))
}

type SUsers struct {
	SUsersHandler
}

func (h *SUsers) Register(ctx context.Context, in *RegisterRequest, out *User) error {
	return h.SUsersHandler.Register(ctx, in, out)
}

func (h *SUsers) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.SUsersHandler.Login(ctx, in, out)
}
