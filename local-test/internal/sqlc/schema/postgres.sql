---Drop tables---
DROP TABLE IF EXISTS accounts CASCADE;
DROP TABLE IF EXISTS blocks CASCADE;
DROP TABLE IF EXISTS conversations CASCADE;
DROP TABLE IF EXISTS follows CASCADE;
DROP TABLE IF EXISTS hashtags CASCADE;
DROP TABLE IF EXISTS interests CASCADE;
DROP TABLE IF EXISTS labels CASCADE;
DROP TABLE IF EXISTS likes CASCADE;
DROP TABLE IF EXISTS messages CASCADE;
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS profiles CASCADE;
DROP TABLE IF EXISTS quotes CASCADE;
DROP TABLE IF EXISTS replies CASCADE;
DROP TABLE IF EXISTS reports CASCADE;
DROP TABLE IF EXISTS retweets CASCADE;
DROP TABLE IF EXISTS settings CASCADE;
DROP TABLE IF EXISTS tweet_hashtags CASCADE;
DROP TABLE IF EXISTS tweets CASCADE;

---Create tables---

-- Table: accounts
CREATE TABLE accounts (
    id CHAR(28) PRIMARY KEY,
    user_id VARCHAR(30) UNIQUE NOT NULL,
    user_name VARCHAR(30) NOT NULL,
    is_suspended BOOLEAN NOT NULL DEFAULT FALSE,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: blocks
CREATE TABLE blocks (
    blocker_account_id CHAR(28) NOT NULL,
    blocked_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (blocker_account_id, blocked_account_id)
);

CREATE INDEX idx_blocks_created_at ON blocks(created_at);

-- Table: conversations
CREATE TABLE conversations (
    id BIGSERIAL PRIMARY KEY,
    account1_id CHAR(28) NOT NULL,
    account2_id CHAR(28) NOT NULL,
    last_message_id BIGINT,
    last_message_time TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (account1_id, account2_id)
);

-- Table: follows
CREATE TYPE follow_status AS ENUM (
    'accepted',
    'pending'
);

CREATE TABLE follows (
    follower_account_id CHAR(28) NOT NULL,
    following_account_id CHAR(28) NOT NULL,
    status follow_status NOT NULL DEFAULT 'accepted',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_account_id, following_account_id)
);

CREATE INDEX idx_follows_created_at ON follows(created_at);

-- Table: hashtags
CREATE TABLE hashtags (
    id BIGSERIAL PRIMARY KEY,
    tag VARCHAR(30) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: interests
CREATE TABLE interests (
    account_id CHAR(28) PRIMARY KEY,
    news_score SMALLINT DEFAULT 0 CHECK (news_score BETWEEN 0 AND 100),
    politics_score SMALLINT DEFAULT 0 CHECK (politics_score BETWEEN 0 AND 100),
    economics_score SMALLINT DEFAULT 0 CHECK (economics_score BETWEEN 0 AND 100),
    health_score SMALLINT DEFAULT 0 CHECK (health_score BETWEEN 0 AND 100),
    sports_score SMALLINT DEFAULT 0 CHECK (sports_score BETWEEN 0 AND 100),
    entertainment_score SMALLINT DEFAULT 0 CHECK (entertainment_score BETWEEN 0 AND 100),
    art_score SMALLINT DEFAULT 0 CHECK (art_score BETWEEN 0 AND 100),
    cooking_score SMALLINT DEFAULT 0 CHECK (cooking_score BETWEEN 0 AND 100),
    travel_score SMALLINT DEFAULT 0 CHECK (travel_score BETWEEN 0 AND 100),
    fashion_score SMALLINT DEFAULT 0 CHECK (fashion_score BETWEEN 0 AND 100),
    beauty_score SMALLINT DEFAULT 0 CHECK (beauty_score BETWEEN 0 AND 100),
    pets_score SMALLINT DEFAULT 0 CHECK (pets_score BETWEEN 0 AND 100),
    parenting_score SMALLINT DEFAULT 0 CHECK (parenting_score BETWEEN 0 AND 100),
    education_score SMALLINT DEFAULT 0 CHECK (education_score BETWEEN 0 AND 100),
    environment_score SMALLINT DEFAULT 0 CHECK (environment_score BETWEEN 0 AND 100),
    climate_score SMALLINT DEFAULT 0 CHECK (climate_score BETWEEN 0 AND 100),
    space_score SMALLINT DEFAULT 0 CHECK (space_score BETWEEN 0 AND 100),
    mental_health_score SMALLINT DEFAULT 0 CHECK (mental_health_score BETWEEN 0 AND 100),
    fitness_score SMALLINT DEFAULT 0 CHECK (fitness_score BETWEEN 0 AND 100),
    reading_score SMALLINT DEFAULT 0 CHECK (reading_score BETWEEN 0 AND 100),
    history_score SMALLINT DEFAULT 0 CHECK (history_score BETWEEN 0 AND 100),
    philosophy_score SMALLINT DEFAULT 0 CHECK (philosophy_score BETWEEN 0 AND 100),
    religion_score SMALLINT DEFAULT 0 CHECK (religion_score BETWEEN 0 AND 100),
    culture_score SMALLINT DEFAULT 0 CHECK (culture_score BETWEEN 0 AND 100),
    volunteering_score SMALLINT DEFAULT 0 CHECK (volunteering_score BETWEEN 0 AND 100),
    social_issues_score SMALLINT DEFAULT 0 CHECK (social_issues_score BETWEEN 0 AND 100),
    law_score SMALLINT DEFAULT 0 CHECK (law_score BETWEEN 0 AND 100),
    taxes_score SMALLINT DEFAULT 0 CHECK (taxes_score BETWEEN 0 AND 100),
    investment_score SMALLINT DEFAULT 0 CHECK (investment_score BETWEEN 0 AND 100),
    real_estate_score SMALLINT DEFAULT 0 CHECK (real_estate_score BETWEEN 0 AND 100),
    diy_score SMALLINT DEFAULT 0 CHECK (diy_score BETWEEN 0 AND 100),
    gardening_score SMALLINT DEFAULT 0 CHECK (gardening_score BETWEEN 0 AND 100),
    interior_design_score SMALLINT DEFAULT 0 CHECK (interior_design_score BETWEEN 0 AND 100),
    automotive_score SMALLINT DEFAULT 0 CHECK (automotive_score BETWEEN 0 AND 100),
    gaming_score SMALLINT DEFAULT 0 CHECK (gaming_score BETWEEN 0 AND 100),
    anime_manga_score SMALLINT DEFAULT 0 CHECK (anime_manga_score BETWEEN 0 AND 100),
    creative_works_score SMALLINT DEFAULT 0 CHECK (creative_works_score BETWEEN 0 AND 100),
    photography_video_score SMALLINT DEFAULT 0 CHECK (photography_video_score BETWEEN 0 AND 100),
    media_score SMALLINT DEFAULT 0 CHECK (media_score BETWEEN 0 AND 100),
    marketing_score SMALLINT DEFAULT 0 CHECK (marketing_score BETWEEN 0 AND 100),
    branding_score SMALLINT DEFAULT 0 CHECK (branding_score BETWEEN 0 AND 100),
    entrepreneurship_score SMALLINT DEFAULT 0 CHECK (entrepreneurship_score BETWEEN 0 AND 100),
    remote_work_score SMALLINT DEFAULT 0 CHECK (remote_work_score BETWEEN 0 AND 100),
    data_science_score SMALLINT DEFAULT 0 CHECK (data_science_score BETWEEN 0 AND 100),
    iot_score SMALLINT DEFAULT 0 CHECK (iot_score BETWEEN 0 AND 100),
    robotics_engineering_score SMALLINT DEFAULT 0 CHECK (robotics_engineering_score BETWEEN 0 AND 100),
    biotechnology_score SMALLINT DEFAULT 0 CHECK (biotechnology_score BETWEEN 0 AND 100),
    nanotechnology_score SMALLINT DEFAULT 0 CHECK (nanotechnology_score BETWEEN 0 AND 100),
    energy_technology_score SMALLINT DEFAULT 0 CHECK (energy_technology_score BETWEEN 0 AND 100),
    archaeology_score SMALLINT DEFAULT 0 CHECK (archaeology_score BETWEEN 0 AND 100),
    psychology_score SMALLINT DEFAULT 0 CHECK (psychology_score BETWEEN 0 AND 100),
    sociology_score SMALLINT DEFAULT 0 CHECK (sociology_score BETWEEN 0 AND 100),
    anthropology_score SMALLINT DEFAULT 0 CHECK (anthropology_score BETWEEN 0 AND 100),
    geography_score SMALLINT DEFAULT 0 CHECK (geography_score BETWEEN 0 AND 100),
    geology_score SMALLINT DEFAULT 0 CHECK (geology_score BETWEEN 0 AND 100),
    meteorology_score SMALLINT DEFAULT 0 CHECK (meteorology_score BETWEEN 0 AND 100),
    disaster_emergency_management_score SMALLINT DEFAULT 0 CHECK (disaster_emergency_management_score BETWEEN 0 AND 100),
    urban_planning_score SMALLINT DEFAULT 0 CHECK (urban_planning_score BETWEEN 0 AND 100),
    architecture_score SMALLINT DEFAULT 0 CHECK (architecture_score BETWEEN 0 AND 100),
    agriculture_score SMALLINT DEFAULT 0 CHECK (agriculture_score BETWEEN 0 AND 100),
    nutrition_science_score SMALLINT DEFAULT 0 CHECK (nutrition_science_score BETWEEN 0 AND 100),
    sleep_science_score SMALLINT DEFAULT 0 CHECK (sleep_science_score BETWEEN 0 AND 100),
    productivity_score SMALLINT DEFAULT 0 CHECK (productivity_score BETWEEN 0 AND 100),
    leadership_score SMALLINT DEFAULT 0 CHECK (leadership_score BETWEEN 0 AND 100),
    international_relations_score SMALLINT DEFAULT 0 CHECK (international_relations_score BETWEEN 0 AND 100),
    future_predictions_score SMALLINT DEFAULT 0 CHECK (future_predictions_score BETWEEN 0 AND 100),
    events_score SMALLINT DEFAULT 0 CHECK (events_score BETWEEN 0 AND 100),
    community_score SMALLINT DEFAULT 0 CHECK (community_score BETWEEN 0 AND 100),
    trends_score SMALLINT DEFAULT 0 CHECK (trends_score BETWEEN 0 AND 100),
    lifestyle_score SMALLINT DEFAULT 0 CHECK (lifestyle_score BETWEEN 0 AND 100),
    software_development_score SMALLINT DEFAULT 0 CHECK (software_development_score BETWEEN 0 AND 100),
    programming_languages_score SMALLINT DEFAULT 0 CHECK (programming_languages_score BETWEEN 0 AND 100),
    web_development_score SMALLINT DEFAULT 0 CHECK (web_development_score BETWEEN 0 AND 100),
    mobile_app_development_score SMALLINT DEFAULT 0 CHECK (mobile_app_development_score BETWEEN 0 AND 100),
    debugging_techniques_score SMALLINT DEFAULT 0 CHECK (debugging_techniques_score BETWEEN 0 AND 100),
    algorithms_mathematics_score SMALLINT DEFAULT 0 CHECK (algorithms_mathematics_score BETWEEN 0 AND 100),
    database_design_score SMALLINT DEFAULT 0 CHECK (database_design_score BETWEEN 0 AND 100),
    cloud_computing_score SMALLINT DEFAULT 0 CHECK (cloud_computing_score BETWEEN 0 AND 100),
    server_management_score SMALLINT DEFAULT 0 CHECK (server_management_score BETWEEN 0 AND 100),
    network_security_score SMALLINT DEFAULT 0 CHECK (network_security_score BETWEEN 0 AND 100),
    cryptography_score SMALLINT DEFAULT 0 CHECK (cryptography_score BETWEEN 0 AND 100),
    artificial_intelligence_score SMALLINT DEFAULT 0 CHECK (artificial_intelligence_score BETWEEN 0 AND 100),
    machine_learning_score SMALLINT DEFAULT 0 CHECK (machine_learning_score BETWEEN 0 AND 100),
    deep_learning_score SMALLINT DEFAULT 0 CHECK (deep_learning_score BETWEEN 0 AND 100),
    computer_vision_score SMALLINT DEFAULT 0 CHECK (computer_vision_score BETWEEN 0 AND 100),
    natural_language_processing_score SMALLINT DEFAULT 0 CHECK (natural_language_processing_score BETWEEN 0 AND 100),
    blockchain_technology_score SMALLINT DEFAULT 0 CHECK (blockchain_technology_score BETWEEN 0 AND 100),
    quantum_computing_score SMALLINT DEFAULT 0 CHECK (quantum_computing_score BETWEEN 0 AND 100),
    edge_computing_score SMALLINT DEFAULT 0 CHECK (edge_computing_score BETWEEN 0 AND 100),
    microservices_architecture_score SMALLINT DEFAULT 0 CHECK (microservices_architecture_score BETWEEN 0 AND 100),
    devops_score SMALLINT DEFAULT 0 CHECK (devops_score BETWEEN 0 AND 100),
    container_technology_score SMALLINT DEFAULT 0 CHECK (container_technology_score BETWEEN 0 AND 100),
    ci_cd_score SMALLINT DEFAULT 0 CHECK (ci_cd_score BETWEEN 0 AND 100),
    test_automation_score SMALLINT DEFAULT 0 CHECK (test_automation_score BETWEEN 0 AND 100),
    ux_ui_design_score SMALLINT DEFAULT 0 CHECK (ux_ui_design_score BETWEEN 0 AND 100),
    agile_development_methodologies_score SMALLINT DEFAULT 0 CHECK (agile_development_methodologies_score BETWEEN 0 AND 100),
    open_source_score SMALLINT DEFAULT 0 CHECK (open_source_score BETWEEN 0 AND 100),
    version_control_score SMALLINT DEFAULT 0 CHECK (version_control_score BETWEEN 0 AND 100),
    api_design_score SMALLINT DEFAULT 0 CHECK (api_design_score BETWEEN 0 AND 100),
    performance_optimization_score SMALLINT DEFAULT 0 CHECK (performance_optimization_score BETWEEN 0 AND 100)
);

-- Table: labels
CREATE TYPE tweet_label AS ENUM (
    'news',
    'politics',
    'economics',
    'health',
    'sports',
    'entertainment',
    'art',
    'cooking',
    'travel',
    'fashion',
    'beauty',
    'pets',
    'parenting',
    'education',
    'environment',
    'climate',
    'space',
    'mental_health',
    'fitness',
    'reading',
    'history',
    'philosophy',
    'religion',
    'culture',
    'volunteering',
    'social_issues',
    'law',
    'taxes',
    'investment',
    'real_estate',
    'diy',
    'gardening',
    'interior_design',
    'automotive',
    'gaming',
    'anime_manga',
    'creative_works',
    'photography_video',
    'media',
    'marketing',
    'branding',
    'entrepreneurship',
    'remote_work',
    'data_science',
    'iot',
    'robotics_engineering',
    'biotechnology',
    'nanotechnology',
    'energy_technology',
    'archaeology',
    'psychology',
    'sociology',
    'anthropology',
    'geography',
    'geology',
    'meteorology',
    'disaster_emergency_management',
    'urban_planning',
    'architecture',
    'agriculture',
    'nutrition_science',
    'sleep_science',
    'productivity',
    'leadership',
    'international_relations',
    'future_predictions',
    'events',
    'community',
    'trends',
    'lifestyle',
    'software_development',
    'programming_languages',
    'web_development',
    'mobile_app_development',
    'debugging_techniques',
    'algorithms_mathematics',
    'database_design',
    'cloud_computing',
    'server_management',
    'network_security',
    'cryptography',
    'artificial_intelligence',
    'machine_learning',
    'deep_learning',
    'computer_vision',
    'natural_language_processing',
    'blockchain_technology',
    'quantum_computing',
    'edge_computing',
    'microservices_architecture',
    'devops',
    'container_technology',
    'ci_cd',
    'test_automation',
    'ux_ui_design',
    'agile_development_methodologies',
    'open_source',
    'version_control',
    'api_design',
    'performance_optimization'
);

CREATE TABLE labels (
    tweet_id BIGINT PRIMARY KEY,
    label1 tweet_label,
    label2 tweet_label,
    label3 tweet_label,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: likes
CREATE TABLE likes (
    liking_account_id CHAR(28) NOT NULL,
    original_tweet_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (liking_account_id, original_tweet_id)
);

CREATE INDEX idx_likes_original_tweet_id ON likes(original_tweet_id);

-- Table: messages
CREATE TABLE messages (
    id BIGSERIAL PRIMARY KEY,
    conversation_id BIGINT NOT NULL,
    sender_account_id CHAR(28) NOT NULL,
    content TEXT NOT NULL,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_messages_conversation_id ON messages (conversation_id);
CREATE INDEX idx_messages_sender_account_id ON messages (sender_account_id);
CREATE INDEX idx_messages_created_at_account_id ON messages (created_at);

-- Table: notifications
CREATE TYPE notification_type AS ENUM (
    'follow',
    'like',
    'retweet',
    'reply',
    'message',
    'quote',
    'follow_request',
    'request_accepted',
    'report',
    'warning',
    'other'
);

CREATE TABLE notifications (
    id BIGSERIAL PRIMARY KEY,
    sender_account_id CHAR(28),
    recipient_account_id CHAR(28) NOT NULL,
    type notification_type NOT NULL,
    content TEXT,
    tweet_id BIGINT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notifications_recipient_account_id ON notifications(recipient_account_id);
CREATE INDEX idx_notifications_created_at ON notifications(created_at);

-- Table: profile
CREATE TABLE profiles (
    account_id CHAR(28) PRIMARY KEY,
    bio TEXT,
    profile_image_url VARCHAR(2083),
    banner_image_url VARCHAR(2083),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: quotes
CREATE TABLE quotes (
    quote_id BIGINT PRIMARY KEY,
    quoting_account_id CHAR(28) NOT NULL,
    original_tweet_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_quotes_original_tweet_id ON quotes (original_tweet_id);
CREATE INDEX idx_quotes_quoting_account_id ON quotes (quoting_account_id);

-- Table: replies
CREATE TABLE replies (
    reply_id BIGINT NOT NULL PRIMARY KEY,
    original_tweet_id BIGINT NOT NULL,
    parent_reply_id BIGINT,
    replying_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_replies_original_tweet_id ON replies (original_tweet_id);
CREATE INDEX idx_replies_parent_reply_id ON replies (parent_reply_id);
CREATE INDEX idx_replies_replying_account_id ON replies (replying_account_id);
CREATE INDEX idx_replies_created_at ON replies (created_at);

-- Table: reports
CREATE TYPE report_reason AS ENUM (
    'spam',
    'harassment',
    'inappropriate_content',
    'other'
);

CREATE TABLE reports (
    id BIGSERIAL PRIMARY KEY,
    reporter_account_id CHAR(28) NOT NULL,
    reported_account_id CHAR(28) NOT NULL,
    reason report_reason NOT NULL,
    content TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_reports_created_at ON reports (created_at);
CREATE INDEX idx_reports_reported_account_id ON reports (reported_account_id);

-- Table: retweets
CREATE TABLE retweets (
    retweeting_account_id CHAR(28) NOT NULL,
    original_tweet_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (retweeting_account_id, original_tweet_id)
);

CREATE INDEX idx_retweets_orginal_tweet_id ON retweets(original_tweet_id);

-- Table: settings
CREATE TABLE settings (
    account_id CHAR(28) PRIMARY KEY,
    is_private BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Table: tweet_hashtags
CREATE TABLE tweet_hashtags (
    tweet_id BIGINT NOT NULL,
    hashtag_id BIGINT NOT NULL,
    PRIMARY KEY (tweet_id, hashtag_id)
);

CREATE INDEX idx_tweet_hashtags_hashtag_id ON tweet_hashtags (hashtag_id);

-- Table: tweets
CREATE TABLE tweets (
    id BIGSERIAL PRIMARY KEY,
    account_id CHAR(28) NOT NULL,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    content TEXT DEFAULT NULL,
    code TEXT DEFAULT NULL,
    likes_count INTEGER NOT NULL DEFAULT 0,
    replies_count INTEGER NOT NULL DEFAULT 0,
    retweets_count INTEGER NOT NULL DEFAULT 0,
    is_reply BOOLEAN NOT NULL DEFAULT FALSE,
    is_quote BOOLEAN NOT NULL DEFAULT FALSE,
    media JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_tweets_account_id ON tweets(account_id);
CREATE INDEX idx_tweets_created_at ON tweets(created_at);
CREATE INDEX idx_tweets_type ON tweets(is_retweet, is_reply, is_quote);

---Define foreign keys---

-- Table: blocks
ALTER TABLE blocks
    ADD CONSTRAINT fk_blocks_blocker_account_id FOREIGN KEY (blocker_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_blocks_blocked_account_id FOREIGN KEY (blocked_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: conversations
ALTER TABLE conversations
    ADD CONSTRAINT fk_conversations_account1_id FOREIGN KEY (account1_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_conversations_account2_id FOREIGN KEY (account2_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_conversations_last_message_id FOREIGN KEY (last_message_id)
        REFERENCES messages(id) ON DELETE SET NULL;

-- Table: follows
ALTER TABLE follows
    ADD CONSTRAINT fk_follows_follower_account_id FOREIGN KEY (follower_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_follows_following_account_id FOREIGN KEY (following_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: interests
ALTER TABLE interests
    ADD CONSTRAINT fk_interests_account_id FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: labels
ALTER TABLE labels
    ADD CONSTRAINT fk_labels_tweet_id FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE;

-- Table: likes
ALTER TABLE likes
    ADD CONSTRAINT fk_likes_liking_account_id FOREIGN KEY (liking_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_likes_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE;

-- Table: messages
ALTER TABLE messages
    ADD CONSTRAINT fk_messages_conversation_id FOREIGN KEY (conversation_id)
        REFERENCES conversations(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_messages_sender_account_id FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: notifications
ALTER TABLE notifications
    ADD CONSTRAINT fk_notifications_sender_account_id FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE SET NULL,
    ADD CONSTRAINT fk_notifications_recipient_account_id FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_notifications_tweet_id FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE SET NULL;

-- Table: profiles
ALTER TABLE profiles
    ADD CONSTRAINT fk_profiles_account_id FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: quotes
ALTER TABLE quotes
    ADD CONSTRAINT fk_quotes_quote_id FOREIGN KEY (quote_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_quotes_quoting_account_id FOREIGN KEY (quoting_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_quotes_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE;

-- Table: replies
ALTER TABLE replies
    ADD CONSTRAINT fk_replies_reply_id FOREIGN KEY (reply_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_replies_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_replies_parent_reply_id FOREIGN KEY (parent_reply_id)
        REFERENCES replies(reply_id) ON DELETE SET NULL,
    ADD CONSTRAINT fk_replies_replying_account_id FOREIGN KEY (replying_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: reports
ALTER TABLE reports
    ADD CONSTRAINT fk_reports_reporter_account_id FOREIGN KEY (reporter_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_reports_reported_account_id FOREIGN KEY (reported_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: retweets
ALTER TABLE retweets
    ADD CONSTRAINT fk_retweets_retweeting_account_id FOREIGN KEY (retweeting_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_retweets_original_tweet_id FOREIGN KEY (original_tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE;

-- Table: settings
ALTER TABLE settings
    ADD CONSTRAINT fk_settings_account_id FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

-- Table: tweet_hashtags
ALTER TABLE tweet_hashtags
    ADD CONSTRAINT fk_tweet_hashtags_tweet_id FOREIGN KEY (tweet_id)
        REFERENCES tweets(id) ON DELETE CASCADE,
    ADD CONSTRAINT fk_tweet_hashtags_hashtag_id FOREIGN KEY (hashtag_id)
        REFERENCES hashtags(id) ON DELETE CASCADE;

-- Table: tweets
ALTER TABLE tweets
    ADD CONSTRAINT fk_tweets_account_id FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE;

---Define functions---
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION increment_like_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET likes_count = likes_count + 1
    WHERE id = NEW.original_tweet_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION decrement_like_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET likes_count = likes_count - 1
    WHERE id = OLD.original_tweet_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION increment_quote_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET retweets_count = retweets_count + 1
    WHERE id = NEW.original_tweet_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION decrement_quote_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET retweets_count = retweets_count - 1
    WHERE id = OLD.original_tweet_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION increment_reply_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET replies_count = replies_count + 1
    WHERE id = COALESCE(NEW.parent_reply_id, NEW.original_tweet_id);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION decrement_reply_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET replies_count = replies_count - 1
    WHERE id = COALESCE(OLD.parent_reply_id, OLD.original_tweet_id);
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION increment_retweet_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET retweets_count = retweets_count + 1
    WHERE id = NEW.original_tweet_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION decrement_retweet_count()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE tweets
    SET retweets_count = retweets_count - 1
    WHERE id = OLD.original_tweet_id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

---Define triggers---

CREATE TRIGGER trigger_update_tweet_timestamp
BEFORE UPDATE ON tweets
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER trigger_increment_like_count
AFTER INSERT ON likes
FOR EACH ROW
EXECUTE FUNCTION increment_like_count();

CREATE TRIGGER trigger_decrement_like_count
AFTER DELETE ON likes
FOR EACH ROW
EXECUTE FUNCTION decrement_like_count();

CREATE TRIGGER trigger_increment_quote_count
AFTER INSERT ON quotes
FOR EACH ROW
EXECUTE FUNCTION increment_quote_count();

CREATE TRIGGER trigger_decrement_quote_count
AFTER DELETE ON quotes
FOR EACH ROW
EXECUTE FUNCTION decrement_quote_count();

CREATE TRIGGER trigger_increment_reply_count
AFTER INSERT ON replies
FOR EACH ROW
EXECUTE FUNCTION increment_reply_count();

CREATE TRIGGER trigger_decrement_reply_count
AFTER DELETE ON replies
FOR EACH ROW
EXECUTE FUNCTION decrement_reply_count();

CREATE TRIGGER trigger_increment_retweet_count
AFTER INSERT ON retweets
FOR EACH ROW
EXECUTE FUNCTION increment_retweet_count();

CREATE TRIGGER trigger_decrement_retweet_count
AFTER DELETE ON retweets
FOR EACH ROW
EXECUTE FUNCTION decrement_retweet_count();
