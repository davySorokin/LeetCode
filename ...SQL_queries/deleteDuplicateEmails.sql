DELETE p1
FROM Person p1
LEFT JOIN (
    SELECT MIN(id) AS min_id
    FROM Person
    GROUP BY email
) p2 ON p1.id = p2.min_id
WHERE p2.min_id IS NULL;
