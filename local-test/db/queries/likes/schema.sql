CREATE TABLE likes (
    liking_account_id CHAR(28) NOT NULL,
    original_tweet_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (liking_account_id, original_tweet_id),
    CONSTRAINT fk_likes_liking_account_id FOREIGN KEY (liking_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_likes_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    INDEX idx_likes_original_tweet_id (original_tweet_id)
);