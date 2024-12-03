package model

import (
	"local-test/pkg/apperrors"
)

type UpdateProfilesParams struct {
	AccountID       string
	UserID          *string
	UserName        *string
	Bio             *string
	ProfileImageURL *string
	BannerImageURL  *string
}

func (p *UpdateProfilesParams) Validate() error {
	if p.ProfileImageURL != nil && len(*p.ProfileImageURL) > 2083 {
		return &apperrors.ErrInvalidInput{
			Message: "profile image URL is too long",
		}
	}

	if p.BannerImageURL != nil && len(*p.BannerImageURL) > 2083 {
		return &apperrors.ErrInvalidInput{
			Message: "banner image URL is too long",
		}
	}

	return nil
}