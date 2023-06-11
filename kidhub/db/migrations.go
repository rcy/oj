package db

var Migrations = []string{
	`
create table messages(
  id integer primary key,
  body text not null,
  sender text not null,
  created_at datetime not null default current_timestamp
)`,
}
