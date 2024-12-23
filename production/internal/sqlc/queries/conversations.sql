-- name: GetConversationID :one
WITH existing_conversation AS (
    SELECT id FROM conversations c
    WHERE (c.account1_id = $1 AND c.account2_id = $2)
       OR (c.account1_id = $2 AND c.account2_id = $1)
),
inserted_conversation AS (
    INSERT INTO conversations (account1_id, account2_id)
    SELECT $1, $2
    WHERE NOT EXISTS (SELECT 1 FROM existing_conversation)
    RETURNING id
)
SELECT id FROM inserted_conversation
UNION ALL
SELECT id FROM existing_conversation
LIMIT 1;


-- name: GetConversationList :many
SELECT
    c.id,
    c.account1_id,
    c.account2_id,
    c.last_message_time,
    m.content,
    a.user_id AS sender_user_id,
    m.is_read
FROM
    conversations c
LEFT JOIN
    messages m ON c.last_message_id = m.id
LEFT JOIN
    accounts a ON m.sender_account_id = a.id
WHERE
    c.account1_id = $1 OR c.account2_id = $1
ORDER BY
    c.last_message_time DESC
LIMIT $2 OFFSET $3;

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
