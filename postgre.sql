-- Database: Todolist

-- DROP DATABASE IF EXISTS "Todolist";

-- VERİ TABANI OLUŞTURMAK İÇİN

CREATE DATABASE "Todolist"
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'Turkish_Turkey.1254'
    LC_CTYPE = 'Turkish_Turkey.1254'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


-- TABLOYU OLUŞTURMAK İÇİN
BEGIN;


CREATE TABLE IF NOT EXISTS public.tasks
(
    token text COLLATE pg_catalog."default",
    task text COLLATE pg_catalog."default",
    status boolean,
    usertoken text COLLATE pg_catalog."default"
);
END;
