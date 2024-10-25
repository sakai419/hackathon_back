package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
)

func (s *Service) PostTweet(ctx context.Context, params *model.PostTweetParams) (error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}


	// Extract hashtags
	var hashtagIDs []int64
	if params.Content != nil {
		hashtags := utils.ExtractHashtags(*params.Content)
		if len(hashtags) > 0 {
			ids, err := s.repo.GetHashtagIDs(ctx, hashtags)
			if err != nil {
				return apperrors.NewInternalAppError("get hashtag IDs", err)
			}
			hashtagIDs = ids
		}
	}

	// Create tweet
	if err := s.repo.CreateTweet(ctx, &model.CreateTweetParams{
		AccountID:  params.AccountID,
		Content:    params.Content,
		Code:       params.Code,
		Media:      params.Media,
		HashtagIDs: hashtagIDs,
	}); err != nil {
		return apperrors.NewInternalAppError("create tweet", err)
	}

	return nil
}