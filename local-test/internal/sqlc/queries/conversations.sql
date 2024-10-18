-- name: GetConversationID :one
WITH existing_conversation AS (
    SELECT id FROM conversations
    WHERE (account1_id = $1 AND account2_id = $2)
       OR (account1_id = $2 AND account2_id = $1)
)
INSERT INTO conversations (account1_id, account2_id)
SELECT $1, $2
WHERE NOT EXISTS (SELECT 1 FROM existing_conversation)
RETURNING id
UNION ALL
SELECT id FROM existing_conversation
LIMIT 1;

-- name: GetConversations :many
SELECT
    c.id,
    c.account1_id,
    c.account2_id,
    c.last_message_time,
    m.content,
    m.sender_account_id,
    m.is_read
FROM
    conversations c
LEFT JOIN
    messages m ON c.last_message_id = m.id
WHERE
    c.account1_id = $1 OR c.account2_id = $1
ORDER BY
    c.last_message_time DESC;

-- name: GetUnreadConversationCount :one
SELECT COUNT(DISTINCT
    CASE
        WHEN c.account1_id = $1 THEN c.account2_id
        ELSE c.account1_id
    END) AS unread_account_count
FROM
    conversations c
JOIN
    messages m ON c.last_message_id = m.id
WHERE
    (c.account1_id = $1 OR c.account2_id = $1)
    AND m.is_read = FALSE
    AND m.sender_account_id <> $1;

-- name: UpdateLastMessage :exec
UPDATE conversations
SET last_message_id = $1, last_message_time = NOW()
WHERE id = $2;
