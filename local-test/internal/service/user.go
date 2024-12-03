package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetClientProfile(ctx context.Context, params *model.GetClientProfileParams) (*model.UserProfile, error) {
	// Get user infos
	userInfo, err := s.repo.GetUserInfo(ctx, &model.GetUserInfoParams{
		TargetAccountID: params.ClientAccountID,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Get tweet count
	tweetCount, err := s.repo.GetTweetCountByAccountID(ctx, params.ClientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet count", "get tweet count by account id", err)
	}

	// Get follower and following count
	followCounts, err := s.repo.GetFollowCounts(ctx, params.ClientAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("follow counts", "get follow counts", err)
	}

	// Convert to response
	resp := convertToUserProfile(userInfo, tweetCount, followCounts)

	return resp, nil
}

func (s *Service) GetUserProfile(ctx context.Context, params *model.GetUserProfileParams) (*model.UserProfile, error) {
	// Check if the client is blocked by the target
	if blocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: params.TargetAccountID,
		BlockedAccountID: params.ClientAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if blocked", err)
	} else if blocked {
		return nil, apperrors.NewBlockedAppError("get user profile", errors.New("client is blocked by target"))
	}

	// Check if the client is blocking the target
	if blocking, err := s.repo.IsBlocking(ctx, &model.IsBlockingParams{
		BlockerAccountID: params.ClientAccountID,
		BlockedAccountID: params.TargetAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if blocking", err)
	} else if blocking {
		return nil, apperrors.NewBlockingAppError("get user profile", errors.New("client is blocking target"))
	}

	// Check if the target is private and the client is not following
	if isPrivateAndNotFollowing, err := s.repo.IsPrivateAndNotFollowing(ctx, &model.IsPrivateAndNotFollowingParams{
		ClientAccountID: params.ClientAccountID,
		TargetAccountID: params.TargetAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if private and not following", err)
	} else if isPrivateAndNotFollowing {
		return nil, apperrors.NewForbiddenAppError("get user profile", errors.New("target is private and client is not following"))
	}

	// Get user infos
	userInfo, err := s.repo.GetUserInfo(ctx, &model.GetUserInfoParams{
		TargetAccountID: params.TargetAccountID,
		ClientAccountID: params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Get tweet count
	tweetCount, err := s.repo.GetTweetCountByAccountID(ctx, params.TargetAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("tweet count", "get tweet count by account id", err)
	}

	// Get follower and following count
	followCounts, err := s.repo.GetFollowCounts(ctx, params.TargetAccountID)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("follow counts", "get follow counts", err)
	}

	// Convert to response
	resp := convertToUserProfile(userInfo, tweetCount, followCounts)

	return resp, nil
}

func (s *Service) GetUserTweets(ctx context.Context, params *model.GetUserTweetsParams) ([]*model.TweetNode, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Check if the client is blocked by the target
	if blocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: params.TargetAccountID,
		BlockedAccountID: params.ClientAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if blocked", err)
	} else if blocked {
		return nil, apperrors.NewBlockedAppError("get user tweets", errors.New("client is blocked by target"))
	}

	// Check if the client is blocking the target
	if blocking, err := s.repo.IsBlocking(ctx, &model.IsBlockingParams{
		BlockerAccountID: params.ClientAccountID,
		BlockedAccountID: params.TargetAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if blocking", err)
	} else if blocking {
		return nil, apperrors.NewBlockingAppError("get user tweets", errors.New("client is blocking target"))
	}

	// Check if the target is private and the client is not following
	if isPrivateAndNotFollowing, err := s.repo.IsPrivateAndNotFollowing(ctx, &model.IsPrivateAndNotFollowingParams{
		ClientAccountID: params.ClientAccountID,
		TargetAccountID: params.TargetAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if private and not following", err)
	} else if isPrivateAndNotFollowing {
		return nil, apperrors.NewForbiddenAppError("get user tweets", errors.New("target is private and client is not following"))
	}

	// Get user tweets
	tweets, err := s.repo.GetTweetInfosByAccountID(ctx, &model.GetTweetInfosByAccountIDParams{
		ClientAccountID: params.ClientAccountID,
		TargetAccountID: params.TargetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get user tweets", err)
	}

	// unset tweet as pinned
	for _, tweet := range tweets {
		tweet.IsPinned = false
	}

	if params.Offset == 0 {
		// Get pinned tweet id
		pinnedTweetID, err := s.repo.GetPinnedTweetID(ctx, params.TargetAccountID)
		if err != nil {
			return nil, apperrors.NewNotFoundAppError("pinned tweet id", "get pinned tweet id", err)
		}

		// Get pinned tweet info
		if pinnedTweetID != nil {
			pinnedTweet, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
				ClientAccountID: params.ClientAccountID,
				TweetIDs:        []int64{*pinnedTweetID},
			})
			if err != nil {
				return nil, apperrors.NewNotFoundAppError("pinned tweet info", "get pinned tweet info", err)
			}

			if len(pinnedTweet) > 0 {
				tweets = append([]*model.TweetInfoInternal{pinnedTweet[0]}, tweets...)
			}
		}
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

	// Get quoted tweet infos
	quotedTweetInfos, err := s.repo.GetQuotedTweetInfos(ctx, &model.GetQuotedTweetInfosParams{
		ClientAccountID: params.ClientAccountID,
		QuotingTweetIDs: quotingTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("quoted tweet infos", "get quoted tweet infos", err)
	}

	// Get reply tweet infos
	replyTweetInfos, err := s.repo.GetRepliedTweetInfos(ctx, &model.GetRepliedTweetInfosParams{
		ClientAccountID:   params.ClientAccountID,
		ReplyingTweetIDs:  replyTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("reply tweet infos", "get reply tweet infos", err)
	}

	// Get account ids of all tweets
	accountIDsMap := make(map[string]bool)
	for _, tweet := range tweets {
		accountIDsMap[tweet.AccountID] = true
	}
	for _, quotedTweetInfo := range quotedTweetInfos {
		accountIDsMap[quotedTweetInfo.QuotedTweet.AccountID] = true
	}
	for _, replyTweetInfo := range replyTweetInfos {
		accountIDsMap[replyTweetInfo.OriginalTweet.AccountID] = true
		if replyTweetInfo.ParentReplyTweet != nil {
			accountIDsMap[replyTweetInfo.ParentReplyTweet.AccountID] = true
		}
	}
	accountIDs := make([]string, 0, len(accountIDsMap))
	for accountID := range accountIDsMap {
		accountIDs = append(accountIDs, accountID)
	}

	// Filter accesible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccesibleAccountIDsParams{
		AccountIDs:       accountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("filter accessible account ids", err)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: accessibleAccountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetNodes(tweets, quotedTweetInfos, replyTweetInfos, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user tweets response", err)
	}

	return responses, nil
}

func (s *Service) GetUserLikes(ctx context.Context, params *model.GetUserLikesParams) ([]*model.TweetNode, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Check if the target and client are the same
	if params.ClientAccountID != params.TargetAccountID {
		return nil, apperrors.NewForbiddenAppError("get user likes", errors.New("client and target account ids do not match"))
	}

	// Get liked tweet ids by account id
	likedTweetIDs, err := s.repo.GetLikedTweetIDsByAccountID(ctx, &model.GetLikedTweetIDsByAccountIDParams{
		AccountID: params.ClientAccountID,
		Limit:     params.Limit,
		Offset:    params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get liked tweet ids by account id", err)
	}

	// Get tweet infos by tweet ids
	tweets, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		ClientAccountID: params.ClientAccountID,
		TweetIDs:        likedTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get tweet infos by tweet ids", err)
	}

	// Extract quoting tweet ids
	quotingTweetIDs := make([]int64, 0, len(tweets))
	for _, tweet := range tweets {
		if tweet.IsQuote {
			quotingTweetIDs = append(quotingTweetIDs, tweet.TweetID)
		}
	}

	// Get quoted tweet infos
	quotedTweetInfos, err := s.repo.GetQuotedTweetInfos(ctx, &model.GetQuotedTweetInfosParams{
		ClientAccountID: params.ClientAccountID,
		QuotingTweetIDs: quotingTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("quoted tweet infos", "get quoted tweet infos", err)
	}

	// Get account ids of all tweets
	accountIDsMap := make(map[string]bool)
	for _, tweet := range tweets {
		accountIDsMap[tweet.AccountID] = true
	}
	for _, quotedTweetInfo := range quotedTweetInfos {
		accountIDsMap[quotedTweetInfo.QuotedTweet.AccountID] = true
	}
	accountIDs := make([]string, 0, len(accountIDsMap))
	for accountID := range accountIDsMap {
		accountIDs = append(accountIDs, accountID)
	}

	// Filter accessible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccesibleAccountIDsParams{
		AccountIDs:       accountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("filter accessible account ids", err)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: accessibleAccountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetNodes(tweets, quotedTweetInfos, nil, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user likes response", err)
	}

	return responses, nil
}

func (s *Service) GetUserRetweets(ctx context.Context, params *model.GetUserRetweetsParams) ([]*model.TweetNode, error) {
	// Validate params
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Check if the client is blocked by the target
	if blocked, err := s.repo.IsBlocked(ctx, &model.IsBlockedParams{
		BlockerAccountID: params.TargetAccountID,
		BlockedAccountID: params.ClientAccountID,
	}); err != nil {
		return nil, apperrors.NewInternalAppError("check if blocked", err)
	} else if blocked {
		return nil, apperrors.NewForbiddenAppError("get user tweets", err)
	}

	// Get retweeted tweet ids by account id
	retweetedTweetIDs, err := s.repo.GetRetweetedTweetIDsByAccountID(ctx, &model.GetRetweetedTweetIDsByAccountIDParams{
		RetweetingAccountID: params.TargetAccountID,
		Limit:               params.Limit,
		Offset:              params.Offset,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get retweeted tweet ids by account id", err)
	}

	// Get tweet infos by tweet IDs
	tweets, err := s.repo.GetTweetInfosByIDs(ctx, &model.GetTweetInfosByIDsParams{
		ClientAccountID: params.ClientAccountID,
		TweetIDs:        retweetedTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("get tweet infos by tweet ids", err)
	}

	// Extract quoting tweet ids
	quotingTweetIDs := make([]int64, 0, len(tweets))
	for _, tweet := range tweets {
		if tweet.IsQuote {
			quotingTweetIDs = append(quotingTweetIDs, tweet.TweetID)
		}
	}

	// Get quoted tweet infos
	quotedTweetInfos, err := s.repo.GetQuotedTweetInfos(ctx, &model.GetQuotedTweetInfosParams{
		ClientAccountID: params.ClientAccountID,
		QuotingTweetIDs: quotingTweetIDs,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("quoted tweet infos", "get quoted tweet infos", err)
	}

	// Get account ids of all tweets
	accountIDsMap := make(map[string]bool)
	for _, tweet := range tweets {
		accountIDsMap[tweet.AccountID] = true
	}
	for _, quotedTweetInfo := range quotedTweetInfos {
		accountIDsMap[quotedTweetInfo.QuotedTweet.AccountID] = true
	}
	accountIDs := make([]string, 0, len(accountIDsMap))
	for accountID := range accountIDsMap {
		accountIDs = append(accountIDs, accountID)
	}

	// Filter accesible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccesibleAccountIDsParams{
		AccountIDs:       accountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("filter accessible account ids", err)
	}

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, &model.GetUserInfosParams{
		TargetAccountIDs: accessibleAccountIDs,
		ClientAccountID:  params.ClientAccountID,
	})
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetNodes(tweets, quotedTweetInfos, nil, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user retweets response", err)
	}

	return responses, nil
}

func convertToUserProfile(userInfo *model.UserInfoInternal, tweetCount int64, followCounts *model.FollowCounts) *model.UserProfile {
	return &model.UserProfile{
		UserInfo:    *convertToUserInfo(userInfo),
		BannerImageURL: userInfo.BannerImageURL,
		TweetCount:     tweetCount,
		FollowerCount:  followCounts.FollowersCount,
		FollowingCount: followCounts.FollowingCount,
		CreatedAt: 	    userInfo.CreatedAt,
	}
}