package messages

import (
	"oj/db"
	"time"
)

type Message struct {
	ID        int64
	Sender    string
	Body      string
	CreatedAt time.Time `db:"created_at"`
}

func Fetch() (messages []Message, err error) {
	err = db.DB.Select(&messages, "select * from messages")
	if err != nil {
		return []Message{}, err
	}
	return messages, nil
}

func Create(body string, username string) (Message, error) {
	var message Message

	result, err := db.DB.Exec("insert into messages(body, sender) values(?,?)", body, username)
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
