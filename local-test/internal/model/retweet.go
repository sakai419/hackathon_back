package model

import "local-test/pkg/apperrors"

type RetweetAndNotifyParams struct {
	RetweetingAccountID string
	OriginalTweetID	    int64
}

type CreateRetweetAndNotifyParams struct {
	RetweetingAccountID string
	RetweetedAccountID  string
	OriginalTweetID     int64
}

type UnretweetParams struct {
	RetweetingAccountID string
	OriginalTweetID     int64
}

type GetRetweetingUserInfosParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

func (p *GetRetweetingUserInfosParams) Validate() error {
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

type GetRetweetingAccountIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}