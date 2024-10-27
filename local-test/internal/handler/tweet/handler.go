package tweet

import (
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

// Post retweet
// (POST /tweets/{tweet_id}/retweet)
func (h *TweetHandler) RetweetTweet(w http.ResponseWriter, r *http.Request, tweetID int64) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Retweet tweet
	if err := h.svc.PostRetweet(r.Context(), &model.PostRetweetParams{
		AccountID: clientAccountID,
		OriginalTweetID:   tweetID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("retweet tweet", err))
		return
	}

	utils.Respond(w, nil)
}