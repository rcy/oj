--! Previous: sha1:be4d1795cc12476bbc67d33ed56515270db5f07c
--! Hash: sha1:1383db6b791428b7aab4fd89d53d155d467ad729

-- Enter migration here

alter table app_public.spaces add column if not exists description text default 'default description';
