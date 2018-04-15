package starter

import (
	"github.com/micro/go-micro"
	"go-micro-demo/users/model"
	"go-micro-demo/users/common"
	"github.com/go-log/log"
)

type Service struct{
	Name string
	Version string
	self micro.Service
}

func NewService() *Service {
	return &Service{
		Name: "go.micro.api.user",
		Version: "v1",
	}
}

func (s *Service) Start() error {
	service := micro.NewService(
		micro.Name(s.Name),
		micro.Version(s.Version),
		micro.AfterStart(func() error {
			err := common.StartConsulClient()
			if err != nil {
				return err
			}
			err = model.Init()
			if err != nil {
				return err
			}
			return nil
		}),
	)
	s.self = service

	RegisterHandler(s.self.Server())

	s.self.Init()

	err := s.self.Run()
	if err != nil {
		log.Log(err)
	}
	return err
}

func (s *Service) Stop()  {
	if s.self != nil {
		s.self.Server().Stop()
	}
}
