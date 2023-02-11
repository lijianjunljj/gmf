package servers

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"gmf/src/common/server"
	"gmf/src/servers/user"
	"golang.org/x/sync/errgroup"
)

var smp *server.Manager

func init() {
	fmt.Println("servers init...........")
	smp = server.Init()
}
func Register(g *errgroup.Group) {
	fmt.Println("Register...........")
	smp.Register(user.NewServer(g))
}

func Client(serverName string, option micro.Option) interface{} {
	return smp.Client(serverName, option)
}
func Clients(option micro.Option) []interface{} {
	return smp.Clients(option)
}
