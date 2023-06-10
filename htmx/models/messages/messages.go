package messages

import (
	"log"
	"time"
)

type Sender struct {
	Username string
}

type Message struct {
	Sender    Sender
	Body      string
	CreatedAt time.Time
}

var messages = []Message{
	// {
	// 	Sender: Sender{Username: "bob"}, Body: "the body", CreatedAt: time.Now(),
	// },
}

func Fetch() ([]Message, error) {
	return messages, nil
}

func Create(body string, username string) (Message, error) {
	message := Message{
		Body:      body,
		Sender:    Sender{Username: username},
		CreatedAt: time.Now(),
	}

	messages = append(messages, message)

	log.Printf("messages=%v", messages)

	return message, nil
}
