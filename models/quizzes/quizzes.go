package quizzes

import (
	"context"
	"oj/db"
	"oj/models/question"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
)

type Quiz struct {
	ID          int64
	CreatedAt   time.Time `db:"created_at"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Published   bool      `db:"published"`
}

func (q *Quiz) Save(ctx context.Context, db *sqlx.DB) (*Quiz, error) {
	var (
		result Quiz
		err    error
	)
	if q.ID == 0 {
		err = db.Get(&result, `insert into quizzes(name, description) values(?,?,?) returning *`, q.Name, q.Description)
	} else {
		err = db.Get(&result, `update quizzes set name = ?, description = ? where id = ? returning *`, q.Name, q.Description, q.ID)
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func FindAll() ([]Quiz, error) {
	var result []Quiz
	err := db.DB.Select(&result, "select * from quizzes order by created_at desc")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindAllPublished() ([]Quiz, error) {
	var result []Quiz
	err := db.DB.Select(&result, "select * from quizzes where published = true order by created_at desc")
	if err != nil {
		return nil, err
	}
	return result, nil
}

func FindByID(id int64) (*Quiz, error) {
	var result Quiz

	err := db.DB.Get(&result, "select * from quizzes where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindByStringID(stringID string) (*Quiz, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return nil, err
	}
	return FindByID(int64(id))
}

func (q *Quiz) FindQuestions() ([]question.Question, error) {
	var result []question.Question
	query := "select id, created_at, quiz_id, text, answer from questions where quiz_id = ?"
	err := db.DB.Select(&result, query, q.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
