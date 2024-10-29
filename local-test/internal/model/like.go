package model

import "local-test/pkg/apperrors"

type LikeTweetAndNotifyParams struct {
	LikingAccountID string
	OriginalTweetID int64
}

type CreateLikeAndNotifyParams struct {
	LikingAccountID string
	LikedAccountID  string
	OriginalTweetID int64
}

type UnlikeTweetParams struct {
	LikingAccountID string
	OriginalTweetID int64
}

type GetLikingUserInfosParams struct {
	ClientAccountID string
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}

func (p *GetLikingUserInfosParams) Validate() error {
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

type GetLikingAccountIDsParams struct {
	OriginalTweetID int64
	Limit           int32
	Offset          int32
}