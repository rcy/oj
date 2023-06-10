package messages

import (
	"log"
	"time"
)

type Message struct {
	Username  string
	Body      string
	CreatedAt time.Time
}

var messages = []Message{
	{
		Username: "bob", Body: "the body", CreatedAt: time.Now(),
	},
}

func Fetch() ([]Message, error) {
	return messages, nil
}

func Create(body string, username string) (Message, error) {
	message := Message{
		Body:      body,
		Username:  username,
		CreatedAt: time.Now(),
	}

	messages = append(messages, message)

	log.Printf("messages=%v", messages)

	return message, nil
}
