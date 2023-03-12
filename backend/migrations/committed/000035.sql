--! Previous: sha1:c17446c01e6d830be78a27916c2b94783178c614
--! Hash: sha1:093f84a203444f3e61536ce76702e3883d879e3d

-- enter migration here
drop function if exists app_public.current_person_id;
create function app_public.current_person_id() returns uuid
    language plpgsql
    as $$
begin
  return current_setting('person.id', true)::uuid;
end;
$$;
