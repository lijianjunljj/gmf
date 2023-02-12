package servers

import (
	"fmt"
	"gmf/src/common/server"
	"golang.org/x/sync/errgroup"
)

var smp *server.Manager

func init() {
	fmt.Println("servers init...........")
	smp = server.Init()
}
func Register(g *errgroup.Group) {
	fmt.Println("Register...........")
	RegisterServers(g)
}

func Client(serverName string) interface{} {
	return smp.Client(serverName)
}
func Clients() []interface{} {
	return smp.Clients()
}
