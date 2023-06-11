CREATE TABLE migration_version (
			version INTEGER
		);
CREATE TABLE messages(
  id integer primary key,
  body text not null,
  sender text not null,
  created_at datetime not null default current_timestamp
);
