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


-- Article

-- name: GetArticle :one
SELECT
	article.id id,
	article.name article_name,
	users.name user_name
FROM
	article
	JOIN users ON article.user_id = users.id
WHERE
	article.id = $1
LIMIT 1;

-- UserArticle

-- name: GetUserArticles :many
SELECT
	users.id AS user_id,
	users.name AS user_name,
	article.id AS article_id,
	article.name AS article_name
FROM
	users
	LEFT JOIN article ON users.id = article.user_id
ORDER BY
	users.id;