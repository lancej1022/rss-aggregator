-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    -- $1, $2, $3, and $4 are parameters that we'll be able to pass into the query in our Go code. 
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

