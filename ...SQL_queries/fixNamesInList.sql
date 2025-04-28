-- Format names with first letter uppercase and rest lowercase, ordered by user_id
SELECT 
    user_id,
    CONCAT(
        UPPER(LEFT(name, 1)),
        LOWER(SUBSTRING(name, 2))
    ) AS name
FROM Users ORDER BY user_id;
