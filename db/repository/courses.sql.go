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
INSERT INTO courses (title, description, price, year, category_id, subject_id, slug)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, title, description, price, year, category_id, subject_id, slug
`

type CreateCourseParams struct {
	Title       string
	Description pgtype.Text
	Price       int32
	Year        pgtype.Int2
	CategoryID  int32
	SubjectID   int32
	Slug        string
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRow(ctx, createCourse,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.Year,
		arg.CategoryID,
		arg.SubjectID,
		arg.Slug,
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
		&i.Slug,
	)
	return i, err
}

const getAllCourses = `-- name: GetAllCourses :many
SELECT c.id, c.title, c.description, c.price, c.year, c.category_id, c.subject_id, c.slug, ct.id, ct.title, sb.id, sb.title FROM courses AS C
LEFT JOIN categories AS CT ON category_id = CT.id
LEFT JOIN subjects AS SB ON subject_id = SB.id
ORDER BY C.id
`

type GetAllCoursesRow struct {
	Course   Course
	Category Category
	Subject  Subject
}

func (q *Queries) GetAllCourses(ctx context.Context) ([]GetAllCoursesRow, error) {
	rows, err := q.db.Query(ctx, getAllCourses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllCoursesRow
	for rows.Next() {
		var i GetAllCoursesRow
		if err := rows.Scan(
			&i.Course.ID,
			&i.Course.Title,
			&i.Course.Description,
			&i.Course.Price,
			&i.Course.Year,
			&i.Course.CategoryID,
			&i.Course.SubjectID,
			&i.Course.Slug,
			&i.Category.ID,
			&i.Category.Title,
			&i.Subject.ID,
			&i.Subject.Title,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCourseById = `-- name: GetCourseById :one
SELECT c.id, c.title, c.description, c.price, c.year, c.category_id, c.subject_id, c.slug, ct.id, ct.title, sb.id, sb.title FROM courses AS C
LEFT JOIN categories AS CT ON category_id = CT.id
LEFT JOIN subjects AS SB ON subject_id = SB.id
WHERE C.id = $1
LIMIT 1
`

type GetCourseByIdRow struct {
	Course   Course
	Category Category
	Subject  Subject
}

func (q *Queries) GetCourseById(ctx context.Context, id int32) (GetCourseByIdRow, error) {
	row := q.db.QueryRow(ctx, getCourseById, id)
	var i GetCourseByIdRow
	err := row.Scan(
		&i.Course.ID,
		&i.Course.Title,
		&i.Course.Description,
		&i.Course.Price,
		&i.Course.Year,
		&i.Course.CategoryID,
		&i.Course.SubjectID,
		&i.Course.Slug,
		&i.Category.ID,
		&i.Category.Title,
		&i.Subject.ID,
		&i.Subject.Title,
	)
	return i, err
}

const getCourseBySlug = `-- name: GetCourseBySlug :one
SELECT c.id, c.title, c.description, c.price, c.year, c.category_id, c.subject_id, c.slug, ct.id, ct.title, sb.id, sb.title FROM courses AS C
INNER JOIN categories AS CT ON category_id = CT.id
INNER JOIN subjects AS SB ON subject_id = SB.id
WHERE C.slug = $1
LIMIT 1
`

type GetCourseBySlugRow struct {
	Course   Course
	Category Category
	Subject  Subject
}

func (q *Queries) GetCourseBySlug(ctx context.Context, slug string) (GetCourseBySlugRow, error) {
	row := q.db.QueryRow(ctx, getCourseBySlug, slug)
	var i GetCourseBySlugRow
	err := row.Scan(
		&i.Course.ID,
		&i.Course.Title,
		&i.Course.Description,
		&i.Course.Price,
		&i.Course.Year,
		&i.Course.CategoryID,
		&i.Course.SubjectID,
		&i.Course.Slug,
		&i.Category.ID,
		&i.Category.Title,
		&i.Subject.ID,
		&i.Subject.Title,
	)
	return i, err
}
