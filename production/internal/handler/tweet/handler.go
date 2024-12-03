package tweet

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type TweetHandler struct {
	svc *service.Service
}

func NewTweetHandler(svc *service.Service) ServerInterface {
	return &TweetHandler{
		svc: svc,
	}
}

// Post tweet
// (POST /tweets)
func (h *TweetHandler) PostTweet(w http.ResponseWriter, r *http.Request) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request body
	var req PostTweetJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Post tweet
	var code *model.Code
	var media *model.Media
	if req.Code != nil {
		code = &model.Code{
			Language: req.Code.Language,
			Content:  req.Code.Content,
		}
	}
	if req.Media != nil {
		media = &model.Media{
			Type: req.Media.Type,
			URL:  req.Media.Url,
		}
	}
	if err := h.svc.PostTweet(r.Context(), &model.PostTweetParams{
		AccountID: clientAccountID,
		Content:   req.Content,
		Code:      code,
		Media:     media,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("post tweet", err))
		return
	}

    utils.Respond(w, nil)
}

// Like tweet and notify poster
// (POST /tweets/{tweet_id}/like)
func (h *TweetHandler) LikeTweetAndNotify(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Like and notify
	if err := h.svc.LikeTweetAndNotify(r.Context(), &model.LikeTweetAndNotifyParams{
		LikingAccountID: clientAccountID,
		OriginalTweetID: tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("like tweet", err))
		return
	}

	utils.Respond(w, nil)
}

// Retweet and notify poster
// (POST /tweets/{tweet_id}/retweet)
func (h *TweetHandler) RetweetAndNotify(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Retweet and notify
	if err := h.svc.RetweetAndNotify(r.Context(), &model.RetweetAndNotifyParams{
		RetweetingAccountID: clientAccountID,
		OriginalTweetID:   tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("retweet", err))
		return
	}

	utils.Respond(w, nil)
}

// Quote tweet
// (POST /tweets/{tweet_id}/quote)
func (h *TweetHandler) PostQuoteAndNotify(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request body
	var req PostQuoteAndNotifyJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Post quote
	var code *model.Code
	var media *model.Media
	if req.Code != nil {
		code = &model.Code{
			Language: req.Code.Language,
			Content:   req.Code.Content,
		}
	}
	if req.Media != nil {
		media = &model.Media{
			Type: req.Media.Type,
			URL:  req.Media.Url,
		}
	}
	if err := h.svc.PostQuoteAndNotify(r.Context(), &model.PostQuoteAndNotifyParams{
		QuotingAccountID: clientAccountID,
		OriginalTweetID:  tweetID,
		Content:          req.Content,
		Code:             code,
		Media:            media,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("quote tweet", err))
		return
	}

	utils.Respond(w, nil)
}

// Post reply
// (POST /tweets/{tweet_id}/reply)
func (h *TweetHandler) PostReplyAndNotify(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request body
	var req PostReplyAndNotifyJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Post reply
	var code *model.Code
	var media *model.Media
	if req.Code != nil {
		code = &model.Code{
			Language: req.Code.Language,
			Content:  req.Code.Content,
		}
	}
	if req.Media != nil {
		media = &model.Media{
			Type: req.Media.Type,
			URL:  req.Media.Url,
		}
	}
	if err := h.svc.PostReplyAndNotify(r.Context(), &model.PostReplyAndNotifyParams{
		ReplyingAccountID: clientAccountID,
		OriginalTweetID:   tweetID,
		Content:           req.Content,
		Code:              code,
		Media:             media,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("reply tweet", err))
		return
	}

	utils.Respond(w, nil)
}

// Unlike tweet
// (DELETE /tweets/{tweet_id}/like)
func (h *TweetHandler) UnlikeTweet(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Unlike
	if err := h.svc.UnlikeTweet(r.Context(), &model.UnlikeTweetParams{
		LikingAccountID: clientAccountID,
		OriginalTweetID: tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("unlike", err))
		return
	}

	utils.Respond(w, nil)
}

// Unretweet
// (DELETE /tweets/{tweet_id}/retweet)
func (h *TweetHandler) Unretweet(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Unretweet
	if err := h.svc.Unretweet(r.Context(), &model.UnretweetParams{
		RetweetingAccountID: clientAccountID,
		OriginalTweetID:   tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("unretweet", err))
		return
	}

	utils.Respond(w, nil)
}

// Set tweet as pinned
// (POST /tweets/{tweet_id}/pin)
func (h *TweetHandler) SetTweetAsPinned(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Set tweet as pinned
	if err := h.svc.SetTweetAsPinned(r.Context(), &model.SetTweetAsPinnedParams{
		ClientAccountID: clientAccountID,
		TweetID:         tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("set tweet as pinned", err))
		return
	}

	utils.Respond(w, nil)
}

// Unset tweet as pinned
// (DELETE /tweets/{tweet_id}/pin)
func (h *TweetHandler) UnsetTweetAsPinned(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Unset tweet as pinned
	if err := h.svc.UnsetTweetAsPinned(r.Context(), &model.UnsetTweetAsPinnedParams{
		ClientAccountID: clientAccountID,
		TweetID:         tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("unset tweet as pinned", err))
		return
	}

	utils.Respond(w, nil)
}

// Get tweet info
// (GET /tweets/{tweet_id})
func (h *TweetHandler) GetTweetNode(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get tweet info
	tweet, err := h.svc.GetTweetInfo(r.Context(), &model.GetTweetInfoParams{
		ClientAccountID: clientAccountID,
		TweetID:         tweetID,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get tweet info", err))
		return
	}

	// Convert to response
	resp := convertToTweetNodes([]*model.TweetNode{tweet})

	utils.Respond(w, resp[0])
}

// Get liking user infos
// (GET /tweets/{tweet_id}/likes)
func (h *TweetHandler) GetLikingUserInfos(w http.ResponseWriter, r *http.Request, tweetID int64, params GetLikingUserInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get liking user infos
	likingUserInfos, err := h.svc.GetLikingUserInfos(r.Context(), &model.GetLikingUserInfosParams{
		ClientAccountID: clientAccountID,
		OriginalTweetID: tweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get liking user infos", err))
		return
	}

	// Convert to response
	resp := convertToUserInfos(likingUserInfos)

	utils.Respond(w, resp)
}

// Get retweeting user infos
// (GET /tweets/{tweet_id}/retweets)
func (h *TweetHandler) GetRetweetingUserInfos(w http.ResponseWriter, r *http.Request, tweetID int64, params GetRetweetingUserInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get retweeting user infos
	retweetingUserInfos, err := h.svc.GetRetweetingUserInfos(r.Context(), &model.GetRetweetingUserInfosParams{
		ClientAccountID: clientAccountID,
		OriginalTweetID: tweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get retweeting user infos", err))
		return
	}

	// Convert to response
	resp := convertToUserInfos(retweetingUserInfos)

	utils.Respond(w, resp)
}

// Get quoting user infos
// (GET /tweets/{tweet_id}/quotes)
func (h *TweetHandler) GetQuotingUserInfos(w http.ResponseWriter, r *http.Request, tweetID int64, params GetQuotingUserInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get quoting user infos
	quotingUserInfos, err := h.svc.GetQuotingUserInfos(r.Context(), &model.GetQuotingUserInfosParams{
		ClientAccountID: clientAccountID,
		OriginalTweetID: tweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get quoting user infos", err))
		return
	}

	// Convert to response
	resp := convertToUserInfos(quotingUserInfos)

	utils.Respond(w, resp)
}

// Get replies for a tweet
// (GET /tweets/{tweet_id}/replies)
func (h *TweetHandler) GetReplyTweetInfos(w http.ResponseWriter, r *http.Request, tweetID int64, params GetReplyTweetInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get replies
	replyTweets, err := h.svc.GetReplyTweetInfos(r.Context(), &model.GetReplyTweetInfosParams{
		ClientAccountID: clientAccountID,
		ParentTweetID:   tweetID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get replies", err))
		return
	}

	// Convert to response
	var resp []*TweetInfo
	for _, replyTweet := range replyTweets {
		info := convertToTweetInfo(replyTweet)
		resp = append(resp, info)
	}

	utils.Respond(w, resp)
}

// Get timeline tweets
// (GET /tweets/timeline)
func (h *TweetHandler) GetTimelineTweetInfos(w http.ResponseWriter, r *http.Request, params GetTimelineTweetInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get timeline tweets
	timelineTweets, err := h.svc.GetTimelineTweetInfos(r.Context(), &model.GetTimelineTweetInfosParams{
		ClientAccountID: clientAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get timeline tweets", err))
		return
	}

	// Convert to response
	resp := convertToTweetNodes(timelineTweets)

	utils.Respond(w, resp)
}

// Get recent tweets
// (GET /tweets/recent)
func (h *TweetHandler) GetRecentTweetInfos(w http.ResponseWriter, r *http.Request, params GetRecentTweetInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get recent tweets
	recentTweets, err := h.svc.GetRecentTweetInfos(r.Context(), &model.GetRecentTweetInfosParams{
		ClientAccountID: clientAccountID,
		Limit:           params.Limit,
		Offset: 		 params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get recent tweets", err))
		return
	}

	// Convert to response
	resp := convertToTweetNodes(recentTweets)

	utils.Respond(w, resp)
}

// Get recent tweet labels
// (GET /tweets/recent/labels)
func (h *TweetHandler) GetRecentLabels(w http.ResponseWriter, r *http.Request, params GetRecentLabelsParams) {
	// Get recent tweet labels
	labels, err := h.svc.GetRecentLabels(r.Context(), params.Limit)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get recent tweet labels", err))
		return
	}

	// convert to response
	resp := make([]*LabelCount, 0, len(labels))
	for _, label := range labels {
		if model.Label(label.Label).Validate() != nil {
			utils.RespondError(w, apperrors.NewUnexpectedError(errors.New("label is invalid")))
			return
		}
		resp = append(resp, &LabelCount{
			Label: string(label.Label),
			Count: label.Count,
		})
	}

	utils.Respond(w, resp)
}

// Delete tweet
// (DELETE /tweets/{tweet_id})
func (h *TweetHandler) DeleteTweet(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Delete tweet
	if err := h.svc.DeleteTweet(r.Context(), &model.DeleteTweetParams{
		ClientAccountID: clientAccountID,
		TweetID:   tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("delete tweet", err))
		return
	}

	utils.Respond(w, nil)
}

// ErrorHandlerFunc is the error handler for tweet handlers
func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	var requiredParamError *RequiredParamError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, apperrors.NewInvalidParamFormatError(
			invalidParamFormatError.ParamName,
			invalidParamFormatError.Err,
		))
		return
	} else if errors.As(err, &requiredParamError) {
		utils.RespondError(w, apperrors.NewRequiredParamError(
			requiredParamError.ParamName,
			requiredParamError,
		))
		return
	}

	utils.RespondError(w, apperrors.NewUnexpectedError(err))
}

func convertToUserInfos(infos []*model.UserInfo) []*UserInfo {
	var resp []*UserInfo
	for _, info := range infos {
		resp = append(resp, &UserInfo{
			UserId:   info.UserID,
			UserName: info.UserName,
			ProfileImageUrl: info.ProfileImageURL,
			Bio:             info.Bio,
			IsPrivate: info.IsPrivate,
			IsAdmin:   info.IsAdmin,
			IsFollowing: info.IsFollowing,
			IsFollowed:  info.IsFollowed,
			IsPending:   info.IsPending,
		})
	}

	return resp
}

func convertToTweetInfo(t *model.TweetInfo) *TweetInfo {
	if t == nil {
		return nil
	}
	tweet := &TweetInfo{
		TweetId:       t.TweetID,
		UserInfo:      UserInfoWithoutBio{
			UserId:          t.UserInfo.UserID,
			UserName:        t.UserInfo.UserName,
			ProfileImageUrl: t.UserInfo.ProfileImageURL,
			IsPrivate:       t.UserInfo.IsPrivate,
			IsAdmin:         t.UserInfo.IsAdmin,
			IsFollowing:     t.UserInfo.IsFollowing,
			IsFollowed:      t.UserInfo.IsFollowed,
			IsPending:       t.UserInfo.IsPending,
		},
		LikesCount:    t.LikesCount,
		RetweetsCount: t.RetweetsCount,
		RepliesCount:  t.RepliesCount,
		IsQuote:      t.IsQuote,
		IsReply:      t.IsReply,
		IsPinned:     t.IsPinned,
		HasLiked:     t.HasLiked,
		HasRetweeted: t.HasRetweeted,
		CreatedAt:    t.CreatedAt,
	}

	if t.Content != nil {
		tweet.Content = t.Content
	}

	if t.Code != nil {
		tweet.Code = &Code{
			Language: t.Code.Language,
			Content:  t.Code.Content,
		}
	}

	if t.Media != nil {
		tweet.Media = &Media{
			Type: t.Media.Type,
			Url:  t.Media.URL,
		}
	}

	return tweet
}

func convertToTweetNodes(tweets []*model.TweetNode) []TweetNode {
	resp := make([]TweetNode, 0, len(tweets))
	for _, t := range tweets {
		resp = append(resp, TweetNode{
			Tweet:             *convertToTweetInfo(&t.Tweet),
			OriginalTweet:     convertToTweetInfo(t.OriginalTweet),
			ParentReply:       convertToTweetInfo(t.ParentReply),
			OmittedReplyExist: t.OmittedReplyExist,
		})
	}
	return resp
}