CREATE TABLE codes(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  code string not null,
  nonce string not null,
  email text check (email like '%@%') not null
);
CREATE TABLE deliveries(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  message_id integer references messages not null,
  room_id integer references rooms not null,
  recipient_id integer references users not null,
  sender_id integer references users not null,
  sent_at datetime,
  unique(message_id, recipient_id)
);
CREATE TABLE friends(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  a_id integer not null references users(id) on delete cascade,
  b_id integer not null references users(id) on delete cascade,
  b_role text not null
);
CREATE TABLE gradients(
  id integer primary key,
  created_at text not null default current_timestamp,
  user_id integer references users not null,
  gradient blob not null
);
CREATE TABLE images(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  url string not null,
  user_id references users not null
);
CREATE TABLE kids_codes(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  code string not null,
  nonce string not null,
  user_id integer references users not null
);
CREATE TABLE kids_parents(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  kid_id integer references users not null,
  parent_id integer references users not null
);
CREATE TABLE messages(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  sender_id integer references users not null,
  room_id text not null references rooms not null,
  body text not null
);
CREATE TABLE migration_version (
			version INTEGER
		);
CREATE TABLE room_users(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  room_id integer references rooms not null,
  user_id integer references users not null,
  unique(room_id, user_id)
);
CREATE TABLE rooms(
  id integer primary key,
  created_at datetime not null default current_timestamp,
  key text not null
);
CREATE TABLE sessions(
  id integer primary key,
  user_id integer references users not null,
  key string unique not null
);
CREATE TABLE users (
  id integer primary key,
  created_at datetime not null default current_timestamp,
  username text not null unique check(length(username) > 0),
  email text check (email like '%@%') unique
, avatar_url not null default 'https://www.gravatar.com/avatar/?d=mp', is_parent bool not null default false, bio text not null default '', become_user_id integer references users, admin bool not null default false);
