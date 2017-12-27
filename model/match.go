package model

import (
	"log"
	"strings"
	"time"

	"github.com/go-pg/pg"
)

// update match info for two groups
func updateMatch(tx *pg.Tx, userID1, userID2 string) error {
	_, err := tx.Exec(`UPDATE relations SET 
		match_state = 'matched',
		update_time = ?
		where user_id=? and wipe_user_id = ? and state='liked'`,
		time.Now(), userID1, userID2)
	if err != nil {
		log.Panicf("wrong in update relations user_id:%v wipe_user_id:%v err:%v", userID1, userID2, err)
		return err
	}
	return nil
}

//UpdateUserMatch when two user like each other, change their state into matched
func UpdateUserMatch(userID1, userID2 string) error {

	// order userid for preventing deadlock
	if strings.Compare(userID1, userID2) > 0 {
		userID1, userID2 = userID2, userID1
	}

	err := DB.RunInTransaction(func(tx *pg.Tx) error {

		//TODO merge into one sql
		relation1, err := queryRelation(tx, userID1, userID2)
		if err != nil {
			return err
		}
		relation2, err := queryRelation(tx, userID2, userID1)
		if err != nil {
			return err
		}

		// if any(relations) is nil , return nil
		if relation1 == nil || relation2 == nil {
			return nil
		}
		//TODO
		if relation1.State == LikedState && relation2.State == LikedState {
			err = updateMatch(tx, userID1, userID2)
			if err != nil {
				return err
			}
			err = updateMatch(tx, userID2, userID1)
			if err != nil {
				return err
			}

			return nil
		}
		return nil
	})

	return err
}
