--! Previous: sha1:669b897f96a1a315d7468cf6ae6dc7f773e89750
--! Hash: sha1:574596cab00acb0c55cc8a3c29a176feb203550e

alter table app_public.family_memberships drop column if exists role_id;
alter table app_public.family_memberships drop column if exists role;
alter table app_public.family_memberships add column role text not null;

create or replace function app_private.create_family() returns trigger
language plpgsql security definer
as $$
declare
  v_family_id uuid;
begin
  -- create the family
  insert into app_public.families(user_id) values(new.id) returning id into v_family_id;
  -- create a family membership for the user person as admin
  insert into app_public.family_memberships(person_id, family_id, role) values(new.person_id, v_family_id, 'admin');
  return new;
end;
$$;
