-- name: CreateSubject :exec
INSERT INTO subjects (title)
VALUES ($1);

-- name: GetSubjects :many
SELECT * FROM subjects;

