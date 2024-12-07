package model

import "local-test/pkg/apperrors"

const (
	SortTypeLatest  SortType = "latest"
	SortTypeOldest  SortType = "oldest"
	SortTypePopular SortType = "popular"
)

type SortType string

func (s SortType) IsValid() bool {
	switch s {
	case SortTypeLatest, SortTypeOldest, SortTypePopular:
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
	ClientAccountID string
	Keyword string
	Offset int32
	Limit int32
}

type SearchTweetsParams struct {
	ClientAccountID string
	Keyword string
	SortType SortType
	Offset int32
	Limit int32
}

func (p *SearchTweetsParams) Validate() error {
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

type SearchTweetsOrderByCreatedAtParams struct {
	ClientAccountID string
	Keyword string
	Offset int32
	Limit int32
}

type SearchTweetsOrderByEngagementScoreParams struct {
	ClientAccountID string
	Keyword string
	Offset int32
	Limit int32
}

type SearchTweetsByLabelsParams struct {
	ClientAccountID string
	Label Label
	SortType SortType
	Offset int32
	Limit int32
}

func (p *SearchTweetsByLabelsParams) Validate() error {
	if err := p.Label.Validate(); err != nil {
		return &apperrors.ErrInvalidInput{
			Message: "Label is invalid",
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

type SearchTweetsByLabelsOrderByCreatedAtParams struct {
	ClientAccountID string
	Label Label
	Offset int32
	Limit int32
}

type SearchTweetsByLabelsOrderByEngagementScoreParams struct {
	ClientAccountID string
	Label Label
	Offset int32
	Limit int32
}

type SearchTweetsByHashtagParams struct {
	ClientAccountID string
	Hashtag string
	SortType SortType
	Offset int32
	Limit int32
}

func (p *SearchTweetsByHashtagParams) Validate() error {
	if p.Hashtag == "" {
		return &apperrors.ErrInvalidInput{
			Message: "Hashtag is required",
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

type SearchTweetsByHashtagOrderByCreatedAtParams struct {
	ClientAccountID string
	Hashtag string
	Offset int32
	Limit int32
}

type SearchTweetsByHashtagOrderByEngagementScoreParams struct {
	ClientAccountID string
	Hashtag string
	Offset int32
	Limit int32
}