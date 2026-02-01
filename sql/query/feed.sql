-- name: AddFeed :one
WITH new_feed AS (
    INSERT INTO feeds (
        id,
        name,
        url,
        user_id
    )
    SELECT
        $1,
        $2,
        $3,
        u.id
    FROM users u
    WHERE u.name = $4
    RETURNING *
),
new_follow AS (
    INSERT INTO feed_follows (
        id,
        user_id,
        feed_id
    )
    SELECT
        $5,
        new_feed.user_id,
        new_feed.id
    FROM new_feed
    ON CONFLICT (user_id, feed_id)
    DO UPDATE SET user_id = EXCLUDED.user_id
    RETURNING id AS feed_follow_id
)
SELECT
    new_feed.*,
    new_follow.feed_follow_id
FROM new_feed
LEFT JOIN new_follow ON true;

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
