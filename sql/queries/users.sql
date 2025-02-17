-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    -- $1, $2, $3, and $4 are parameters that we'll be able to pass into the query in our Go code. 
    $1,
    $2,
    $3,
    $4
)
RETURNING *;


-- name: DeleteUsers :exec
DELETE FROM users;


-- name: GetUser :one
SELECT * FROM users WHERE name = $1;