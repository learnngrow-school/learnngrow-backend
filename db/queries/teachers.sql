-- name: CreateTeacher :one
WITH new_user AS (
	INSERT INTO users (email, password, is_teacher, first_name, middle_name, last_name, slug)
		VALUES ($1, $2, true, $3, $4, $5, $6)
	RETURNING *
)
INSERT INTO teachers (user_id)
	VALUES ((SELECT id FROM new_user))
RETURNING user_id;

-- name: GetAllTeachers :many
SELECT sqlc.embed(U), sqlc.embed(T)
FROM teachers AS T
INNER JOIN users AS U ON T.user_id = U.id
ORDER BY T.user_id;

-- name: GetTeacherByID :one
SELECT sqlc.embed(U), sqlc.embed(T)
FROM teachers AS T
INNER JOIN users AS U ON T.user_id = U.id
WHERE T.user_id = $1
LIMIT 1;

-- name: GetTeacherBySlug :one
SELECT sqlc.embed(U), sqlc.embed(T)
FROM teachers AS T
INNER JOIN users AS U ON T.user_id = U.id
WHERE U.slug = $1
LIMIT 1;
