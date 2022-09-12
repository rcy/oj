--! Previous: sha1:1d0357e4dedc8c4aeefb8b84b3b93512a137fda9
--! Hash: sha1:669b897f96a1a315d7468cf6ae6dc7f773e89750

alter table app_public.family_memberships drop column if exists id;
alter table app_public.family_memberships add column id uuid primary key default gen_random_uuid();
