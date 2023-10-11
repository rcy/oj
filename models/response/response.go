package response

import (
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
