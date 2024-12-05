package repository

import (
	"context"
	"local-test/internal/model"
	"local-test/internal/sqlc/sqlcgen"
	"local-test/pkg/apperrors"
)

func (r *Repository) LabelTweet(ctx context.Context, params *model.LabelTweetParams) error {
	// Label tweet
	if err := r.q.CreateLabel(ctx, convertToCreateLabelParams(params)); err != nil {
		return apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "label tweet",
				Err: err,
			},
		)
	}

	return nil
}

func (r *Repository) GetRecentLabels(ctx context.Context, limit int32) ([]*model.LabelCount, error) {
	// Get recent labels
	labelCounts, err := r.q.GetRecentLabels(ctx, limit)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get recent labels",
				Err: err,
			},
		)
	}

	// Convert to model
	ret := make([]*model.LabelCount, 0)
	for _, labelCount := range labelCounts {
		if labelCount.Label.Valid {
			if model.Label(labelCount.Label.TweetLabel).Validate() != nil {
				continue
			}
			ret = append(ret, &model.LabelCount{
				Label: model.Label(labelCount.Label.TweetLabel),
				Count: labelCount.LabelCount,
			})
		}
	}

	return ret, nil
}

func (r *Repository) GetLikedTweetLabelsCount(ctx context.Context, accountID string) ([]*model.LabelCount, error) {
	// Get liked tweet labels count
	labelCounts, err := r.q.GetLikedTweetLabelsCount(ctx, accountID)
	if err != nil {
		return nil, apperrors.WrapRepositoryError(
			&apperrors.ErrOperationFailed{
				Operation: "get liked tweet labels count",
				Err: err,
			},
		)
	}

	// Convert to model
	ret := make([]*model.LabelCount, 0)
	for _, labelCount := range labelCounts {
		if labelCount.Label.Valid {
			if model.Label(labelCount.Label.TweetLabel).Validate() != nil {
				continue
			}
			ret = append(ret, &model.LabelCount{
				Label: model.Label(labelCount.Label.TweetLabel),
				Count: labelCount.LabelCount,
			})
		}
	}

	return ret, nil
}

func convertToCreateLabelParams(params *model.LabelTweetParams) sqlcgen.CreateLabelParams {
	ret := sqlcgen.CreateLabelParams{
		TweetID: params.TweetID,
	}

	if params.Label1 != nil {
		ret.Label1 = sqlcgen.NullTweetLabel{
			TweetLabel: sqlcgen.TweetLabel(*params.Label1),
			Valid:  true,
		}
	}

	if params.Label2 != nil {
		ret.Label2 = sqlcgen.NullTweetLabel{
			TweetLabel: sqlcgen.TweetLabel(*params.Label2),
			Valid:  true,
		}
	}

	if params.Label3 != nil {
		ret.Label3 = sqlcgen.NullTweetLabel{
			TweetLabel: sqlcgen.TweetLabel(*params.Label3),
			Valid:  true,
		}
	}

	return ret
}