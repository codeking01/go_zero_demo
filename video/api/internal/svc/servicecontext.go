package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go_zero_demo/user/rpc/userclient"
	"go_zero_demo/video/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
