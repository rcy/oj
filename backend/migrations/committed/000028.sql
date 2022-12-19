--! Previous: sha1:c46453fd872de1cd3daffc5c9d1d290fecdc0a6f
--! Hash: sha1:777478f2583af5654fc7077b16d577ff6a3aaedf

alter table app_public.family_memberships drop constraint if exists family_memberships_person_id_unq;

alter table app_public.family_memberships
  add constraint family_memberships_person_id_unq unique (person_id);
