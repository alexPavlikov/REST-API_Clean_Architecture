SELECT 
    b.id, b.name,
    array((SELECT ba.author_id FROM book_authors ba WHERE ba.book_id = b.id)) 
    AS authors
FROM book b
GROUP BY b.id, b.name;
 