package main

import (
	"flag"
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

	var port = flag.String("port", ":8080", "http listen port")
	flag.Parse()

	fmt.Printf("begin to start server on port:%v\n", *port)
	server.StartServer(*port, config)
	fmt.Println("hello world")
}
