package model

import (
	"local-test/pkg/apperrors"
	"time"
)

type Notification struct {
	ID                 int64
	SenderAccountID    *string
	Type               string
	Content            *string
	TweetID            *int64
	IsRead             bool
	CreatedAt          time.Time
}

type NotificationResponse struct {
	ID                 int64
	SenderInfo         *UserInfo
	Type               string
	Content            *string
	RelatedTweet       *TweetInfo
	IsRead             bool
	CreatedAt          time.Time
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