package main

import (
	"fmt"

	"github.com/p1cn/onboard/liuzhenzhong/config"
	"github.com/p1cn/onboard/liuzhenzhong/server"
)

func main() {
	config, err := config.NewConfig("conf/config.yml")
	if err != nil {
		fmt.Printf("error in load config %v\n", err)
		return
	}
	port := ":8080"
	server.StartServer(port, config)
	fmt.Println("hello world")
}
