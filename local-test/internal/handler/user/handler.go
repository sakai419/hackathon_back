package user

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type UserHandler struct {
	svc *service.Service
}

func NewUserHandler(svc *service.Service) ServerInterface {
	return &UserHandler{
		svc: svc,
	}
}

// Get user's tweets
// (GET /users/{user_id}/tweets)
func (h *UserHandler) GetUserTweets(w http.ResponseWriter, r *http.Request, _ string, params GetUserTweetsParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Check if target account ID is suspended
	if utils.IsTargetSuspended(w, r) {
		utils.RespondError(w, apperrors.NewForbiddenAppError("get user tweets", errors.New("target account is suspended")))
	}

	// Get tweets
	tweets, err := h.svc.GetUserTweets(r.Context(), &model.GetUserTweetsParams{
		ClientAccountID: clientAccountID,
		TargetAccountID: targetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get user tweets", err))
		return
	}

	// Convert to response
	resp := convertToGetUserTweetsResponse(tweets)

	utils.Respond(w, resp)
}

// Get user's likes
// (GET /users/{user_id}/likes)
func (h *UserHandler) GetUserLikes(w http.ResponseWriter, r *http.Request, _ string, params GetUserLikesParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Get likes
	likes, err := h.svc.GetUserLikes(r.Context(), &model.GetUserLikesParams{
		ClientAccountID: clientAccountID,
		TargetAccountID: targetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get user likes", err))
		return
	}

	// Convert to response
	resp := convertToGetUserLikesResponse(likes)

	utils.Respond(w, resp)
}

// Get user's retweets
// (GET /users/{user_id}/retweets)
func (h *UserHandler) GetUserRetweets(w http.ResponseWriter, r *http.Request, _ string, params GetUserRetweetsParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Check if target account ID is suspended
	if utils.IsTargetSuspended(w, r) {
		utils.RespondError(w, apperrors.NewForbiddenAppError("get user retweets", errors.New("target account is suspended")))
		return
	}

	// Get retweets
	retweets, err := h.svc.GetUserRetweets(r.Context(), &model.GetUserRetweetsParams{
		ClientAccountID: clientAccountID,
		TargetAccountID: targetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get user retweets", err))
		return
	}

	// Convert to response
	resp := convertToGetUserRetweetsResponse(retweets)

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

func convertToTweetInfo(t *model.TweetInfo) *TweetInfo {
	if t == nil {
		return nil
	}
	tweet := &TweetInfo{
		TweetID:       t.TweetID,
		UserInfo:      UserInfoWithoutBio{
			UserId:          t.UserInfo.UserID,
			UserName:        t.UserInfo.UserName,
			ProfileImageUrl: t.UserInfo.ProfileImageURL,
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
		tweet.Code = t.Code
	}

	if t.Media != nil {
		tweet.Media = &Media{
			Type: t.Media.Type,
			Url:  t.Media.URL,
		}
	}

	return tweet
}

func convertToGetUserTweetsResponse(tweets []*model.GetUserTweetsResponse) []GetUserTweetsResponse {
	resp := make([]GetUserTweetsResponse, 0, len(tweets))
	for _, t := range tweets {
		resp = append(resp, GetUserTweetsResponse{
			Tweet:             *convertToTweetInfo(&t.Tweet),
			OriginalTweet:     convertToTweetInfo(t.OriginalTweet),
			ParentReply:       convertToTweetInfo(t.ParentReply),
			OmittedReplyExist: t.OmittedReplyExist,
		})
	}
	return resp
}

func convertToGetUserLikesResponse(likes []*model.TweetInfo) []TweetInfo {
	resp := make([]TweetInfo, 0, len(likes))
	for _, l := range likes {
		resp = append(resp, *convertToTweetInfo(l))
	}
	return resp
}

func convertToGetUserRetweetsResponse(retweets []*model.TweetInfo) []TweetInfo {
	resp := make([]TweetInfo, 0, len(retweets))
	for _, r := range retweets {
		resp = append(resp, *convertToTweetInfo(r))
	}
	return resp
}