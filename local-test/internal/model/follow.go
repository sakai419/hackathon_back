package model

import "local-test/pkg/utils"

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

type UnfollowParams struct {
	FollowerAccountID  string
	FollowingAccountID  string
}

func (p *UnfollowParams) Validate() error {
	if p.FollowerAccountID == p.FollowingAccountID {
		return &utils.ErrInvalidInput{
			Message: "Follower account ID and following account ID must be different",
		}
	}
	return nil
}

type GetFollowerAccountIDsParams struct {
	FollowingAccountID string
	Limit			   int32
	Offset			   int32
}