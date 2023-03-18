alter table app_public.people drop column if exists meteor_id;
alter table app_public.people add column meteor_id text unique;

drop table if exists app_public.events;
create table app_public.events (
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       name text not null,
       person_id uuid references app_public.people,
       space_id uuid references app_public.spaces,
       payload jsonb not null default '{}'::jsonb
);
