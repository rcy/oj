// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package api

import (
	"context"
	"time"
)

const allQuizzes = `-- name: AllQuizzes :many
select id, created_at, name, description, published from quizzes order by created_at desc
`

func (q *Queries) AllQuizzes(ctx context.Context) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, allQuizzes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quiz
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Published,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const attemptNextQuestion = `-- name: AttemptNextQuestion :one
select questions.id, questions.created_at, questions.quiz_id, questions.text, questions.answer from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = ?)
and
  questions.quiz_id = ?
order by random()
`

type AttemptNextQuestionParams struct {
	AttemptID interface{}
	QuizID    int64
}

func (q *Queries) AttemptNextQuestion(ctx context.Context, arg AttemptNextQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, attemptNextQuestion, arg.AttemptID, arg.QuizID)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.Text,
		&i.Answer,
	)
	return i, err
}

const attemptResponseIDs = `-- name: AttemptResponseIDs :many
select id from responses where attempt_id = ?
`

func (q *Queries) AttemptResponseIDs(ctx context.Context, attemptID interface{}) ([]int64, error) {
	rows, err := q.db.QueryContext(ctx, attemptResponseIDs, attemptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createAttempt = `-- name: CreateAttempt :one
insert into attempts(quiz_id, user_id) values(?,?) returning id, created_at, quiz_id, user_id
`

type CreateAttemptParams struct {
	QuizID int64
	UserID int64
}

func (q *Queries) CreateAttempt(ctx context.Context, arg CreateAttemptParams) (Attempt, error) {
	row := q.db.QueryRowContext(ctx, createAttempt, arg.QuizID, arg.UserID)
	var i Attempt
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.UserID,
	)
	return i, err
}

const createQuestion = `-- name: CreateQuestion :one
insert into questions(quiz_id, text, answer) values(?,?,?) returning id, created_at, quiz_id, text, answer
`

type CreateQuestionParams struct {
	QuizID int64
	Text   string
	Answer string
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion, arg.QuizID, arg.Text, arg.Answer)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.Text,
		&i.Answer,
	)
	return i, err
}

const createQuiz = `-- name: CreateQuiz :one
insert into quizzes(name,description) values(?,?) returning id, created_at, name, description, published
`

type CreateQuizParams struct {
	Name        interface{}
	Description interface{}
}

func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, createQuiz, arg.Name, arg.Description)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Published,
	)
	return i, err
}

const createResponse = `-- name: CreateResponse :one
insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(?,?,?,?,?) returning id, created_at, quiz_id, user_id, attempt_id, question_id, text
`

type CreateResponseParams struct {
	QuizID     interface{}
	UserID     interface{}
	AttemptID  interface{}
	QuestionID interface{}
	Text       interface{}
}

func (q *Queries) CreateResponse(ctx context.Context, arg CreateResponseParams) (Response, error) {
	row := q.db.QueryRowContext(ctx, createResponse,
		arg.QuizID,
		arg.UserID,
		arg.AttemptID,
		arg.QuestionID,
		arg.Text,
	)
	var i Response
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.UserID,
		&i.AttemptID,
		&i.QuestionID,
		&i.Text,
	)
	return i, err
}

const createRoom = `-- name: CreateRoom :one
insert into rooms(key) values(?) returning id, created_at, "key"
`

func (q *Queries) CreateRoom(ctx context.Context, key string) (Room, error) {
	row := q.db.QueryRowContext(ctx, createRoom, key)
	var i Room
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Key)
	return i, err
}

const createRoomUser = `-- name: CreateRoomUser :one
insert into room_users(room_id, user_id) values(?, ?) returning id, created_at, room_id, user_id
`

type CreateRoomUserParams struct {
	RoomID int64
	UserID int64
}

func (q *Queries) CreateRoomUser(ctx context.Context, arg CreateRoomUserParams) (RoomUser, error) {
	row := q.db.QueryRowContext(ctx, createRoomUser, arg.RoomID, arg.UserID)
	var i RoomUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.RoomID,
		&i.UserID,
	)
	return i, err
}

const delivery = `-- name: Delivery :one
select id, created_at, message_id, room_id, recipient_id, sender_id, sent_at from deliveries where id = ?
`

func (q *Queries) Delivery(ctx context.Context, id int64) (Delivery, error) {
	row := q.db.QueryRowContext(ctx, delivery, id)
	var i Delivery
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.MessageID,
		&i.RoomID,
		&i.RecipientID,
		&i.SenderID,
		&i.SentAt,
	)
	return i, err
}

const getAttemptByID = `-- name: GetAttemptByID :one
select id, created_at, quiz_id, user_id from attempts where id = ?
`

func (q *Queries) GetAttemptByID(ctx context.Context, id int64) (Attempt, error) {
	row := q.db.QueryRowContext(ctx, getAttemptByID, id)
	var i Attempt
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.UserID,
	)
	return i, err
}

const publishedQuizzes = `-- name: PublishedQuizzes :many
select id, created_at, name, description, published from quizzes where published = true order by created_at desc
`

func (q *Queries) PublishedQuizzes(ctx context.Context) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, publishedQuizzes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Quiz
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Name,
			&i.Description,
			&i.Published,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const question = `-- name: Question :one
select id, created_at, quiz_id, text, answer from questions where id = ?
`

func (q *Queries) Question(ctx context.Context, id int64) (Question, error) {
	row := q.db.QueryRowContext(ctx, question, id)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.Text,
		&i.Answer,
	)
	return i, err
}

const questionCount = `-- name: QuestionCount :one
select count(*) from questions where quiz_id = ?
`

func (q *Queries) QuestionCount(ctx context.Context, quizID int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, questionCount, quizID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const quiz = `-- name: Quiz :one
select id, created_at, name, description, published from quizzes where id = ?
`

func (q *Queries) Quiz(ctx context.Context, id int64) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, quiz, id)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Published,
	)
	return i, err
}

const quizQuestions = `-- name: QuizQuestions :many
select id, created_at, quiz_id, text, answer from questions where quiz_id = ?
`

func (q *Queries) QuizQuestions(ctx context.Context, quizID int64) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, quizQuestions, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.QuizID,
			&i.Text,
			&i.Answer,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const recentMessages = `-- name: RecentMessages :many
select id, created_at, sender_id, room_id, body, sender_avatar_url from (
  select m.id, m.created_at, m.sender_id, m.room_id, m.body, sender.avatar_url as sender_avatar_url
   from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = ?
   order by m.created_at desc
   limit 128
  ) t
order by created_at asc
`

type RecentMessagesRow struct {
	ID              int64
	CreatedAt       time.Time
	SenderID        int64
	RoomID          string
	Body            string
	SenderAvatarURL interface{}
}

func (q *Queries) RecentMessages(ctx context.Context, roomID string) ([]RecentMessagesRow, error) {
	rows, err := q.db.QueryContext(ctx, recentMessages, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RecentMessagesRow
	for rows.Next() {
		var i RecentMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.SenderID,
			&i.RoomID,
			&i.Body,
			&i.SenderAvatarURL,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const responseCount = `-- name: ResponseCount :one
select count(*) from responses where attempt_id = ?
`

func (q *Queries) ResponseCount(ctx context.Context, attemptID interface{}) (int64, error) {
	row := q.db.QueryRowContext(ctx, responseCount, attemptID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const responses = `-- name: Responses :many
select
   responses.id, responses.created_at, responses.quiz_id, responses.user_id, responses.attempt_id, responses.question_id, responses.text,
   questions.answer question_answer,
   questions.text question_text
from responses
 join questions on responses.question_id = questions.id
 where attempt_id = ?
order by responses.created_at
`

type ResponsesRow struct {
	ID             int64
	CreatedAt      time.Time
	QuizID         interface{}
	UserID         interface{}
	AttemptID      interface{}
	QuestionID     interface{}
	Text           interface{}
	QuestionAnswer string
	QuestionText   string
}

func (q *Queries) Responses(ctx context.Context, attemptID interface{}) ([]ResponsesRow, error) {
	rows, err := q.db.QueryContext(ctx, responses, attemptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ResponsesRow
	for rows.Next() {
		var i ResponsesRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.QuizID,
			&i.UserID,
			&i.AttemptID,
			&i.QuestionID,
			&i.Text,
			&i.QuestionAnswer,
			&i.QuestionText,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const roomByKey = `-- name: RoomByKey :one
select id, created_at, "key" from rooms where key = ?
`

func (q *Queries) RoomByKey(ctx context.Context, key string) (Room, error) {
	row := q.db.QueryRowContext(ctx, roomByKey, key)
	var i Room
	err := row.Scan(&i.ID, &i.CreatedAt, &i.Key)
	return i, err
}

const updateQuestion = `-- name: UpdateQuestion :one
update questions set text = ?, answer = ? where id = ? returning id, created_at, quiz_id, text, answer
`

type UpdateQuestionParams struct {
	Text   string
	Answer string
	ID     int64
}

func (q *Queries) UpdateQuestion(ctx context.Context, arg UpdateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, updateQuestion, arg.Text, arg.Answer, arg.ID)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.QuizID,
		&i.Text,
		&i.Answer,
	)
	return i, err
}

const updateQuiz = `-- name: UpdateQuiz :one
update quizzes set name = ?, description = ? where id = ? returning id, created_at, name, description, published
`

type UpdateQuizParams struct {
	Name        interface{}
	Description interface{}
	ID          int64
}

func (q *Queries) UpdateQuiz(ctx context.Context, arg UpdateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, updateQuiz, arg.Name, arg.Description, arg.ID)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Name,
		&i.Description,
		&i.Published,
	)
	return i, err
}

const userGradient = `-- name: UserGradient :one
select id, created_at, user_id, gradient from gradients where user_id = ? order by created_at desc limit 1
`

func (q *Queries) UserGradient(ctx context.Context, userID int64) (Gradient, error) {
	row := q.db.QueryRowContext(ctx, userGradient, userID)
	var i Gradient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UserID,
		&i.Gradient,
	)
	return i, err
}
