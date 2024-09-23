DROP TABLE IF EXISTS users;

CREATE TABLE accounts (
    id CHAR(28) PRIMARY KEY,
    user_id VARCHAR(28) UNIQUE NOT NULL,
    user_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE INDEX idx_accounts_user_id ON accounts (user_id);
CREATE INDEX idx_accounts_user_name ON accounts (user_name);

INSERT INTO accounts (id, user_id, user_name) VALUES ('1', '1', 'user1');