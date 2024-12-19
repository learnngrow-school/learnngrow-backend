-- name: CreateSubject :one
INSERT INTO subjects (title)
	VALUES ($1)
RETURNING *;

-- name: GetAllSubjects :many
SELECT *
FROM subjects;

-- name: CreateCategory :one
INSERT INTO categories (title)
	VALUES ($1)
RETURNING *;

-- name: GetAllCategories :many
SELECT *
FROM categories;

