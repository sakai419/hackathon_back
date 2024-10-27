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