-- name: CreateProfilesWithDefaultValues :exec
INSERT INTO profiles (account_id)
VALUES ($1);

-- name: GetProfilesByAccountID :one
SELECT * FROM profiles
WHERE account_id = $1;

-- name: UpdateProfiles :execresult
UPDATE profiles
SET bio = COALESCE($1, bio),
    profile_image_url = COALESCE($2, profile_image_url),
    banner_image_url = COALESCE($3, banner_image_url)
WHERE account_id = $4;

-- name: DeleteProfiles :exec
DELETE FROM profiles
WHERE account_id = $1;

-- name: CheckProfilesExists :one
SELECT EXISTS(SELECT 1 FROM profiles WHERE account_id = $1);