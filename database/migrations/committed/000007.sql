--! Previous: sha1:41bf81890929a9a4d8b23d844a9290a490f96ae7
--! Hash: sha1:b0735dca145fdb4c63a68b4e6fc02cd1d42aa830

drop function if exists app_public.current_user();
create function app_public.current_user() returns app_public.users as $$
  select users.* from app_public.users where id = app_public.user_id();
$$ language sql stable;
