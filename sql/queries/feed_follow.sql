-- name: CreateFeedFollow :many
WITH inserted_feeds_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
) SELECT
    inserted_feeds_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feeds_follow
INNER JOIN feeds ON feeds.ID = inserted_feeds_follow.feed_id
INNER JOIN users ON users.ID = inserted_feeds_follow.user_id;

-- name: GetFeedFollowsForUser :many
SELECT
  feed_follows.*,
  feeds.name AS feed_name,
  users.name AS user_name
FROM feed_follows
JOIN feeds AS feeds ON feeds.id = feed_follows.feed_id
JOIN users AS users ON users.id = feed_follows.user_id
WHERE feed_follows.user_id = $1;
