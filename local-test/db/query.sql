-- name: CreateAccount :exec
INSERT INTO accounts (id, user_id, user_name)
VALUES (?, ?, ?);

-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = ?;

-- name: GetAccountByUserId :one
SELECT * FROM accounts
WHERE user_id = ?;

-- name: GetAccountByUserName :one
SELECT * FROM accounts
WHERE user_name = ?;

-- name: UpdateAccountUserName :exec
UPDATE accounts
SET user_name = ?
WHERE id = ?;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?;

-- name: SearchAccountsByUserName :many
SELECT * FROM accounts
WHERE user_name LIKE ?
ORDER BY user_name
LIMIT ? OFFSET ?;

-- name: GetAccountCreationDate :one
SELECT created_at FROM accounts
WHERE id = ?;

-- name: CountAccounts :one
SELECT COUNT(*) FROM accounts;

-- name: CheckUserNameExists :one
SELECT EXISTS(SELECT 1 FROM accounts WHERE user_name = ?);

-- name: CheckUserIdExists :one
SELECT EXISTS(SELECT 1 FROM accounts WHERE user_id = ?);

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

-- name: CreateSettings :exec
INSERT INTO settings (account_id, is_private)
VALUES (?, ?);

-- name: GetSettingsByAccountId :one
SELECT * FROM settings
WHERE account_id = ?;

-- name: UpdateSettingsPrivacy :exec
UPDATE settings
SET is_private = ?
WHERE account_id = ?;

-- name: DeleteSettings :exec
DELETE FROM settings
WHERE account_id = ?;

-- name: CheckSettingsExist :one
SELECT EXISTS(SELECT 1 FROM settings WHERE account_id = ?);

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

-- name: CreateNotification :exec
INSERT INTO notifications (sender_account_id, recipient_account_id, type, content)
VALUES (?, ?, ?, ?);

-- name: GetNotificationById :one
SELECT * FROM notifications
WHERE id = ?;

-- name: GetNotificationsByRecipientId :many
SELECT * FROM notifications
WHERE recipient_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetUnreadNotificationsByRecipientId :many
SELECT * FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: MarkNotificationAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE id = ? AND recipient_account_id = ?;

-- name: MarkAllNotificationsAsRead :exec
UPDATE notifications
SET is_read = TRUE
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: DeleteNotification :exec
DELETE FROM notifications
WHERE id = ? AND recipient_account_id = ?;

-- name: DeleteAllNotificationsForRecipient :exec
DELETE FROM notifications
WHERE recipient_account_id = ?;

-- name: GetNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ?;

-- name: GetUnreadNotificationCountByRecipientId :one
SELECT COUNT(*) FROM notifications
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: GetNotificationsByType :many
SELECT * FROM notifications
WHERE recipient_account_id = ? AND type = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: DeleteOldReadNotifications :exec
DELETE FROM notifications
WHERE recipient_account_id = ? AND is_read = TRUE AND created_at < ?;

-- name: CreateMessage :exec
INSERT INTO messages (sender_account_id, recipient_account_id, content)
VALUES (?, ?, ?);

-- name: GetMessageById :one
SELECT * FROM messages
WHERE id = ?;

-- name: GetMessagesBetweenUsers :many
SELECT * FROM messages
WHERE (sender_account_id = ? AND recipient_account_id = ?)
    OR (sender_account_id = ? AND recipient_account_id = ?)
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetUnreadMessagesForUser :many
SELECT * FROM messages
WHERE recipient_account_id = ? AND is_read = FALSE
ORDER BY created_at DESC;

-- name: MarkMessageAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE id = ? AND recipient_account_id = ?;

-- name: MarkAllMessagesAsRead :exec
UPDATE messages
SET is_read = TRUE
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: DeleteMessage :exec
DELETE FROM messages
WHERE id = ? AND (sender_account_id = ? OR recipient_account_id = ?);

-- name: GetMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE sender_account_id = ? OR recipient_account_id = ?;

-- name: GetUnreadMessageCountForUser :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = ? AND is_read = FALSE;

-- name: GerUnreadMessageCountBetweenUsers :one
SELECT COUNT(*) FROM messages
WHERE recipient_account_id = ? AND sender_account_id = ? AND is_read = FALSE;

-- name: GetLatestMessageForEachConversation :many
SELECT m.*
FROM messages m
INNER JOIN (
    SELECT
        CASE
            WHEN sender_account_id < recipient_account_id
            THEN sender_account_id
            ELSE recipient_account_id
        END AS user1,
        CASE
            WHEN sender_account_id < recipient_account_id
            THEN recipient_account_id
            ELSE sender_account_id
        END AS user2,
        MAX(created_at) AS max_created_at
    FROM messages
    WHERE sender_account_id = ? OR recipient_account_id = ?
    GROUP BY user1, user2
) latest ON (
    (m.sender_account_id = latest.user1 AND m.recipient_account_id = latest.user2) OR
    (m.sender_account_id = latest.user2 AND m.recipient_account_id = latest.user1)
) AND m.created_at = latest.max_created_at
ORDER BY m.created_at DESC;

-- name: SearchMessages :many
SELECT * FROM messages
WHERE (sender_account_id = ? OR recipient_account_id = ?)
    AND content LIKE ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: DeleteOldMessages :exec
DELETE FROM messages
WHERE created_at < ? AND is_read = TRUE;

-- name: CreateFollow :exec
INSERT INTO follows (follower_account_id, following_account_id)
VALUES (?, ?);

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE follower_account_id = ? AND following_account_id = ?;

-- name: CheckFollowExists :one
SELECT EXISTS(
    SELECT 1 FROM follows
    WHERE follower_account_id = ? AND following_account_id = ?
) AS is_following;

-- name: GetFollowers: :many
SELECT follower_account_id
FROM follows
WHERE following_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetFollowing :many
SELECT following_account_id
FROM follows
WHERE follower_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetFollowerCount :one
SELECT COUNT(*) FROM follows
WHERE following_account_id = ?;

-- name: GetFollowingCount :one
SELECT COUNT(*) FROM follows
WHERE follower_account_id = ?;

-- name: GetMutualFollows :many
SELECT f1.following_account_id,
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = ? AND f2.following_account_id = ?
LIMIT ? OFFSET ?;

-- name: GetFollowSuggestions :many
SELECT DISTINCT f2.following_account_id
FROM follows f1
JOIN follows f2 ON f1.following_account_id = f2.follower_account_id
WHERE f1.follower_account_id = ?
    AND f2.following_account_id != f1.follower_account_id
    AND NOT EXISTS (
        SELECT 1 FROM follows f3
        WHERE f3.follower_account_id = f1.follower_account_id
            AND f3.following_account_id = f2.following_account_id
    )
LIMIT ?;

-- name: CreateFollowRequest :exec
INSERT INTO follow_requests (requester_account_id, requestee_account_id)
VALUES (?, ?);

-- name: DeleteFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = ? AND requestee_account_id = ?;

-- name: GetPendingFollowRequests :many
SELECT requester_account_id
FROM follow_requests
WHERE requestee_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetSentFollowRequests :many
SELECT requestee_account_id
FROM follow_requests
WHERE requester_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetPendingFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requestee_account_id = ?;

-- name: GetSentFollowRequestCount :one
SELECT COUNT(*) FROM follow_requests
WHERE requester_account_id = ?;

-- name: DeleteOldFollowRequests :exec
DELETE FROM follow_requests
WHERE created_at < ? AND requestee_account_id = ?;

-- name: AcceptFollowRequest :exec
INSERT INTO follows (follower_account_id, following_account_id)
SELECT requester_account_id, requestee_account_id
FROM follow_requests
WHERE requester_account_id = ? AND requestee_account_id = ?;

-- name: RejectFollowRequest :exec
DELETE FROM follow_requests
WHERE requester_account_id = ? AND requestee_account_id = ?;

-- name: CreateBlock :exec
INSERT INTO blocks (blocker_account_id, blocked_account_id)
VALUES (?, ?);

-- name: DeleteBlock :exec
DELETE FROM blocks
WHERE blocker_account_id = ? AND blocked_account_id = ?;

-- name: CheckBlockExists :one
SELECT EXISTS(
    SELECT 1 FROM blocks
    WHERE blocker_account_id = ? AND blocked_account_id = ?
) AS is_blocked;

-- name: GetBlockedUsers :many
SELECT blocked_account_id
FROM blocks
WHERE blocker_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetBlockersOfUser :many
SELECT blocker_account_id
FROM blocks
WHERE blocked_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetBlockCount :one
SELECT COUNT(*) FROM blocks
WHERE blocker_account_id = ?;

-- name: GetBlockedByCount :one
SELECT COUNT(*) FROM blocks
WHERE blocked_account_id = ?;

-- name: DeleteAllBlocksForUser :exec
DELETE FROM blocks
WHERE blocker_account_id = ? OR blocked_account_id = ?;

-- name: CreateTweet :exec
INSERT INTO tweets (
    account_id, content, code, is_retweet, is_reply,
    original_tweet_id, quoted_tweet_id, reply_to_tweet_id,
    reply_to_account_id, media
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetTweetById :one
SELECT * FROM tweets WHERE id = ?;

-- name: GetTweetsByAccountId :many
SELECT * FROM tweets
WHERE account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: UpdateTweetContent :exec
UPDATE tweets
SET content = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND account_id = ?;

-- name: UpdateTweetCode :exec
UPDATE tweets
SET code = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ? AND account_id = ?;

-- name: DeleteTweet :exec
DELETE FROM tweets WHERE id = ? AND account_id = ?;

-- name: IncrementLikesCount :exec
UPDATE tweets SET likes_count = likes_count + 1 WHERE id = ?;

-- name: IncrementRepliesCount :exec
UPDATE tweets SET replies_count = replies_count + 1 WHERE id = ?;

-- name: IncrementRetweetsCount :exec
UPDATE tweets SET retweets_count = retweets_count + 1 WHERE id = ?;

-- name: UpdateEngagementScore :exec
UPDATE tweets
SET engagement_score = likes_count + replies_count + retweets_count
WHERE id = ?;

-- name: GetRepliesForTweet :many
SELECT * FROM tweets
WHERE reply_to_tweet_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetRetweetsForTweet :many
SELECT * FROM tweets
WHERE original_tweet_id = ? AND is_retweet = TRUE
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetQuotesForTweet :many
SELECT * FROM tweets
WHERE quoted_tweet_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTrendingTweets :many
SELECT * FROM tweets
ORDER BY engagement_score DESC
LIMIT ?;

-- name: GetPinnedTweetForAccount :one
SELECT * FROM tweets
WHERE account_id = ? AND is_pinned = TRUE
LIMIT 1;

-- name: SetTweetAsPinned :exec
UPDATE tweets
SET is_pinned = TRUE
WHERE id = ? AND account_id = ?;

-- name: UnpinTweet :exec
UPDATE tweets
SET is_pinned = FALSE
WHERE id = ? AND account_id = ?;

-- name: SearchTweets :many
SELECT * FROM tweets
WHERE content LIKE ? OR code LIKE ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTweetCountByAccountId :one
SELECT COUNT(*) FROM tweets WHERE account_id = ?;

-- name: CreateLabel :exec
INSERT INTO labels (tweet_id, label1, label2, label3)
VALUES (?, ?, ?, ?);

-- name: GetLabelsByTweetId :one
SELECT * FROM labels
WHERE tweet_id = ?;

-- name: UpdateLabels :exec
UPDATE labels
SET label1 = ?, label2 = ?, label3 = ?
WHERE tweet_id = ?;

-- name: DeleteLabel :exec
DELETE FROM labels
WHERE tweet_id = ?;

-- name: GetTweetsByLabel :many
SELECT t.* FROM tweets t
JOIN labels l ON t.id = l.tweet_id
WHERE l.label1 = ? OR l.label2 = ? OR l.label3 = ?
ORDER BY t.created_at DESC
LIMIT ? OFFSET ?;

-- name: GetTweetsWithoutLabels :many
SELECT t.* FROM tweets t
LEFT JOIN labels l ON t.id = l.tweet_id
WHERE l.tweet_id IS NULL
ORDER BY t.created_at DESC
LIMIT ? OFFSET ?;

-- name: CreateLike :exec
INSERT INTO likes (liking_account_id, original_tweet_id)
VALUES (?, ?);

-- name: DeleteLike :exec
DELETE FROM likes
WHERE liking_account_id = ? AND original_tweet_id = ?;

-- name: GetLikesByTweetId :many
SELECT * FROM likes
WHERE original_tweet_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetLikesByAccountId :many
SELECT * FROM likes
WHERE liking_account_id = ?
ORDER BY created_at DESC
LIMIT ? OFFSET ?;

-- name: GetLikeCount :one
SELECT COUNT(*) FROM likes
WHERE original_tweet_id = ?;