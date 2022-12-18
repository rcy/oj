--! Previous: sha1:68b6eece5668b2e98ae0710a69596feace422955
--! Hash: sha1:c46453fd872de1cd3daffc5c9d1d290fecdc0a6f

-- only return the person_id if the current user manages the current person
-- OR the current user *is* the current person
create or replace function app_public.current_person_id() returns uuid as
$$
declare
  v_person_id uuid;
begin
  -- first check to see if person is managed by user
  select person_id
  from app_public.managed_people
  where
    user_id = app_public.user_id() and
    person_id = current_setting('person.id', true)::uuid
  into v_person_id;

  if v_person_id is null then
    -- check if person *is* user
    select person_id
    from app_public.users
    where
      id = app_public.user_id() and
      person_id = current_setting('person.id', true)::uuid
    into v_person_id;
  end if;

  return v_person_id;
end;
$$
language plpgsql;

create or replace function app_public.current_person() returns app_public.people as $$
  select people.*
  from app_public.people
  where id = app_public.current_person_id();
$$ language sql stable;
