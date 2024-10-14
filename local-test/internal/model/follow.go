package model

import "local-test/pkg/utils"

type FollowUserParams struct {
	FollowerAccountID   string
	FollowingUserID 	string
}

type FollowAndNotifyParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

func (p *FollowAndNotifyParams) Validate() error {
	if p.FollowerAccountID == p.FollowingAccountID {
		return &utils.ErrInvalidInput{
			Message: "Follower account ID and following account ID must be different",
		}
	}
	return nil
}

type UnfollowUserParams struct {
	FollowerAccountID   string
	FollowingUserID 	string
}

type DeleteFollowParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

func (p *DeleteFollowParams) Validate() error {
	if p.FollowerAccountID == p.FollowingAccountID {
		return &utils.ErrInvalidInput{
			Message: "Follower account ID and following account ID must be different",
		}
	}
	return nil
}

type GetFollowerInfosParams struct {
	FollowingUserID    string
	Limit			   int32
	Offset			   int32
}

func (p *GetFollowerInfosParams) Validate() error {
	if p.Limit < 1 {
		return &utils.ErrInvalidInput{
			Message: "Limit must be greater than 0",
		}
	}
	if p.Offset < 0 {
		return &utils.ErrInvalidInput{
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
	FollowerUserID    string
	Limit			   int32
	Offset			   int32
}

func (p *GetFollowingInfosParams) Validate() error {
	if p.Limit < 1 {
		return &utils.ErrInvalidInput{
			Message: "Limit must be greater than 0",
		}
	}
	if p.Offset < 0 {
		return &utils.ErrInvalidInput{
			Message: "Offset must be greater than or equal to 0",
		}
	}
	return nil
}

type GetFollowingAccountIDsParams struct {
	FollowerAccountID string
	Limit			   int32
	Offset			   int32
}

type RequestFollowParams struct {
	RequesterAccountID string
	RequestedUserID string
}

type RequestFollowAndNotifyParams struct {
	RequesterAccountID string
	RequestedAccountID string
}

func (p *RequestFollowAndNotifyParams) Validate() error {
	if p.RequesterAccountID == p.RequestedAccountID {
		return &utils.ErrInvalidInput{
			Message: "Requester account ID and requested user ID must be different",
		}
	}
	return nil
}

type AcceptFollowRequestParams struct {
	RequesterUserID string
	RequestedAccountID string
}

type AcceptFollowRequestAndNotifyParams struct {
	RequestedAccountID string
	RequesterAccountID string
}

func (p *AcceptFollowRequestAndNotifyParams) Validate() error {
	if p.RequestedAccountID == p.RequesterAccountID {
		return &utils.ErrInvalidInput{
			Message: "Requested account ID and requester account ID must be different",
		}
	}
	return nil
}

type RejectFollowRequestParams struct {
	RequesterUserID string
	RequestedAccountID string
}

type DeleteFollowRequestParams struct {
	RequesterAccountID string
	RequestedAccountID string
}

func (p *DeleteFollowRequestParams) Validate() error {
	if p.RequesterAccountID == p.RequestedAccountID {
		return &utils.ErrInvalidInput{
			Message: "Requester account ID and requested account ID must be different",
		}
	}
	return nil
}