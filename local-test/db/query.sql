-- name: GetAccountById :one
SELECT * FROM accounts WHERE id = ?;

-- name: GetAccountByUserId :one
SELECT * FROM accounts WHERE user_id = ?;

-- name: SearchAccountsByUserId :many
SELECT * FROM accounts WHERE user_id LIKE "%?%";

-- name: GetAccountByUserName :one
SELECT * FROM accounts WHERE user_name = ?;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts WHERE user_name LIKE "%?%";

-- name: CreateAccount :one
INSERT INTO accounts (id, user_id, user_name) VALUES (?, ?, ?);

-- name: UpdateAccountUserId :one
UPDATE accounts SET user_id = ? WHERE id = ?;

-- name: UpdateAccountUserName :one
UPDATE accounts SET user_name = ? WHERE id = ?;

-- name: DeleteAccount :one
DELETE FROM accounts WHERE id = ?;

-- name: GetProfileByAccountId :one
SELECT * FROM profiles WHERE account_id = ?;

-- name: CreateProfile :one
INSERT INTO profiles (account_id, bio, profile_image_url, banner_image_url) VALUES (?, ?, ?, ?);

-- name: UpdateProfileBio :one
UPDATE profiles SET bio = ? WHERE account_id = ?;

-- name: UpdateProfileProfileImageUrl :one
UPDATE profiles SET profile_image_url = ? WHERE account_id = ?;

-- name: UpdateProfileBannerImageUrl :one
UPDATE profiles SET banner_image_url = ? WHERE account_id = ?;

-- name: DeleteProfile :one
DELETE FROM profiles WHERE account_id = ?;