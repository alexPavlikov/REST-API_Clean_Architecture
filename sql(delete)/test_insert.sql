
INSERT INTO public.author (id, name) VALUES ('08d2ef86-edf0-4361-8742-11fa9acea942', 'Alex');
INSERT INTO public.author (id, name) VALUES ('3f749155-1f75-408d-a0a4-37dbff42c1cb', 'Danil');
INSERT INTO public.author (id, name) VALUES ('78d19e5b-1280-4f9a-8e6a-176eed7086d5', 'Vasya');


INSERT INTO public.book (id, name) VALUES ('5dc2d424-9c62-4259-96d1-da670dd52761', 'Book1');
INSERT INTO public.book (id, name) VALUES ('b2c048d6-d939-451a-8d8b-559db158897e', 'Book2');


INSERT INTO public.book_authors (book_id, author_id) VALUES ('5dc2d424-9c62-4259-96d1-da670dd52761', '08d2ef86-edf0-4361-8742-11fa9acea942');
INSERT INTO public.book_authors (book_id, author_id) VALUES ('5dc2d424-9c62-4259-96d1-da670dd52761', '3f749155-1f75-408d-a0a4-37dbff42c1cb');
INSERT INTO public.book_authors (book_id, author_id) VALUES ('b2c048d6-d939-451a-8d8b-559db158897e', '78d19e5b-1280-4f9a-8e6a-176eed7086d5');
