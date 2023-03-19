--! Previous: sha1:26a072ecc7e393739600cd02253c4cd758a6919a
--! Hash: sha1:0b57f285d7e87b4533cc861ca775a6597f6174d8

-- Enter migration here
drop function if exists app_public.current_family_membership_id;
CREATE FUNCTION app_public.current_family_membership_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
  select m2.id
  from app_public.family_memberships as m1
  join app_public.family_memberships as m2 on m1.family_id = m2.family_id
  join app_public.users on users.person_id = m1.person_id
  where
    users.id = app_public.user_id() and
    m1.role = 'admin' and
    m2.person_id = app_public.current_person_id();
$$;
