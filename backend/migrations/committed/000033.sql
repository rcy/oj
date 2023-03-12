--! Previous: sha1:ead6262d24bda06e44c1dddaad3cb0a23cc48980
--! Hash: sha1:3fb8bf83b741da6bcc4ce1ae2deb6bc2ff680ce3

-- Enter migration here
alter table app_public.people drop column if exists username;
alter table app_public.people add column username text unique;
