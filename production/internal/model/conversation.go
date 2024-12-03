package model

import (
	"local-test/pkg/apperrors"
	"time"
)

type GetConversationIDParams struct {
	Account1ID    string
	Account2ID  string
}

type GetConversationsParams struct {
	ClientAccountID string
	Limit     int32
	Offset    int32
}

func (p *GetConversationsParams) Validate() error {
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

type GetConversationListParams struct {
	AccountID string
	Limit     int32
	Offset    int32
}

type Conversation struct {
	ID              int64
	OpponentID      string
	LastMessageTime time.Time
	Content         string
	SenderUserID    string
	IsRead          bool
}

type ConversationResponse struct {
	ID              int64
	OpponentInfo    UserInfoWithoutBio
	LastMessageTime time.Time
	Content         string
	SenderUserID    string
	IsRead	        bool
}