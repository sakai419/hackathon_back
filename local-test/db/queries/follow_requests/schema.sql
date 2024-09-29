CREATE TABLE follow_requests (
    requester_account_id CHAR(28) NOT NULL,
    requestee_account_id CHAR(28) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (requester_account_id, requestee_account_id),
    CONSTRAINT fk_follow_requests_requester_account_id FOREIGN KEY (requester_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    CONSTRAINT fk_follow_requests_requestee_account_id FOREIGN KEY (requestee_account_id)
        REFERENCES accounts(id) ON DELETE CASCADE,
    INDEX idx_follow_requests_created_at (created_at)
);