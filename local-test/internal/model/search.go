package model

import "local-test/pkg/apperrors"

const (
	SortTypeLatest  SortType = "latest"
	SortTypeOldest  SortType = "oldest"
)

type SortType string

func (s SortType) IsValid() bool {
	switch s {
	case SortTypeLatest, SortTypeOldest:
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