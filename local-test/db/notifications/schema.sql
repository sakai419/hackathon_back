    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_account_id CHAR(28),
    recipient_account_id CHAR(28) NOT NULL,
    type VARCHAR(50) NOT NULL,
    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_notifications_sender_account_id FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE SET NULL,
    CONSTRAINT fk_notifications_recipient_account_id FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_notifications_recipient_account_id (recipient_account_id),
    INDEX idx_notifications_created_at (created_at)
);