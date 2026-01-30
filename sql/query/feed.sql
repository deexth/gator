-- name: AddFeed :one
INSERT INTO feeds (
    id,
    name,
    url,
    user_id
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE url = $1 AND user_id = $2;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetUsersAndFeeds :many
SELECT
    f.name AS feed_name,
    f.url,
    u.name AS username
FROM feeds f
INNER JOIN users u
    ON f.user_id = u.id;


-- name: DeleteFeed :exec
DELETE FROM feeds WHERE url = $1 AND user_id = $2;

-- name: UpdateFeed :exec
UPDATE feeds
    SET updated_at = $1, name = $2, url = $3
    WHERE user_id = $4;
