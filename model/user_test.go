package model

import (
	. "gopkg.in/check.v1"
)

func (t *DBTest) TestCreateUser(c *C) {

	users, err := GetUsers(-1)
	c.Assert(err, IsNil)
	c.Assert(len(users), Equals, 0)

	userID := "test_user_id"
	err = CreateNewUser(userID, "liu", "user")
	c.Assert(err, IsNil)

	users, err = GetUsers(-1)
	c.Assert(err, IsNil)
	c.Assert(len(users), Equals, 1)

	user, err := GetUserByID(userID)
	c.Assert(err, IsNil)
	c.Assert(*user, Equals, *users[0])

}
