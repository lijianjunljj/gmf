package server

import (
	"errors"
	"fmt"
	"github.com/micro/go-micro/v2"
	"gmf/src/common/config"
)

type Manager struct {
	servers map[string]ServerIFace
}

func (sm *Manager) Init() *Manager {
	sm.servers = make(map[string]ServerIFace)
	return sm
}

func (sm *Manager) Client(serverName string, option micro.Option) interface{} {
	for _, server := range sm.servers {
		fmt.Println("Client", server.GetName(), serverName)
		if server.GetName() == serverName {
			fmt.Println("ServiceName", server.GetServiceName())
			fmt.Println("server", server)
			fmt.Printf("server type %T\n", server)
			return server.ServiceClient(option)
		}
	}
	return nil
}

func (sm *Manager) Clients(option micro.Option) []interface{} {
	var clients []interface{}
	for _, server := range sm.servers {
		client := server.ServiceClient(option)
		clients = append(clients, client)
	}
	return clients
}

func (sm *Manager) Register(servers ...ServerIFace) *Manager {
	for _, server := range servers {
		serverName := server.GetName()
		serviceName := server.GetServiceName()
		fmt.Println("serverName:", serverName)
		fmt.Println("serviceName:", serviceName)
		fmt.Println("server:", server)
		if _, ok := sm.servers[serverName]; !ok {
			sm.servers[serverName] = server
			fmt.Println("sm.servers[serverName]:", sm.servers[serverName].GetServiceName())
		}
	}
	return sm
}
func (sm *Manager) Run(name string, config *config.Config) error {
	hasFound := false
	for _, s := range sm.servers {
		if s.GetName() == name {
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
