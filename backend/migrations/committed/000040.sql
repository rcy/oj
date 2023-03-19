--! Previous: sha1:c2731c52091e511ff10d792caf73606722b41734
--! Hash: sha1:f3574ed6b57b0330ece73351add6307d12451234

drop function if exists app_public.create_login_code;
create function app_public.create_login_code(username text) returns uuid
    language plpgsql strict security definer
    as $$
declare
v_person_id uuid;
v_result uuid;
begin
  select id
  into v_person_id
  from app_public.people p
  where p.username = create_login_code.username;

  if v_person_id is not null then
    -- remove all existing codes
    delete from app_private.login_codes where person_id = v_person_id;

    -- generate and insert new code
    insert into app_private.login_codes(person_id, code)
      values(v_person_id, app_public.gen_random_code(4))
      returning id into v_result;
  end if;

  return v_result;
end;
$$;
comment on function app_public.create_login_code(text) IS '@resultFieldName loginCodeId';
