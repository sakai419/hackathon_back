package model

import (
	"local-test/pkg/apperrors"
	"time"
)

type NotificationResponse struct {
	ID                 int64      `json:"id"`
	SenderAccountID    *string    `json:"sender_account_id"`
	Type               string     `json:"notification_type"`
	Content            *string    `json:"content"`
	TweetID            *int64	  `json:"tweet_id"`
	IsRead             bool       `json:"is_read"`
	CreatedAt          time.Time  `json:"created_at"`
}

type GetNotificationsParams struct {
	RecipientAccountID string
	Limit			   int32
	Offset			   int32
}

func (p GetNotificationsParams) Validate() error {
	if p.Limit <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "Limit must be greater than 0",
		}
	}

	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "Offset must be greater than or equal to 0",
		}
	}

	return nil
}

type GetUnreadNotificationsParams struct {
	RecipientAccountID string
	Limit			   int32
	Offset			   int32
}

func (p GetUnreadNotificationsParams) Validate() error {
	if p.Limit <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "Limit must be greater than 0",
		}
	}

	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "Offset must be greater than or equal to 0",
		}
	}

	return nil
}

type MarkNotificationAsReadParams struct {
	ID                 int64
	RecipientAccountID string
}