-- name: GetHashtagIDs :many
WITH new_hashtags AS (
    INSERT INTO hashtags (tag)
    SELECT UNNEST(@tags::VARCHAR(30)[])
    ON CONFLICT (tag) DO NOTHING
    RETURNING id, tag
)
SELECT id FROM new_hashtags
UNION
SELECT id FROM hashtags WHERE tag = ANY(@tags::VARCHAR(30)[])
ORDER BY id;