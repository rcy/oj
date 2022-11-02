--! Previous: sha1:f6f2b2e19faa254887256d90583ad87b3b03fe02
--! Hash: sha1:08dbdb80d0a9402e8a53058878a8e4511ca64c11

set search_path to app_public;

alter table users
      drop column if exists family_id;

alter table users
      add column family_id uuid references families;

update users
       set family_id = families.id
       from families
       where families.user_id = users.id;

-- next migration drops families.user_id
