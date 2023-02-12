package servers

import (
	"gmf/src/servers/user"
	"golang.org/x/sync/errgroup"
)

func RegisterServers(g *errgroup.Group) {
	smp.Register(user.NewServer(g))
}
