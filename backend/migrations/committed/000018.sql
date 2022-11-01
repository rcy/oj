--! Previous: sha1:1383db6b791428b7aab4fd89d53d155d467ad729
--! Hash: sha1:602c804a14e22121cc941a169904a19b40effa19

-- Enter migration here

alter table app_public.space_memberships drop constraint if exists space_memberships_person_id_space_id_unq;

alter table app_public.space_memberships
add constraint space_memberships_person_id_space_id_unq unique(person_id, space_id);
