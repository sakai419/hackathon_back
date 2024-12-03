package model

import (
	"local-test/pkg/apperrors"
)

type FollowAndNotifyParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

func (p *FollowAndNotifyParams) Validate() error {
	if p.FollowerAccountID == p.FollowingAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Follower account ID and following account ID must be different",
		}
	}
	return nil
}

type UnfollowParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

func (p *UnfollowParams) Validate() error {
	if p.FollowerAccountID == p.FollowingAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Follower account ID and following account ID must be different",
		}
	}
	return nil
}

type GetFollowerInfosParams struct {
	ClientAccountID	      string
	FollowingAccountID    string
	Limit			      int32
	Offset			      int32
}

func (p *GetFollowerInfosParams) Validate() error {
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

type GetFollowerAccountIDsParams struct {
	FollowingAccountID string
	Limit			   int32
	Offset			   int32
}

type GetFollowingInfosParams struct {
	ClientAccountID    string
	FollowerAccountID  string
	Limit			   int32
	Offset			   int32
}

func (p *GetFollowingInfosParams) Validate() error {
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

type GetFollowingAccountIDsParams struct {
	FollowerAccountID  string
	Limit			   int32
	Offset			   int32
}

type RequestFollowAndNotifyParams struct {
	RequesterAccountID string
	RequestedAccountID string
}

func (p *RequestFollowAndNotifyParams) Validate() error {
	if p.RequesterAccountID == p.RequestedAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Requester account ID and requested user ID must be different",
		}
	}
	return nil
}

type AcceptFollowRequestAndNotifyParams struct {
	RequestedAccountID string
	RequesterAccountID string
}

func (p *AcceptFollowRequestAndNotifyParams) Validate() error {
	if p.RequestedAccountID == p.RequesterAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Requested account ID and requester account ID must be different",
		}
	}
	return nil
}

type RejectFollowRequestParams struct {
	RequesterAccountID string
	RequestedAccountID string
}

func (p *RejectFollowRequestParams) Validate() error {
	if p.RequesterAccountID == p.RequestedAccountID {
		return &apperrors.ErrInvalidInput{
			Message: "Requester account ID and requested account ID must be different",
		}
	}
	return nil
}

type GetFollowCountsParams struct {
	AccountID string
}

type FollowCounts struct {
	FollowersCount int64
	FollowingCount int64
}

type CheckIsFollowedParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

type IsPrivateAndNotFollowingParams struct {
	ClientAccountID  string
	TargetAccountID  string
}