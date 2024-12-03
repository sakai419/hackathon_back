package model

import (
	"local-test/pkg/apperrors"
	"time"
)

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

type GetUserLikesParams struct {
	ClientAccountID string
	TargetAccountID string
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

type GetUserRetweetsParams struct {
	ClientAccountID string
	TargetAccountID string
	Limit           int32
	Offset          int32
}

func (p *GetUserRetweetsParams) Validate() error {
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

type GetClientProfileParams struct {
	ClientAccountID string
}

type GetUserProfileParams struct {
	ClientAccountID string
	TargetAccountID string
}

type UserProfile struct {
	UserInfo       UserInfo
	BannerImageURL string
	TweetCount     int64
	FollowerCount  int64
	FollowingCount int64
	IsFollowed     bool
	CreatedAt	   time.Time
}