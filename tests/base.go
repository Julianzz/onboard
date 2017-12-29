package tests

import "github.com/go-pg/pg"
import . "gopkg.in/check.v1"

type BaseTest struct {
	DB *pg.DB
}

func (t *BaseTest) SetUpTest(c *C) {

	t.DB = pg.Connect(&pg.Options{
		User:     "liuzhenzhong",
		Database: "test",
	})

	queries := []string{
		`DROP TABLE IF EXISTS relations`,
		`DROP TABLE IF EXISTS users`,
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
		_, err := t.DB.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func (t *BaseTest) TearDownTest(c *C) {
	c.Assert(t.DB.Close(), IsNil)
}
