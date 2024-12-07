package service

import (
	"context"
	"errors"
	"local-test/internal/model"
	"local-test/pkg/apperrors"
	"log"
)

func (s *Service) SearchUsers(ctx context.Context, params *model.SearchUsersParams) ([]*model.UserInfo, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Search users
    var users []*model.UserInfoInternal
	switch params.SortType {
		case model.SortTypeLatest, "":
			temp, err := s.repo.SearchUsersOrderByCreatedAt(ctx, &model.SearchUsersOrderByCreatedAtParams{
				ClientAccountID: params.ClientAccountID,
				Keyword: params.Keyword,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search users", err)
			}
			users = temp
		default:
			return nil, apperrors.NewInternalAppError("SortType is not supported", errors.New("SortType is not supported"))
	}

	// extract accountIDs
	accountIDs := make([]string, 0)
	for _, u := range users {
		accountIDs = append(accountIDs, u.ID)
	}

	// Filter accessible account ids
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDsByBlockStatus(ctx, &model.FilterAccessibleAccountIDsByBlockStatusParams{
		ClientAccountID: params.ClientAccountID,
		AccountIDs: accountIDs,
	})
	if err != nil {
		return nil, apperrors.NewInternalAppError("Failed to filter accessible account ids", err)
	}

	// Filter users
	userInfos := filterUsers(users, accessibleAccountIDs)

	return userInfos, nil
}

func (s *Service) SearchTweets(ctx context.Context, params *model.SearchTweetsParams) ([]*model.TweetNode, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Search tweets
	var tweets []*model.TweetInfoInternal
	switch params.SortType {
		case model.SortTypeLatest, "":
			temp, err := s.repo.SearchTweetsOrderByCreatedAt(ctx, &model.SearchTweetsOrderByCreatedAtParams{
				ClientAccountID: params.ClientAccountID,
				Keyword: params.Keyword,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search tweets", err)
			}
			tweets = temp
		case model.SortTypePopular:
			temp, err := s.repo.SearchTweetsOrderByEngagementScore(ctx, &model.SearchTweetsOrderByEngagementScoreParams{
				ClientAccountID: params.ClientAccountID,
				Keyword: params.Keyword,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search tweets", err)
			}
			tweets = temp
		default:
			return nil, apperrors.NewInternalAppError("SortType is not supported", errors.New("SortType is not supported"))
	}

	// unset tweet as pinned
	for _, t := range tweets {
		t.IsPinned = false
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
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccessibleAccountIDsParams{
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

func (s *Service) SearchTweetsByLabels(ctx context.Context, params *model.SearchTweetsByLabelsParams) ([]*model.TweetNode, error) {
	// Validate input
	if err := params.Validate(); err != nil {
		return nil, apperrors.NewValidateAppError(err)
	}

	// Search tweets
	var tweets []*model.TweetInfoInternal
	switch params.SortType {
		case model.SortTypeLatest, "":
			temp, err := s.repo.SearchTweetsByLabelsOrderByCreatedAt(ctx, &model.SearchTweetsByLabelsOrderByCreatedAtParams{
				ClientAccountID: params.ClientAccountID,
				Label: params.Label,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search tweets", err)
			}
			tweets = temp
		case model.SortTypePopular:
			temp, err := s.repo.SearchTweetsByLabelsOrderByEngagementScore(ctx, &model.SearchTweetsByLabelsOrderByEngagementScoreParams{
				ClientAccountID: params.ClientAccountID,
				Label: params.Label,
				Offset: params.Offset,
				Limit: params.Limit,
			})
			if err != nil {
				return nil, apperrors.NewInternalAppError("Failed to search tweets", err)
			}
			tweets = temp
		default:
			return nil, apperrors.NewInternalAppError("SortType is not supported", errors.New("SortType is not supported"))
	}

	log.Println("tweets: ", len(tweets))

	// unset tweet as pinned
	for _, t := range tweets {
		t.IsPinned = false
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
	accessibleAccountIDs, err := s.repo.FilterAccessibleAccountIDs(ctx, &model.FilterAccessibleAccountIDsParams{
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

	log.Println("responses: ", len(responses))

	return responses, nil
}

func filterUsers(users []*model.UserInfoInternal, accessibleAccountIDs []string) []*model.UserInfo {
	// create map of accessible account ids
	accessibleAccountIDsMap := make(map[string]bool)
	for _, id := range accessibleAccountIDs {
		accessibleAccountIDsMap[id] = true
	}

	filteredUsers := make([]*model.UserInfo, 0)
	for _, u := range users {
		if _, ok := accessibleAccountIDsMap[u.ID]; ok {
			filteredUsers = append(filteredUsers, convertToUserInfo(u))
		}
	}
	return filteredUsers
}