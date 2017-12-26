package main

import (
	"fmt"

	"github.com/go-pg/migrations"
)

func init() {
	fmt.Println("inside init 0")
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating table tantan_onboard...")
		_, err := db.Exec(`CREATE TABLE tantan_onboard(
			user_id varchar(100),
			
			)`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping table my_table...")
		_, err := db.Exec(`DROP TABLE tantan_onboard`)
		return err
	})
}
