// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package api

import (
	"database/sql"
	"time"
)

type Attempt struct {
	ID        int64
	CreatedAt time.Time
	QuizID    int64
	UserID    int64
}

type Code struct {
	ID        int64
	CreatedAt time.Time
	Code      interface{}
	Nonce     interface{}
	Email     string
}

type Delivery struct {
	ID          int64
	CreatedAt   time.Time
	MessageID   int64
	RoomID      int64
	RecipientID int64
	SenderID    int64
	SentAt      sql.NullTime
}

type Friend struct {
	ID        int64
	CreatedAt time.Time
	AID       int64
	BID       int64
	BRole     string
}

type Gradient struct {
	ID        int64
	CreatedAt string
	UserID    int64
	Gradient  []byte
}

type Image struct {
	ID        int64
	CreatedAt time.Time
	Url       interface{}
	UserID    interface{}
}

type KidsCode struct {
	ID        int64
	CreatedAt time.Time
	Code      interface{}
	Nonce     interface{}
	UserID    int64
}

type KidsParent struct {
	ID        int64
	CreatedAt time.Time
	KidID     int64
	ParentID  int64
}

type Message struct {
	ID        int64
	CreatedAt time.Time
	SenderID  int64
	RoomID    string
	Body      string
}

type MigrationVersion struct {
	Version sql.NullInt64
}

type Question struct {
	ID        int64
	CreatedAt string
	QuizID    int64
	Text      string
	Answer    string
}

type Quiz struct {
	ID          int64
	CreatedAt   time.Time
	Name        interface{}
	Description interface{}
	Published   sql.NullBool
}

type Response struct {
	ID         int64
	CreatedAt  time.Time
	QuizID     interface{}
	UserID     interface{}
	AttemptID  interface{}
	QuestionID interface{}
	Text       interface{}
}

type Room struct {
	ID        int64
	CreatedAt time.Time
	Key       string
}

type RoomUser struct {
	ID        int64
	CreatedAt time.Time
	RoomID    int64
	UserID    int64
}

type Session struct {
	ID     int64
	UserID int64
	Key    interface{}
}

type User struct {
	ID           int64
	CreatedAt    time.Time
	Username     string
	Email        sql.NullString
	AvatarUrl    interface{}
	IsParent     bool
	Bio          string
	BecomeUserID sql.NullInt64
	Admin        bool
}
