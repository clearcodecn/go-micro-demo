package controller

import (
	"context"
	"go-micro-demo/protocol"
	"go-micro-demo/users/service"
)

type UserHandler struct {}

func (u *UserHandler)Register(ctx context.Context, req *protocol.RegisterRequest,rsp *protocol.User) error {
	//todo 使用 govalidator 进行验证
	var err error
	err = service.Register(req, rsp)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserHandler)Login(ctx context.Context, req *protocol.LoginRequest, rsp *protocol.LoginResponse) error {
	return nil
}


