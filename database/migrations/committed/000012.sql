--! Previous: sha1:674f19f0ef3a42b2df4e416666ccca771bd01163
--! Hash: sha1:1d0357e4dedc8c4aeefb8b84b3b93512a137fda9

create or replace function app_private.create_person() returns trigger
language plpgsql security definer
as $$
declare 
  v_person_id uuid;
begin
  -- create a person using name from the user
  insert into app_public.people(name) values(new.name) returning id into v_person_id;
  new.person_id = v_person_id;
  return new;
end;
$$;

create or replace trigger _100_create_person before insert on app_public.users
for each row
execute function app_private.create_person();



create or replace function app_private.create_family() returns trigger
language plpgsql security definer
as $$
declare
  v_family_id uuid;
begin
  -- create the family
  insert into app_public.families(user_id) values(new.id) returning id into v_family_id;
  -- create a family membership for the user person as owner
  insert into app_public.family_memberships(person_id, family_id, role_id) values(new.person_id, v_family_id, 1);
  return new;
end;
$$;

create or replace trigger _200_create_family after insert on app_public.users
for each row
execute function app_private.create_family();
