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

func (t *DBTest) TestUpdateBatchRelation(c *C) {

	mockDatas := [][]string{
		[]string{"test1", "test2", "liked", "liked"},
		[]string{"test2", "test3", "liked", "disliked"},
		[]string{"test3", "test4", "disliked", "disliked"},
	}

	for _, relation := range mockDatas {
		userID1, userID2, state1, state2 := relation[0], relation[1], relation[2], relation[3]
		err := CreateUserRelation(userID1, userID2, "relation", state1)
		c.Assert(err, IsNil)

		rel, err := GetRelationsByUserIDs(userID1, userID2)
		c.Assert(rel, NotNil)
		c.Assert(err, IsNil)
		c.Assert(rel.State, Equals, state1)
		c.Assert(rel.WipeUserID, Equals, userID2)

		err = CreateUserRelation(userID2, userID1, "relation", state2)
		c.Assert(err, IsNil)

		matchedState := UnMatchedState
		if state1 == state2 && state1 == "liked" {
			matchedState = MatchedState
		}
		rel, err = GetRelationsByUserIDs(userID2, userID1)
		c.Assert(rel, NotNil)
		c.Assert(err, IsNil)
		c.Assert(rel.State, Equals, state2)
		c.Assert(rel.MatchState, Equals, matchedState)
		c.Assert(rel.WipeUserID, Equals, userID1)

		rel, err = GetRelationsByUserIDs(userID1, userID2)
		c.Assert(rel, NotNil)
		c.Assert(err, IsNil)
		c.Assert(rel.State, Equals, state1)
		c.Assert(rel.MatchState, Equals, matchedState)
		c.Assert(rel.WipeUserID, Equals, userID2)
		fmt.Println("inside test update relation")
	}

	userID := mockDatas[1][0]
	relations, err := GetRelationsByUserID(userID)
	c.Assert(err, IsNil)
	c.Assert(len(relations), Equals, 2)
}
