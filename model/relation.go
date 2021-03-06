package model

import (
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
)

const (
	// LikedState like another
	LikedState = "liked"
	// DisLikeState dislike
	DisLikeState = "disliked"
	// MatchedState match each other
	MatchedState = "matched"
	// not match
	UnMatchedState = "unmatched"
)

//Relation struture info
type Relation struct {
	//add
	tableName struct{} `sql:"relations"`

	ID         int    `sql:"id"`
	UserID     string `sql:"user_id"`
	WipeUserID string `sql:"wipe_user_id"`
	Type       string `sql:"type"`
	State      string `sql:"state"`
	MatchState string `sql:"match_state"`

	// match time
	CreateTime time.Time `sql:"create_time,default:now()"`
	UpdateTime time.Time `sql:"update_time,default:now()"`
}

func (r *Relation) String() string {
	return fmt.Sprintf("userid:%v, wipe_user_id:%v, type:%v state:%v ", r.UserID, r.WipeUserID, r.Type, r.State)
}

// CreateUserRelation update relations
func CreateUserRelation(userid1, userid2 string, t string, state string) error {
	relation := Relation{
		UserID:     userid1,
		WipeUserID: userid2,
		Type:       t,
		State:      state,
		MatchState: UnMatchedState,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	err := DB.Insert(&relation)
	if err != nil {
		log.Printf("error in insert relation into db: %v error:%v", relation, err)
		return err
	}

	// dislike will not go into math logic
	if state != LikedState {
		return nil
	}

	err = UpdateUserMatch(userid1, userid2)
	if err != nil {
		log.Printf("error in update relation match %v\n", err)
		return err
	}

	return nil
}

// GetRelationsByUserID get relation by id
func GetRelationsByUserID(userID string) ([]Relation, error) {

	var relations []Relation
	// ignore no rows error, return empty relation
	_, err := DB.Query(&relations, `select * from relations where user_id=?`, userID)
	if err != nil && err != pg.ErrNoRows {
		return relations, err
	}

	return relations, nil
}

// GetRelationsByUserIDs get relation by id
func GetRelationsByUserIDs(userID, wipeUserID string) (*Relation, error) {

	var relation Relation
	// ignore no rows error, return empty relation
	_, err := DB.QueryOne(&relation, `select * from relations where user_id=? and wipe_user_id=?`,
		userID, wipeUserID)
	if err != nil && err != pg.ErrNoRows {
		return nil, err
	}

	return &relation, nil
}

// be cautious , when user not exist, will return nil relation, but err != nil
func queryRelation(tx *pg.Tx, userID1, userID2 string) (*Relation, error) {
	var relation Relation
	_, err := tx.QueryOne(&relation, `select * from relations where user_id=? and wipe_user_id=?`, userID1, userID2)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		log.Printf("error in query user_id: %v wipe_user_id: %v error:%v \n", userID1, userID2, err)
		return nil, err
	}
	return &relation, nil
}
