-- name: CreateLabel :exec
INSERT INTO labels (tweet_id, label1, label2, label3)
VALUES (?, ?, ?, ?);

-- name: GetLabelsByTweetId :one
SELECT * FROM labels
WHERE tweet_id = ?;

-- name: UpdateLabels :exec
UPDATE labels
SET label1 = ?, label2 = ?, label3 = ?
WHERE tweet_id = ?;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE tweet_id = ?;

-- name: GetTweetsByLabel :many
SELECT t.* FROM tweets t
JOIN labels l ON t.id = l.tweet_id
WHERE l.label1 = ? OR l.label2 = ? OR l.label3 = ?
ORDER BY t.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTweetsWithoutLabels :many
SELECT t.* FROM tweets t
LEFT JOIN labels l ON t.id = l.tweet_id
WHERE l.tweet_id IS NULL
ORDER BY t.created_at DESC
LIMIT ? OFFSET ?;