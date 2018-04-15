package service

import (
	"go-micro-demo/protocol"
	"go-micro-demo/users/model"
)

func Register(request *protocol.RegisterRequest, rsp *protocol.User) (error){
	u := &model.User{
		Username:request.GetUsername(),
		Password:request.GetPassword(),
	}
	if err := model.UserTable().Create(u).Error ; err != nil {
		return err
	}

	rsp.Id = uint32(u.ID)
	rsp.Username = u.Username
	rsp.Password = u.Password
	return nil
}

func Login()  {

}
