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
		_, err := t.DB.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func (t *BaseTest) TearDownTest(c *C) {
	c.Assert(t.DB.Close(), IsNil)
}
