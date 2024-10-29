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
	var media *model.Media
	if req.Media != nil {
		media = &model.Media{
			Type: req.Media.Type,
			URL:  req.Media.Url,
		}
	}
	if err := h.svc.PostTweet(r.Context(), &model.PostTweetParams{
		AccountID: clientAccountID,
		Content:   req.Content,
		Code:      req.Code,
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
	var media *model.Media
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
		Code:             req.Code,
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
	var media *model.Media
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
		Code:              req.Code,
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
	resp := convertToUserInfoWithoutBios(likingUserInfos)

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
	resp := convertToUserInfoWithoutBios(retweetingUserInfos)

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
	resp := convertToUserInfoWithoutBios(quotingUserInfos)

	utils.Respond(w, resp)
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

func convertToUserInfoWithoutBios(infos []*model.UserInfoWithoutBio) []*UserInfoWithoutBio {
	var resp []*UserInfoWithoutBio
	for _, info := range infos {
		resp = append(resp, &UserInfoWithoutBio{
			UserId:   info.UserID,
			UserName: info.UserName,
			ProfileImageUrl: info.ProfileImageURL,
		})
	}

	return resp
}