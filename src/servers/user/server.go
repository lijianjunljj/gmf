package user

import (
	"github.com/micro/go-micro/v2"
	"gmf/src/common/server"
	"gmf/src/servers/user/core"
	"gmf/src/servers/user/services"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group) *Server {
	s := new(Server)
	s.Server.Name = "userServer"
	s.Server.ServiceName = "rpcUserService"
	s.G = g
	return s
}

type Server struct {
	server.Server
}

func (s *Server) GetName() string {
	return s.Server.Name
}
func (s *Server) GetServiceName() string {
	return s.Server.ServiceName
}
func (s *Server) RegisterServiceHandlerFunc() error {
	s.Server.ServiceCallFunc = func(microService micro.Service) {
		_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	}
	return nil
}
