-- name: CreateCourse :one
INSERT INTO courses (title, description, price, year, category_id, subject_id)
	VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllCourses :many
SELECT sqlc.embed(C), sqlc.embed(CT), sqlc.embed(SB) FROM courses AS C
LEFT JOIN categories AS CT ON category_id = CT.id
LEFT JOIN subjects AS SB ON subject_id = SB.id
ORDER BY C.id;

