package server

import (
	"fmt"
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
func StartAll() {
	//fmt.Println(configFile)
	sm.RunAll()
}
func Start(name string) {
	err := sm.Run(name)

	if err != nil {
		fmt.Println(err.Error())
	}

}
