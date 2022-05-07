-- User

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1)
RETURNING id, name;

-- name: UpdateUser :execrows
UPDATE users SET name = $2
WHERE id = $1;

-- name: DeleteUser :execrows
DELETE FROM users WHERE id = $1;