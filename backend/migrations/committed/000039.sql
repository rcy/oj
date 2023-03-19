--! Previous: sha1:b96666fe513d0e949a43eebb58d192f45323f309
--! Hash: sha1:c2731c52091e511ff10d792caf73606722b41734

drop function if exists app_public.become_person;
create function app_public.become_person(id uuid) returns uuid
    language plpgsql strict security definer
    as $$
declare
v_id uuid;
v_result uuid;
begin
  -- check that person_id is managed by the current user
  select mp.id 
  from app_public.managed_people mp
  into v_id
  where mp.person_id = become_person.id
  and user_id = app_public.user_id();

  if v_id is not null then
    insert into app_private.sessions(person_id) values (become_person.id) returning sessions.id into v_result;
    return v_result;
  else
    raise exception 'person % not managed by user %', id, app_public.user_id();
  end if;
end;
$$;

comment on function app_public.become_person(id uuid) is '@resultFieldName sessionKey';
