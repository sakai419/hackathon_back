CREATE TABLE messages (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_account_id CHAR(28) NOT NULL,
    recipient_account_id CHAR(28) NOT NULL,
    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_messages_sender_account_id FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_recipient_account_id FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_messages_sender_account_id (sender_account_id),
    INDEX idx_messages_recipient_account_id (recipient_account_id),
    INDEX idx_messages_created_at_account_id (created_at)
);