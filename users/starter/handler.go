package starter

import (
	"go-micro-demo/protocol"
	"github.com/micro/go-micro/server"
	"go-micro-demo/users/controller"
)

func RegisterHandler(s server.Server)  {

	protocol.RegisterSUsersHandler(s,new(controller.UserHandler))

}
