CREATE TABLE tweets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    account_id CHAR(28) NOT NULL,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    content TEXT DEFAULT NULL,
    code TEXT DEFAULT NULL,
    likes_count INT UNSIGNED NOT NULL DEFAULT 0,
    replies_count INT UNSIGNED NOT NULL DEFAULT 0,
    retweets_count INT UNSIGNED NOT NULL DEFAULT 0,
    is_retweet BOOLEAN NOT NULL DEFAULT FALSE,
    is_reply BOOLEAN NOT NULL DEFAULT FALSE,
    is_quote BOOLEAN NOT NULL DEFAULT FALSE,
    engagement_score INT UNSIGNED NOT NULL DEFAULT 0,
    media JSON,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_tweets_account_id FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_tweets_account_id (account_id),
    INDEX idx_tweets_engagement_score (engagement_score),
    INDEX idx_tweets_created_at (created_at),
    INDEX idx_tweets_type (is_retweet, is_reply, is_quote),
    CHECK (JSON_VALID(media))
);