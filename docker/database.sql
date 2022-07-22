CREATE TABLE IF NOT EXISTS public.users
(
    user_id integer NOT NULL,
    username text COLLATE pg_catalog."default",
    language text COLLATE pg_catalog."default",
    first_name text COLLATE pg_catalog."default",
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to admin;