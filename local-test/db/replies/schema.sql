CREATE TABLE replies (
    reply_id BIGINT UNSIGNED NOT NULL,
    original_tweet_id BIGINT UNSIGNED NOT NULL,
    parent_reply_id BIGINT UNSIGNED,
    replying_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (reply_id),
    CONSTRAINT fk_replies_reply_id FOREIGN KEY (reply_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_replies_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_replies_parent_reply_id FOREIGN KEY (parent_reply_id)
        REFERENCES replies(id) ON DELETE SET NULL,
    CONSTRAINT fk_replies_replying_account_id FOREIGN KEY (replying_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_replies_original_tweet_id (original_tweet_id),
    INDEX idx_replies_parent_reply_id (parent_reply_id),
    INDEX idx_replies_replying_account_id (replying_account_id),
    INDEX idx_replies_created_at (created_at)
);