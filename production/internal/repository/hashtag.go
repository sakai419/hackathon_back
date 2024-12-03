package repository

import (
	"context"
	"local-test/pkg/apperrors"
)

func (r *Repository) GetHashtagIDs(ctx context.Context, hashtags []string) ([]int64, error) {
	// Get hashtag IDs
	hashtagIDs, err := r.q.GetHashtagIDs(ctx, hashtags)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get hashtag IDs",
				Err: err,
			},
		)
	}

	return hashtagIDs, nil
}