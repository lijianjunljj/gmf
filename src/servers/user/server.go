package user

import (
	"github.com/lijianjunljj/gmfcommon/server"
	"github.com/micro/go-micro/v2"
	"gmf/src/servers/user/core"
	"gmf/src/servers/user/router"
	"gmf/src/servers/user/services"
	"gmf/src/servers/user/wrappers"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group) *server.Server {
	s := new(server.Server)
	s.Name = "userServer"
	s.ServiceName = "UserService"
	s.ClientService = NewClientService(s)
	s.WebRouter = new(router.Router)
	s.ServiceCallFunc = ServiceCallFunc()
	s.ServiceClient()
	s.Service = services.NewUserService(s.ServiceName, s.ClientService.Client())
	s.G = g
	return s
}

func ServiceCallFunc() func(microService micro.Service) {
	return func(microService micro.Service) {
		_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	}
}

func NewClientService(s *server.Server) micro.Service {
	return micro.NewService(
		micro.Name(s.ServiceName+".client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
}
