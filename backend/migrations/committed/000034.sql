--! Previous: sha1:3fb8bf83b741da6bcc4ce1ae2deb6bc2ff680ce3
--! Hash: sha1:c17446c01e6d830be78a27916c2b94783178c614

drop table if exists app_private.login_codes;
create table app_private.login_codes (
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       person_id uuid references app_public.people not null,
       attempts int not null default 0,
       code text not null
);

drop table if exists app_private.sessions;
create table app_private.sessions (
       id uuid primary key default gen_random_uuid(),
       created_at timestamptz default now() not null,
       person_id uuid references app_public.people not null
);

drop function if exists app_public.create_login_code;
create function app_public.create_login_code(username text) returns uuid
  language plpgsql strict
  security definer
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
    insert into app_private.login_codes(person_id, code)
    values(v_person_id, app_public.gen_random_code(6))
    returning id into v_result;
  end if;

  return v_result;
end;
$$;
comment on function app_public.create_login_code(text) IS '@resultFieldName loginCodeId';

drop function if exists app_public.exchange_code;
create function app_public.exchange_code(login_code_id uuid, code text) returns uuid
  language plpgsql strict
  security definer
  as $$
declare
v_login_code_id uuid;
v_person_id uuid;
v_attempts int;
v_result uuid;
begin
  select id, person_id
    into v_login_code_id, v_person_id
    from app_private.login_codes t
    where id = login_code_id and t.code = exchange_code.code;
  
  if v_person_id is not null then
    insert into app_private.sessions(person_id) values (v_person_id) returning id into v_result;
    delete from app_private.login_codes where id = v_login_code_id;
  else
    update app_private.login_codes set attempts = attempts + 1 where id = login_code_id returning attempts into v_attempts;
    if v_attempts >= 3 then
      delete from app_private.login_codes where id = login_code_id;
    end if;
  end if;

  return v_result;
end;
$$;
comment on function app_public.exchange_code(uuid, text) IS '@resultFieldName sessionKey';


drop function if exists app_public.gen_random_code;
create function app_public.gen_random_code(len int) returns text
  language plpgsql strict
  as $$
declare
v_result text;
begin
  select into v_result array_to_string(array(select substr('0123456789',((random()*9+1)::integer),1) from generate_series(1,len)),'');
  return v_result;
end;     
$$;
