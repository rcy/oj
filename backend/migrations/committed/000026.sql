--! Previous: sha1:0ca0121ac257cf18e608c20c5648834e1fd71475
--! Hash: sha1:68b6eece5668b2e98ae0710a69596feace422955

drop table if exists app_public.managed_people;
create table app_public.managed_people (
  id uuid default gen_random_uuid() not null,
  created_at timestamp with time zone default now() not null,
  user_id uuid references app_public.users not null,
  person_id uuid references app_public.people not null
);

grant all on table app_public.managed_people to visitor;
