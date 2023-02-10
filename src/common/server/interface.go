package server

import "golang.org/x/sync/errgroup"

type ServerIFace interface {
	Run() error
	Name() string
	ErrGroup() *errgroup.Group
}
