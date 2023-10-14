package response

import (
	"oj/db"
	"time"
)

type response struct {
	ID         int64
	CreatedAt  time.Time `db:"created_at"`
	QuizID     int64     `db:"quiz_id"`
	UserID     int64     `db:"user_id"`
	AttemptID  int64     `db:"attempt_id"`
	QuestionID int64     `db:"question_id"`
	Text       string    `db:"text"`
}

type Response struct {
	response
	Question string `db:"question_text"`
	Answer   string `db:"question_answer"`
}

func (r *Response) IsCorrect() bool {
	return r.Text == r.Answer
}

func FindResponses(attemptID int64) ([]Response, error) {
	query := `
select
   responses.*,
   questions.answer question_answer,
   questions.text question_text
from responses
 join questions on responses.question_id = questions.id
 where attempt_id = ?`
	var responses []Response
	err := db.DB.Select(&responses, query, attemptID)
	if err != nil {
		return nil, err
	}
	return responses, nil
}
