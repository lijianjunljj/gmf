package user

import (
	"github.com/micro/go-micro/v2"
	"gmf/src/common/server"
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

//type Server struct {
//	server.Server
//}

//func (s *Server) GetName() string {
//	return s.Server.Name
//}
//
//// 用户服务调用实例
//
//func (s *Server) GetServiceName() string {
//	return s.Server.ServiceName
//}
