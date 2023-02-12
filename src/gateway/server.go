package gateway

import (
	"fmt"
	"github.com/micro/go-micro/v2/web"
	"gmf/src/common/config"
	"gmf/src/common/server"
	web2 "gmf/src/gateway/web"
	servers "gmf/src/servers"
	"golang.org/x/sync/errgroup"
	"sync"
	"time"
)

var gatewayServer *Server

//var lock sync.Mutex
//var inited uint32 = 0
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

//
//func GetInstance(g *errgroup.Group) *Server {
//	fmt.Println("GetInstance1")
//	if atomic.LoadUint32(&inited) > 0 {
//		return gatewayServer
//	}
//	fmt.Println("GetInstance2")
//	lock.Lock()
//	defer lock.Unlock()
//	if gatewayServer == nil {
//		fmt.Println("init gatewayServer.....")
//		gatewayServer = NewServer(g)
//		atomic.StoreUint32(&inited, 1)
//	}
//	return gatewayServer
//}

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
