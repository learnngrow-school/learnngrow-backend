-- name: GetAllStudents :many
SELECT *
FROM users AS U
WHERE U.is_teacher = false
  AND U.is_superuser = false
ORDER BY U.id;

