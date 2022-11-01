--! Previous: sha1:4a7cfdb7298896fb48a880bc81ebc3c49a802af0
--! Hash: sha1:a95d5e83eda77cf80bf426ad3fe6efbe3f2f5357

-- adapted from https://github.com/voxpelli/node-connect-pg-simple/blob/head/table.sql
drop table if exists app_private.passport_sessions;

create table app_private.passport_sessions (
  "sid" varchar not null collate "default",
  "sess" json not null,
  "expire" timestamp(6) not null
)
with (oids=false);

alter table app_private.passport_sessions add constraint "session_pkey" primary key ("sid") not deferrable initially immediate;

create index "idx_session_expire" on app_private.passport_sessions ("expire");
