-- name: CreateFile :one
INSERT INTO files (
    name,
    email
) VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetFiles :many
SELECT * FROM files;

-- name: GetFileByID :one
SELECT * FROM files
WHERE id = $1;