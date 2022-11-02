--! Previous: sha1:e6246032d3592fc9beef0d31473667ea373f6e22
--! Hash: sha1:119c3e0afc52952f9bd88d5accaeb81d5971d116

drop trigger if exists _200_create_family on app_public.users;
drop function if exists app_private.create_family();

create function app_private.create_family() returns trigger
    language plpgsql security definer
    as $$
declare
  v_family_id uuid;
begin
  -- create the family
  insert into app_public.families default values returning id into v_family_id;

  -- create a family membership for the user person as admin
  insert into app_public.family_memberships(person_id, family_id, role) values(new.person_id, v_family_id, 'admin');

  -- update the new user
  update app_public.users set family_id = v_family_id where id = new.id;

  return new;
end;
$$;

create trigger _200_create_family after insert on app_public.users for each row execute function app_private.create_family();
