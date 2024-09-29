-- name: CreateHashtag :exec
INSERT INTO hashtags (tag) VALUES (?)
ON DUPLICATE KEY UPDATE id = LAST_INSERT_ID(id);

-- name: GetHashtagById :one
SELECT * FROM hashtags WHERE id = ?;

-- name: GetHashtagByTag :one
SELECT * FROM hashtags WHERE tag = ?;

-- name: DeleteHashtag :exec
DELETE FROM hashtags WHERE id = ?;

-- name: GetAllHashtags :many
SELECT * FROM hashtags
ORDER BY tag ASC
LIMIT ? OFFSET ?;

-- name: SearchHashtags :many
SELECT * FROM hashtags
WHERE tag LIKE CONCAT('%', ?, '%')
ORDER BY tag ASC
LIMIT ?;

-- name: GetHashtagCount :one
SELECT COUNT(*) FROM hashtags;

-- name: UpdateHashtagCreatedAt :exec
UPDATE hashtags
SET created_at = CURRENT_TIMESTAMP
WHERE id = ?;