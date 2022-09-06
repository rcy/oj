--! Previous: sha1:00a38a506b4394fce10aff1ed697976c56e47dd4
--! Hash: sha1:b7fdef2e4a0093412523ad0ba1705c901351ed19

drop table if exists app_public.families;
create table app_public.families (
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       user_id uuid references app_public.users not null unique
);
alter table app_public.families enable row level security;

grant
  select,
  insert(user_id)
on app_public.families
to visitor;

create policy
  select_own
on app_public.families
for select using (user_id = app_public.user_id());
