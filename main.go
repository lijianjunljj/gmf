package main

import (
	"flag"
	"fmt"
	"gmf/src/common"
	"gmf/src/user"
	"golang.org/x/sync/errgroup"
)

var sm common.ServerManager
var configFile = flag.String("f", "./config/development.yaml", "配置文件")
var g errgroup.Group

func main() {
	fmt.Println(configFile)
	sm.Init()
	sm.Register(user.Server{}).RunAll(&g)
	if err := g.Wait(); err != nil {
		fmt.Printf("错误退出:%s\n", err)
	}
}
