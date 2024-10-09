-- name: CreateProfilesWithDefaultValues :exec
INSERT INTO profiles (account_id)
VALUES (?);

-- name: GetProfilesByAccountId :one
SELECT * FROM profiles
WHERE account_id = ?;

-- name: UpdateProfilesBio :exec
UPDATE profiles
SET bio = ?
WHERE account_id = ?;

-- name: UpdateProfilesImageUrl :exec
UPDATE profiles
SET profile_image_url = ?
WHERE account_id = ?;

-- name: UpdateBannerImageUrl :exec
UPDATE profiles
SET banner_image_url = ?
WHERE account_id = ?;

-- name: DeleteProfiles :exec
DELETE FROM profiles
WHERE account_id = ?;

-- name: CheckProfilesExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE account_id = ?);