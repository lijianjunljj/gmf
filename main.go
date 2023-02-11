package main

import (
	"fmt"
	"gmf/cmd"
	"gmf/src/gateway"
	"gmf/src/servers"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group
	gateway.GetInstance(&g)
	servers.Register(&g)
	cmd.Execute()
	if err := g.Wait(); err != nil {
		fmt.Printf("错误退出:%s\n", err)
	}

	fmt.Println("服务退出")
}
