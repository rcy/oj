--! Previous: sha1:db35d00b04a47bfac97b17831d5c3ef34b07fde1
--! Hash: sha1:be4d1795cc12476bbc67d33ed56515270db5f07c

drop table if exists app_public.posts;
drop table if exists app_public.space_memberships;
drop table if exists app_public.interests;
drop table if exists app_public.space_topics;
drop table if exists app_public.spaces;
drop table if exists app_public.topics;

create table app_public.topics(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       name text not null       
);
grant all privileges on app_public.topics to visitor;

create table app_public.spaces(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       name text not null

);
grant all privileges on app_public.spaces to visitor;

create table app_public.space_topics(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       topic_id uuid references app_public.topics,
       space_id uuid references app_public.spaces
);
grant all privileges on app_public.topics to visitor;

-- aka people_topics
create table app_public.interests(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       topic_id uuid references app_public.topics,
       person_id uuid references app_public.people
);
grant all privileges on app_public.interests to visitor;

create table app_public.space_memberships(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       person_id uuid references app_public.people not null,
       space_id uuid references app_public.spaces not null,
       role_id text not null
);
grant all privileges on app_public.space_memberships to visitor;

create table app_public.posts(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       membership_id uuid references app_public.space_memberships not null,
       body text not null
);
grant all privileges on app_public.posts to visitor;
