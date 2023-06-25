package messages

import (
	"html/template"
	"oj/db"
	"oj/md"
	"time"
)

type Message struct {
	ID              int64
	CreatedAt       time.Time `db:"created_at"`
	RoomID          int64     `db:"room_id"`
	SenderID        int64     `db:"sender_id"`
	SenderAvatarURL string    `db:"sender_avatar_url"` // joined
	Body            string
}

func (m Message) BodyHTML() template.HTML {
	return md.RenderString(m.Body)
}

func FindByRoomID(roomID int64) (messages []Message, err error) {
	err = db.DB.Select(&messages, "select * from (select * from messages where room_id = ? order by created_at desc limit 1000) order by created_at asc", roomID)
	if err != nil {
		return []Message{}, err
	}
	return messages, nil
}

func Create(roomID int, body string, senderID int64) (Message, error) {
	var message *Message

	result, err := db.DB.Exec("insert into messages(room_id, body, sender_id) values(?,?,?)", roomID, body, senderID)
	if err != nil {
		return Message{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return Message{}, err
	}
	message, err = FindById(id)
	if err != nil {
		return Message{}, err
	}

	return *message, nil
}

func FindById(id int64) (*Message, error) {
	var message Message

	err := db.DB.Get(&message, "select * from messages where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
