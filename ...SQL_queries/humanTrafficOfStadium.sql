SELECT id, visit_date, people
FROM (
    SELECT s1.id, s1.visit_date, s1.people,
           IF(s1.people >= 100, 1, 0) AS valid,
           ROW_NUMBER() OVER (ORDER BY s1.id) -
           ROW_NUMBER() OVER (PARTITION BY s1.people >= 100 ORDER BY s1.id) AS grp
    FROM Stadium s1
) t
WHERE valid = 1
  AND grp IN (
      SELECT grp
      FROM (
          SELECT grp, COUNT(*) as cnt
          FROM (
              SELECT s1.id, s1.people,
                     IF(s1.people >= 100, 1, 0) AS valid,
                     ROW_NUMBER() OVER (ORDER BY s1.id) -
                     ROW_NUMBER() OVER (PARTITION BY s1.people >= 100 ORDER BY s1.id) AS grp
              FROM Stadium s1
          ) x
          WHERE valid = 1
          GROUP BY grp
          HAVING COUNT(*) >= 3
      ) y
)
ORDER BY visit_date ASC;
