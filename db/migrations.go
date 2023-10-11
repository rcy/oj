package db

var Migrations = []string{
	`--1
create table messages(
  id integer primary key,
  body text not null,
  sender text not null,
  created_at datetime not null default current_timestamp
)`,
	`--2
drop table if exists users;
create table users (
  id integer primary key,
  created_at datetime not null default current_timestamp,
  username text not null unique check(length(username) > 0),
  email text check (email like '%@%') unique
);

drop table if exists codes;
create table codes(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  code string not null,
  nonce string not null,
  email text check (email like '%@%') not null
);
`,
	`--3
drop table if exists sessions;
create table sessions(
  id integer primary key,
  user_id integer references users not null,
  key string unique not null
);
`,
	`--4
drop table if exists kids_codes;
create table kids_codes(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  code string not null,
  nonce string not null,
  user_id integer references users not null
);
`,
	`--5
drop table if exists kids_parents;
create table kids_parents(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  kid_id integer references users not null,
  parent_id integer references users not null
);
`,
	`--6
drop table if exists gradients;
create table gradients(
  id integer primary key,
  created_at text not null default current_timestamp,
  user_id integer references users not null,
  gradient blob not null
);
`,
	`--7
drop table if exists messages;
create table messages(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  room_id text not null,
  body text not null,
  sender text not null
)`,
	`--8 add rooms
drop table if exists rooms;
create table rooms(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  key text not null
);`,
	`--9 update messages room_id
drop table if exists messages;
create table messages(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  sender_id integer references users not null,
  room_id text not null references rooms not null,
  body text not null
);`,
	`--10 add room_users
drop table if exists room_users;
create table room_users(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  room_id integer references rooms not null,
  user_id integer references users not null,
  unique(room_id, user_id)
);
`,
	`--11 add deliveries
drop table if exists deliveries;
create table deliveries(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  message_id integer references messages not null,
  recipient_id integer references users not null,
  sent_at datetime,
  read_at datetime,
  unique(message_id, recipient_id)
);`,
	`--12 add bios table
drop table if exists bios;
create table bios(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  user_id references users not null,
  text string not null default ''
);`,
	`--13 add avatar_url
alter table users add column avatar_url not null default 'http://www.gravatar.com/avatar/?d=mp'
`,
	`--14 fix default avatar_url domain
alter table users drop column avatar_url;
alter table users add column avatar_url not null default 'https://www.gravatar.com/avatar/?d=mp';
`,
	`-- 15 fix deliveries
drop table if exists reads;
drop table if exists deliveries;
create table deliveries(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  message_id integer references messages not null,
  room_id integer references rooms not null,
  recipient_id integer references users not null,
  sender_id integer references users not null,
  sent_at datetime,
  unique(message_id, recipient_id)
);`,
	`-- 16 add friends
drop table if exists friends;
create table friends(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  a_id integer not null references users(id) on delete cascade,
  b_id integer not null references users(id) on delete cascade,
  b_role text not null
);
insert into friends(a_id, b_id, b_role) select parent_id, kid_id, 'child' from kids_parents;
insert into friends(a_id, b_id, b_role) select kid_id, parent_id, 'parent' from kids_parents;
`,
	`-- 17 add is_parent to users
alter table users add column is_parent bool not null default false;
update users set is_parent = true where email is not null;
`,
	`-- 18 add unique constraint to friends
create unique index uidx_friends_a_b on friends (a_id, b_id)
`,
	`-- 19 remove bios table
alter table users add column bio text not null default '';

update users
set bio = (
  select coalesce(
    (select text from bios where users.id = bios.user_id), ''
  )
);

drop table bios;
`,
	` --20 add become_user_id to users
alter table users add column become_user_id integer references users;
`,
	` --21 add admin to users
alter table users add column admin bool not null default false;
`,
	` --22 add images
drop table if exists images;
create table images(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  url string not null,
  user_id references users not null
);
`,
	`--23 add quizzes
drop table if exists quizzes;
create table quizzes(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  name string not null,
  description string not null
);
`,
	`--24 add questions
drop table if exists questions;
create table questions(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  quiz_id references quizzes not null,
  text string not null,
  answer string not null
);
`,
	`--25 add attempts
drop table if exists attempts;
create table attempts(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  quiz_id references quizzes not null,
  user_id references users not null
);
`,
	`--26 add responses
drop table if exists responses;
create table responses(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  quiz_id references quizzes not null,
  user_id references users not null,
  attempt_id references attempts not null,
  question_id references questions not null,
  text string not null,

  unique(attempt_id, question_id)
);
`,
}
