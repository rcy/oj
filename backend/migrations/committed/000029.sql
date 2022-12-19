--! Previous: sha1:777478f2583af5654fc7077b16d577ff6a3aaedf
--! Hash: sha1:dfea05713d55236117820fd5b35321980d5a533f

create or replace function app_public.current_family_membership_id() returns uuid
    language sql stable
    as $$
  select m2.id
  from app_public.family_memberships as m1
  join app_public.family_memberships as m2 on m1.family_id = m2.family_id
  join app_public.users on users.person_id = m1.person_id
  where
    users.id = app_public.user_id() and
    m1.role = 'admin' and
    m2.person_id = current_setting('person.id', true)::uuid;
$$;
