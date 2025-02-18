-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    -- The name of the feed (like "The Changelog, or "The Boot.dev Blog")
    name TEXT NOT NULL UNIQUE,
    -- The URL of the feed
    url TEXT NOT NULL UNIQUE,
    -- The ID of the user who added this feed
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;

-- psql "postgres://lancejeffers:@localhost:5432/gator"