--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
-- Dumped by pg_dump version 14.7 (Ubuntu 14.7-0ubuntu0.22.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: app_hidden; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_hidden;


--
-- Name: app_private; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_private;


--
-- Name: app_public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA app_public;


--
-- Name: create_family(); Type: FUNCTION; Schema: app_private; Owner: -
--

CREATE FUNCTION app_private.create_family() RETURNS trigger
    LANGUAGE plpgsql SECURITY DEFINER
    AS $$
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


--
-- Name: create_person(); Type: FUNCTION; Schema: app_private; Owner: -
--

CREATE FUNCTION app_private.create_person() RETURNS trigger
    LANGUAGE plpgsql SECURITY DEFINER
    AS $$
declare 
  v_person_id uuid;
begin
  -- create a person using name from the user
  insert into app_public.people(name) values(new.name) returning id into v_person_id;
  new.person_id = v_person_id;
  return new;
end;
$$;


--
-- Name: create_user_authentication(text, text, text, jsonb); Type: FUNCTION; Schema: app_private; Owner: -
--

CREATE FUNCTION app_private.create_user_authentication(name text, service text, identifier text, details jsonb) RETURNS uuid
    LANGUAGE plpgsql STRICT
    AS $$
declare
  user_id uuid;
begin
  insert into app_public.users(name) values(name) returning id into user_id;
  insert into app_public.authentications(service, identifier, user_id, details) values(service, identifier, user_id, details);
  return user_id;
end;
$$;


--
-- Name: become_person(uuid); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.become_person(id uuid) RETURNS uuid
    LANGUAGE plpgsql STRICT SECURITY DEFINER
    AS $$
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


--
-- Name: FUNCTION become_person(id uuid); Type: COMMENT; Schema: app_public; Owner: -
--

COMMENT ON FUNCTION app_public.become_person(id uuid) IS '@resultFieldName sessionKey';


--
-- Name: create_login_code(text); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.create_login_code(username text) RETURNS uuid
    LANGUAGE plpgsql STRICT SECURITY DEFINER
    AS $$
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


--
-- Name: FUNCTION create_login_code(username text); Type: COMMENT; Schema: app_public; Owner: -
--

COMMENT ON FUNCTION app_public.create_login_code(username text) IS '@resultFieldName loginCodeId';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: family_memberships; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.family_memberships (
    family_id uuid NOT NULL,
    person_id uuid NOT NULL,
    title text,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    role text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: create_new_family_member(text, text); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.create_new_family_member(name text, role text) RETURNS app_public.family_memberships
    LANGUAGE plpgsql STRICT
    AS $$
declare
  v_person_id uuid;
  v_result app_public.family_memberships;
  v_family_id uuid;
begin
  -- TODO: verify current user is admin in family
  select family_id into v_family_id from app_public.family_memberships where id = app_public.current_family_membership_id();
  insert into app_public.people(name) values(name) returning id into v_person_id;
  insert into app_public.family_memberships(person_id, family_id, role) values(v_person_id, v_family_id, role) returning * into v_result;
  return v_result;
end;
$$;


--
-- Name: current_family_membership(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.current_family_membership() RETURNS app_public.family_memberships
    LANGUAGE sql STABLE
    AS $$
  select family_memberships.* from app_public.family_memberships where id = app_public.current_family_membership_id();
$$;


--
-- Name: current_family_membership_id(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.current_family_membership_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
  select m2.id
  from app_public.family_memberships as m1
  join app_public.family_memberships as m2 on m1.family_id = m2.family_id
  join app_public.users on users.person_id = m1.person_id
  where
    users.id = app_public.user_id() and
    m1.role = 'admin' and
    m2.person_id = app_public.current_person_id();
$$;


--
-- Name: people; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.people (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    avatar_url text DEFAULT 'https://www.gravatar.com/avatar/DEFAULT?f=y&d=mp'::text NOT NULL,
    username text
);


--
-- Name: current_person(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.current_person() RETURNS app_public.people
    LANGUAGE sql STABLE
    AS $$
  select people.*
  from app_public.people
  where id = app_public.current_person_id();
$$;


--
-- Name: current_person_id(); Type: FUNCTION; Schema: app_public; Owner: -
--

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


--
-- Name: users; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    avatar_url text,
    person_id uuid,
    family_id uuid,
    CONSTRAINT users_avatar_url_check CHECK ((avatar_url ~ '^https?://[^/]+'::text))
);


--
-- Name: current_user(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public."current_user"() RETURNS app_public.users
    LANGUAGE sql STABLE
    AS $$
  select users.* from app_public.users where id = app_public.user_id();
$$;


--
-- Name: exchange_code(uuid, text); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.exchange_code(login_code_id uuid, code text) RETURNS uuid
    LANGUAGE plpgsql STRICT SECURITY DEFINER
    AS $$
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


--
-- Name: FUNCTION exchange_code(login_code_id uuid, code text); Type: COMMENT; Schema: app_public; Owner: -
--

COMMENT ON FUNCTION app_public.exchange_code(login_code_id uuid, code text) IS '@resultFieldName sessionKey';


--
-- Name: gen_random_code(integer); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.gen_random_code(len integer) RETURNS text
    LANGUAGE plpgsql STRICT
    AS $$
declare
v_result text;
begin
  select into v_result array_to_string(array(select substr('0123456789',((random()*9+1)::integer),1) from generate_series(1,len)),'');
  return v_result;
end;     
$$;


--
-- Name: notify_space_post_created(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.notify_space_post_created() RETURNS trigger
    LANGUAGE plpgsql
    SET search_path TO 'app_public'
    AS $$
begin
  perform pg_notify('graphql:spaceposts:' || new.space_id, json_build_object('event', 'postCreated', 'subject', new.id)::text);
  return new;
end
$$;


--
-- Name: posts; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.posts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    membership_id uuid NOT NULL,
    body text NOT NULL,
    space_id uuid NOT NULL
);


--
-- Name: post_message(uuid, text); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.post_message(space_membership_id uuid, body text) RETURNS app_public.posts
    LANGUAGE plpgsql STRICT SECURITY DEFINER
    AS $$
declare
        v_space_id uuid;
        v_result app_public.posts;
begin
        select space_id
        into v_space_id
        from app_public.space_memberships
        where id = space_membership_id;

        insert
                into app_public.posts(membership_id, space_id, body)
                values(space_membership_id, v_space_id, body)
                returning * into v_result;

        return v_result;
end;
$$;


--
-- Name: trigger_job(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.trigger_job() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
  PERFORM graphile_worker.add_job(TG_ARGV[0], json_build_object(
    'schema', TG_TABLE_SCHEMA,
    'table', TG_TABLE_NAME,
    'op', TG_OP,
    'id', (CASE WHEN TG_OP = 'DELETE' THEN OLD.id ELSE NEW.id END)
  ));
  RETURN NEW;
END;
$$;


--
-- Name: user_id(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.user_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
  select nullif(current_setting('user.id', true), '')::uuid;
$$;


--
-- Name: login_codes; Type: TABLE; Schema: app_private; Owner: -
--

CREATE TABLE app_private.login_codes (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    person_id uuid NOT NULL,
    attempts integer DEFAULT 0 NOT NULL,
    code text NOT NULL
);


--
-- Name: passport_sessions; Type: TABLE; Schema: app_private; Owner: -
--

CREATE TABLE app_private.passport_sessions (
    sid character varying NOT NULL,
    sess json NOT NULL,
    expire timestamp(6) without time zone NOT NULL
);


--
-- Name: sessions; Type: TABLE; Schema: app_private; Owner: -
--

CREATE TABLE app_private.sessions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    person_id uuid NOT NULL
);


--
-- Name: authentications; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.authentications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL,
    service text NOT NULL,
    identifier text NOT NULL,
    details jsonb DEFAULT '{}'::jsonb NOT NULL
);


--
-- Name: families; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.families (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: family_roles; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.family_roles (
    id integer NOT NULL,
    name text NOT NULL
);


--
-- Name: family_roles_id_seq; Type: SEQUENCE; Schema: app_public; Owner: -
--

CREATE SEQUENCE app_public.family_roles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: family_roles_id_seq; Type: SEQUENCE OWNED BY; Schema: app_public; Owner: -
--

ALTER SEQUENCE app_public.family_roles_id_seq OWNED BY app_public.family_roles.id;


--
-- Name: interests; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.interests (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    topic_id uuid,
    person_id uuid
);


--
-- Name: managed_people; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.managed_people (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL,
    person_id uuid NOT NULL
);


--
-- Name: notifications; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.notifications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    post_id uuid NOT NULL,
    membership_id uuid NOT NULL
);


--
-- Name: space_memberships; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.space_memberships (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    person_id uuid NOT NULL,
    space_id uuid NOT NULL,
    role_id text NOT NULL
);


--
-- Name: space_topics; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.space_topics (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    topic_id uuid,
    space_id uuid
);


--
-- Name: spaces; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.spaces (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    description text DEFAULT 'default description'::text
);


--
-- Name: topics; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.topics (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    name text NOT NULL
);


--
-- Name: family_roles id; Type: DEFAULT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_roles ALTER COLUMN id SET DEFAULT nextval('app_public.family_roles_id_seq'::regclass);


--
-- Name: login_codes login_codes_pkey; Type: CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.login_codes
    ADD CONSTRAINT login_codes_pkey PRIMARY KEY (id);


--
-- Name: passport_sessions session_pkey; Type: CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.passport_sessions
    ADD CONSTRAINT session_pkey PRIMARY KEY (sid);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: authentications authentications_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_pkey PRIMARY KEY (id);


--
-- Name: authentications authentications_service_identifier_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_service_identifier_key UNIQUE (service, identifier);


--
-- Name: families families_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.families
    ADD CONSTRAINT families_pkey PRIMARY KEY (id);


--
-- Name: family_memberships family_memberships_person_id_unq; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_person_id_unq UNIQUE (person_id);


--
-- Name: family_memberships family_memberships_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_pkey PRIMARY KEY (id);


--
-- Name: family_roles family_roles_name_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_roles
    ADD CONSTRAINT family_roles_name_key UNIQUE (name);


--
-- Name: family_roles family_roles_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_roles
    ADD CONSTRAINT family_roles_pkey PRIMARY KEY (id);


--
-- Name: interests interests_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.interests
    ADD CONSTRAINT interests_pkey PRIMARY KEY (id);


--
-- Name: people people_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


--
-- Name: people people_username_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.people
    ADD CONSTRAINT people_username_key UNIQUE (username);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: space_memberships space_memberships_person_id_space_id_unq; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_memberships
    ADD CONSTRAINT space_memberships_person_id_space_id_unq UNIQUE (person_id, space_id);


--
-- Name: space_memberships space_memberships_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_memberships
    ADD CONSTRAINT space_memberships_pkey PRIMARY KEY (id);


--
-- Name: space_topics space_topics_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_topics
    ADD CONSTRAINT space_topics_pkey PRIMARY KEY (id);


--
-- Name: spaces spaces_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.spaces
    ADD CONSTRAINT spaces_pkey PRIMARY KEY (id);


--
-- Name: topics topics_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.topics
    ADD CONSTRAINT topics_pkey PRIMARY KEY (id);


--
-- Name: users users_person_id_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.users
    ADD CONSTRAINT users_person_id_key UNIQUE (person_id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_session_expire; Type: INDEX; Schema: app_private; Owner: -
--

CREATE INDEX idx_session_expire ON app_private.passport_sessions USING btree (expire);


--
-- Name: users _100_create_person; Type: TRIGGER; Schema: app_public; Owner: -
--

CREATE TRIGGER _100_create_person BEFORE INSERT ON app_public.users FOR EACH ROW EXECUTE FUNCTION app_private.create_person();


--
-- Name: users _200_create_family; Type: TRIGGER; Schema: app_public; Owner: -
--

CREATE TRIGGER _200_create_family AFTER INSERT ON app_public.users FOR EACH ROW EXECUTE FUNCTION app_private.create_family();


--
-- Name: posts _500_notify; Type: TRIGGER; Schema: app_public; Owner: -
--

CREATE TRIGGER _500_notify AFTER INSERT ON app_public.posts FOR EACH ROW EXECUTE FUNCTION app_public.notify_space_post_created();


--
-- Name: posts create_post_notifications; Type: TRIGGER; Schema: app_public; Owner: -
--

CREATE TRIGGER create_post_notifications AFTER INSERT ON app_public.posts FOR EACH ROW EXECUTE FUNCTION app_public.trigger_job('create_post_notifications');


--
-- Name: login_codes login_codes_person_id_fkey; Type: FK CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.login_codes
    ADD CONSTRAINT login_codes_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: sessions sessions_person_id_fkey; Type: FK CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.sessions
    ADD CONSTRAINT sessions_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: authentications authentications_user_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_public.users(id);


--
-- Name: family_memberships family_memberships_family_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_family_id_fkey FOREIGN KEY (family_id) REFERENCES app_public.families(id);


--
-- Name: family_memberships family_memberships_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id) ON DELETE CASCADE;


--
-- Name: interests interests_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.interests
    ADD CONSTRAINT interests_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: interests interests_topic_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.interests
    ADD CONSTRAINT interests_topic_id_fkey FOREIGN KEY (topic_id) REFERENCES app_public.topics(id);


--
-- Name: managed_people managed_people_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.managed_people
    ADD CONSTRAINT managed_people_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: managed_people managed_people_user_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.managed_people
    ADD CONSTRAINT managed_people_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_public.users(id);


--
-- Name: notifications notifications_membership_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.notifications
    ADD CONSTRAINT notifications_membership_id_fkey FOREIGN KEY (membership_id) REFERENCES app_public.space_memberships(id);


--
-- Name: notifications notifications_post_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.notifications
    ADD CONSTRAINT notifications_post_id_fkey FOREIGN KEY (post_id) REFERENCES app_public.posts(id);


--
-- Name: posts posts_membership_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.posts
    ADD CONSTRAINT posts_membership_id_fkey FOREIGN KEY (membership_id) REFERENCES app_public.space_memberships(id);


--
-- Name: posts posts_space_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.posts
    ADD CONSTRAINT posts_space_id_fkey FOREIGN KEY (space_id) REFERENCES app_public.spaces(id);


--
-- Name: space_memberships space_memberships_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_memberships
    ADD CONSTRAINT space_memberships_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: space_memberships space_memberships_space_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_memberships
    ADD CONSTRAINT space_memberships_space_id_fkey FOREIGN KEY (space_id) REFERENCES app_public.spaces(id);


--
-- Name: space_topics space_topics_space_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_topics
    ADD CONSTRAINT space_topics_space_id_fkey FOREIGN KEY (space_id) REFERENCES app_public.spaces(id);


--
-- Name: space_topics space_topics_topic_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.space_topics
    ADD CONSTRAINT space_topics_topic_id_fkey FOREIGN KEY (topic_id) REFERENCES app_public.topics(id);


--
-- Name: users users_family_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.users
    ADD CONSTRAINT users_family_id_fkey FOREIGN KEY (family_id) REFERENCES app_public.families(id);


--
-- Name: users users_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.users
    ADD CONSTRAINT users_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


--
-- Name: authentications; Type: ROW SECURITY; Schema: app_public; Owner: -
--

ALTER TABLE app_public.authentications ENABLE ROW LEVEL SECURITY;

--
-- Name: families; Type: ROW SECURITY; Schema: app_public; Owner: -
--

ALTER TABLE app_public.families ENABLE ROW LEVEL SECURITY;

--
-- Name: family_memberships insert_as_admin; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY insert_as_admin ON app_public.family_memberships WITH CHECK ((EXISTS ( SELECT app_public.current_family_membership() AS current_family_membership
  WHERE (family_memberships.role = 'admin'::text))));


--
-- Name: users select_all; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY select_all ON app_public.users FOR SELECT USING (true);


--
-- Name: authentications select_own; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY select_own ON app_public.authentications FOR SELECT USING ((user_id = app_public.user_id()));


--
-- Name: families select_own; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY select_own ON app_public.families FOR SELECT USING ((id = ( SELECT "current_user".family_id
   FROM app_public."current_user"() "current_user"(id, created_at, updated_at, name, avatar_url, person_id, family_id))));


--
-- Name: users update_own; Type: POLICY; Schema: app_public; Owner: -
--

CREATE POLICY update_own ON app_public.users FOR UPDATE USING ((id = app_public.user_id()));


--
-- Name: users; Type: ROW SECURITY; Schema: app_public; Owner: -
--

ALTER TABLE app_public.users ENABLE ROW LEVEL SECURITY;

--
-- Name: SCHEMA app_public; Type: ACL; Schema: -; Owner: -
--

GRANT ALL ON SCHEMA app_public TO visitor;


--
-- Name: TABLE family_memberships; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.family_memberships TO visitor;


--
-- Name: TABLE people; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.people TO visitor;


--
-- Name: TABLE users; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT,UPDATE ON TABLE app_public.users TO visitor;


--
-- Name: TABLE posts; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.posts TO visitor;


--
-- Name: TABLE authentications; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.authentications TO visitor;


--
-- Name: TABLE families; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.families TO visitor;


--
-- Name: TABLE family_roles; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.family_roles TO visitor;


--
-- Name: TABLE interests; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.interests TO visitor;


--
-- Name: TABLE managed_people; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.managed_people TO visitor;


--
-- Name: TABLE notifications; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.notifications TO visitor;


--
-- Name: TABLE space_memberships; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.space_memberships TO visitor;


--
-- Name: TABLE spaces; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.spaces TO visitor;


--
-- Name: TABLE topics; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.topics TO visitor;


--
-- PostgreSQL database dump complete
--

