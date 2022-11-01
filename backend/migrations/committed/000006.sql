--! Previous: sha1:a6ffffa2b07638496306989447a5fcaa99d782b5
--! Hash: sha1:41bf81890929a9a4d8b23d844a9290a490f96ae7

revoke all privileges on table app_public.users from visitor;

grant select on table app_public.users to visitor;
grant update on table app_public.users to visitor;

alter table app_public.users enable row level security;

drop policy if exists select_all on app_public.users;
create policy select_all on app_public.users for select using (true);

drop policy if exists update_own on app_public.users;
create policy update_own on app_public.users for update using (id = app_public.user_id());
