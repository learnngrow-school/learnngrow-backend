// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: courses.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO courses (title, description, price, year, category_id, subject_id)
	VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, title, description, price, year, category_id, subject_id
`

type CreateCourseParams struct {
	Title       string
	Description pgtype.Text
	Price       int32
	Year        pgtype.Int2
	CategoryID  int32
	SubjectID   int32
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRow(ctx, createCourse,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.Year,
		arg.CategoryID,
		arg.SubjectID,
	)
	var i Course
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.Year,
		&i.CategoryID,
		&i.SubjectID,
	)
	return i, err
}
