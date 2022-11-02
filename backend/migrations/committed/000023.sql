--! Previous: sha1:08dbdb80d0a9402e8a53058878a8e4511ca64c11
--! Hash: sha1:e6246032d3592fc9beef0d31473667ea373f6e22

set search_path to app_public;

drop policy if exists select_own on families;

create policy select_own
on families
for select
using (id = (select family_id from app_public.current_user()));

alter table families drop column if exists user_id;
