-- name: RecentMessages :many
select * from (
  select m.*, sender.avatar_url as sender_avatar_url
   from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = ?
   order by m.created_at desc
   limit 128
  ) t
order by created_at asc;

-- name: GetAttemptByID :one
select * from attempts where id = ?;

-- name: AttemptNextQuestion :one
select questions.* from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = ?)
and
  questions.quiz_id = ?
order by random();

-- name: QuestionCount :one
select count(*) from questions where quiz_id = ?;

-- name: ResponseCount :one
select count(*) from responses where attempt_id = ?;

-- name: AttemptResponseIDs :many
select id from responses where attempt_id = ?;

-- name: CreateResponse :one
insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(?,?,?,?,?) returning *;

-- name: CreateAttempt :one
insert into attempts(quiz_id, user_id) values(?,?) returning *;

-- name: Delivery :one
select * from deliveries where id = ?;

-- name: UserGradient :one
select * from gradients where user_id = ? order by created_at desc limit 1;
