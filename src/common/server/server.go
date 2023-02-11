package server

import (
	"fmt"
	"gmf/src/common/config"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	G *errgroup.Group
}

func (s Server) ErrGroup() *errgroup.Group {
	return s.G
}

var sm Manager

func Init() *Manager {
	sm.Init()
	return &sm
}

func StartAll(config *config.Config) {
	sm.RunAll(config)
}
func Start(name string, config *config.Config) {
	err := sm.Run(name, config)
	if err != nil {
		fmt.Println(err.Error())
	}

}
