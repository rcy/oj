--
-- PostgreSQL database dump
--

-- Dumped from database version 14.4 (Debian 14.4-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)

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
  insert into app_public.families(user_id) values(new.id) returning id into v_family_id;
  -- create a family membership for the user person as admin
  insert into app_public.family_memberships(person_id, family_id, role) values(new.person_id, v_family_id, 'admin');
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


SET default_tablespace = '';

SET default_table_access_method = heap;

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
-- Name: user_id(); Type: FUNCTION; Schema: app_public; Owner: -
--

CREATE FUNCTION app_public.user_id() RETURNS uuid
    LANGUAGE sql STABLE
    AS $$
  select nullif(current_setting('user.id', true), '')::uuid;
$$;


--
-- Name: passport_sessions; Type: TABLE; Schema: app_private; Owner: -
--

CREATE TABLE app_private.passport_sessions (
    sid character varying NOT NULL,
    sess json NOT NULL,
    expire timestamp(6) without time zone NOT NULL
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
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    user_id uuid NOT NULL
);


--
-- Name: family_memberships; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.family_memberships (
    family_id uuid NOT NULL,
    person_id uuid NOT NULL,
    title text,
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    role text NOT NULL
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
-- Name: people; Type: TABLE; Schema: app_public; Owner: -
--

CREATE TABLE app_public.people (
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
-- Name: passport_sessions session_pkey; Type: CONSTRAINT; Schema: app_private; Owner: -
--

ALTER TABLE ONLY app_private.passport_sessions
    ADD CONSTRAINT session_pkey PRIMARY KEY (sid);


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
-- Name: families families_user_id_key; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.families
    ADD CONSTRAINT families_user_id_key UNIQUE (user_id);


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
-- Name: people people_pkey; Type: CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


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
-- Name: authentications authentications_user_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.authentications
    ADD CONSTRAINT authentications_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_public.users(id);


--
-- Name: families families_user_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.families
    ADD CONSTRAINT families_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_public.users(id);


--
-- Name: family_memberships family_memberships_family_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_family_id_fkey FOREIGN KEY (family_id) REFERENCES app_public.families(id);


--
-- Name: family_memberships family_memberships_person_id_fkey; Type: FK CONSTRAINT; Schema: app_public; Owner: -
--

ALTER TABLE ONLY app_public.family_memberships
    ADD CONSTRAINT family_memberships_person_id_fkey FOREIGN KEY (person_id) REFERENCES app_public.people(id);


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

CREATE POLICY select_own ON app_public.families FOR SELECT USING ((user_id = app_public.user_id()));


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
-- Name: TABLE users; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT,UPDATE ON TABLE app_public.users TO visitor;


--
-- Name: TABLE authentications; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.authentications TO visitor;


--
-- Name: TABLE families; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.families TO visitor;


--
-- Name: COLUMN families.user_id; Type: ACL; Schema: app_public; Owner: -
--

GRANT INSERT(user_id) ON TABLE app_public.families TO visitor;


--
-- Name: TABLE family_memberships; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.family_memberships TO visitor;


--
-- Name: TABLE family_roles; Type: ACL; Schema: app_public; Owner: -
--

GRANT SELECT ON TABLE app_public.family_roles TO visitor;


--
-- Name: TABLE people; Type: ACL; Schema: app_public; Owner: -
--

GRANT ALL ON TABLE app_public.people TO visitor;


--
-- PostgreSQL database dump complete
--

