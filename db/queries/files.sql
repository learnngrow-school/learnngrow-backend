-- name: UploadFile :exec
INSERT INTO files (slug, fname, fsize, fdata)
VALUES ($1, $2, $3, $4);

-- name: GetFile :one
SELECT * FROM files
WHERE slug == $1;

-- -- name: GetFileInfo :one
-- SELECT (slug, fname, fsize) FROM files
-- WHERE slug == $1;
