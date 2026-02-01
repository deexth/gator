-- name: CreateFeedFollow :one
WITH ff AS (
    INSERT INTO feed_follows (
        id,
        user_id,
        feed_id
    )
    SELECT
        $1,          -- feed_follows.id
        u.id,        -- user_id resolved from username
        f.id         -- feed_id resolved from feed URL
    FROM users u
    JOIN feeds f
        ON f.url = $3
    WHERE u.name = $2
    ON CONFLICT (user_id, feed_id)
    DO UPDATE SET user_id = EXCLUDED.user_id
    RETURNING *
)
SELECT
    ff.*,
    u.name AS user_name,
    f.name AS feed_name
FROM ff
JOIN users u
    ON u.id = ff.user_id
JOIN feeds f
    ON f.id = ff.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT f.name
FROM feed_follows ff
INNER JOIN users u
    ON u.id = ff.user_id
INNER JOIN feeds f
    ON f.id = ff.feed_id
WHERE u.name = $1
ORDER BY ff.updated_at DESC;

-- name: UnfollowFeed :exec
DELETE FROM feed_follows ff 
USING users u, feeds f 
WHERE ff.user_id = u.id
AND ff.feed_id = f.id
AND u.name = $1 AND f.url = $2;
