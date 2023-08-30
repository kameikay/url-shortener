-- name: InsertUrl :exec

INSERT INTO
    urls (id, url, code, created_at)
VALUES (DEFAULT, $1, $2, DEFAULT);

-- name: GetUrlByCode :one

SELECT id, url FROM urls WHERE code = $1;