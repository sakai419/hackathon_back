package model

import (
	"errors"
	"local-test/pkg/apperrors"
	"regexp"
	"time"
)

type CreateAccountParams struct {
	ID       string
	UserID   string
	UserName string
}

func validateUserID(userID string) error {
	var validPattern = regexp.MustCompile(`^[a-zA-Z0-9._-]+$`)
	if !validPattern.MatchString(userID) {
		return errors.New("invalid user_id: only alphanumeric characters, '-', '_', and '.' are allowed")
	}
	return nil
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
	if err := validateUserID(p.UserID); err != nil {
		return &apperrors.ErrInvalidInput{
			Message: err.Error(),
		}
	}
	if len(p.UserName) > 30 {
		return &apperrors.ErrInvalidInput{
			Message: "UserName must be less than 30 characters",
		}
	}
	return nil
}

type FilterAccesibleAccountIDsParams struct {
	AccountIDs []string
	ClientAccountID string
}

type UserInfoInternal struct {
	ID              string
	UserID          string
	UserName        string
	Bio	            string
	ProfileImageURL string
	BannerImageURL  string
	IsPrivate       bool
	IsAdmin		    bool
	IsFollowing     bool
	IsFollowed      bool
	IsPending       bool
	CreatedAt	    time.Time
}

type GetUserInfoParams struct {
	TargetAccountID string
	ClientAccountID string
}

type GetUserInfosParams struct {
	TargetAccountIDs []string
	ClientAccountID string
}

type UserInfo struct {
	UserID          string
	UserName        string
	Bio	            string
	ProfileImageURL string
	IsPrivate       bool
	IsAdmin		    bool
	IsFollowing     bool
	IsFollowed      bool
	IsPending       bool
}

type UserInfoWithoutBio struct {
	UserID          string
	UserName        string
	ProfileImageURL string
	IsPrivate       bool
	IsAdmin		    bool
	IsFollowing     bool
	IsFollowed      bool
	IsPending       bool
}

type AccountInfo struct {
	IsAdmin	    bool
	IsSuspended bool
	IsPrivate   bool
}
