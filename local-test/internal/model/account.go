package model

import (
	"local-test/pkg/apperrors"
)

// CreateAccount
type CreateAccountParams struct {
	ID       string
	UserID   string
	UserName string
}

func (p *CreateAccountParams) Validate() error {
	if len(p.ID) != 28 {
		return &apperrors.ErrInvalidInput{
			Message: "ID must be 28 characters",
		}
	}
	if len(p.UserID) > 30 {
		return &apperrors.ErrInvalidInput{
			Message: "UserID must be less than 30 characters",
		}
	}
	if len(p.UserName) > 30 {
		return &apperrors.ErrInvalidInput{
			Message: "UserName must be less than 30 characters",
		}
	}
	return nil
}

type UserInfoInternal struct {
	ID              string
	UserID          string
	UserName        string
	Bio	            string
	ProfileImageURL string
	IsPrivate       bool
	IsAdmin		    bool
}

type UserInfo struct {
	UserID          string
	UserName        string
	Bio	            string
	ProfileImageURL string
	IsPrivate       bool
	IsAdmin		    bool
}

type UserInfoWithoutBio struct {
	UserID          string
	UserName        string
	ProfileImageURL string
	IsPrivate       bool
	IsAdmin		    bool
}

type AccountInfo struct {
	IsAdmin	    bool
	IsSuspended bool
	IsPrivate   bool
}