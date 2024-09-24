DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
    id CHAR(28) PRIMARY KEY,
    user_id VARCHAR(28) UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY idx_accounts_user_name (user_name)
);

DROP TABLE IF EXISTS profiles;

CREATE TABLE profiles (
    account_id CHAR(28) PRIMARY KEY,
    bio TEXT,
    profile_image_url VARCHAR(2083),
    banner_image_url VARCHAR(2083),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_profiles_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS settings;

CREATE TABLE settings (
    account_id CHAR(28) PRIMARY KEY,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_settings_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS interests;

CREATE TABLE interests (
    account_id CHAR(28) PRIMARY KEY,
    news_score TINYINT UNSIGNED DEFAULT 0,
    politics_score TINYINT UNSIGNED DEFAULT 0,
    economics_score TINYINT UNSIGNED DEFAULT 0,
    health_score TINYINT UNSIGNED DEFAULT 0,
    sports_score TINYINT UNSIGNED DEFAULT 0,
    entertainment_score TINYINT UNSIGNED DEFAULT 0,
    art_score TINYINT UNSIGNED DEFAULT 0,
    cooking_score TINYINT UNSIGNED DEFAULT 0,
    travel_score TINYINT UNSIGNED DEFAULT 0,
    fashion_score TINYINT UNSIGNED DEFAULT 0,
    beauty_score TINYINT UNSIGNED DEFAULT 0,
    pets_score TINYINT UNSIGNED DEFAULT 0,
    parenting_score TINYINT UNSIGNED DEFAULT 0,
    education_score TINYINT UNSIGNED DEFAULT 0,
    environment_score TINYINT UNSIGNED DEFAULT 0,
    climate_score TINYINT UNSIGNED DEFAULT 0,
    space_score TINYINT UNSIGNED DEFAULT 0,
    mental_health_score TINYINT UNSIGNED DEFAULT 0,
    fitness_score TINYINT UNSIGNED DEFAULT 0,
    reading_score TINYINT UNSIGNED DEFAULT 0,
    history_score TINYINT UNSIGNED DEFAULT 0,
    philosophy_score TINYINT UNSIGNED DEFAULT 0,
    religion_score TINYINT UNSIGNED DEFAULT 0,
    culture_score TINYINT UNSIGNED DEFAULT 0,
    volunteering_score TINYINT UNSIGNED DEFAULT 0,
    social_issues_score TINYINT UNSIGNED DEFAULT 0,
    law_score TINYINT UNSIGNED DEFAULT 0,
    taxes_score TINYINT UNSIGNED DEFAULT 0,
    investment_score TINYINT UNSIGNED DEFAULT 0,
    real_estate_score TINYINT UNSIGNED DEFAULT 0,
    diy_score TINYINT UNSIGNED DEFAULT 0,
    gardening_score TINYINT UNSIGNED DEFAULT 0,
    interior_design_score TINYINT UNSIGNED DEFAULT 0,
    automotive_score TINYINT UNSIGNED DEFAULT 0,
    gaming_score TINYINT UNSIGNED DEFAULT 0,
    anime_manga_score TINYINT UNSIGNED DEFAULT 0,
    creative_works_score TINYINT UNSIGNED DEFAULT 0,
    photography_video_score TINYINT UNSIGNED DEFAULT 0,
    media_score TINYINT UNSIGNED DEFAULT 0,
    marketing_score TINYINT UNSIGNED DEFAULT 0,
    branding_score TINYINT UNSIGNED DEFAULT 0,
    entrepreneurship_score TINYINT UNSIGNED DEFAULT 0,
    remote_work_score TINYINT UNSIGNED DEFAULT 0,
    data_science_score TINYINT UNSIGNED DEFAULT 0,
    iot_score TINYINT UNSIGNED DEFAULT 0,
    robotics_engineering_score TINYINT UNSIGNED DEFAULT 0,
    biotechnology_score TINYINT UNSIGNED DEFAULT 0,
    nanotechnology_score TINYINT UNSIGNED DEFAULT 0,
    energy_technology_score TINYINT UNSIGNED DEFAULT 0,
    archaeology_score TINYINT UNSIGNED DEFAULT 0,
    psychology_score TINYINT UNSIGNED DEFAULT 0,
    sociology_score TINYINT UNSIGNED DEFAULT 0,
    anthropology_score TINYINT UNSIGNED DEFAULT 0,
    geography_score TINYINT UNSIGNED DEFAULT 0,
    geology_score TINYINT UNSIGNED DEFAULT 0,
    meteorology_score TINYINT UNSIGNED DEFAULT 0,
    disaster_emergency_management_score TINYINT UNSIGNED DEFAULT 0,
    urban_planning_score TINYINT UNSIGNED DEFAULT 0,
    architecture_score TINYINT UNSIGNED DEFAULT 0,
    agriculture_score TINYINT UNSIGNED DEFAULT 0,
    nutrition_science_score TINYINT UNSIGNED DEFAULT 0,
    sleep_science_score TINYINT UNSIGNED DEFAULT 0,
    productivity_score TINYINT UNSIGNED DEFAULT 0,
    leadership_score TINYINT UNSIGNED DEFAULT 0,
    international_relations_score TINYINT UNSIGNED DEFAULT 0,
    future_predictions_score TINYINT UNSIGNED DEFAULT 0,
    events_score TINYINT UNSIGNED DEFAULT 0,
    community_score TINYINT UNSIGNED DEFAULT 0,
    trends_score TINYINT UNSIGNED DEFAULT 0,
    lifestyle_score TINYINT UNSIGNED DEFAULT 0,
    software_development_score TINYINT UNSIGNED DEFAULT 0,
    programming_languages_score TINYINT UNSIGNED DEFAULT 0,
    web_development_score TINYINT UNSIGNED DEFAULT 0,
    mobile_app_development_score TINYINT UNSIGNED DEFAULT 0,
    debugging_techniques_score TINYINT UNSIGNED DEFAULT 0,
    algorithms_mathematics_score TINYINT UNSIGNED DEFAULT 0,
    database_design_score TINYINT UNSIGNED DEFAULT 0,
    cloud_computing_score TINYINT UNSIGNED DEFAULT 0,
    server_management_score TINYINT UNSIGNED DEFAULT 0,
    network_security_score TINYINT UNSIGNED DEFAULT 0,
    cryptography_score TINYINT UNSIGNED DEFAULT 0,
    artificial_intelligence_score TINYINT UNSIGNED DEFAULT 0,
    machine_learning_score TINYINT UNSIGNED DEFAULT 0,
    deep_learning_score TINYINT UNSIGNED DEFAULT 0,
    computer_vision_score TINYINT UNSIGNED DEFAULT 0,
    natural_language_processing_score TINYINT UNSIGNED DEFAULT 0,
    blockchain_technology_score TINYINT UNSIGNED DEFAULT 0,
    quantum_computing_score TINYINT UNSIGNED DEFAULT 0,
    edge_computing_score TINYINT UNSIGNED DEFAULT 0,
    microservices_architecture_score TINYINT UNSIGNED DEFAULT 0,
    devops_score TINYINT UNSIGNED DEFAULT 0,
    container_technology_score TINYINT UNSIGNED DEFAULT 0,
    ci_cd_score TINYINT UNSIGNED DEFAULT 0,
    test_automation_score TINYINT UNSIGNED DEFAULT 0,
    ux_ui_design_score TINYINT UNSIGNED DEFAULT 0,
    agile_development_methodologies_score TINYINT UNSIGNED DEFAULT 0,
    open_source_score TINYINT UNSIGNED DEFAULT 0,
    version_control_score TINYINT UNSIGNED DEFAULT 0,
    api_design_score TINYINT UNSIGNED DEFAULT 0,
    performance_optimization_score TINYINT UNSIGNED DEFAULT 0,
    CONSTRAINT fk_interests_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CHECK (news_score BETWEEN 0 AND 100),
    CHECK (politics_score BETWEEN 0 AND 100),
    CHECK (economics_score BETWEEN 0 AND 100),
    CHECK (health_score BETWEEN 0 AND 100),
    CHECK (sports_score BETWEEN 0 AND 100),
    CHECK (entertainment_score BETWEEN 0 AND 100),
    CHECK (art_score BETWEEN 0 AND 100),
    CHECK (cooking_score BETWEEN 0 AND 100),
    CHECK (travel_score BETWEEN 0 AND 100),
    CHECK (fashion_score BETWEEN 0 AND 100),
    CHECK (beauty_score BETWEEN 0 AND 100),
    CHECK (pets_score BETWEEN 0 AND 100),
    CHECK (parenting_score BETWEEN 0 AND 100),
    CHECK (education_score BETWEEN 0 AND 100),
    CHECK (environment_score BETWEEN 0 AND 100),
    CHECK (climate_score BETWEEN 0 AND 100),
    CHECK (space_score BETWEEN 0 AND 100),
    CHECK (mental_health_score BETWEEN 0 AND 100),
    CHECK (fitness_score BETWEEN 0 AND 100),
    CHECK (reading_score BETWEEN 0 AND 100),
    CHECK (history_score BETWEEN 0 AND 100),
    CHECK (philosophy_score BETWEEN 0 AND 100),
    CHECK (religion_score BETWEEN 0 AND 100),
    CHECK (culture_score BETWEEN 0 AND 100),
    CHECK (volunteering_score BETWEEN 0 AND 100),
    CHECK (social_issues_score BETWEEN 0 AND 100),
    CHECK (law_score BETWEEN 0 AND 100),
    CHECK (taxes_score BETWEEN 0 AND 100),
    CHECK (investment_score BETWEEN 0 AND 100),
    CHECK (real_estate_score BETWEEN 0 AND 100),
    CHECK (diy_score BETWEEN 0 AND 100),
    CHECK (gardening_score BETWEEN 0 AND 100),
    CHECK (interior_design_score BETWEEN 0 AND 100),
    CHECK (automotive_score BETWEEN 0 AND 100),
    CHECK (gaming_score BETWEEN 0 AND 100),
    CHECK (anime_manga_score BETWEEN 0 AND 100),
    CHECK (creative_works_score BETWEEN 0 AND 100),
    CHECK (photography_video_score BETWEEN 0 AND 100),
    CHECK (media_score BETWEEN 0 AND 100),
    CHECK (marketing_score BETWEEN 0 AND 100),
    CHECK (branding_score BETWEEN 0 AND 100),
    CHECK (entrepreneurship_score BETWEEN 0 AND 100),
    CHECK (remote_work_score BETWEEN 0 AND 100),
    CHECK (data_science_score BETWEEN 0 AND 100),
    CHECK (iot_score BETWEEN 0 AND 100),
    CHECK (robotics_engineering_score BETWEEN 0 AND 100),
    CHECK (biotechnology_score BETWEEN 0 AND 100),
    CHECK (nanotechnology_score BETWEEN 0 AND 100),
    CHECK (energy_technology_score BETWEEN 0 AND 100),
    CHECK (archaeology_score BETWEEN 0 AND 100),
    CHECK (psychology_score BETWEEN 0 AND 100),
    CHECK (sociology_score BETWEEN 0 AND 100),
    CHECK (anthropology_score BETWEEN 0 AND 100),
    CHECK (geography_score BETWEEN 0 AND 100),
    CHECK (geology_score BETWEEN 0 AND 100),
    CHECK (meteorology_score BETWEEN 0 AND 100),
    CHECK (disaster_emergency_management_score BETWEEN 0 AND 100),
    CHECK (urban_planning_score BETWEEN 0 AND 100),
    CHECK (architecture_score BETWEEN 0 AND 100),
    CHECK (agriculture_score BETWEEN 0 AND 100),
    CHECK (nutrition_science_score BETWEEN 0 AND 100),
    CHECK (sleep_science_score BETWEEN 0 AND 100),
    CHECK (productivity_score BETWEEN 0 AND 100),
    CHECK (leadership_score BETWEEN 0 AND 100),
    CHECK (international_relations_score BETWEEN 0 AND 100),
    CHECK (future_predictions_score BETWEEN 0 AND 100),
    CHECK (events_score BETWEEN 0 AND 100),
    CHECK (community_score BETWEEN 0 AND 100),
    CHECK (trends_score BETWEEN 0 AND 100),
    CHECK (lifestyle_score BETWEEN 0 AND 100),
    CHECK (software_development_score BETWEEN 0 AND 100),
    CHECK (programming_languages_score BETWEEN 0 AND 100),
    CHECK (web_development_score BETWEEN 0 AND 100),
    CHECK (mobile_app_development_score BETWEEN 0 AND 100),
    CHECK (debugging_techniques_score BETWEEN 0 AND 100),
    CHECK (algorithms_mathematics_score BETWEEN 0 AND 100),
    CHECK (database_design_score BETWEEN 0 AND 100),
    CHECK (cloud_computing_score BETWEEN 0 AND 100),
    CHECK (server_management_score BETWEEN 0 AND 100),
    CHECK (network_security_score BETWEEN 0 AND 100),
    CHECK (cryptography_score BETWEEN 0 AND 100),
    CHECK (artificial_intelligence_score BETWEEN 0 AND 100),
    CHECK (machine_learning_score BETWEEN 0 AND 100),
    CHECK (deep_learning_score BETWEEN 0 AND 100),
    CHECK (computer_vision_score BETWEEN 0 AND 100),
    CHECK (natural_language_processing_score BETWEEN 0 AND 100),
    CHECK (blockchain_technology_score BETWEEN 0 AND 100),
    CHECK (quantum_computing_score BETWEEN 0 AND 100),
    CHECK (edge_computing_score BETWEEN 0 AND 100),
    CHECK (microservices_architecture_score BETWEEN 0 AND 100),
    CHECK (devops_score BETWEEN 0 AND 100),
    CHECK (container_technology_score BETWEEN 0 AND 100),
    CHECK (ci_cd_score BETWEEN 0 AND 100),
    CHECK (test_automation_score BETWEEN 0 AND 100),
    CHECK (ux_ui_design_score BETWEEN 0 AND 100),
    CHECK (agile_development_methodologies_score BETWEEN 0 AND 100),
    CHECK (open_source_score BETWEEN 0 AND 100),
    CHECK (version_control_score BETWEEN 0 AND 100),
    CHECK (api_design_score BETWEEN 0 AND 100),
    CHECK (performance_optimization_score BETWEEN 0 AND 100)
);

DROP TABLE IF EXISTS notifications;

CREATE TABLE notifications (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_account_id CHAR(28),
    recipient_account_id CHAR(28) NOT NULL,
    type VARCHAR(50) NOT NULL,
    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_notifications_resipient FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_notifications_sender FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE SET NULL,
    INDEX idx_notifications_recipient (recipient_account_id),
    INDEX idx_notifications_created_at (created_at)
);

DROP TABLE IF EXISTS messages;

CREATE TABLE messages (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    sender_account_id CHAR(28) NOT NULL,
    recipient_account_id CHAR(28) NOT NULL,
    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_messages_sender FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_messages_recipient FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_messages_sender (sender_account_id),
    INDEX idx_messages_recipient (recipient_account_id),
    INDEX idx_messages_created_at (created_at)
);

DROP TABLE IF EXISTS follows;

CREATE TABLE follows (
    follower_account_id CHAR(28) NOT NULL,
    following_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_account_id, following_account_id),
    CONSTRAINT fk_follows_follower FOREIGN KEY (follower_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_follows_following FOREIGN KEY (following_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_follows_created_at (created_at)
);

DROP TABLE IF EXISTS follow_requests;

CREATE TABLE follow_requests (
    requester_account_id CHAR(28) NOT NULL,
    requestee_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (requester_account_id, requestee_account_id),
    CONSTRAINT fk_follow_requests_requester FOREIGN KEY (requester_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_follow_requests_requestee FOREIGN KEY (requestee_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_follow_requests_created_at (created_at)
);

DROP TABLE IF EXISTS blocks;

CREATE TABLE blocks (
    blocker_account_id CHAR(28) NOT NULL,
    blocked_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (blocker_account_id, blocked_account_id),
    CONSTRAINT fk_blocks_blocker FOREIGN KEY (blocker_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_blocks_blocked FOREIGN KEY (blocked_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_blocks_created_at (created_at)
);

DROP TABLE IF EXISTS tweets;

CREATE TABLE tweets (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    account_id CHAR(28) NOT NULL,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    content TEXT,
    code TEXT,
    likes_count INT UNSIGNED NOT NULL DEFAULT 0,
    replies_count INT UNSIGNED NOT NULL DEFAULT 0,
    retweets_count INT UNSIGNED NOT NULL DEFAULT 0,
    is_retweet BOOLEAN NOT NULL DEFAULT FALSE,
    is_reply BOOLEAN NOT NULL DEFAULT FALSE,
    engagement_score INT UNSIGNED NOT NULL DEFAULT 0,
    original_tweet_id BIGINT UNSIGNED,
    quoted_tweet_id BIGINT UNSIGNED,
    reply_to_tweet_id BIGINT UNSIGNED,
    reply_to_account_id CHAR(28),
    media JSON,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_tweets_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_tweets_original_tweet FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE SET NULL,
    CONSTRAINT fk_tweets_quoted_tweet FOREIGN KEY (quoted_tweet_id)
        REFERENCES tweets(id) ON DELETE SET NULL,
    CONSTRAINT fk_tweets_reply_to_tweet FOREIGN KEY (reply_to_tweet_id)
        REFERENCES tweets(id) ON DELETE SET NULL,
    CONSTRAINT fk_tweets_reply_to_account FOREIGN KEY (reply_to_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_tweets_account (account_id),
    INDEX idx_tweets_engagement_score (engagement_score),
    INDEX idx_tweets_created_at (created_at),
    CHECK (JSON_VALID(media))
);

DROP TABLE IF EXISTS labels;

CREATE TABLE labels (
    tweet_id BIGINT UNSIGNED PRIMARY KEY,
    label1 VARCHAR(50) NOT NULL,
    label2 VARCHAR(50),
    label3 VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_labels_tweet FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS likes;

CREATE TABLE likes (
    account_id CHAR(28) NOT NULL,
    tweet_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, tweet_id),
    CONSTRAINT fk_likes_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_likes_tweet FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    INDEX idx_tweets_id (tweet_id)
);

DROP TABLE IF EXISTS retweets;

CREATE TABLE retweets (
    account_id CHAR(28) NOT NULL,
    tweet_id BIGINT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (account_id, tweet_id),
    CONSTRAINT fk_retweets_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_retweets_tweet FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    INDEX idx_retweets_tweets_id (tweet_id)
);

DROP TABLE IF EXISTS replies;

CREATE TABLE replies (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tweet_id BIGINT UNSIGNED NOT NULL,
    reply_tweet_id BIGINT UNSIGNED NOT NULL,
    parent_reply_id BIGINT UNSIGNED,
    account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_reply_tweet (reply_tweet_id),
    INDEX idx_tweet_id (tweet_id),
    INDEX idx_parent_reply_id (parent_reply_id),
    INDEX idx_account_id (account_id),
    INDEX idx_created_at (created_at),
    CONSTRAINT fk_replies_tweet FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_replies_reply_tweet FOREIGN KEY (reply_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_replies_parent_reply FOREIGN KEY (parent_reply_id)
        REFERENCES replies(id) ON DELETE SET NULL,
    CONSTRAINT fk_replies_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS hashtags;

CREATE TABLE hashtags (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tag VARCHAR(30) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    INDEX idx_hashtags_tag (tag)
);

CREATE TABLE tweet_hashtags (
    tweet_id BIGINT UNSIGNED NOT NULL,
    hashtag_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (tweet_id, hashtag_id),
    CONSTRAINT fk_tweet_hashtags_tweet FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    CONSTRAINT fk_tweet_hashtags_hashtag FOREIGN KEY (hashtag_id)
        REFERENCES hashtags(id) ON DELETE CASCADE,
    INDEX idx_tweet_hashtags_hashtag_id (hashtag_id)
);