--! Previous: sha1:b7fdef2e4a0093412523ad0ba1705c901351ed19
--! Hash: sha1:674f19f0ef3a42b2df4e416666ccca771bd01163

-- Enter migration here

alter table app_public.users drop column if exists person_id;
drop table if exists app_public.family_memberships;
drop table if exists app_public.people;
drop table if exists app_public.family_memberships;
drop table if exists app_public.family_roles;

create table app_public.people(
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       updated_at timestamptz default now() not null,
       name text not null
);
grant all privileges on app_public.people to visitor;
alter table app_public.users add column person_id uuid references app_public.people unique;

-- roles can be "parent", "child", "friend of family", "grandparent", "extended family", etc
create table app_public.family_roles(
       id serial primary key,
       name text not null unique
);
grant select on app_public.family_roles to visitor;

create table app_public.family_memberships(
       family_id uuid references app_public.families not null,
       person_id uuid references app_public.people not null,
       role_id int references app_public.family_roles not null,
       -- how someone is named in the family, mom, dad, by name, etc
       title text
);
grant all privileges on app_public.family_memberships to visitor;
