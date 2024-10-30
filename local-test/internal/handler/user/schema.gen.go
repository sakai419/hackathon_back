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

// GetUserTweetsResponse defines model for GetUserTweetsResponse.
type GetUserTweetsResponse struct {
	OmittedReplyExist *bool      `json:"OmittedReplyExist"`
	OriginalTweet     *TweetInfo `json:"OriginalTweet,omitempty"`
	ParentReply       *TweetInfo `json:"ParentReply,omitempty"`
	Tweet             TweetInfo  `json:"Tweet"`
}

// GetUserTweetsResponses defines model for GetUserTweetsResponses.
type GetUserTweetsResponses = []GetUserTweetsResponse

// Media defines model for Media.
type Media struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

// TweetInfo defines model for TweetInfo.
type TweetInfo struct {
	Code          *string            `json:"Code"`
	Content       *string            `json:"Content"`
	CreatedAt     time.Time          `json:"CreatedAt"`
	HasLiked      bool               `json:"HasLiked"`
	HasRetweeted  bool               `json:"HasRetweeted"`
	IsPinned      bool               `json:"IsPinned"`
	IsQuote       bool               `json:"IsQuote"`
	IsReply       bool               `json:"IsReply"`
	LikesCount    int32              `json:"LikesCount"`
	Media         *Media             `json:"Media,omitempty"`
	RepliesCount  int32              `json:"RepliesCount"`
	RetweetsCount int32              `json:"RetweetsCount"`
	TweetID       int64              `json:"TweetID"`
	UserInfo      UserInfoWithoutBio `json:"UserInfo"`
}

// TweetInfos defines model for TweetInfos.
type TweetInfos = []TweetInfo

// UserInfoWithoutBio defines model for UserInfoWithoutBio.
type UserInfoWithoutBio struct {
	ProfileImageUrl string `json:"profile_image_url"`
	UserId          string `json:"user_id"`
	Username        string `json:"username"`
}

// GetUserLikesParams defines parameters for GetUserLikes.
type GetUserLikesParams struct {
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
	// Get liked tweets by user
	// (GET /users/{user_id}/likes)
	GetUserLikes(w http.ResponseWriter, r *http.Request, userId string, params GetUserLikesParams)
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

	r.HandleFunc(options.BaseURL+"/users/{user_id}/likes", wrapper.GetUserLikes).Methods("GET")

	r.HandleFunc(options.BaseURL+"/users/{user_id}/tweets", wrapper.GetUserTweets).Methods("GET")

	return r
}
