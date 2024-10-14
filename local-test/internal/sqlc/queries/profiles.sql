-- name: CreateProfilesWithDefaultValues :exec
INSERT INTO profiles (account_id)
VALUES ($1);

-- name: GetProfilesByAccountId :one
SELECT * FROM profiles
WHERE account_id = $1;

-- name: UpdateProfilesBio :exec
UPDATE profiles
SET bio = $1
WHERE account_id = $2;

-- name: UpdateProfilesImageUrl :exec
UPDATE profiles
SET profile_image_url = $1
WHERE account_id = $2;

-- name: UpdateBannerImageUrl :exec
UPDATE profiles
SET banner_image_url = $1
WHERE account_id = $2;

-- name: DeleteProfiles :exec
DELETE FROM profiles
WHERE account_id = $1;

-- name: CheckProfilesExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE account_id = $1);