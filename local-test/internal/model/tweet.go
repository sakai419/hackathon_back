package model

import (
	"local-test/pkg/apperrors"
)

const (
	MediaTypeJPG = "jpg"
)

type Media struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (m *Media) Validate() error {
	switch m.Type {
	case MediaTypeJPG:
		if m.URL == "" {
			return &apperrors.ErrInvalidInput{
				Message: "media url is missing",
			}
		}
	default:
		return &apperrors.ErrInvalidInput{
			Message: "media type is invalid",
		}
	}

	return nil
}

type PostTweetParams struct {
	AccountID       string
	Content         *string
	Code            *string
	Media           *Media
}

func (p *PostTweetParams) Validate() error {
	if p.Content == nil && p.Code == nil && p.Media == nil {
		return &apperrors.ErrInvalidInput{
			Message: "content, code, and media are all missing",
		}
	}

	if p.Media != nil {
		if err := p.Media.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type CreateTweetParams struct {
	AccountID   string
	Content     *string
	Code        *string
	Media       *Media
	HashtagIDs  []int64
}

type GetTweetLabelsParams struct {
	Content *string
	Code    *string
	Media   *Media
}