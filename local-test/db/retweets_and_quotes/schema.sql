CREATE TABLE retweets_and_quotes (
    retweet_id BIGINT UNSIGNED NOT NULL,
    retweeting_account_id CHAR(28) NOT NULL,
    original_tweet_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (retweet_id),
    CONSTRAINT fk_retweets_and_quotes_retweet_id FOREIGN KEY (retweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_retweets_and_quotes_retweeting_account_id FOREIGN KEY (retweeting_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_retweets_and_quotes_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    INDEX idx_retweets_retweet_id (retweet_id)
);