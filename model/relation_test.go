package model

import (
	"fmt"

	. "gopkg.in/check.v1"
)

func (t *DBTest) TestRelationsFetch(c *C) {
	userID := "test_relation_userd1"

	relations, err := GetRelationsByUserID(userID)
	c.Assert(err, IsNil)
	c.Assert(len(relations), Equals, 0)

	userID2 := userID + "_match"
	err = CreateUserRelation(userID, userID2, "relation", "disliked")
	c.Assert(err, IsNil)

	relations, err = GetRelationsByUserID(userID)
	c.Assert(err, IsNil)
	c.Assert(len(relations), Equals, 1)

	c.Assert(relations[0].State, Equals, "disliked")
	c.Assert(relations[0].WipeUserID, Equals, userID2)
}

func (t *DBTest) TestUpdateRelation(c *C) {
	var userID1 = "test1"
	var userID2 = "test2"

	fmt.Println("inside test update relation")

	err := CreateUserRelation(userID1, userID2, "relation", "liked")
	c.Assert(err, IsNil)

	relations, err := GetRelationsByUserID(userID1)
	c.Assert(len(relations), Equals, 1)
	c.Assert(relations[0].State, Equals, "liked")
	c.Assert(relations[0].WipeUserID, Equals, userID2)

	err = CreateUserRelation(userID2, userID1, "relation", "liked")
	c.Assert(err, IsNil)

	relations, err = GetRelationsByUserID(userID2)
	c.Assert(len(relations), Equals, 1)
	c.Assert(relations[0].State, Equals, "matched")
	c.Assert(relations[0].WipeUserID, Equals, userID1)

}
