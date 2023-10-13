package question

import (
	"context"
	"oj/db"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type Question struct {
	ID        int64
	CreatedAt string `db:"created_at"`
	QuizID    int64  `db:"quiz_id"`
	Text      string `db:"text"`
	Answer    string `db:"answer"`
}

func (q *Question) Save(ctx context.Context, db *sqlx.DB) (*Question, error) {
	var (
		result Question
		err    error
	)
	if q.ID == 0 {
		err = db.Get(&result, `insert into questions(quiz_id, text, answer) values(?,?,?) returning *`, q.QuizID, q.Text, q.Answer)
	} else {
		err = db.Get(&result, `update questions set text = ?, answer = ? where id = ? returning *`, q.Text, q.Answer, q.ID)
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
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
