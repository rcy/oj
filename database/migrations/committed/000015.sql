--! Previous: sha1:574596cab00acb0c55cc8a3c29a176feb203550e
--! Hash: sha1:db35d00b04a47bfac97b17831d5c3ef34b07fde1

-- only return the id if the current user logged in is an admin in the same family
create or replace function app_public.current_family_membership_id() returns uuid as $$
  select m2.id
  from app_public.family_memberships as m1
  join app_public.family_memberships as m2 on m1.family_id = m2.family_id
  join app_public.users on users.person_id = m1.person_id
  where
    users.id = app_public.user_id() and
    m1.role = 'admin' and
    m2.id = current_setting('family_membership.id', true)::uuid;
$$ language sql stable;

-- return the current family membership if an admin is logged in from the same family
create or replace function app_public.current_family_membership() returns app_public.family_memberships as $$
  select family_memberships.* from app_public.family_memberships where id = app_public.current_family_membership_id();
$$ language sql stable;

create or replace function app_public.create_new_family_member(name text, role text)
returns app_public.family_memberships
language plpgsql strict security invoker
as $$
declare
  v_person_id uuid;
  v_result app_public.family_memberships;
  v_family_id uuid;
begin
  -- TODO: verify current user is admin in family
  select family_id into v_family_id from app_public.family_memberships where id = app_public.current_family_membership_id();
  insert into app_public.people(name) values(name) returning id into v_person_id;
  insert into app_public.family_memberships(person_id, family_id, role) values(v_person_id, v_family_id, role) returning * into v_result;
  return v_result;
end;
$$;

drop policy if exists insert_as_admin on app_public.family_memberships;

create policy insert_as_admin on app_public.family_memberships
  with check (exists (select app_public.current_family_membership() where role = 'admin'));
