--! Previous: sha1:be1904ec9b1ce61f5e50bedc38b73adeed190ed0
--! Hash: sha1:00a38a506b4394fce10aff1ed697976c56e47dd4

-- move function from to private schema

drop function if exists app_public.create_user_authentication(text, text, text, jsonb);

drop function if exists app_private.create_user_authentication(text, text, text, jsonb);
create function app_private.create_user_authentication(name text, service text, identifier text, details jsonb) returns uuid as $$
declare
  user_id uuid;
begin
  insert into app_public.users(name) values(name) returning id into user_id;
  insert into app_public.authentications(service, identifier, user_id, details) values(service, identifier, user_id, details);
  return user_id;
end;
$$ language plpgsql strict volatile;
