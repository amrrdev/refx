-- name: CreateShortCode :one
INSERT INTO urls (short_code, long_url)
VALUES ($1, $2)
RETURNING short_code, long_url;

-- name: GetLongUrl :one
SELECT long_url FROM urls WHERE short_code = $1;

-- name: GetByLongUrl :one
SELECT short_code FROM urls WHERE long_url = $1;
