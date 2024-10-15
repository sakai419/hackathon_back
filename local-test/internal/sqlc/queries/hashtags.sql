-- name: CreateHashtag :exec
INSERT INTO hashtags (tag) VALUES ($1)
ON CONFLICT (tag) DO NOTHING;

-- name: GetHashtagByID :one
SELECT * FROM hashtags WHERE id = $1;

-- name: GetHashtagByTag :one
SELECT * FROM hashtags WHERE tag = $1;

-- name: DeleteHashtag :exec
DELETE FROM hashtags WHERE id = $1;

-- name: GetAllHashtags :many
SELECT * FROM hashtags
ORDER BY tag ASC
LIMIT $1 OFFSET $2;

-- name: SearchHashtags :many
SELECT * FROM hashtags
WHERE tag LIKE CONCAT('%', $1, '%')
ORDER BY tag ASC
LIMIT $2;

-- name: GetHashtagCount :one
SELECT COUNT(*) FROM hashtags;

-- name: UpdateHashtagCreatedAt :exec
UPDATE hashtags
SET created_at = CURRENT_TIMESTAMP
WHERE id = $1;