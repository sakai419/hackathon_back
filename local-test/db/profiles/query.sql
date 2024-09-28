-- name: GetProfileByAccountId :one
SELECT * FROM profiles
WHERE account_id = ?;

-- name: CreateProfile :exec
INSERT INTO profiles (account_id, bio, profile_image_url, banner_image_url)
VALUES (?, ?, ?, ?);

-- name: GetProfileByAccountId :one
SELECT * FROM profiles
WHERE account_id = ?;

-- name: UpdateProfileBio :exec
UPDATE profiles
SET bio = ?
WHERE account_id = ?;

-- name: UpdateProfileImageUrl :exec
UPDATE profiles
SET profile_image_url = ?
WHERE account_id = ?;

-- name: UpdateBannerImageUrl :exec
UPDATE profiles
SET banner_image_url = ?
WHERE account_id = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE account_id = ?;

-- name: CheckProfileExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE account_id = ?);