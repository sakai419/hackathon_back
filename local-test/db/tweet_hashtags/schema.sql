CREATE TABLE tweet_hashtags (
    tweet_id BIGINT UNSIGNED NOT NULL,
    hashtag_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (tweet_id, hashtag_id),
    CONSTRAINT fk_tweet_hashtags_twee_id FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_tweet_hashtags_hashtag_id FOREIGN KEY (hashtag_id)
        REFERENCES hashtags(id) ON DELETE CASCADE,
    INDEX idx_tweet_hashtags_hashtag_id (hashtag_id)
);