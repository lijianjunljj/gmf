package server

import (
	"gmf/src/common/config"
	"golang.org/x/sync/errgroup"
)

type ServerIFace interface {
	Run(*config.Config) error
	Name() string
	ErrGroup() *errgroup.Group
}
