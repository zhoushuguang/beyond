package svc

import (
	"beyond/application/user/rpc/internal/config"
	"beyond/application/user/rpc/internal/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel *model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
