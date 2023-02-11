package user

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"gmf/src/common/config"
	"gmf/src/common/server"
	"gmf/src/servers/user/core"
	"gmf/src/servers/user/services"
	"golang.org/x/sync/errgroup"
)

func NewServer(g *errgroup.Group) *Server {
	s := &Server{}
	s.G = g
	return s
}

type Server struct {
	server.Server
}

func (s Server) Name() string {
	return "userServer"
}

func (s Server) Run(config *config.Config) error {
	fmt.Println("config1111:", config)
	// etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"), // 微服务名字
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()
	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	_ = microService.Run()
	return nil
}