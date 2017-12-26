package model

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

// User struture info
type User struct {
	tableName struct{} `sql:"users"`

	UserID string `sql:"user_id"`
	Name   string `sql:"name"`
	Type   string `sql:"type"`

	// create and update time
	CreateTime time.Time `sql:"create_time,default:now()"`
	UpdateTime time.Time `sql:"update_time,default:now()"`
}

// CreateNewUser create new user in db
func CreateNewUser(userID, name, t string) error {
	user := User{
		UserID:     userID,
		Name:       name,
		Type:       t,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	err := DB.Insert(&user)
	if err != nil {
		fmt.Printf("error in create user %v\n", err)
		return err
	}

	return nil
}

// GetUsers list users
// TODO add paging function
func GetUsers(limit int) ([]*User, error) {
	var users = make([]*User, 0)
	// ignore no rows error, return empty relation
	_, err := DB.Query(&users, `select * from users`)
	if err != nil && err != pg.ErrNoRows {
		log.Printf("error in query users error:%v", err)
		return users, err
	}

	return users, nil
}

// GetUserByID get user info by user_id
func GetUserByID(userID string) (*User, error) {
	var user User
	_, err := DB.QueryOne(&user, `select * from users where user_id =?`, userID)
	if err != nil {
		fmt.Printf("error in query user %v err:%v \n", userID, err)
		return nil, err
	}
	return &user, nil
}
