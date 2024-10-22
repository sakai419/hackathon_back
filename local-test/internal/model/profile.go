package model

import "errors"

type UpdateProfilesParams struct {
	AccountID       string
	Bio             *string
	ProfileImageURL *string
	BannerImageURL  *string
}

func (p *UpdateProfilesParams) Validate() error {
	if p.ProfileImageURL != nil && len(*p.ProfileImageURL) > 2083 {
		return errors.New("profile image URL is too long")
	}

	if p.BannerImageURL != nil && len(*p.BannerImageURL) > 2083 {
		return errors.New("banner image URL is too long")
	}

	return nil
}