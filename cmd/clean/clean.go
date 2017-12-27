package main

import (
	"fmt"

	"github.com/p1cn/onboard/liuzhenzhong/config"
	"github.com/p1cn/onboard/liuzhenzhong/model"
)

func main() {

	config, err := config.NewConfig("conf/config.yml")
	if err != nil {
		fmt.Printf("error in load config %v\n", err)
		return
	}

	db, err := model.InitDB(config)
	if err != nil {
		fmt.Printf("error in initing db %v\n", err)
		return
	}

	queries := []string{
		`DROP TABLE IF EXISTS relations`,
		`DROP TABLE IF EXISTS users`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}
