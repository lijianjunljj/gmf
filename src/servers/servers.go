package servers

import (
	"gmf/src/common/server"
	"gmf/src/gateway"
	"gmf/src/servers/user"
	"golang.org/x/sync/errgroup"
)

var smp *server.Manager

func init() {
	smp = server.Init()
}
func Register(g *errgroup.Group) {
	smp.Register(user.NewServer(g))
	smp.Register(gateway.NewServer(g))
}
