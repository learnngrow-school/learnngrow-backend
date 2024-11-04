-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    email = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, PASSWORD, is_teacher, first_name, middle_name, last_name)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

