-- name: CreateLabel :exec
INSERT INTO labels (tweet_id, label1, label2, label3)
VALUES ($1, $2, $3, $4);

-- name: GetLabelsByTweetID :one
SELECT * FROM labels
WHERE tweet_id = $1;

-- name: UpdateLabels :exec
UPDATE labels
SET label1 = $1, label2 = $2, label3 = $3
WHERE tweet_id = $4;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE tweet_id = $1;

-- name: GetTweetsByLabel :many
SELECT t.* FROM tweets t
JOIN labels l ON t.id = l.tweet_id
WHERE l.label1 = $1 OR l.label2 = $2 OR l.label3 = $3
ORDER BY t.created_at DESC
LIMIT $4 OFFSET $5;

-- name: GetTweetsWithoutLabels :many
SELECT t.* FROM tweets t
LEFT JOIN labels l ON t.id = l.tweet_id
WHERE l.tweet_id IS NULL
ORDER BY t.created_at DESC
LIMIT $1 OFFSET $2;