-- name: CreateTeacher :one
WITH new_user AS (
	INSERT INTO users (email, password, is_teacher, first_name, middle_name, last_name)
		VALUES ($1, $2, true, $3, $4, $5)
	RETURNING *
)
INSERT INTO teachers (user_id)
	VALUES ((SELECT id FROM new_user))
RETURNING user_id;

