package servers

import (
	"gmf/src/servers/user"
	"golang.org/x/sync/errgroup"
)

func RegisterServers(g *errgroup.Group) {
	serverManager.Register(user.NewServer(g))
}
