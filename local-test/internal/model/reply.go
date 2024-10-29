package model

import "local-test/pkg/apperrors"

type ReplyTweetAndNotifyParams struct {
	ReplyingAccountID string
	OriginalTweetID   int64
	Content           *string
	Code 			  *string
	Media             *Media
}

func (p *ReplyTweetAndNotifyParams) Validate() error {
	if p.ReplyingAccountID == "" {
		return &apperrors.ErrInvalidInput{
			Message: "replying account id is missing",
		}
	}

	if p.OriginalTweetID <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "original tweet id is missing",
		}
	}

	if p.Content == nil && p.Code == nil && p.Media == nil {
		return &apperrors.ErrInvalidInput{
			Message: "content, code, and media are all missing",
		}
	}

	if p.Media != nil {
		if err := p.Media.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type CreateReplyAndNotifyParams struct {
	ReplyingAccountID string
	RepliedAccountID  string
	OriginalTweetID   int64
	Content           *string
	Code 			  *string
	Media             *Media
	HashtagIDs		  []int64
}