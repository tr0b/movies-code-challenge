// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: movie.sql

package db

import (
	"context"
)

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (
  title
) VALUES (
  $1
) RETURNING id, title, likes, created_at
`

func (q *Queries) CreateMovie(ctx context.Context, title string) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie, title)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Likes,
		&i.CreatedAt,
	)
	return i, err
}

const getMovie = `-- name: GetMovie :one
SELECT id, title, likes, created_at FROM movies
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Likes,
		&i.CreatedAt,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
SELECT id, title, likes, created_at FROM movies
ORDER BY id
`

func (q *Queries) ListMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, listMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Movie{}
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Likes,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE movies
SET likes = likes + 1
WHERE id = $1
RETURNING id, title, likes, created_at
`

func (q *Queries) UpdateMovie(ctx context.Context, id int64) (Movie, error) {
	row := q.db.QueryRowContext(ctx, updateMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Likes,
		&i.CreatedAt,
	)
	return i, err
}
