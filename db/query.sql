-- name: UserBySessionKey :one
select users.* from sessions join users on sessions.user_id = users.id where sessions.key = ?;

-- name: UserByID :one
select * from users where id = ?;

-- name: CreateParent :one
insert into users(email, username, is_parent) values(?, ?, true) returning *;

-- name: ParentByID :one
select * from users where id = ? and is_parent = true;

-- name: UserByEmail :one
select * from users where email = ?;

-- name: UpdateAvatar :one
update users set avatar_url = ? where id = ? returning *;

-- name: RecentRoomMessages :many
select * from (
  select m.*, sender.avatar_url as sender_avatar_url
   from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = ?
   order by m.created_at desc
   limit 128
  ) t
order by created_at asc;

-- name: AdminRecentMessages :many
select
        m.*,
        sender.username as sender_username,
        sender.avatar_url as sender_avatar_url
 from messages m
 join users sender on m.sender_id = sender.id
 order by m.created_at desc
 limit 128;

-- name: AdminDeleteMessage :one
update messages
set body = '+++ deleted +++'
where id = ?1
returning *;

-- name: UsersWithUnreadCounts :many
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = ? and sent_at is null
group by users.username;


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

-- name: SetQuizPublished :one
update quizzes set published = ? where id = ? returning *;

-- name: QuestionCount :one
select count(*) from questions where quiz_id = ?;

-- name: ResponseCount :one
select count(*) from responses where attempt_id = ?;

-- name: UserByUsername :one
select * from users where username = ?;

-- name: CreateUser :one
insert into users(username) values(?) returning *;

-- name: CreateKidParent :one
insert into kids_parents(kid_id, parent_id) values(?, ?) returning *;

-- name: CreateFriend :one
insert into friends(a_id, b_id, b_role) values(?, ?, ?) returning *;

-- name: AllUsers :many
select * from users order by created_at desc;

-- name: ParentsByKidID :many
select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = ?;

-- name: KidsByParentID :many
select users.* from kids_parents join users on kids_parents.kid_id = users.id where kids_parents.parent_id = ? order by kids_parents.created_at desc;

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

-- name: Question :one
select * from questions where id = ?;

-- name: CreateQuestion :one
insert into questions(quiz_id, text, answer) values(?,?,?) returning *;

-- name: UpdateQuestion :one
update questions set text = ?, answer = ? where id = ? returning *;

-- name: QuizQuestions :many
select * from questions where quiz_id = ?;

-- name: AllQuizzes :many
select * from quizzes order by created_at desc;

-- name: PublishedQuizzes :many
select * from quizzes where published = true order by created_at desc;

-- name: Quiz :one
select * from quizzes where id = ?;

-- name: UpdateQuiz :one
update quizzes set name = ?, description = ? where id = ? returning *;

-- name: CreateQuiz :one
insert into quizzes(name,description) values(?,?) returning *;

-- name: Responses :many
select
   responses.*,
   questions.answer question_answer,
   questions.text question_text,
   questions.answer = responses.text is_correct
from responses
 join questions on responses.question_id = questions.id
 where attempt_id = ?
order by responses.created_at;

-- name: RoomByKey :one
select * from rooms where key = ?;

-- name: CreateRoom :one
insert into rooms(key) values(?) returning *;

-- name: CreateRoomUser :one
insert into room_users(room_id, user_id) values(?, ?) returning *;

-- name: GetConnection :one
select u.*,
       case
           when f1.a_id = ?1 then f1.b_role
           else ""
       end as role_out,
       case
           when f2.b_id = ?1 then f2.b_role
           else ""
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = ?1
left join friends f2 on f2.a_id = u.id and f2.b_id = ?1
where
  u.id = ?2;

-- name: GetCurrentAndPotentialParentConnections :many
select u.*,
       case
           when f1.a_id = ?1 then f1.b_role
           else ""
       end as role_out,
       case
           when f2.b_id = ?1 then f2.b_role
           else ""
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = ?1
left join friends f2 on f2.a_id = u.id and f2.b_id = ?1
where
  u.id != ?1
and
  is_parent = 1
order by role_in desc
limit 128;

-- name: GetFriends :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1 and f1.b_role = 'friend'
join friends f2 on f2.a_id = u.id and f2.b_id = ?1 and f2.b_role = 'friend';

-- name: GetConnections :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1
join friends f2 on f2.a_id = u.id and f2.b_id = ?1
where f1.b_role <> '' and f2.b_role <> '';

-- name: GetFriendsWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1
join friends f2 on f2.a_id = u.id and f2.b_id = ?1
left outer join gradients g
on g.user_id = f1.b_id
where f1.b_role = 'friend'
group by u.id;

-- name: GetFamilyWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1
join friends f2 on f2.a_id = u.id and f2.b_id = ?1
left outer join gradients g
on g.user_id = f1.b_id
where f1.b_role <> 'friend'
group by u.id;

-- name: GetConnectionsWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1
join friends f2 on f2.a_id = u.id and f2.b_id = ?1
left outer join gradients g
on g.user_id = f1.b_id
group by u.id;

-- name: GetKids :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1 and f1.b_role = 'child'
join friends f2 on f2.a_id = u.id and f2.b_id = ?1 and f2.b_role = 'parent';

-- name: GetParents :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = ?1 and f1.b_role = 'parent'
join friends f2 on f2.a_id = u.id and f2.b_id = ?1 and f2.b_role = 'child';

-- name: UserPostcardsSent :many
select p.*, r.username, r.avatar_url
from postcards p
join users r on p.recipient = r.id
where sender = ?
order by p.created_at desc;

-- name: UserPostcardsReceived :many
select p.*, s.username, s.avatar_url
from postcards p
join users s on p.sender = s.id
where recipient = ?
order by p.created_at desc;

-- name: CreatePostcard :one
insert into postcards(sender, recipient, subject, body, state) values(?,?,?,?,?) returning *;

-- name: CreateThread :one
insert into threads(user_id, thread_id, assistant_id) values(?,?,?) returning *;

-- name: AssistantThreads :many
select * from threads where assistant_id = ? and user_id = ?;

-- name: UserThreadByID :one
select * from threads where user_id = ? and thread_id = ?;

-- name: CreateBot :one
insert into bots(owner_id, assistant_id, name, description) values(?,?,?,?) returning *;

-- name: AllBots :many
select * from bots;

-- name: PublishedBots :many
select * from bots where published = 1;

-- name: UserVisibleBots :many
select * from bots where owner_id = ? or published = 1;

-- name: Bot :one
select * from bots where id = ?;
