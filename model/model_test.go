package model

import (
	"testing"

	"github.com/go-pg/pg"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&DBTest{})

type DBTest struct {
	db *pg.DB
}

func (t *DBTest) SetUpTest(c *C) {
	t.db = pg.Connect(&pg.Options{
		User:     "liuzhenzhong",
		Database: "test",
	})

	DB = t.db
	queries := []string{
		`DROP TABLE IF EXISTS relations`,
		`DROP TABLE IF EXISTS users`,
		`CREATE TABLE IF NOT EXISTS users(
			user_id varchar(64),
			name varchar(128),
			type varchar(16),
			create_time timestamp default now(),  
			update_time timestamp default now()
		) `,
		`CREATE TABLE IF NOT EXISTS relations(
			user_id varchar(64), 
			wipe_user_id varchar(64), 
			type varchar(16), 
			state varchar(16), 
			create_time timestamp default now(), 
			update_time timestamp default now()
			)`,
	}

	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			panic(err)
		}
	}
}

func (t *DBTest) TearDownTest(c *C) {
	c.Assert(t.db.Close(), IsNil)
}
