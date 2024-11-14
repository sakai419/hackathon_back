package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"log"
	"sort"
	"time"

	"golang.org/x/exp/rand"
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

func (s *Service) SetTweetAsPinned(ctx context.Context, params *model.SetTweetAsPinnedParams) error {
	// Get account id of tweet
	accountID, err := s.repo.GetAccountIDByTweetID(ctx, params.TweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if client is authorized
	if accountID != params.ClientAccountID {
		return apperrors.NewForbiddenAppError("Set tweet as pinned", nil)
	}

	// Get pinned tweet id
	pinnedTweetID, err := s.repo.GetPinnedTweetID(ctx, params.ClientAccountID)
	if err != nil {
		return apperrors.NewInternalAppError("get pinned tweet id", err)
	} else if pinnedTweetID != nil {
		return apperrors.NewDuplicateEntryAppError("pinned tweet", "set tweet as pinned", &apperrors.ErrDuplicateEntry{
			Entity: "pinned tweet",
			Err:    errors.New("pinned tweet already exists"),
		})
	}

	// Set tweet as pinned
	if err := s.repo.SetTweetAsPinned(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("tweet", "set tweet as pinned", err)
	}

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

func (s *Service) UnsetTweetAsPinned(ctx context.Context, params *model.UnsetTweetAsPinnedParams) error {
	// Get account id of tweet
	accountID, err := s.repo.GetAccountIDByTweetID(ctx, params.TweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if client is authorized
	if accountID != params.ClientAccountID {
		return apperrors.NewForbiddenAppError("Unset tweet as pinned", nil)
	}

	// Get pinned tweet id
	pinnedTweetID, err := s.repo.GetPinnedTweetID(ctx, params.ClientAccountID)
	if err != nil {
		return apperrors.NewInternalAppError("get pinned tweet id", err)
	} else if pinnedTweetID == nil {
		return apperrors.NewNotFoundAppError("pinned tweet", "unset tweet as pinned", &apperrors.ErrRecordNotFound{
			Condition: "pinned tweet id",
		})
	} else if *pinnedTweetID != params.TweetID {
		return apperrors.NewForbiddenAppError("Unset tweet as pinned", nil)
	}

	// Unset tweet as pinned
	if err := s.repo.UnsetTweetAsPinned(ctx, params); err != nil {
		return apperrors.NewNotFoundAppError("tweet", "unset tweet as pinned", err)
	}

	return nil
}

func (s *Service) GetLikingUserInfos(ctx context.Context, params *model.GetLikingUserInfosParams) ([]*model.UserInfoWithoutBio, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get poster account id
	posterAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if poster and client are same
	if posterAccountID != params.ClientAccountID {
		return nil, apperrors.NewForbiddenAppError("Get liking user info", nil)
	}

	// Get liking account ids
	likingAccountIDs, err := s.repo.GetLikingAccountIDs(ctx, &model.GetLikingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get liking account ids", err)
	}

	// Get user infos
	infos, err := s.repo.GetUserInfos(ctx, likingAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("liking user info", "get liking user infos", err)
	}

	likingUserInfos := sortUserInfoWithoutBios(infos, likingAccountIDs)

	return likingUserInfos, nil
}

func (s *Service) GetRetweetingUserInfos(ctx context.Context, params *model.GetRetweetingUserInfosParams) ([]*model.UserInfoWithoutBio, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get poster account id
	posterAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if poster and client are same
	if posterAccountID != params.ClientAccountID {
		return nil, apperrors.NewForbiddenAppError("Get retweeting user info", nil)
	}

	// Get retweeting account ids
	retweetingAccountIDs, err := s.repo.GetRetweetingAccountIDs(ctx, &model.GetRetweetingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get retweeting account ids", err)
	}

	// Get user infos
	infos, err := s.repo.GetUserInfos(ctx, retweetingAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("retweeting user info", "get retweeting user infos", err)
	}

	retweetingUserInfos := sortUserInfoWithoutBios(infos, retweetingAccountIDs)

	return retweetingUserInfos, nil
}

func (s *Service) GetQuotingUserInfos(ctx context.Context, params *model.GetQuotingUserInfosParams) ([]*model.UserInfoWithoutBio, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get poster account id
	posterAccountID, err := s.repo.GetAccountIDByTweetID(ctx, params.OriginalTweetID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if quoted and client are same
	if posterAccountID != params.ClientAccountID {
		return nil, apperrors.NewForbiddenAppError("Get quoting user info", nil)
	}

	// Get quoting account ids
	quotingAccountIDs, err := s.repo.GetQuotingAccountIDs(ctx, &model.GetQuotingAccountIDsParams{
		OriginalTweetID: params.OriginalTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get quoting account ids", err)
	}

	// Get user infos
	infos, err := s.repo.GetUserInfos(ctx, quotingAccountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("quoting user info", "get quoting user infos", err)
	}

	quotingUserInfos := sortUserInfoWithoutBios(infos, quotingAccountIDs)

	return quotingUserInfos, nil
}

func (s *Service) GetReplyTweetInfos(ctx context.Context, params *model.GetReplyTweetInfosParams) ([]*model.TweetInfo, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Get reply tweet ids
	replyTweetIDs, err := s.repo.GetReplyIDs(ctx, &model.GetReplyIDsParams{
		OriginalTweetID: params.ParentTweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get reply ids", err)
	}

	// Get tweet infos by tweet IDs
	tweets, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		ClientAccountID: params.ClientAccountID,
		TweetIDs:        replyTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get tweet infos by ids", err)
	}

	// Get account IDs of all tweets
	accountIDsMap := make(map[string]bool)
	for _, tweet := range tweets {
		accountIDsMap[tweet.AccountID] = true
	}
	accountIDs := make([]string, 0, len(accountIDsMap))
	for accountID := range accountIDsMap {
		accountIDs = append(accountIDs, accountID)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, accountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user info", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetInfos(replyTweetIDs, tweets, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to tweet info", err)
	}

	return responses, nil
}

func (s *Service) GetTimelineTweetInfos(ctx context.Context, params *model.GetTimelineTweetInfosParams) ([]*model.GetTimelineTweetInfosResponse, error) {
	// Get recent tweet metadatas
	tweetMetadatas, err := s.repo.GetRecentTweetMetadatas(ctx, &model.GetRecentTweetMetadatasParams{
		Limit:  params.Limit,
		Offset: params.Offset,
        ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get recent tweet metadatas", err)
	}

	// Get client interest scores
	interestScores, err := s.repo.GetInterestScores(ctx, params.ClientAccountID)
    if err != nil {
        return nil, apperrors.NewNotFoundAppError("interest scores", "get interest scores", err)
    }

	// Calculate tweet scores with labels
	tweetScoresWithLabel := make(map[int64]int64)
	for _, metadata := range tweetMetadatas {
		tweetScoresWithLabel[metadata.TweetID] = max(calculateTweetScore(metadata, interestScores), 1)
	}

	// Calcluate tweet scores with engagement
	tweetScoresWithEngagement := make(map[int64]int64)
	for _, metadata := range tweetMetadatas {
		tweetScoresWithEngagement[metadata.TweetID] = max(int64(30 * metadata.LikesCount + 20 * metadata.RetweetsCount + metadata.RepliesCount), 1)
	}

	// Get top scoring tweets with labels
	topScoringTweetIDsWithLabel := weightedRandomSample(tweetScoresWithLabel, min(10, len(tweetScoresWithLabel)))

	// Get top scoring tweets with engagement
	topScoringTweetIDsWithEngagement := weightedRandomSample(tweetScoresWithEngagement, min(10, len(tweetScoresWithEngagement)))

	// Remove duplicates
	tweetIDsMap := make(map[int64]struct{})
	for _, tweetID := range topScoringTweetIDsWithLabel {
		tweetIDsMap[tweetID] = struct{}{}
	}
	for _, tweetID := range topScoringTweetIDsWithEngagement {
		tweetIDsMap[tweetID] = struct{}{}
	}
	tweetIDs := make([]int64, 0, len(tweetIDsMap))
	for tweetID := range tweetIDsMap {
		tweetIDs = append(tweetIDs, tweetID)
	}

	// Get tweet infos internal
	tweets, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		ClientAccountID: params.ClientAccountID,
		TweetIDs:        tweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get tweet infos by ids", err)
	}

	// Extract quoting and reply tweet ids
	quotingTweetIDs := make([]int64, 0, len(tweets))
	replyTweetIDs := make([]int64, 0, len(tweets))
	for _, tweet := range tweets {
		if tweet.IsQuote {
			quotingTweetIDs = append(quotingTweetIDs, tweet.TweetID)
		}
		if tweet.IsReply {
			replyTweetIDs = append(replyTweetIDs, tweet.TweetID)
		}
	}

	// Get quoting tweet infos
	quotingTweetInfos, err := s.repo.GetQuotedTweetInfos(ctx, &model.GetQuotedTweetInfosParams{
		ClientAccountID: params.ClientAccountID,
		QuotingTweetIDs:  quotingTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("quoted tweet infos", "get quoted tweet infos", err)
	}

	// Get reply tweet infos
	replyTweetInfos, err := s.repo.GetRepliedTweetInfos(ctx, &model.GetRepliedTweetInfosParams{
		ClientAccountID: params.ClientAccountID,
		ReplyingTweetIDs:   replyTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("replied tweet infos", "get replied tweet infos", err)
	}

	// Get account ids of all tweets
	accountIDs := make([]string, 0, len(tweets))
	for _, tweet := range tweets {
		accountIDs = append(accountIDs, tweet.AccountID)
	}
	for _, quotingTweetInfo := range quotingTweetInfos {
		accountIDs = append(accountIDs, quotingTweetInfo.QuotedTweet.AccountID)
	}
	for _, replyTweetInfo := range replyTweetInfos {
		accountIDs = append(accountIDs, replyTweetInfo.OriginalTweet.AccountID)
		accountIDs = append(accountIDs, replyTweetInfo.ParentReplyTweet.AccountID)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, accountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user info", "get user infos", err)
	}

	// Convert to model
	responses, err := convertToGetTimelineTweetInfosResponse(tweets, quotingTweetInfos, replyTweetInfos, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get timeline tweet infos", err)
	}

	return responses, nil
}

func (s *Service) DeleteTweet(ctx context.Context, params *model.DeleteTweetParams) error {
	// Get account id of tweet
	accountID, err := s.repo.GetAccountIDByTweetID(ctx, params.TweetID)
	if err != nil {
		return apperrors.NewNotFoundAppError("tweet id", "get account id by tweet id", err)
	}

	// Check if client is authorized
	if accountID != params.ClientAccountID {
		return apperrors.NewForbiddenAppError("Delete tweet", nil)
	}

	// Delete tweet
	if err := s.repo.DeleteTweet(ctx, params.TweetID); err != nil {
		return apperrors.NewNotFoundAppError("tweet", "delete tweet", err)
	}

	return nil
}

func convertToTweetInfos(tweetIDs []int64, tweets []*model.TweetInfoInternal, userInfos []*model.UserInfoInternal) ([]*model.TweetInfo, error) {
	// Create map of user info
	userInfoMap := make(map[string]*model.UserInfoInternal)
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Create map of tweet info
	tweetInfoMap := make(map[int64]*model.TweetInfoInternal)
	for _, tweet := range tweets {
		tweetInfoMap[tweet.TweetID] = tweet
	}

	// Convert to model
	var ret []*model.TweetInfo
	for _, id := range tweetIDs {
		tweet, ok := tweetInfoMap[id]
		if !ok {
			return nil, apperrors.NewNotFoundAppError("tweet info", "convert to tweet info", nil)
		}

		userInfo, ok := userInfoMap[tweet.AccountID]
		if !ok {
			return nil, apperrors.NewNotFoundAppError("user info", "convert to tweet info", nil)
		}

		info := &model.TweetInfo{
			TweetID:       tweet.TweetID,
			Content:       tweet.Content,
			Code:          tweet.Code,
			LikesCount:    tweet.LikesCount,
			RepliesCount:  tweet.RepliesCount,
			RetweetsCount: tweet.RetweetsCount,
			IsReply:       tweet.IsReply,
			IsQuote:       tweet.IsQuote,
			IsPinned:      tweet.IsPinned,
			HasLiked:      tweet.HasLiked,
			HasRetweeted:  tweet.HasRetweeted,
			Media:         tweet.Media,
			UserInfo:      model.UserInfoWithoutBio{
				UserID:          userInfo.UserID,
				UserName:        userInfo.UserName,
				ProfileImageURL: userInfo.ProfileImageURL,
			},
			CreatedAt:    tweet.CreatedAt,
		}

		ret = append(ret, info)
	}

	return ret, nil
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

func calculateTweetScore(metadata *model.TweetMetadata, interestScores *model.InterestScores) int64 {
	score := int64(0)
	if metadata.Label1 != nil {
		score += int64(10 * (interestScores.GetScore(metadata.Label1) + 1))
	}

	if metadata.Label2 != nil {
		score += int64(5 * (interestScores.GetScore(metadata.Label2) + 1))
	}

	if metadata.Label3 != nil {
		score += int64(3 * (interestScores.GetScore(metadata.Label3) + 1))
	}

	return score
}

func weightedRandomSample(scores map[int64]int64, m int) []int64 {
	// Calculate total score
	var totalScore float64
	for _, score := range scores {
		totalScore += float64(score)
	}

	// Calculate cumulative probabilities
	cumulativeProbs := make([]float64, len(scores))
	tweetIDs := make([]int64, 0, len(scores))
	cumulativeSum := 0.0
	i := 0
	for tweetID, score := range scores {
		cumulativeSum += float64(score) / totalScore
		cumulativeProbs[i] = cumulativeSum
		tweetIDs = append(tweetIDs, tweetID)
		i++
	}

	// Select m tweet_ids
	rand.Seed(uint64(time.Now().UnixNano()))
	selectedTweetIDs := make(map[int64]struct{})
	for len(selectedTweetIDs) < m {		r := rand.Float64()
		idx := sort.SearchFloat64s(cumulativeProbs, r)
		selectedTweetIDs[tweetIDs[idx]] = struct{}{}
	}

	// Convert map to slice
	result := make([]int64, 0, len(selectedTweetIDs))
	for tweetID := range selectedTweetIDs {
		result = append(result, tweetID)
	}

	return result
}

func convertToGetTimelineTweetInfosResponse(tweets []*model.TweetInfoInternal, quotedTweetInfos []*model.QuotedTweetInfoInternal, replyTweetInfos []*model.RepliedTweetInfoInternal, userInfos []*model.UserInfoInternal) ([]*model.GetTimelineTweetInfosResponse, error) {
	// Create map of user infos
	userInfoMap := make(map[string]*model.UserInfoInternal)
	for _, userInfo := range userInfos {
		userInfoMap[userInfo.ID] = userInfo
	}

	// Create map of quoted tweet infos
	quotedTweetInfoMap := make(map[int64]*model.QuotedTweetInfoInternal)
	for _, quotedTweetInfo := range quotedTweetInfos {
		quotedTweetInfoMap[quotedTweetInfo.QuotingTweetID] = quotedTweetInfo
	}

	// Create map of reply tweet infos
	replyTweetInfoMap := make(map[int64]*model.RepliedTweetInfoInternal)
	for _, replyTweetInfo := range replyTweetInfos {
		replyTweetInfoMap[replyTweetInfo.ReplyingTweetID] = replyTweetInfo
	}

	// Create response
	responses := make([]*model.GetTimelineTweetInfosResponse, 0, len(tweets))
	for _, tweet := range tweets {
		// Get user info
		userInfo, ok := userInfoMap[tweet.AccountID]
		if !ok {
			return nil, errors.New("user info not found")
		}

		tweetInfo := model.TweetInfo{
			TweetID:       tweet.TweetID,
			Content:       tweet.Content,
			Code:          tweet.Code,
			Media:         tweet.Media,
			LikesCount:    tweet.LikesCount,
			RetweetsCount: tweet.RetweetsCount,
			RepliesCount:  tweet.RepliesCount,
			IsQuote:       tweet.IsQuote,
			IsReply:       tweet.IsReply,
			IsPinned:      tweet.IsPinned,
			HasLiked:      tweet.HasLiked,
			HasRetweeted:  tweet.HasRetweeted,
			CreatedAt:     tweet.CreatedAt,
			UserInfo:      model.UserInfoWithoutBio{
				UserID:          userInfo.UserID,
				UserName:        userInfo.UserName,
				ProfileImageURL: userInfo.ProfileImageURL,
			},
		}

		response := &model.GetTimelineTweetInfosResponse{
			Tweet: tweetInfo,
		}

		// Get quoted tweet info
		quotedTweetInfo, ok := quotedTweetInfoMap[tweet.TweetID]
		if ok {
			userInfo, ok := userInfoMap[quotedTweetInfo.QuotedTweet.AccountID]
			if !ok {
				return nil, errors.New("user info not found")
			}

			quotedTweet := &model.TweetInfo{
				TweetID:       quotedTweetInfo.QuotedTweet.TweetID,
				Content:       quotedTweetInfo.QuotedTweet.Content,
				Code:          quotedTweetInfo.QuotedTweet.Code,
				Media:         quotedTweetInfo.QuotedTweet.Media,
				LikesCount:    quotedTweetInfo.QuotedTweet.LikesCount,
				RetweetsCount: quotedTweetInfo.QuotedTweet.RetweetsCount,
				RepliesCount:  quotedTweetInfo.QuotedTweet.RepliesCount,
				IsQuote:       quotedTweetInfo.QuotedTweet.IsQuote,
				IsReply:       quotedTweetInfo.QuotedTweet.IsReply,
				IsPinned:      quotedTweetInfo.QuotedTweet.IsPinned,
				HasLiked:      quotedTweetInfo.QuotedTweet.HasLiked,
				HasRetweeted:  quotedTweetInfo.QuotedTweet.HasRetweeted,
				CreatedAt:     quotedTweetInfo.QuotedTweet.CreatedAt,
				UserInfo:      model.UserInfoWithoutBio{
					UserID:          userInfo.UserID,
					UserName:        userInfo.UserName,
					ProfileImageURL: userInfo.ProfileImageURL,
				},
			}
			response.OriginalTweet = quotedTweet
		}

		// Get reply tweet info
		replyTweetInfo, ok := replyTweetInfoMap[tweet.TweetID]
		if ok {
			if replyTweetInfo.ParentReplyTweet != nil {
				userInfo, ok := userInfoMap[replyTweetInfo.ParentReplyTweet.AccountID]
				if !ok {
					return nil, errors.New("user info not found")
				}

				parentReplyTweet := &model.TweetInfo{
					TweetID:       replyTweetInfo.ParentReplyTweet.TweetID,
					Content:       replyTweetInfo.ParentReplyTweet.Content,
					Code:          replyTweetInfo.ParentReplyTweet.Code,
					Media:         replyTweetInfo.ParentReplyTweet.Media,
					LikesCount:    replyTweetInfo.ParentReplyTweet.LikesCount,
					RetweetsCount: replyTweetInfo.ParentReplyTweet.RetweetsCount,
					RepliesCount:  replyTweetInfo.ParentReplyTweet.RepliesCount,
					IsQuote:       replyTweetInfo.ParentReplyTweet.IsQuote,
					IsReply:       replyTweetInfo.ParentReplyTweet.IsReply,
					IsPinned:      replyTweetInfo.ParentReplyTweet.IsPinned,
					HasLiked:      replyTweetInfo.ParentReplyTweet.HasLiked,
					HasRetweeted:  replyTweetInfo.ParentReplyTweet.HasRetweeted,
					CreatedAt:     replyTweetInfo.ParentReplyTweet.CreatedAt,
					UserInfo:      model.UserInfoWithoutBio{
						UserID:          userInfo.UserID,
						UserName:        userInfo.UserName,
						ProfileImageURL: userInfo.ProfileImageURL,
					},
				}

				response.ParentReply = parentReplyTweet
			}

			userInfo, ok := userInfoMap[replyTweetInfo.OriginalTweet.AccountID]
			if !ok {
				return nil, errors.New("user info not found")
			}

			originalTweet := &model.TweetInfo{
				TweetID:       replyTweetInfo.OriginalTweet.TweetID,
				Content:       replyTweetInfo.OriginalTweet.Content,
				Code:          replyTweetInfo.OriginalTweet.Code,
				Media:         replyTweetInfo.OriginalTweet.Media,
				LikesCount:    replyTweetInfo.OriginalTweet.LikesCount,
				RetweetsCount: replyTweetInfo.OriginalTweet.RetweetsCount,
				RepliesCount:  replyTweetInfo.OriginalTweet.RepliesCount,
				IsQuote:       replyTweetInfo.OriginalTweet.IsQuote,
				IsReply:       replyTweetInfo.OriginalTweet.IsReply,
				IsPinned:      replyTweetInfo.OriginalTweet.IsPinned,
				HasLiked:      replyTweetInfo.OriginalTweet.HasLiked,
				HasRetweeted:  replyTweetInfo.OriginalTweet.HasRetweeted,
				CreatedAt:     replyTweetInfo.OriginalTweet.CreatedAt,
				UserInfo:      model.UserInfoWithoutBio{
					UserID:          userInfo.UserID,
					UserName:        userInfo.UserName,
					ProfileImageURL: userInfo.ProfileImageURL,
				},
			}

			response.OriginalTweet = originalTweet
		}

		responses = append(responses, response)
	}

	return responses, nil
}