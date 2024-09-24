DROP TABLE IF EXISTS accounts;

CREATE TABLE accounts (
    id CHAR(28) PRIMARY KEY,
    user_id VARCHAR(28) UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    KEY idx_accounts_user_name (user_name)
);

DROP TABLE IF EXISTS profiles;

CREATE TABLE profiles (
    account_id CHAR(28) PRIMARY KEY,
    bio TEXT,
    profile_image_url VARCHAR(2083),
    banner_image_url VARCHAR(2083),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_profiles_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS settings;

CREATE TABLE settings (
    account_id CHAR(28) PRIMARY KEY,
    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_settings_account FOREIGN KEY (account_id)
        REFERENCES accounts(id) ON DELETE CASCADE
);

DROP TABLE IF EXISTS notifications;

CREATE TABLE notifications (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    recipient_account_id CHAR(28) NOT NULL,
    sender_account_id CHAR(28),
    type VARCHAR(50) NOT NULL,
    content TEXT,
    is_read BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT fk_notifications_resipient FOREIGN KEY (recipient_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_notifications_sender FOREIGN KEY (sender_account_id)
        REFERENCES accounts(id) ON DELETE SET NULL,
    INDEX idx_notifications_recipient (recipient_account_id),
    INDEX idx_notifications_created_at (created_at)
);