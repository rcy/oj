--! Previous: sha1:119c3e0afc52952f9bd88d5accaeb81d5971d116
--! Hash: sha1:0ca0121ac257cf18e608c20c5648834e1fd71475

alter table app_public.family_memberships
drop constraint family_memberships_person_id_fkey,
add constraint family_memberships_person_id_fkey
  foreign key (person_id)
  references app_public.people(id)
  on delete cascade;
