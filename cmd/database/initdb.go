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
			id serial primary key ,
			user_id varchar(64) not null unique,
			name varchar(128) not null,
			type varchar(16) not null,
			create_time timestamp default now(),  
			update_time timestamp default now()
		) `,
		`CREATE TABLE IF NOT EXISTS relations(
			id serial primary key,
			user_id varchar(64) not null , 
			wipe_user_id varchar(64) not null, 
			type varchar(16) not null, 
			state varchar(16) not null, 
			match_state varchar(16) not null,
			create_time timestamp default now(), 
			update_time timestamp default now(),
			constraint user_wipe_user_uniq unique(user_id,wipe_user_id)
		)`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}
