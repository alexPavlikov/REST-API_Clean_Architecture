CREATE TABLE public.author (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL
);

CREATE TABLE public.book (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(100) NOT NULL 
);

CREATE TABLE public.book_authors (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	book_id UUID NOT NULL,
	author_id UUID NOT NULL,
	CONSTRAINT book_fk FOREIGN KEY (book_id) REFERENCES public.book(id), 
	CONSTRAINT author_fk FOREIGN KEY (author_id) REFERENCES public.author(id), 
	CONSTRAINT book_author_unique UNIQUE (book_id, author_id)
);

CREATE TABLE public.user (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	firstname VARCHAR(30) NOT NULL,
	lastname VARCHAR(50) NOT NULL,
	age smallint NOT NULL,
	email VARCHAR(50) NOT NULL,
	password_hash VARCHAR(100) NOT NULL
)

CREATE TABLE public.worker (
	id  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	firstname VARCHAR(30) NOT NULL,
	lastname VARCHAR(50) NOT NULL, 
	age smallint NOT NULL,
	experiens smallint NOT NULL,
	number VARCHAR(30) NOT NULL,
	address VARCHAR(50) NOT NULL,
	email VARCHAR(50) NOT NULL,
	password_hash VARCHAR(100) NOT NULL
);
