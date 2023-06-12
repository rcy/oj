package db

var Migrations = []string{
	`
create table messages(
  id integer primary key,
  body text not null,
  sender text not null,
  created_at datetime not null default current_timestamp
)`, `
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
`, `
drop table if exists sessions;
create table sessions(
  id integer primary key,
  user_id integer references users not null,
  key string unique not null
);
`,
}

var Current = ``
