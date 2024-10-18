package model

import (
	"local-test/pkg/apperrors"
	"time"
)

type SendMessageParams struct {
	ClientAccountID    string
	TargetAccountID    string
	Content			   string
}

func (p *SendMessageParams) Validate() error {
	if p.ClientAccountID == p.TargetAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Sender and recipient account id must be different",
		}
	}
	return nil
}

type CreateMessageParams struct {
	ConversationID  int64
	SenderAccountID string
	Content		    string
}

type GetMessagesParams struct {
	ClientAccountID string
	TargetAccountID string
	Limit		  int32
	Offset		  int32
}

func (p *GetMessagesParams) Validate() error {
	if p.ClientAccountID == p.TargetAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Sender and recipient account id must be different",
		}
	}

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

type GetMessageListParams struct {
	ConversationID  int64
	Limit		  int32
	Offset		  int32
}

type MessageResponse struct {
	SenderAccountID string
	Content string
	IsRead bool
	CreatedAt time.Time
}

type MarkMessagesAsReadParams struct {
	ClientAccountID string
	TargetAccountID string
}

func (p *MarkMessagesAsReadParams) Validate() error {
	if p.ClientAccountID == p.TargetAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Sender and recipient account id must be different",
		}
	}
	return nil
}

type MarkMessageListAsReadParams struct {
	ConversationID  int64
	SenderAccountID string
}