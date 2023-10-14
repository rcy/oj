package response

import (
	"oj/db"
	"time"
)

type Response struct {
	ID         int64
	CreatedAt  time.Time `db:"created_at"`
	QuizID     int64     `db:"quiz_id"`
	UserID     int64     `db:"user_id"`
	AttemptID  int64     `db:"attempt_id"`
	QuestionID int64     `db:"question_id"`
	Text       string    `db:"text"`
}

func FindByAttemptID(attemptID int64) ([]Response, error) {
	query := `select * from responses where attempt_id = ?`
	var responses []Response
	err := db.DB.Select(&responses, query, attemptID)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
