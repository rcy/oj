--! Previous: sha1:0b57f285d7e87b4533cc861ca775a6597f6174d8
--! Hash: sha1:b96666fe513d0e949a43eebb58d192f45323f309

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

  v_person_id = nullif(current_setting('person.id', true), '')::uuid;

  if v_person_id is null then
    select person_id from app_public.current_user() into v_person_id;
  end if;

  return v_person_id;
end;
$$;
