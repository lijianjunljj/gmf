package server

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"gmf/src/common/config"
	"golang.org/x/sync/errgroup"
)

type ServerIFace interface {
	BeforeRun(*config.Config) micro.Service
	Run(*config.Config) error
	GetName() string
	GetServiceName() string
	ServiceClient(micro.Option) interface{}
	ErrGroup() *errgroup.Group
	EtcdReg(*config.EtcdOptions) registry.Registry
}
