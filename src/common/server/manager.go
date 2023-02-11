package server

import (
	"errors"
	"fmt"
	"gmf/src/common/config"
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
func (sm *Manager) Run(name string, config *config.Config) error {
	hasFound := false
	for _, s := range sm.servers {
		if s.Name() == name {
			s.ErrGroup().Go(func() error {
				err := s.Run(config)
				if err != nil {
					fmt.Println("Run Server Error:", err)
				}
				return err
			})
			hasFound = true
		}
	}
	if !hasFound {
		return errors.New("server not found")
	}
	return nil
}
func (sm *Manager) RunAll(config *config.Config) {
	for _, s := range sm.servers {

		s.ErrGroup().Go(func() error {
			return s.Run(config)
		})
	}
}
