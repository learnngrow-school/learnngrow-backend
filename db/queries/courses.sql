-- name: CreateCourse :one
INSERT INTO courses (title, description, price, year, category_id, subject_id)
	VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

