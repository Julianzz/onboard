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
		`CREATE TABLE IF NOT EXISTS users(
			user_id varchar(64),
			name varchar(128),
			type varchar(16),
			create_time timestamp default now(),  
			update_time timestamp default now(),
			PRIMARY KEY (user_id)
		) `,
		`CREATE TABLE IF NOT EXISTS relations(
			user_id varchar(64), 
			wipe_user_id varchar(64), 
			type varchar(16), 
			state varchar(16), 
			create_time timestamp default now(), 
			update_time timestamp default now(),
			PRIMARY KEY (user_id,wipe_user_id)
			)`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}
