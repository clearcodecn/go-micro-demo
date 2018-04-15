package common

import "github.com/hashicorp/consul/api"

var ConsulClient *api.Client

func StartConsulClient() error {
	var err error
	ConsulClient , err = api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	return nil
}

