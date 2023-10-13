package attempts

import (
	"oj/db"
	"oj/models/question"
	"oj/models/response"
	"strconv"
	"time"
)

type Attempt struct {
	ID        int64
	CreatedAt time.Time `db:"created_at"`
	QuizID    int64     `db:"quiz_id"`
	UserID    int64     `db:"user_id"`
}

func Create(quizID int64, userID int64) (*Attempt, error) {
	var result Attempt
	err := db.DB.Get(&result, "insert into attempts(quiz_id, user_id) values(?,?) returning *", quizID, userID)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateResponse(attemptID int64, questionID int64, text string) (*response.Response, error) {
	a, err := FindByID(attemptID)
	if err != nil {
		return nil, err
	}
	q, err := question.FindByID(questionID)
	if err != nil {
		return nil, err
	}
	var result response.Response
	err = db.DB.Get(&result, `insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(?,?,?,?,?) returning *`, q.QuizID, a.UserID, a.ID, q.ID, text)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func FindByID(id int64) (*Attempt, error) {
	var result Attempt

	err := db.DB.Get(&result, "select * from attempts where id = ?", id)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindByStringID(stringID string) (*Attempt, error) {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return nil, err
	}
	return FindByID(int64(id))
}

func (a *Attempt) NextQuestion() (*question.Question, error) {
	var result question.Question

	query := `
select questions.* from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = ?)
and
  questions.quiz_id = ?
order by random()`
	err := db.DB.Get(&result, query, a.ID, a.QuizID)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *Attempt) QuestionCount() (int, error) {
	var result int
	err := db.DB.Get(&result, "select count(*) from questions where quiz_id = ?", a.QuizID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (a *Attempt) ResponseCount() (int, error) {
	var result int
	err := db.DB.Get(&result, "select count(*) from responses where attempt_id = ?", a.ID)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (a *Attempt) ResponseIDs() ([]int64, error) {
	var result []int64
	//	var responses []response.Response
	err := db.DB.Select(&result, "select id from responses where attempt_id = ?", a.ID)
	if err != nil {
		return nil, err
	}
	// var result []int64
	// for _, response := range responses {
	// 	result = append(result, response.ID)
	// }
	return result, nil
}
