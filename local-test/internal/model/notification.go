package model

import "time"

type Notification struct {
	ID                 int64     `json:"id"`
	SenderAccountID    string    `json:"sender_account_id"`
	RecipientAccountID string    `json:"recipient_account_id"`
	Type               string    `json:"notification_type"`
	Content            string    `json:"content"`
	IsRead             bool      `json:"is_read"`
	CreatedAt          time.Time `json:"created_at"`
}

type GetNotificationsParams struct {
	RecipientAccountID string
	Limit			   int32
	Offset			   int32
}

type GetUnreadNotificationParams struct {
	RecipientAccountID string
	Limit			   int32
	Offset			   int32
}

type MarkNotificationAsReadParams struct {
	ID                 int64
	RecipientAccountID string
}

type DeleteNotificationParams struct {
	ID                 int64
	RecipientAccountID string
}