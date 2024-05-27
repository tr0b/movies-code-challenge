-- name: CreateMovie :one
INSERT INTO movies (
  title
) VALUES (
  $1
) RETURNING *;

-- name: GetMovie :one
SELECT * FROM movies
WHERE id = $1 LIMIT 1;

-- name: ListMovies :many
SELECT * FROM movies
ORDER BY id;

-- name: UpdateMovie :one
UPDATE movies
SET likes = likes + 1
WHERE id = $1
RETURNING *;
