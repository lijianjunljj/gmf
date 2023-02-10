package main

import (
	"fmt"
	"gmf/cmd"
	"gmf/src/common/server"
	"gmf/src/gateway"
	"gmf/src/user"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	server.Init().Register(user.NewServer(&g), gateway.NewServer(&g))
	cmd.Execute()
	if err := g.Wait(); err != nil {
		fmt.Printf("错误退出:%s\n", err)
	}
}
