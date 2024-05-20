-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, surname, email, password)
VALUES ($1, $2, $3, $4);

-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE($2, name),
    surname = COALESCE($3, surname),
    email = COALESCE($4, email),
    password = COALESCE($5, password)
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;