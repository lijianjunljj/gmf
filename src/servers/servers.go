package servers

import (
	"gmf/src/common/router"
	"gmf/src/common/server"
	"golang.org/x/sync/errgroup"
)

var smp *server.Manager

func init() {
	smp = server.Init()
}
func Register(g *errgroup.Group) {
	RegisterServers(g)
}

func Client(serverName string) interface{} {
	return smp.Client(serverName)
}

func Clients() []interface{} {
	return smp.Clients()
}
func Routers() []router.AbstractRouter {
	return smp.Routers()
}
