package service

import (
	"context"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"log"
	"time"
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
	tweetID, err := s.repo.CreateTweet(ctx, &model.CreateTweetParams{
		AccountID:  params.AccountID,
		Content:    params.Content,
		Code:       params.Code,
		Media:      params.Media,
		HashtagIDs: hashtagIDs,
	})
	if err != nil {
		return apperrors.NewInternalAppError("create tweet", err)
	}

	go func(params *model.PostTweetParams) {
		// Get tweet label
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		labels := getTweetLabels(timeoutCtx, params)
		if err := s.repo.LabelTweet(timeoutCtx, &model.LabelTweetParams{
			TweetID: tweetID,
			Label1: &labels[0],
			Label2: &labels[1],
			Label3: &labels[2],
		}); err != nil {
			log.Println("error labelling tweet:", err)
		}
	}(params)

	return nil
}

func (s *Service) PostRetweet(ctx context.Context, params *model.PostRetweetParams) (error) {
	// Post retweet
	if err := s.repo.PostRetweet(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("original tweet id", "post retweet", err)
	}

	return nil
}

func getTweetLabels(_ context.Context, _ *model.PostTweetParams) []model.Label{
	// Temporary function to get the label of a tweet
	// This function should be implemented in the future
	// For now, it returns a static label
	labels := make([]model.Label, 3)
	labels[0] = model.LabelNews
	labels[1] = model.LabelPolitics
	labels[2] = model.LabelEconomics

	return labels
}