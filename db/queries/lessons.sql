-- name: CreateLesson :one
WITH student AS (
	SELECT *
	FROM users AS U
	WHERE U.slug = $1
	LIMIT 1
), teacher AS (
	SELECT *
	FROM
		teachers AS T
		INNER JOIN users AS U ON U.id = T.user_id
	WHERE U.slug = $2
	LIMIT 1
)
INSERT INTO lessons (
	student_id
,	teacher_id
,	timestamp_m
,	teacher_notes
,	homework
) VALUES (
	(SELECT id FROM student)
,	(SELECT id FROM teacher)
,	$3
,	$4
,	$5
)
RETURNING *;

-- name: GetLessonsByTeacher :many
SELECT L.*
FROM lessons AS L
INNER JOIN teachers AS T ON L.teacher_id = T.user_id
INNER JOIN users AS U ON T.user_id = U.id
WHERE U.email = $1; --TODO replace with slugs

