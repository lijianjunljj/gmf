package main

import (
	"fmt"
	"gmf/src/common"
	config "gmf/src/common/config"
	"gmf/src/user"
	"golang.org/x/sync/errgroup"
)

var sm common.ServerManager
var g errgroup.Group

func Start(conf *config.MysqlOption) {
	//fmt.Println(configFile)
	sm.Init()
	sm.Register(user.Server{}).RunAll(&g)
	if err := g.Wait(); err != nil {
		fmt.Printf("错误退出:%s\n", err)
	}
}
