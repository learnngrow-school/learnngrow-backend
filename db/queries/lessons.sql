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
,	ts
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

-- name: GetLessonsByUser :many
SELECT L.*
FROM
	lessons AS L
	INNER JOIN users AS S ON L.student_id = S.id
	INNER JOIN teachers AS T ON L.teacher_id = T.user_id
	INNER JOIN users AS TU ON T.user_id = TU.id
WHERE
	(TU.email = $1 OR S.email = $1) --TODO replace with slugs
	AND L.ts BETWEEN
		    date_trunc('week', CURRENT_DATE) + make_interval(weeks => $2)
		AND date_trunc('week', CURRENT_DATE) + make_interval(weeks => $2, days => 6)
ORDER BY L.ts;

