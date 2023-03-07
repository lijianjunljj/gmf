package servers

import (
	"github.com/lijianjunljj/gmfcommon/router"
	"github.com/lijianjunljj/gmfcommon/server"
	"golang.org/x/sync/errgroup"
)

var serverManager *server.Manager

func init() {
	serverManager = server.Init()
}
func Register(g *errgroup.Group) {
	RegisterServers(g)
}

func Clients() []interface{} {
	return serverManager.Clients()
}
func Routers() []router.AbstractRouter {
	return serverManager.Routers()
}
