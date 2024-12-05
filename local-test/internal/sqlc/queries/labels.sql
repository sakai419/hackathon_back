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

-- name: GetRecentLabels :many
WITH all_labels AS (
    SELECT label1 AS label FROM labels
    UNION ALL
    SELECT label2 AS label FROM labels
    UNION ALL
    SELECT label3 AS label FROM labels
)
SELECT
    label,
    COUNT(*) AS label_count
FROM
    all_labels
WHERE
    label IS NOT NULL
GROUP BY
    label
ORDER BY
    label_count DESC
LIMIT $1;

-- name: GetLikedTweetLabelsCount :many
WITH liked_tweets AS (
    SELECT l.original_tweet_id
    FROM likes l
    WHERE l.liking_account_id = $1
    ORDER BY l.created_at DESC
    LIMIT 100
)
SELECT
    label,
    COUNT(*) AS label_count
FROM (
    SELECT label1 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
    UNION ALL
    SELECT label2 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
    UNION ALL
    SELECT label3 AS label FROM labels WHERE tweet_id IN (SELECT original_tweet_id FROM liked_tweets)
) labels_combined
GROUP BY label;
