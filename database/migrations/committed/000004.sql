--! Previous: sha1:c07ec65a8601d448169ace83aefe7143bbdc3a82
--! Hash: sha1:9f1f2bca52283665cde704ca7bb9f38b317d35a9

-- prepare fol rls for visitor role based access to tables

drop function if exists app_public.user_id();
create function app_public.user_id() returns uuid as $$
  select nullif(current_setting('user.id', true), '')::uuid;
$$ language sql stable;

grant all privileges on schema app_public to visitor;
grant all privileges on table app_public.users to visitor;
grant all privileges on table app_public.authentications to visitor;
