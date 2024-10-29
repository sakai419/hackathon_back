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

	// Label tweet
	go func(params *model.PostTweetParams) {
		// Get tweet label
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		labels := getTweetLabels(timeoutCtx, &model.GetTweetLabelsParams{
            Content: params.Content,
            Code:    params.Code,
            Media:   params.Media,
        })
		if err := s.repo.LabelTweet(timeoutCtx, &model.LabelTweetParams{
			TweetID: tweetID,
			Label1: &labels[0],
			Label2: &labels[1],
			Label3: &labels[2],
		}); err != nil {
			log.Println(apperrors.NewInternalAppError("label tweet", err))
		}
	}(params)

	return nil
}

func (s *Service) LikeTweetAndNotify(ctx context.Context, params *model.LikeTweetAndNotifyParams) error {
	// Get poster account id
	posterAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if blocked
	isBlocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: posterAccountID,
		BlockedAccountID: params.LikingAccountID,
	}); if err != nil {
		return apperrors.NewInternalAppError("check if blocked", err)
	} else if isBlocked {
		return apperrors.NewForbiddenAppError("Like request", err)
	}

	// Create like
	if err := s.repo.CreateLikeAndNotify(ctx, &model.CreateLikeAndNotifyParams{
		LikingAccountID: params.LikingAccountID,
		OriginalTweetID: params.OriginalTweetID,
		LikedAccountID:  posterAccountID,
	}); err != nil {
		return apperrors.NewDuplicateEntryAppError("like", "create like and notify", err)
	}

	return nil
}

func (s *Service) RetweetAndNotify(ctx context.Context, params *model.RetweetAndNotifyParams) error {
	// Get poster account id
	posterAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if blocked
	isBlocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: posterAccountID,
		BlockedAccountID: params.RetweetingAccountID,
	}); if err != nil {
		return apperrors.NewInternalAppError("check if blocked", err)
	} else if isBlocked {
		return apperrors.NewForbiddenAppError("Retweet request", err)
	}

	// Create retweet
	if err := s.repo.CreateRetweetAndNotify(ctx, &model.CreateRetweetAndNotifyParams{
		RetweetingAccountID: params.RetweetingAccountID,
		OriginalTweetID:     params.OriginalTweetID,
		RetweetedAccountID:  posterAccountID,
	}); err != nil {
		return apperrors.NewDuplicateEntryAppError("retweet", "create retweet and notify", err)
	}

	return nil
}

func (s *Service) PostQuoteAndNotify(ctx context.Context, params *model.PostQuoteAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Get quoted account id
	quotedAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if blocked
	isBlocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: quotedAccountID,
		BlockedAccountID: params.QuotingAccountID,
	}); if err != nil {
		return apperrors.NewInternalAppError("check if blocked", err)
	} else if isBlocked {
		return apperrors.NewForbiddenAppError("Quote request", err)
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

	// Create quote
	tweetID, err := s.repo.CreateQuoteAndNotify(ctx, &model.CreateQuoteAndNotifyParams{
		QuotingAccountID: params.QuotingAccountID,
		QuotedAccountID:  quotedAccountID,
		OriginalTweetID:  params.OriginalTweetID,
		Content:          params.Content,
		Code:             params.Code,
		Media:            params.Media,
		HashtagIDs:       hashtagIDs,
	})
	if err != nil {
		return apperrors.NewInternalAppError("create quote and notify", err)
	}

	// Label tweet
	go func(params *model.PostQuoteAndNotifyParams) {
		// Get tweet label
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		labels := getTweetLabels(timeoutCtx, &model.GetTweetLabelsParams{
			Content: params.Content,
			Code:    params.Code,
			Media:   params.Media,
		})
		if err := s.repo.LabelTweet(timeoutCtx, &model.LabelTweetParams{
			TweetID: tweetID,
			Label1: &labels[0],
			Label2: &labels[1],
			Label3: &labels[2],
		}); err != nil {
			log.Println(apperrors.NewInternalAppError("label tweet", err))
		}
	}(params)

	return nil
}

func (s *Service) PostReplyAndNotify(ctx context.Context, params *model.PostReplyAndNotifyParams) error {
	// Validate params
	if err := params.Validate(); err != nil {
		return apperrors.NewValidateAppError(err)
	}

	// Get replied account id
	repliedAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if blocked
	isBlocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: repliedAccountID,
		BlockedAccountID: params.ReplyingAccountID,
	}); if err != nil {
		return apperrors.NewInternalAppError("check if blocked", err)
	} else if isBlocked {
		return apperrors.NewForbiddenAppError("Reply request", err)
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

	// Create reply
	tweetID, err := s.repo.CreateReplyAndNotify(ctx, &model.CreateReplyAndNotifyParams{
		ReplyingAccountID: params.ReplyingAccountID,
        RepliedAccountID:  repliedAccountID,
		OriginalTweetID:   params.OriginalTweetID,
		Content:           params.Content,
		Code:              params.Code,
		Media:             params.Media,
		HashtagIDs:        hashtagIDs,
	})
	if err != nil {
		return apperrors.NewInternalAppError("create reply and notify", err)
	}

	// Label tweet
	go func(params *model.PostReplyAndNotifyParams) {
		// Get tweet label
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		labels := getTweetLabels(timeoutCtx, &model.GetTweetLabelsParams{
            Content: params.Content,
            Code:    params.Code,
            Media:   params.Media,
        })
        if err := s.repo.LabelTweet(timeoutCtx, &model.LabelTweetParams{
            TweetID: tweetID,
            Label1: &labels[0],
            Label2: &labels[1],
            Label3: &labels[2],
        }); err != nil {
            log.Println(apperrors.NewInternalAppError("label tweet", err))
        }
    }(params)

	return nil
}

func (s *Service) UnlikeTweet(ctx context.Context, params *model.UnlikeTweetParams) error {
	// Unlike tweet
	if err := s.repo.UnlikeTweet(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("like", "unlike", err)
	}

	return nil
}

func (s *Service) Unretweet(ctx context.Context, params *model.UnretweetParams) error {
	// Unretweet
	if err := s.repo.Unretweet(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("retweet", "unretweet", err)
	}

	return nil
}

func getTweetLabels(_ context.Context, _ *model.GetTweetLabelsParams) []model.Label{
	// Temporary function to get the label of a tweet
	// This function should be implemented in the future
	// For now, it returns a static label
	labels := make([]model.Label, 3)
	labels[0] = model.LabelNews
	labels[1] = model.LabelPolitics
	labels[2] = model.LabelEconomics

	return labels
}