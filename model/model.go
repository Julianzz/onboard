package model

import (
	"github.com/go-pg/pg"
	"github.com/p1cn/onboard/liuzhenzhong/config"
)

//DB for db connection
var DB *pg.DB

// InitDB init database connection
func InitDB(config *config.Config) (*pg.DB, error) {

	db := pg.Connect(&pg.Options{
		Addr:     config.DBSetting.Host,
		User:     config.DBSetting.User,
		Password: config.DBSetting.Password,
		Database: config.DBSetting.Database,
	})
	DB = db
	return db, nil
}
