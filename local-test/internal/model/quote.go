package model

import "local-test/pkg/apperrors"

type PostQuoteAndNotifyParams struct {
	QuotingAccountID string
	OriginalTweetID  int64
	Content          *string
	Code             *string
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