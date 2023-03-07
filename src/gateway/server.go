package gateway

import (
	"fmt"
	"github.com/lijianjunljj/gmfcommon/config"
	"github.com/lijianjunljj/gmfcommon/server"
	"github.com/micro/go-micro/v2/web"
	web2 "gmf/src/gateway/web"
	servers "gmf/src/servers"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

var gatewayServer *Server

var once sync.Once

func Register(g *errgroup.Group) *Server {
	if gatewayServer == nil {
		once.Do(func() {
			gatewayServer = NewServer(g)
		})
	}
	return gatewayServer
}

func Start(cfg *config.Config) {
	gatewayServer.Run(cfg)
}

func NewServer(g *errgroup.Group) *Server {
	s := &Server{}
	s.G = g
	return s
}

type Server struct {
	server.Server
}

func (s Server) Name() string {
	return "gateway"
}
func (s Server) Run(config *config.Config) error {
	config.InitService(s.Name())
	fmt.Println("config", config)
	fmt.Println("config.Etcd.RegistryAddr:", config.Etcd.RegistryAddr)
	etcdReg := s.EtcdReg(config.Etcd)
	services := servers.Clients()
	fmt.Println("service3:", services)
	//创建微服务实例，使用gin暴露http接口并注册到etcd
	server := web.NewService(
		web.Name("httpService"),
		web.Address(config.Web.Addr),
		//将服务调用实例使用gin处理
		web.Handler(web2.NewRouter(services)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": config.Web.Protocol}),
	)
	//接收命令行参数
	_ = server.Init()
	_ = server.Run()
	return nil
}
