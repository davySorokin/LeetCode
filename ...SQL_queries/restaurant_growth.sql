WITH DailyTotals AS (
    SELECT
        visited_on,
        SUM(amount) AS amount
    FROM
        Customer
    GROUP BY
        visited_on
),
RunningTotals AS (
    SELECT
        visited_on,
        SUM(amount) OVER (ORDER BY visited_on ROWS BETWEEN 6 PRECEDING AND CURRENT ROW) AS amount,
        ROUND(AVG(amount) OVER (ORDER BY visited_on ROWS BETWEEN 6 PRECEDING AND CURRENT ROW), 2) AS average_amount
    FROM
        DailyTotals
)
SELECT
    visited_on,
    amount,
    average_amount
FROM
    RunningTotals
WHERE
    visited_on >= (SELECT MIN(visited_on) FROM DailyTotals) + INTERVAL 6 DAY
ORDER BY
    visited_on;
