// Package user provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package user

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

// Media defines model for Media.
type Media struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// TweetInfo defines model for TweetInfo.
type TweetInfo struct {
	Code          *Code              `json:"code,omitempty"`
	Content       *string            `json:"content,omitempty"`
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
	OmittedReplyExist *bool      `json:"omitted_reply_exist,omitempty"`
	OriginalTweet     *TweetInfo `json:"original_tweet,omitempty"`
	ParentReply       *TweetInfo `json:"parent_reply,omitempty"`
	Tweet             TweetInfo  `json:"tweet"`
}

// TweetNodes defines model for TweetNodes.
type TweetNodes = []TweetNode

// UserInfo defines model for UserInfo.
type UserInfo struct {
	Bio             string `json:"bio"`
	IsAdmin         bool   `json:"is_admin"`
	IsPrivate       bool   `json:"is_private"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserId          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

// UserInfoWithoutBio defines model for UserInfoWithoutBio.
type UserInfoWithoutBio struct {
	IsAdmin         bool   `json:"is_admin"`
	IsPrivate       bool   `json:"is_private"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserId          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

// UserProfile defines model for UserProfile.
type UserProfile struct {
	BannerImageUrl string    `json:"banner_image_url"`
	CreatedAt      time.Time `json:"created_at"`
	FollowerCount  int64     `json:"follower_count"`
	FollowingCount int64     `json:"following_count"`
	IsFollowed     bool      `json:"is_followed"`
	TweetCount     int64     `json:"tweet_count"`
	UserInfo       UserInfo  `json:"user_info"`
}

// GetUserLikesParams defines parameters for GetUserLikes.
type GetUserLikesParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetUserRetweetsParams defines parameters for GetUserRetweets.
type GetUserRetweetsParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetUserTweetsParams defines parameters for GetUserTweets.
type GetUserTweetsParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get client profile
	// (GET /users/me)
	GetClientProfile(w http.ResponseWriter, r *http.Request)
	// Get user profile
	// (GET /users/{user_id})
	GetUserProfile(w http.ResponseWriter, r *http.Request, userId string)
	// Get liked tweets by user
	// (GET /users/{user_id}/likes)
	GetUserLikes(w http.ResponseWriter, r *http.Request, userId string, params GetUserLikesParams)
	// Get retweeted tweets by user
	// (GET /users/{user_id}/retweets)
	GetUserRetweets(w http.ResponseWriter, r *http.Request, userId string, params GetUserRetweetsParams)
	// Get tweets by user
	// (GET /users/{user_id}/tweets)
	GetUserTweets(w http.ResponseWriter, r *http.Request, userId string, params GetUserTweetsParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetClientProfile operation middleware
func (siw *ServerInterfaceWrapper) GetClientProfile(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetClientProfile(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserProfile operation middleware
func (siw *ServerInterfaceWrapper) GetUserProfile(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserProfile(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserLikes operation middleware
func (siw *ServerInterfaceWrapper) GetUserLikes(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserLikesParams

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
		siw.Handler.GetUserLikes(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserRetweets operation middleware
func (siw *ServerInterfaceWrapper) GetUserRetweets(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserRetweetsParams

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
		siw.Handler.GetUserRetweets(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUserTweets operation middleware
func (siw *ServerInterfaceWrapper) GetUserTweets(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserTweetsParams

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
		siw.Handler.GetUserTweets(w, r, userId, params)
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

	r.HandleFunc(options.BaseURL+"/users/me", wrapper.GetClientProfile).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{user_id}", wrapper.GetUserProfile).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{user_id}/likes", wrapper.GetUserLikes).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{user_id}/retweets", wrapper.GetUserRetweets).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{user_id}/tweets", wrapper.GetUserTweets).Methods("GET")

	return r
}