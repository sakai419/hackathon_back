CREATE TABLE blocks (
    blocker_account_id CHAR(28) NOT NULL,
    blocked_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (blocker_account_id, blocked_account_id),
    CONSTRAINT fk_blocks_blocker_account_id FOREIGN KEY (blocker_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_blocks_blocked_account_id FOREIGN KEY (blocked_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_blocks_created_at (created_at)
);