package user

import (
	"github.com/micro/go-micro/v2"
	"gmf/src/common/server"
	"gmf/src/servers/user/core"
	"gmf/src/servers/user/services"
	"gmf/src/servers/user/wrappers"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group) *Server {
	s := new(Server)
	s.Server.Name = "userServer"
	s.Server.ServiceName = "UserService"
	s.Server.ClientService = micro.NewService(
		micro.Name(s.Server.ServiceName+".client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	s.Server.ServiceCallFunc = func(microService micro.Service) {
		_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	}
	s.G = g
	return s
}

type Server struct {
	server.Server
}

func (s *Server) GetName() string {
	return s.Server.Name
}

// 用户服务调用实例

func (s *Server) GetServiceName() string {
	return s.Server.ServiceName
}
