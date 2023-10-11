package question

import (
	"oj/db"
	"strconv"
	"time"
)

type Question struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	QuizID    int64     `db:"quiz_id"`
	Text      string    `db:"text"`
	Answer    string    `db:"answer"`
}

func FindByID(id int64) (*Question, error) {
	var result Question

	err := db.DB.Get(&result, "select * from questions where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindByStringID(stringID string) (*Question, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return nil, err
	}
	return FindByID(int64(id))
}
