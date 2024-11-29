package model

import (
	"local-test/pkg/apperrors"
	"time"
)

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
	CreatedAt	    time.Time
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


const (
	SortTypeNewest  SortType = "newest"
	SortTypeOldest  SortType = "oldest"
)

type SortType string

func (s SortType) IsValid() bool {
	switch s {
	case SortTypeNewest, SortTypeOldest:
		return true
	}
	return false
}

type SearchUsersParams struct {
	ClientAccountID string
	Keyword string
	SortType SortType
	Offset int32
	Limit int32
}

func (p *SearchUsersParams) Validate() error {
	if p.Keyword == "" {
		return &apperrors.ErrInvalidInput{
			Message: "Keyword is required",
		}
	}
	if !p.SortType.IsValid() {
		return &apperrors.ErrInvalidInput{
			Message: "SortType is invalid",
		}
	}
	if p.Offset < 0 {
		return &apperrors.ErrInvalidInput{
			Message: "Offset must be greater than or equal to 0",
		}
	}
	if p.Limit < 1 {
		return &apperrors.ErrInvalidInput{
			Message: "Limit must be greater than or equal to 1",
		}
	}

	return nil
}

type SearchUsersOrderByCreatedAtParams struct {
	Keyword string
	Offset int32
	Limit int32
}