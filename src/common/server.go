package common

import (
	"golang.org/x/sync/errgroup"
)

type ServerIFace interface {
	Run() error
	Name() string
}

type ServerManager struct {
	servers map[string]ServerIFace
}

func (sm *ServerManager) Init() *ServerManager {
	sm.servers = make(map[string]ServerIFace)
	return sm
}

func (sm *ServerManager) Register(servers ...ServerIFace) *ServerManager {
	for _, server := range servers {
		serverName := server.Name()
		if _, ok := sm.servers[serverName]; !ok {
			sm.servers[serverName] = server
		}
	}
	return sm

}

func (sm *ServerManager) RunAll(g *errgroup.Group) {
	for _, s := range sm.servers {

		g.Go(func() error {
			return s.Run()
		})
	}
}
