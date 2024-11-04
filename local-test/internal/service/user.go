package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
)

func (s *Service) GetUserTweets(ctx context.Context, params *model.GetUserTweetsParams) ([]*model.GetUserTweetsResponse, error) {
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

	// Get user infos
	userInfos, err := s.repo.GetUserInfos(ctx, accountIDs)
	if err != nil {
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToGetUserTweetsResponse(tweets, quotedTweetInfos, replyTweetInfos, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user tweets response", err)
	}

	return responses, nil
}

func (s *Service) GetUserLikes(ctx context.Context, params *model.GetUserLikesParams) ([]*model.TweetInfo, error) {
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

	// Get account ids of all tweets
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
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetInfos(likedTweetIDs, tweets, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user likes response", err)
	}

	return responses, nil
}

func (s *Service) GetUserRetweets(ctx context.Context, params *model.GetUserRetweetsParams) ([]*model.TweetInfo, error) {
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

	// Get account ids of all tweets
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
		return nil, apperrors.NewNotFoundAppError("user infos", "get user infos", err)
	}

	// Convert to response
	responses, err := convertToTweetInfos(retweetedTweetIDs, tweets, userInfos)
	if err != nil {
		return nil, apperrors.NewInternalAppError("convert to get user retweets response", err)
	}

	return responses, nil
}

func convertToGetUserTweetsResponse(tweets []*model.TweetInfoInternal, quotedTweetInfos []*model.QuotedTweetInfoInternal, replyTweetInfos []*model.RepliedTweetInfoInternal, userInfos []*model.UserInfoInternal) ([]*model.GetUserTweetsResponse, error) {
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
	responses := make([]*model.GetUserTweetsResponse, 0, len(tweets))
	for _, tweet := range tweets {
		// Get user info
		userInfo, ok := userInfoMap[tweet.AccountID]
		if !ok {
			return nil, apperrors.NewInternalAppError("user info not found", nil)
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

		response := &model.GetUserTweetsResponse{
			Tweet: tweetInfo,
		}

		// Get quoted tweet info
		quotedTweetInfo, ok := quotedTweetInfoMap[tweet.TweetID]
		if ok {
			userInfo, ok := userInfoMap[quotedTweetInfo.QuotedTweet.AccountID]
			if !ok {
				return nil, apperrors.NewInternalAppError("user info not found", nil)
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
					return nil, apperrors.NewInternalAppError("user info not found", nil)
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
				return nil, apperrors.NewInternalAppError("user info not found", nil)
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