CREATE TABLE labels (
    tweet_id BIGINT UNSIGNED PRIMARY KEY,
    label1 VARCHAR(50) NOT NULL,
    label2 VARCHAR(50),
    label3 VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_labels_tweet_id FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE
);