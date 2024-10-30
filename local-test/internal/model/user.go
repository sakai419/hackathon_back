package model

import "local-test/pkg/apperrors"

type GetUserTweetsParams struct {
	TargetAccountID string
	ClientAccountID string
	Limit           int32
	Offset          int32
}

func (p *GetUserTweetsParams) Validate() error {
	if p.Limit <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "limit must be greater than 0",
		}
	}

	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "offset must be greater than or equal to 0",
		}
	}

	return nil
}

type GetUserTweetsResponse struct {
	Tweet             TweetInfo
	OriginalTweet     *TweetInfo
	ParentReply       *TweetInfo
	OmittedReplyExist *bool
}

type GetUserLikesParams struct {
	ClientAccountID string
	Limit           int32
	Offset          int32
}

func (p *GetUserLikesParams) Validate() error {
	if p.Limit <= 0 {
		return &apperrors.ErrInvalidInput{
			Message: "limit must be greater than 0",
		}
	}

	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "offset must be greater than or equal to 0",
		}
	}

	return nil
}