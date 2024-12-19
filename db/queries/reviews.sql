-- name: CreateSchoolReview :one
INSERT INTO school_reviews (details, author_name)
	VALUES ($1, $2)
RETURNING *;

-- name: GetAllSchoolReviews :many
SELECT *
FROM school_reviews
ORDER BY id;

-- name: CreateTeacherReview :one
INSERT INTO teacher_reviews (details, author_name, teacher_id, rating)
	VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllTeacherReviews :many
SELECT *
FROM teacher_reviews
WHERE teacher_id = $1
ORDER BY id;

-- name: CreateCourseReview :one
INSERT INTO course_reviews (details, author_name, course_id, rating)
	VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAllCourseReviews :many
SELECT *
FROM course_reviews
WHERE course_id = $1
ORDER BY id;

