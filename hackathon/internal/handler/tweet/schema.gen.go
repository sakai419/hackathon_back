// Package tweet provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package tweet

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Code defines model for Code.
type Code struct {
	Content  string `json:"content"`
	Language string `json:"language"`
}

// LabelCount defines model for LabelCount.
type LabelCount struct {
	Count int64  `json:"count"`
	Label string `json:"label"`
}

// LabelCounts defines model for LabelCounts.
type LabelCounts = []LabelCount

// Media defines model for Media.
type Media struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// PostTweetRequest defines model for PostTweetRequest.
type PostTweetRequest struct {
	Code    *Code   `json:"code,omitempty"`
	Content *string `json:"content,omitempty"`
	Media   *Media  `json:"media,omitempty"`
}

// TweetInfo defines model for TweetInfo.
type TweetInfo struct {
	Code          *Code              `json:"code,omitempty"`
	Content       *string            `json:"content"`
	CreatedAt     time.Time          `json:"created_at"`
	HasLiked      bool               `json:"has_liked"`
	HasRetweeted  bool               `json:"has_retweeted"`
	IsPinned      bool               `json:"is_pinned"`
	IsQuote       bool               `json:"is_quote"`
	IsReply       bool               `json:"is_reply"`
	LikesCount    int32              `json:"likes_count"`
	Media         *Media             `json:"media,omitempty"`
	RepliesCount  int32              `json:"replies_count"`
	RetweetsCount int32              `json:"retweets_count"`
	TweetId       int64              `json:"tweet_id"`
	UserInfo      UserInfoWithoutBio `json:"user_info"`
}

// TweetNode defines model for TweetNode.
type TweetNode struct {
	OmittedReplyExist *bool      `json:"omitted_reply_exist"`
	OriginalTweet     *TweetInfo `json:"original_tweet,omitempty"`
	ParentReply       *TweetInfo `json:"parent_reply,omitempty"`
	Tweet             TweetInfo  `json:"tweet"`
}

// TweetNodes defines model for TweetNodes.
type TweetNodes = []TweetNode

// UserInfoWithoutBio defines model for UserInfoWithoutBio.
type UserInfoWithoutBio struct {
	IsAdmin         bool   `json:"is_admin"`
	IsPrivate       bool   `json:"is_private"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserId          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

// UserInfoWithoutBios defines model for UserInfoWithoutBios.
type UserInfoWithoutBios = []UserInfoWithoutBio

// GetRecentTweetInfosParams defines parameters for GetRecentTweetInfos.
type GetRecentTweetInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetRecentLabelsParams defines parameters for GetRecentLabels.
type GetRecentLabelsParams struct {
	Limit int32 `form:"limit" json:"limit"`
}

// GetTimelineTweetInfosParams defines parameters for GetTimelineTweetInfos.
type GetTimelineTweetInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetLikingUserInfosParams defines parameters for GetLikingUserInfos.
type GetLikingUserInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetQuotingUserInfosParams defines parameters for GetQuotingUserInfos.
type GetQuotingUserInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetReplyTweetInfosParams defines parameters for GetReplyTweetInfos.
type GetReplyTweetInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetRetweetingUserInfosParams defines parameters for GetRetweetingUserInfos.
type GetRetweetingUserInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// PostTweetJSONRequestBody defines body for PostTweet for application/json ContentType.
type PostTweetJSONRequestBody = PostTweetRequest

// PostQuoteAndNotifyJSONRequestBody defines body for PostQuoteAndNotify for application/json ContentType.
type PostQuoteAndNotifyJSONRequestBody = PostTweetRequest

// PostReplyAndNotifyJSONRequestBody defines body for PostReplyAndNotify for application/json ContentType.
type PostReplyAndNotifyJSONRequestBody = PostTweetRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a tweet
	// (POST /tweets)
	PostTweet(w http.ResponseWriter, r *http.Request)
	// Get recent tweets
	// (GET /tweets/recent)
	GetRecentTweetInfos(w http.ResponseWriter, r *http.Request, params GetRecentTweetInfosParams)
	// Get recent tweet labels
	// (GET /tweets/recent/labels)
	GetRecentLabels(w http.ResponseWriter, r *http.Request, params GetRecentLabelsParams)
	// Get timeline for user
	// (GET /tweets/timeline)
	GetTimelineTweetInfos(w http.ResponseWriter, r *http.Request, params GetTimelineTweetInfosParams)
	// Delete a tweet
	// (DELETE /tweets/{tweet_id})
	DeleteTweet(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Get tweet by ID
	// (GET /tweets/{tweet_id})
	GetTweetNode(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Unlike a tweet
	// (DELETE /tweets/{tweet_id}/like)
	UnlikeTweet(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Like a tweet and notify poster
	// (POST /tweets/{tweet_id}/like)
	LikeTweetAndNotify(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Get likes for a tweet
	// (GET /tweets/{tweet_id}/likes)
	GetLikingUserInfos(w http.ResponseWriter, r *http.Request, tweetId int64, params GetLikingUserInfosParams)
	// Unpin a tweet
	// (DELETE /tweets/{tweet_id}/pin)
	UnsetTweetAsPinned(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Pin a tweet
	// (PATCH /tweets/{tweet_id}/pin)
	SetTweetAsPinned(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Quote retweet and notify poster
	// (POST /tweets/{tweet_id}/quote)
	PostQuoteAndNotify(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Get quotes for a tweet
	// (GET /tweets/{tweet_id}/quotes)
	GetQuotingUserInfos(w http.ResponseWriter, r *http.Request, tweetId int64, params GetQuotingUserInfosParams)
	// Get replies for a tweet
	// (GET /tweets/{tweet_id}/replies)
	GetReplyTweetInfos(w http.ResponseWriter, r *http.Request, tweetId int64, params GetReplyTweetInfosParams)
	// Reply to a tweet
	// (POST /tweets/{tweet_id}/reply)
	PostReplyAndNotify(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Delete retweet
	// (DELETE /tweets/{tweet_id}/retweet)
	Unretweet(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Retweet and notify poster
	// (POST /tweets/{tweet_id}/retweet)
	RetweetAndNotify(w http.ResponseWriter, r *http.Request, tweetId int64)
	// Get retweets for a tweet
	// (GET /tweets/{tweet_id}/retweets)
	GetRetweetingUserInfos(w http.ResponseWriter, r *http.Request, tweetId int64, params GetRetweetingUserInfosParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostTweet operation middleware
func (siw *ServerInterfaceWrapper) PostTweet(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTweet(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRecentTweetInfos operation middleware
func (siw *ServerInterfaceWrapper) GetRecentTweetInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRecentTweetInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRecentTweetInfos(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRecentLabels operation middleware
func (siw *ServerInterfaceWrapper) GetRecentLabels(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRecentLabelsParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRecentLabels(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetTimelineTweetInfos operation middleware
func (siw *ServerInterfaceWrapper) GetTimelineTweetInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTimelineTweetInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTimelineTweetInfos(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteTweet operation middleware
func (siw *ServerInterfaceWrapper) DeleteTweet(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTweet(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetTweetNode operation middleware
func (siw *ServerInterfaceWrapper) GetTweetNode(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTweetNode(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UnlikeTweet operation middleware
func (siw *ServerInterfaceWrapper) UnlikeTweet(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UnlikeTweet(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// LikeTweetAndNotify operation middleware
func (siw *ServerInterfaceWrapper) LikeTweetAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.LikeTweetAndNotify(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetLikingUserInfos operation middleware
func (siw *ServerInterfaceWrapper) GetLikingUserInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLikingUserInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLikingUserInfos(w, r, tweetId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// UnsetTweetAsPinned operation middleware
func (siw *ServerInterfaceWrapper) UnsetTweetAsPinned(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.UnsetTweetAsPinned(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// SetTweetAsPinned operation middleware
func (siw *ServerInterfaceWrapper) SetTweetAsPinned(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SetTweetAsPinned(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostQuoteAndNotify operation middleware
func (siw *ServerInterfaceWrapper) PostQuoteAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostQuoteAndNotify(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetQuotingUserInfos operation middleware
func (siw *ServerInterfaceWrapper) GetQuotingUserInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetQuotingUserInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetQuotingUserInfos(w, r, tweetId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetReplyTweetInfos operation middleware
func (siw *ServerInterfaceWrapper) GetReplyTweetInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetReplyTweetInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetReplyTweetInfos(w, r, tweetId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostReplyAndNotify operation middleware
func (siw *ServerInterfaceWrapper) PostReplyAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostReplyAndNotify(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Unretweet operation middleware
func (siw *ServerInterfaceWrapper) Unretweet(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Unretweet(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RetweetAndNotify operation middleware
func (siw *ServerInterfaceWrapper) RetweetAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RetweetAndNotify(w, r, tweetId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetRetweetingUserInfos operation middleware
func (siw *ServerInterfaceWrapper) GetRetweetingUserInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "tweet_id" -------------
	var tweetId int64

	err = runtime.BindStyledParameterWithOptions("simple", "tweet_id", mux.Vars(r)["tweet_id"], &tweetId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "tweet_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRetweetingUserInfosParams

	// ------------- Required query parameter "limit" -------------

	if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "limit"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "limit", r.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "limit", Err: err})
		return
	}

	// ------------- Required query parameter "offset" -------------

	if paramValue := r.URL.Query().Get("offset"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "offset"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "offset", r.URL.Query(), &params.Offset)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "offset", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRetweetingUserInfos(w, r, tweetId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/tweets", wrapper.PostTweet).Methods("POST")

	r.HandleFunc(options.BaseURL+"/tweets/recent", wrapper.GetRecentTweetInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/recent/labels", wrapper.GetRecentLabels).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/timeline", wrapper.GetTimelineTweetInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}", wrapper.DeleteTweet).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}", wrapper.GetTweetNode).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/like", wrapper.UnlikeTweet).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/like", wrapper.LikeTweetAndNotify).Methods("POST")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/likes", wrapper.GetLikingUserInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/pin", wrapper.UnsetTweetAsPinned).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/pin", wrapper.SetTweetAsPinned).Methods("PATCH")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/quote", wrapper.PostQuoteAndNotify).Methods("POST")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/quotes", wrapper.GetQuotingUserInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/replies", wrapper.GetReplyTweetInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/reply", wrapper.PostReplyAndNotify).Methods("POST")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/retweet", wrapper.Unretweet).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/retweet", wrapper.RetweetAndNotify).Methods("POST")

	r.HandleFunc(options.BaseURL+"/tweets/{tweet_id}/retweets", wrapper.GetRetweetingUserInfos).Methods("GET")

	return r
}
