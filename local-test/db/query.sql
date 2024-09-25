-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = ?;

-- name: GetAccountByUserId :one
SELECT * FROM accounts
WHERE user_id = ?;

-- name: SearchAccountsByUserId :many
SELECT * FROM accounts
WHERE user_id LIKE "%?%";

-- name: GetAccountByUserName :one
SELECT * FROM accounts
WHERE user_name = ?;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts
WHERE user_name LIKE "%?%";

-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES (?, ?, ?);

-- name: UpdateAccountUserId :exec
UPDATE accounts
SET user_id = ?
WHERE id = ?;

-- name: UpdateAccountUserName :exec
UPDATE accounts
SET user_name = ?
WHERE id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;

-- name: GetProfileByAccountId :one
SELECT * FROM profiles
WHERE account_id = ?;

-- name: CreateProfile :exec
INSERT INTO profiles (account_id, bio, profile_image_url, banner_image_url)
VALUES (?, ?, ?, ?);

-- name: UpdateProfileBio :exec
UPDATE profiles
SET bio = ?
WHERE account_id = ?;

-- name: UpdateProfileProfileImageUrl :exec
UPDATE profiles
SET profile_image_url = ?
WHERE account_id = ?;

-- name: UpdateProfileBannerImageUrl :exec
UPDATE profiles
SET banner_image_url = ?
WHERE account_id = ?;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE account_id = ?;

-- name: GetSettingByAccountId :one
SELECT * FROM settings
WHERE account_id = ?;

-- name: CreateSetting :exec
INSERT INTO settings (account_id, is_private)
VALUES (?, ?);

-- name: UpdateSettingIsPrivate :exec
UPDATE settings
SET is_private = ?
WHERE account_id = ?;

-- name: DeleteSetting :exec
DELETE FROM settings
WHERE account_id = ?;

-- name: GetInterestByAccountId :one
SELECT * FROM interests
WHERE account_id = ?;

-- name: CreateInterestsWithDefaultScores :exec
INSERT INTO interests (account_id)
VALUES (?);

-- name: UpdateInterestScore :exec
UPDATE interests
SET
    news_score = COALESCE(?, news_score),
    politics_score = COALESCE(?, politics_score),
    economics_score = COALESCE(?, economics_score),
    health_score = COALESCE(?, health_score),
    sports_score = COALESCE(?, sports_score),
    entertainment_score = COALESCE(?, entertainment_score),
    art_score = COALESCE(?, art_score),
    cooking_score = COALESCE(?, cooking_score),
    travel_score = COALESCALE(?, travel_score),
    fashion_score = COALESCE(?, fashion_score),
    beauty_score = COALESCE(?, beauty_score),
    pets_score = COALESCE(?, pets_score),
    parenting_score = COALESCE(?, parenting_score),
    education_score = COALESCE(?, education_score),
    environment_score = COALESCE(?, environment_score),
    climate_score = COALESCE(?, climate_score),
    space_score = COALESCE(?, space_score),
    mental_health_score = COALESCE(?, mental_health_score),
    fitness_score = COALESCE(?, fitness_score),
    reading_score = COALESCE(?, reading_score),
    history_score = COALESCE(?, history_score),
    philosophy_score = COALESCE(?, philosophy_score),
    religion_score = COALESCE(?, religion_score),
    culture_score = COALESCE(?, culture_score),
    volunteer_score = COALESCE(?, volunteer_score),
    social_issues_score = COALESCE(?, social_issues_score),
    law_score = COALESCE(?, law_score),
    taxes_score = COALESCE(?, taxes_score),
    investment_score = COALESCE(?, investment_score),
    real_estate_score = COALESCE(?, real_estate_score),
    diy_score = COALESCE(?, diy_score),
    gardening_score = COALESCE(?, gardening_score),
    interior_design_score = COALESCE(?, interior_design_score),
    automotive_score = COALESCE(?, automotive_score),
    gaming_score = COALESCE(?, gaming_score),
    anime_manga_score = COALESCE(?, anime_manga_score),
    creative_works_score = COALESCE(?, creative_works_score),
    photography_video_score = COALESCE(?, photography_video_score),
    media_score = COALESCE(?, media_score),
    marketing_score = COALESCE(?, marketing_score),
    branding_score = COALESCE(?, branding_score),
    entrepreneurship_score = COALESCE(?, entrepreneurship_score),
    remote_work_score = COALESCE(?, remote_work_score),
    data_science_score = COALESCE(?, data_science_score),
    iot_score = COALESCE(?, iot_score),
    robotics_engineering_score = COALESCE(?, robotics_engineering_score),
    biotechnology_score = COALESCE(?, biotechnology_score),
    nanotechnology_score = COALESCE(?, nanotechnology_score),
    energy_technology_score = COALESCE(?, energy_technology_score),
    archaeology_score = COALESCE(?, archaeology_score),
    psychology_score = COALESCE(?, psychology_score),
    sociology_score = COALESCE(?, sociology_score),
    anthropology_score = COALESCE(?, anthropology_score),
    geography_score = COALESCE(?, geography_score),
    geology_score = COALESCE(?, geology_score),
    meteorology_score = COALESCE(?, meteorology_score),
    disaster_emergency_management_score = COALESCE(?, disaster_emergency_management_score),
    urban_planning_score = COALESCE(?, urban_planning_score),
    architecture_score = COALESCE(?, architecture_score),
    agriculture_score = COALESCE(?, agriculture_score),
    nutrition_science_score = COALESCE(?, nutrition_science_score),
    sleep_science_score = COALESCE(?, sleep_science_score),
    productivity_score = COALESCE(?, productivity_score),
    leadership_score = COALESCE(?, leadership_score),
    international_relations_score = COALESCE(?, international_relations_score),
    future_predictions_score = COALESCE(?, future_predictions_score),
    events_score = COALESCE(?, events_score),
    community_score = COALESCE(?, community_score),
    trends_score = COALESCE(?, trends_score),
    lifestyle_score = COALESCE(?, lifestyle_score),
    software_development_score = COALESCE(?, software_development_score),
    programming_languages_score = COALESCE(?, programming_languages_score),
    web_development_score = COALESCE(?, web_development_score),
    mobile_app_development_score = COALESCE(?, mobile_app_development_score),
    debugging_techniques_score = COALESCE(?, debugging_techniques_score),
    algorithms_mathematics_score = COALESCE(?, algorithms_mathematics_score),
    database_design_score = COALESCE(?, database_design_score),
    cloud_computing_score = COALESCE(?, cloud_computing_score),
    server_management_score = COALESCE(?, server_management_score),
    network_security_score = COALESCE(?, network_security_score),
    cryptography_score = COALESCE(?, cryptography_score),
    artificial_intelligence_score = COALESCE(?, artificial_intelligence_score),
    machine_learning_score = COALESCE(?, machine_learning_score),
    deep_learning_score = COALESCE(?, deep_learning_score),
    computer_vision_score = COALESCE(?, computer_vision_score),
    natural_language_processing_score = COALESCE(?, natural_language_processing_score),
    blockchain_technology_score = COALESCE(?, blockchain_technology_score),
    quantum_computing_score = COALESCE(?, quantum_computing_score),
    edge_computing_score = COALESCE(?, edge_computing_score),
    microservices_architecture_score = COALESCE(?, microservices_architecture_score),
    devops_score = COALESCE(?, devops_score),
    container_technology_score = COALESCE(?, container_technology_score),
    ci_cd_score = COALESCE(?, ci_cd_score),
    test_automation_score = COALESCE(?, test_automation_score),
    ux_ui_design_score = COALESCE(?, ux_ui_design_score),
    agile_development_methodologies_score = COALESCE(?, agile_development_methodologies_score),
    open_source_score = COALESCE(?, open_source_score),
    version_control_score = COALESCE(?, version_control_score),
    api_design_score = COALESCE(?, api_design_score),
    performance_optimization_score = COALESCE(?, performance_optimization_score),
WHERE account_id = ?;

-- name: DeleteInterests :exec
DELETE FROM interests
WHERE account_id = ?;

-- name: GetNotificationByAccountId :many
SELECT * FROM notifications
WHERE recipient_account_id = ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;

-- name: GetUnreadNotificationsByAccountId :many
SELECT * FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;

-- name: GetUnreadNotificationCount :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: CreateNotification :exec
INSERT INTO notifications (recipient_account_id, sender_account_id, type, content, is_read)
VALUES (?, ?, ?, ?, ?);

-- name: UpdateNotificationsIsRead :exec
UPDATE notifications
SET is_read = ?
WHERE id IN (?) AND recipient_account_id = ?;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = ?;

-- name: GetNotificationById :one
SELECT * FROM notifications
WHERE id = ? AND recipient_account_id = ?;

-- name: DeleteOldNotifications :exec
DELETE FROM notifications
WHERE created_at < ? AND is_read = TRUE;
