CREATE TABLE follows (
    follower_account_id CHAR(28) NOT NULL,
    following_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_account_id, following_account_id),
    CONSTRAINT fk_follows_follower_account_id FOREIGN KEY (follower_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_follows_following_account_id FOREIGN KEY (following_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_follows_created_at (created_at)
);