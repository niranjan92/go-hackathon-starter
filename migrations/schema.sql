--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.1
-- Dumped by pg_dump version 10.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: contacts; Type: TABLE; Schema: public; Owner: niranjan
--

CREATE TABLE contacts (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE contacts OWNER TO niranjan;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: niranjan
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO niranjan;

--
-- Name: user_profiles; Type: TABLE; Schema: public; Owner: niranjan
--

CREATE TABLE user_profiles (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    data character varying(255) NOT NULL,
    user_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE user_profiles OWNER TO niranjan;

--
-- Name: users; Type: TABLE; Schema: public; Owner: niranjan
--

CREATE TABLE users (
    id uuid NOT NULL,
    name character varying(255),
    email character varying(255),
    provider character varying(255) NOT NULL,
    provider_id character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    gender character varying(255) DEFAULT ''::character varying,
    location character varying(255) DEFAULT ''::character varying,
    website character varying(255) DEFAULT ''::character varying,
    gravatar character varying(255) DEFAULT ''::character varying
);


ALTER TABLE users OWNER TO niranjan;

--
-- Name: widgets; Type: TABLE; Schema: public; Owner: niranjan
--

CREATE TABLE widgets (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE widgets OWNER TO niranjan;

--
-- Name: contacts contacts_pkey; Type: CONSTRAINT; Schema: public; Owner: niranjan
--

ALTER TABLE ONLY contacts
    ADD CONSTRAINT contacts_pkey PRIMARY KEY (id);


--
-- Name: user_profiles user_profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: niranjan
--

ALTER TABLE ONLY user_profiles
    ADD CONSTRAINT user_profiles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: niranjan
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: widgets widgets_pkey; Type: CONSTRAINT; Schema: public; Owner: niranjan
--

ALTER TABLE ONLY widgets
    ADD CONSTRAINT widgets_pkey PRIMARY KEY (id);


--
-- Name: user_profiles_provider_provider_id_idx; Type: INDEX; Schema: public; Owner: niranjan
--

CREATE UNIQUE INDEX user_profiles_provider_provider_id_idx ON user_profiles USING btree (provider, provider_id);


--
-- Name: users_provider_provider_id_idx; Type: INDEX; Schema: public; Owner: niranjan
--

CREATE UNIQUE INDEX users_provider_provider_id_idx ON users USING btree (provider, provider_id);


--
-- Name: version_idx; Type: INDEX; Schema: public; Owner: niranjan
--

CREATE UNIQUE INDEX version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

