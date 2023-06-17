package messages

import (
	"oj/db"
	"time"
)

type Message struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	RoomID    string    `db:"room_id"`
	Sender    string
	Body      string
}

func Fetch(roomID string) (messages []Message, err error) {
	err = db.DB.Select(&messages, "select * from messages where room_id = ? order by created_at asc", roomID)
	if err != nil {
		return []Message{}, err
	}
	return messages, nil
}

func Create(roomID string, body string, username string) (Message, error) {
	var message Message

	result, err := db.DB.Exec("insert into messages(room_id, body, sender) values(?,?,?)", roomID, body, username)
	if err != nil {
		return Message{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Message{}, err
	}
	err = db.DB.Get(&message, "select * from messages where id = ?", id)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}
