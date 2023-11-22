package room

import (
	"database/sql"
	"fmt"
	"math"
	"oj/db"
	"time"
)

type Room struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	Key       string
}

func FindOrCreateByUserIDs(id1, id2 int64) (*Room, error) {
	var room Room

	key := makeRoomKey(id1, id2)

	err := db.DB.Get(&room, "select * from rooms where key = ?", key)
	if err != nil {
		if err == sql.ErrNoRows {
			return create(id1, id2)
		}
		return nil, err
	}
	return &room, nil
}

func create(userID1, userID2 int64) (*Room, error) {
	key := makeRoomKey(userID1, userID2)

	tx, err := db.DB.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	result, err := tx.Exec("insert into rooms(key) values(?)", key)
	if err != nil {
		return nil, err
	}
	roomID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// add users to room
	_, err = tx.Exec("insert into room_users(room_id, user_id) values(?, ?)", roomID, userID1)
	if err != nil {
		return nil, err
	}

	if userID1 != userID2 {
		_, err = tx.Exec("insert into room_users(room_id, user_id) values(?, ?)", roomID, userID2)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return findById(roomID)
}

func findById(id int64) (*Room, error) {
	var room Room

	err := db.DB.Get(&room, "select * from rooms where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

// Generate a string to be used as a roomKey given 2 user IDs
func makeRoomKey(id1, id2 int64) string {
	min := int64(math.Min(float64(id1), float64(id2)))
	max := int64(math.Max(float64(id1), float64(id2)))

	return fmt.Sprintf("dm-%d-%d", min, max)
}
