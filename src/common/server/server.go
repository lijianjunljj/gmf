package server

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"gmf/src/common/config"
	"gmf/src/servers/user/services"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	ServerIFace
	G               *errgroup.Group
	ServiceCallFunc func(microService micro.Service)
	Config          *config.Config
	Name            string
	ServiceName     string
}

func (s *Server) BeforeRun(config *config.Config) micro.Service {
	s.Config = config
	s.Config.InitService(s.Name)
	fmt.Println("config1111:", s.Config)
	// etcd注册件
	etcdReg := s.EtcdReg(s.Config.Etcd)
	// 得到一个微服务实例
	microService := micro.NewService(
		micro.Name(s.Config.Service.ServiceName), // 微服务名字
		micro.Address(s.Config.Service.Address),
		micro.Registry(etcdReg), // etcd注册件
	)
	// 结构命令行参数，初始化
	microService.Init()

	return microService
}

func (s *Server) Run(config *config.Config) error {
	microService := s.BeforeRun(config)
	// 服务注册

	if s.ServiceCallFunc != nil {
		s.ServiceCallFunc(microService)
	}

	// 启动微服务
	_ = microService.Run()
	return nil
}
func (s *Server) ErrGroup() *errgroup.Group {
	return s.G
}
func (s *Server) GetConfig() *config.Config {
	return s.Config
}

func (s *Server) ServiceClient(option micro.Option) interface{} {
	fmt.Println("ServiceClient...", s)
	fmt.Printf("s type %T \n", s)
	serverName := s.Name
	fmt.Println("serverName:", serverName)
	serviceName := s.ServiceName
	fmt.Println("serviceName:", serviceName)
	userMicroService := micro.NewService(
		micro.Name(serviceName+".client"),
		option, //micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := services.NewUserService(s.ServiceName, userMicroService.Client())
	return userService
}

func (s *Server) EtcdReg(options *config.EtcdOptions) registry.Registry {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(options.RegistryAddr),
	)
	return etcdReg
}

var sm Manager

func Init() *Manager {
	sm.Init()
	return &sm
}

func StartAll(config *config.Config) {
	sm.RunAll(config)
}
func Start(name string, config *config.Config) {
	err := sm.Run(name, config)
	if err != nil {
		fmt.Println(err.Error())
	}

}
