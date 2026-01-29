-- name: CreateUser :one
INSERT INTO users (id, name)
VALUES (
    $1,
    $2
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE name = $1;

-- name: ResetDb :exec
DELETE FROM users;

-- name: GetUsers :many
SELECT * FROM users;
