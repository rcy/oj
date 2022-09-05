--! Previous: sha1:a95d5e83eda77cf80bf426ad3fe6efbe3f2f5357
--! Hash: sha1:c07ec65a8601d448169ace83aefe7143bbdc3a82

drop table if exists app_public.authentications;
drop table if exists app_public.users;

create table app_public.users(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       name text not null,
       avatar_url text check(avatar_url ~ '^https?://[^/]+')
);

create table app_public.authentications(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       user_id uuid references app_public.users not null,
       service text not null,
       identifier text not null,
       details jsonb default '{}'::jsonb not null,
       unique(service, identifier)
);

drop function if exists app_public.create_user_authentication(text, text, text, jsonb);
create function app_public.create_user_authentication(name text, service text, identifier text, details jsonb) returns uuid as $$
declare
  user_id uuid;
begin
  insert into app_public.users(name) values(name) returning id into user_id;
  insert into app_public.authentications(service, identifier, user_id, details) values(service, identifier, user_id, details);
  return user_id;
end;
$$ language plpgsql strict volatile;
