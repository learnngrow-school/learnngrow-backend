-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    email = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, PASSWORD, is_teacher, first_name, middle_name, last_name, slug)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING
    *;

-- name: CreateSuperuser :exec
INSERT INTO users (email, PASSWORD, is_teacher, is_superuser, first_name, last_name, slug)
    VALUES ('admin', $1, FALSE, TRUE, 'Admin', 'Admin', 'Admin')
ON CONFLICT(email) DO UPDATE SET password = $1;

