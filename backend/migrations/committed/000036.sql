--! Previous: sha1:093f84a203444f3e61536ce76702e3883d879e3d
--! Hash: sha1:26a072ecc7e393739600cd02253c4cd758a6919a

drop function if exists app_public.current_person_id;
CREATE FUNCTION app_public.current_person_id() RETURNS uuid
    LANGUAGE plpgsql
    AS $$
declare
v_person_id uuid;
begin
  -- if the person.id session variable is set, return that
  -- otherwise, return the person_id associated with the user.id, if that is set
  -- otherwise, return null

  v_person_id = current_setting('person.id', true)::uuid;

  if v_person_id is null then
    select person_id from app_public.current_user() into v_person_id;
  end if;

  return v_person_id;
end;
$$;
