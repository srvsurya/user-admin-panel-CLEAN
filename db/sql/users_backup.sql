--
-- PostgreSQL database dump
--

\restrict FcrM9UfX2Qq2gnxDtbR0cntjSGWsJOxCTaJ72IgMHqa3yFxcETzeicyyjSkP7y2

-- Dumped from database version 16.13 (Ubuntu 16.13-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.13 (Ubuntu 16.13-0ubuntu0.24.04.1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: saurav
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    password text NOT NULL,
    role character varying(50)
);


ALTER TABLE public.users OWNER TO saurav;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: saurav
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.users_user_id_seq OWNER TO saurav;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: saurav
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: saurav
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: saurav
--

COPY public.users (user_id, name, email, password, role) FROM stdin;
3	Jason Red	jasonred@gmail.com	$2a$10$SuGAf0GWmOD/e5nf3qETj.O/iLGcdY7V.aQKLlA849lJdbJQMh2L6	admin
2	Saurav Jayasurya	themednobility@gmail.com	$2a$10$1/V1me2TTmKE3dt03YyXnOEyjql2PTCzBvVftSNjFwwzhD/jc/wze	admin
14	Salmon Moriey	salmon@gmail.com	$2a$10$d0TTg7QoMUirh/jgMBwUk.yUMFSa0YpM/XSdIkYOKIT35Bfcrz1ei	app_user
24	Friuger Terrence	fterrence@gmail.com	$2a$10$h1D3x5SlPhfUffXvaf74IOCxXf2Z12c8EhbbyPRWRuPq0zFRUX/4O	app_user
26	Josh Nathan	jnathan@gmail.com	$2a$10$T.e7SnfIvfJkoWP8WX9rL.CGqUIs8PwuHC5a1KnszUzEqhDAx8UB2	app_user
27	Josh Andrew	jandrew@gmail.com	$2a$10$Mn.dR68.zKDgIA8HY.mA/.WPl30ghaFncZFC6UonEOYi9EjKWocuC	app_user
65	Redrix Mouther	redrix@gmail.com	$2a$10$gViIHEkMZTf4Op9BTSPJcOLw4cbIai3Md1aBEEK3qlaa/.ssBSpjy	app_user
29	Jacob Stevens	jstevens@gmail.com	$2a$10$MuN4Xh0Pim06K45iGWV5kOHEhXzgHr4KgWIta/MQGy4XabpgoNZWK	app_user
30	Guteres Vilova	gvilova@gmail.com	$2a$10$eOjiRA51AWC9oKwzovMMceMdswPj/QfJoe9sEXDT/j3Ex836X2Y5a	app_user
31	Jeffrey Neyman	neyman@gmail.com	$2a$10$U.gA3EQ8TpLnK/Y99BfYQeZOo.PdsOOHsWburLqwYxCx.lsWHLSw6	app_user
32	Steven Seagull	seagull@gmail.com	$2a$10$/3fe1dTInap73JwIk3bRR.DTqIZBTfLdq6uOOLJO6f0E2nhh3xP4K	app_user
33	Polyester Pineman	pineman@gmail.com	$2a$10$nbQCq3Bt1CHobCYm/jM0ouIhwAMsQegqmEYG4VC5G7jPzth9D/HhS	app_user
34	Jeff Junior	junior@gmail.com	$2a$10$FFudqGP5D7eayoz0c3q5z.cgQkkwYhpoxGTt6YD4ETaiqSyCVcEd2	app_user
35	Logal Jimmy	ljimmy@gmail.com	$2a$10$/ez9ucXSKTRN16NFHm8MV.gF6zgU4DmRHTVmHRcAMzbMFBFISTleK	app_user
36	Derick Stulinger	dstulinger@gmail.com	$2a$10$3z0s4S8rx4A1Z.NLpaZNM..MjhDWJPP5IjX1LgUYC9G9cr67YhDGa	app_user
37	Primerus Drake	pdrake@gmail.com	$2a$10$ec9/momvwMG2GHtUhsKUDOJdHmlK498zWXU5i4CgGQ6CZrF2bI8pi	app_user
38	Frieza Ferry	ferry@gmail.com	$2a$10$xr982.Is5DiwhF.Se./EJeyGpUkeW3whyvcMGZRnqXYy444CFzEQO	app_user
39	Zuvo Yula	yulazuvo@gmail.com	$2a$10$9ofC8iHqNX5mYjiuY3eR.e9SoGoDhOGS3SldIrtd5d5Nh8BaUkrOK	app_user
40	Ticko Sever	sever@gmail.com	$2a$10$QErGXJN7MT3JVJfyLL74nuX6Mke0DajfguPn5wqxLyOwc16tJIGwO	app_user
41	Faderuno Red	fred@gmail.com	$2a$10$zfjipMlLURcSFaGnbD1qO.fygVUMcrgG9zl7cu47bPkdbnYzqJj8.	app_user
42	Jira Jacobs	jjacobs@gmail.com	$2a$10$YZ7CvuJ2dwnSl/AavJSTVuuccCywgw5EudgS3X9lUlou9DmU0ZlN2	app_user
43	Casey Edwards	edwards@gmail.com	$2a$10$LLCl.S.4grkBIhCMFBd39udR4P0u1p29eADAZeQaqGGm2FuGHcYBe	app_user
44	Quno Tiles	qtiles@gmail.com	$2a$10$St/enqnDfC64gpbVSHmj3e44EltAXC2kKml1r9mCJBIowlHe090HC	app_user
45	James Juro	jamesjuro@gmail.com	$2a$10$i/1eiA2dNP2hI3zOvHJWNuk.K2Dc1ZSi18nO01eW2Pmr4q75sUQc6	app_user
46	Namek Goku	namekgoku@gmail.com	$2a$10$V4KL5v8VVDwsXFf0GI4wPu5g1gQ0MrX.F6.Jt4CMg.UCypF1ap.De	app_user
47	Diva Yostova	dyostova@gmail.com	$2a$10$ZelHyicwkw.VHF42wtdOyOqMSdmwipL8env/EL0E8drvYeGRccnEm	app_user
48	Parry Summers	psummers@gmail.com	$2a$10$0pz/fhbiNHDnVjEIKQ7.BOFjgMVpOhbZ6b9q48Ouh2OAJIALE2ecS	app_user
49	Kipor Tisa	ktisa@gmail.com	$2a$10$MSMHqf9ZO4ZV4qSPmMcgVeVHWM4JQT/wC7yx31p1Pi97ARIgEC3I2	app_user
66	Karol Leeway	lee@gmail.com	$2a$10$BvdIF67grvLM4j0/34uNpeYBmVPloGilTGFIryQWaYboov/5mogsa	app_user
67	Kimmy Terons	teroni@gmail.com	$2a$10$TFTRSe4nKKf.HLBxHIFryu.ocJad1zVOb2tQqJiMuFGCb59VfhoIC	app_user
13	Teech Teer	teechteer@gmail.com	$2a$10$7./BfWcwyRlIYMsRqxPdjOBUKX2J4igk7PVbuchapKIr48n214mbC	app_user
54	test2	test2@gmail.com	$2a$10$AZzu2TjEyPPpgnZ44DUqGuVR.a12Wec8GXBvf7QEyj2lu5bfrVdZi	app_user
55	Frieza Koren	koren@gmail.com	$2a$10$.5fR1AykOOeM4WRl0MNMwuTLF./QjUk8nI0RAwUuhsrCemtZIaIhi	app_user
59	Telen	telen@gmail.com	$2a$10$zmDTdSviF1M9sIa613u0pegU5lw0rVMTHb5CJonhLxqgHqvmncAuy	app_user
60	Jeremy Doe	doe@gmail.com	$2a$10$8DnQIeADCVTmVuksCUHU2e2mvVHBCRwUr16P4mbDbHjwx.k1xrqdO	app_user
62	Mel Yoru	mel@gmail.com	$2a$10$d0OPsYw9W3zIjijWrYciHOi/FoJMaGyiPd/bzFkxq4NZUUtgecyR2	app_user
63	Terry Andrews	terry@gmail.com	$2a$10$zS7Ln0HuCOw7P/7Be4JCpOLgY.XfUp6vPqdL7L0euNePY8iaTa4ZG	app_user
0	Jerry Phisher	jerry@gmail.com	$2a$10$5UFtUcjWxVGLBfoQRoGKdOkQJc14kN6bNh/S5oXlluTZSdpTUT7wK	app_user
\.


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: saurav
--

SELECT pg_catalog.setval('public.users_user_id_seq', 67, true);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: saurav
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: saurav
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: users log_trigger; Type: TRIGGER; Schema: public; Owner: saurav
--

CREATE TRIGGER log_trigger AFTER INSERT ON public.users FOR EACH ROW EXECUTE FUNCTION public.log_user_creation();


--
-- Name: users users_role_fkey; Type: FK CONSTRAINT; Schema: public; Owner: saurav
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_role_fkey FOREIGN KEY (role) REFERENCES public.roles(name);


--
-- PostgreSQL database dump complete
--

\unrestrict FcrM9UfX2Qq2gnxDtbR0cntjSGWsJOxCTaJ72IgMHqa3yFxcETzeicyyjSkP7y2

