package model

import "local-test/pkg/apperrors"

type PostQuoteAndNotifyParams struct {
	QuotingAccountID string
	OriginalTweetID  int64
	Content          *string
	Code             *Code
	Media            *Media
}

func (p *PostQuoteAndNotifyParams) Validate() error {
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

type CreateQuoteAndNotifyParams struct {
	QuotingAccountID string
	QuotedAccountID  string
	OriginalTweetID  int64
	Content          *string
	Code			 *Code
	Media            *Media
	HashtagIDs		 []int64
}

type GetQuotingUserInfosParams struct {
	ClientAccountID string
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

func (p *GetQuotingUserInfosParams) Validate() error {
	if p.Limit < 1 {
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

type GetQuotingAccountIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

type GetQuotedTweetInfosParams struct {
	ClientAccountID string
	QuotingTweetIDs []int64
}

type QuotedTweetInfoInternal struct {
	QuotedTweet    TweetInfoInternal
	QuotingTweetID int64
}