--
-- Name: book_stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.book_stocks (
    book_id character varying(36) NOT NULL,
    code character varying(50) NOT NULL,
    status character varying(50) NOT NULL,
    borrower_id character varying(36),
    borrowed_at timestamp(6) without time zone
);


--
-- Name: books; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.books (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    isbn character varying(100) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    cover_id character varying(36)
);


--
-- Name: charges; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.charges (
    id character varying(36) NOT NULL,
    journal_id character varying(36) NOT NULL,
    days_late integer DEFAULT 1 NOT NULL,
    daily_late_fee integer NOT NULL,
    total integer NOT NULL,
    user_id character varying(36) NOT NULL,
    created_at timestamp(6) without time zone
);


--
-- Name: customers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.customers (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    code character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone
);


--
-- Name: journals; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.journals (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    book_id character varying(36) NOT NULL,
    stock_code character varying(255) NOT NULL,
    customer_id character varying(36) NOT NULL,
    status character varying(50) NOT NULL,
    borrowed_at timestamp(6) without time zone NOT NULL,
    returned_at timestamp(6) without time zone,
    due_at timestamp(6) without time zone
);


--
-- Name: media; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.media (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    path text,
    created_at timestamp(6) without time zone NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id character varying(36) DEFAULT gen_random_uuid() NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL
);


--
-- Data for Name: book_stocks; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.book_stocks (book_id, code, status, borrower_id, borrowed_at) VALUES ('a3a5bb3d-4d89-4adf-9bae-82701247415b', 'BP-002', 'AVAILABLE', NULL, NULL);
INSERT INTO public.book_stocks (book_id, code, status, borrower_id, borrowed_at) VALUES ('a3a5bb3d-4d89-4adf-9bae-82701247415b', 'BP-001', 'AVAILABLE', NULL, NULL);


--
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.books (id, title, description, isbn, created_at, updated_at, deleted_at, cover_id) VALUES ('a3a5bb3d-4d89-4adf-9bae-82701247415b', 'Thing and grow rich', 'Book about thingking and critial', '0000001', '2024-07-11 15:16:37.218247', '2024-07-11 15:17:26.976471', NULL, NULL);
INSERT INTO public.books (id, title, description, isbn, created_at, updated_at, deleted_at, cover_id) VALUES ('22155604-07e7-4ad9-bcbd-4eccf36539db', 'Atomic Habit', 'The power of atomic habit', '10100129112', '2024-07-14 14:30:11.452375', NULL, NULL, 'e9e30364-6f3d-401a-8696-97f300c1f7c3');


--
-- Data for Name: charges; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.charges (id, journal_id, days_late, daily_late_fee, total, user_id, created_at) VALUES ('1a9a2565-2fcd-4c50-9e20-e9db1f37f749', '34dbcc4c-38f5-44e3-8130-22e20b12a026', 6, 5000, 30000, '4246fb58-ff45-4d2a-8946-93e541fc39fd', '2024-07-15 15:38:25.615904');
INSERT INTO public.charges (id, journal_id, days_late, daily_late_fee, total, user_id, created_at) VALUES ('7b3308d1-4f71-434f-976e-fccc7bef73df', '3f5b3c36-eba6-4faf-91a2-3c9f92070c7a', 1, 5000, 5000, '4246fb58-ff45-4d2a-8946-93e541fc39fd', '2024-07-15 15:40:04.678628');


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.customers (id, code, name, created_at, updated_at, deleted_at) VALUES ('efee29c8-7a47-49a7-a053-a296d07e8d1a', 'M-0001', 'Louis Fernando', '2024-07-10 21:54:06', '2024-07-10 21:54:09', NULL);
INSERT INTO public.customers (id, code, name, created_at, updated_at, deleted_at) VALUES ('99fee649-5ca3-46da-bba6-b90ff06bb4ad', 'MT-0004', 'Lofers', '2024-07-11 00:26:34.029608', '2024-07-11 00:38:49.537663', NULL);


--
-- Data for Name: journals; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.journals (id, book_id, stock_code, customer_id, status, borrowed_at, returned_at, due_at) VALUES ('34dbcc4c-38f5-44e3-8130-22e20b12a026', 'a3a5bb3d-4d89-4adf-9bae-82701247415b', 'BP-001', '99fee649-5ca3-46da-bba6-b90ff06bb4ad', 'COMPLETED', '2024-07-15 15:36:10.927671', '2024-07-15 15:38:25.614094', '2024-07-22 15:36:10.927669');
INSERT INTO public.journals (id, book_id, stock_code, customer_id, status, borrowed_at, returned_at, due_at) VALUES ('804d396f-3537-4116-b895-ffe3724a5a20', 'a3a5bb3d-4d89-4adf-9bae-82701247415b', 'BP-001', '99fee649-5ca3-46da-bba6-b90ff06bb4ad', 'COMPLETED', '2024-07-15 15:39:17.611751', '2024-07-15 15:39:27.282173', '2024-07-22 15:39:17.61175');
INSERT INTO public.journals (id, book_id, stock_code, customer_id, status, borrowed_at, returned_at, due_at) VALUES ('3f5b3c36-eba6-4faf-91a2-3c9f92070c7a', 'a3a5bb3d-4d89-4adf-9bae-82701247415b', 'BP-001', '99fee649-5ca3-46da-bba6-b90ff06bb4ad', 'COMPLETED', '2024-07-15 15:39:37.84623', '2024-07-15 15:40:04.677834', '2024-07-14 15:39:37.846');


--
-- Data for Name: media; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.media (id, path, created_at) VALUES ('e9e30364-6f3d-401a-8696-97f300c1f7c3', '52d39a55-580a-4f4a-a889-4dcc6bb6f41e9786020633176_.Atomic_Habit.jpg', '2024-07-14 14:28:11.010225');


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO public.users (id, email, password) VALUES ('4246fb58-ff45-4d2a-8946-93e541fc39fd', 'admin@gmail.com', '$2a$12$HvB9/JsVFO.S1WlNNFXHaO9dEKBvqdDmPAoP33zlBaUyGNa9pxK4G');


--
-- Name: book_stocks book_stocks_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.book_stocks
    ADD CONSTRAINT book_stocks_pk PRIMARY KEY (code);


--
-- Name: books books_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pk PRIMARY KEY (id);


--
-- Name: charges charges_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.charges
    ADD CONSTRAINT charges_pk PRIMARY KEY (id);


--
-- Name: customers customers_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pk PRIMARY KEY (id);


--
-- Name: journals journals_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.journals
    ADD CONSTRAINT journals_pk PRIMARY KEY (id);


--
-- Name: media media_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.media
    ADD CONSTRAINT media_pk PRIMARY KEY (id);


--
-- Name: users users_pk; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pk PRIMARY KEY (id);