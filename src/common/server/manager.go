package server

import (
	"errors"
)

type Manager struct {
	servers map[string]ServerIFace
}

func (sm *Manager) Init() *Manager {
	sm.servers = make(map[string]ServerIFace)
	return sm
}

func (sm *Manager) Register(servers ...ServerIFace) *Manager {
	for _, server := range servers {
		serverName := server.Name()
		if _, ok := sm.servers[serverName]; !ok {
			sm.servers[serverName] = server
		}
	}
	return sm

}
func (sm *Manager) Run(name string) error {
	hasFound := false
	for _, s := range sm.servers {
		if s.Name() == name {
			s.Run()
			s.ErrGroup().Go(func() error {
				return s.Run()
			})
			hasFound = true
		}
	}
	if !hasFound {
		return errors.New("server not found")
	}
	return nil
}
func (sm *Manager) RunAll() {
	for _, s := range sm.servers {

		s.ErrGroup().Go(func() error {
			return s.Run()
		})
	}
}
