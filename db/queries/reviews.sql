-- name: CreateReview :one
INSERT INTO school_reviews (details, author_name)
	VALUES ($1, $2)
RETURNING *;

