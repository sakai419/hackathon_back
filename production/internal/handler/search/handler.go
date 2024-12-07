package search

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"log"
	"net/http"
)

type SearchHandler struct {
	svc *service.Service
}

func NewSearchHandler(svc *service.Service) ServerInterface {
	return &SearchHandler{
		svc: svc,
	}
}

// Search for users
// (GET /search/users)
func (h *SearchHandler) SearchUsers(w http.ResponseWriter, r *http.Request, params SearchUsersParams) {
	// get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// search users
	userInfos, err := h.svc.SearchUsers(r.Context(), &model.SearchUsersParams{
		ClientAccountID: clientAccountID,
		SortType:        model.SortType(params.SortType),
		Keyword:         params.Keyword,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("Failed to search users", err))
		return
	}

	// convert to response
	resp := make(UserInfos, len(userInfos))
	for i, u := range userInfos {
		resp[i] = UserInfo{
			UserId:          u.UserID,
			UserName:        u.UserName,
			Bio:             u.Bio,
			ProfileImageUrl: u.ProfileImageURL,
			IsPrivate:       u.IsPrivate,
			IsAdmin:         u.IsAdmin,
			IsFollowing:     u.IsFollowing,
			IsFollowed:      u.IsFollowed,
			IsPending:       u.IsPending,
		}
	}

	utils.Respond(w, resp)
}

// Search for tweets
// (GET /search/tweets)
func (h *SearchHandler) SearchTweets(w http.ResponseWriter, r *http.Request, params SearchTweetsParams) {
	// get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// search tweets
	var tweetNodes []*model.TweetNode
	if params.Keyword != nil && params.Label != nil {
		if *params.Keyword == "" && *params.Label == "" {
			utils.RespondError(w, apperrors.NewHandlerError("search tweets", errors.New("keyword or label is required")))
			return
		}
		if (*params.Keyword != "" && *params.Label != "") {
			utils.RespondError(w, apperrors.NewHandlerError("search tweets", errors.New("keyword and label cannot be used together")))
			return
		}
		if *params.Keyword != "" {
			log.Println("searching tweets by keyword")
			temp, err := h.svc.SearchTweets(r.Context(), &model.SearchTweetsParams{
				ClientAccountID: clientAccountID,
				Keyword:         *params.Keyword,
				SortType:        model.SortType(params.SortType),
				Limit:           params.Limit,
				Offset:          params.Offset,
			})
			if err != nil {
				utils.RespondError(w, apperrors.NewHandlerError("search tweets", err))
				return
			}
			tweetNodes = temp
		} else if *params.Label != "" {
			log.Println("searching tweets by label")
			temp, err := h.svc.SearchTweetsByLabels(r.Context(), &model.SearchTweetsByLabelsParams{
				ClientAccountID: clientAccountID,
				Label:           model.Label(*params.Label),
				SortType:        model.SortType(params.SortType),
				Limit:           params.Limit,
				Offset:          params.Offset,
			})
			if err != nil {
				utils.RespondError(w, apperrors.NewHandlerError("search tweets by labels", err))
				return
			}
			log.Println("tweets: ", len(temp))
			tweetNodes = temp
		}
	} else {
		utils.RespondError(w, apperrors.NewHandlerError("seatch tweets", errors.New("keyword or label is required")))
		return
	}

	// convert to response
	resp := convertToTweetNodes(tweetNodes)

	log.Println("tweets: ", len(resp))

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
		TweetId:       t.TweetID,
		UserInfo:      UserInfoWithoutBio{
			UserId:          t.UserInfo.UserID,
			UserName:        t.UserInfo.UserName,
			ProfileImageUrl: t.UserInfo.ProfileImageURL,
			IsPrivate: 	 t.UserInfo.IsPrivate,
			IsAdmin: 	 t.UserInfo.IsAdmin,
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
			Content: t.Code.Content,
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