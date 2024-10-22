// Package follow provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package follow

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Count defines model for Count.
type Count struct {
	// Count The count of notifications.
	Count int64 `json:"count"`
}

// FollowCounts defines model for FollowCounts.
type FollowCounts struct {
	// FollowersCount The count of followers.
	FollowersCount int64 `json:"followers_count"`

	// FollowingCount The count of following.
	FollowingCount int64 `json:"following_count"`
}

// UserInfo defines model for UserInfo.
type UserInfo struct {
	// Bio The bio of the user.
	Bio string `json:"bio"`

	// ProfileImageUrl URL of the user's profile image.
	ProfileImageUrl string `json:"profile_image_url"`

	// UserId The ID of the user.
	UserId string `json:"user_id"`

	// UserName The name of the user.
	UserName string `json:"user_name"`
}

// UserInfos defines model for UserInfos.
type UserInfos = []UserInfo

// GetFollowerInfosParams defines parameters for GetFollowerInfos.
type GetFollowerInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetFollowingInfosParams defines parameters for GetFollowingInfos.
type GetFollowingInfosParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get the number of followers and following of a user
	// (GET /follows/count/{user_id})
	GetFollowCounts(w http.ResponseWriter, r *http.Request, userId string)
	// Get followers of a user
	// (GET /follows/followers/{user_id})
	GetFollowerInfos(w http.ResponseWriter, r *http.Request, userId string, params GetFollowerInfosParams)
	// Get users followed by a user
	// (GET /follows/following/{user_id})
	GetFollowingInfos(w http.ResponseWriter, r *http.Request, userId string, params GetFollowingInfosParams)
	// Get the number of follow requests sent to the user
	// (GET /follows/requests/received/count)
	GetFollowRequestCount(w http.ResponseWriter, r *http.Request)
	// Reject a follow request
	// (DELETE /follows/requests/received/{user_id})
	RejectFollowRequest(w http.ResponseWriter, r *http.Request, userId string)
	// Accept a follow request
	// (PATCH /follows/requests/received/{user_id}/accept)
	AcceptFollowRequestAndNotify(w http.ResponseWriter, r *http.Request, userId string)
	// Send a follow request to a user
	// (POST /follows/requests/{user_id})
	RequestFollowAndNotify(w http.ResponseWriter, r *http.Request, userId string)
	// Unfollow a user
	// (DELETE /follows/{user_id})
	Unfollow(w http.ResponseWriter, r *http.Request, userId string)
	// Follow a user
	// (POST /follows/{user_id})
	FollowAndNotify(w http.ResponseWriter, r *http.Request, userId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetFollowCounts operation middleware
func (siw *ServerInterfaceWrapper) GetFollowCounts(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFollowCounts(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetFollowerInfos operation middleware
func (siw *ServerInterfaceWrapper) GetFollowerInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFollowerInfosParams

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
		siw.Handler.GetFollowerInfos(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetFollowingInfos operation middleware
func (siw *ServerInterfaceWrapper) GetFollowingInfos(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetFollowingInfosParams

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
		siw.Handler.GetFollowingInfos(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetFollowRequestCount operation middleware
func (siw *ServerInterfaceWrapper) GetFollowRequestCount(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFollowRequestCount(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RejectFollowRequest operation middleware
func (siw *ServerInterfaceWrapper) RejectFollowRequest(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RejectFollowRequest(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// AcceptFollowRequestAndNotify operation middleware
func (siw *ServerInterfaceWrapper) AcceptFollowRequestAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.AcceptFollowRequestAndNotify(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// RequestFollowAndNotify operation middleware
func (siw *ServerInterfaceWrapper) RequestFollowAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.RequestFollowAndNotify(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// Unfollow operation middleware
func (siw *ServerInterfaceWrapper) Unfollow(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.Unfollow(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// FollowAndNotify operation middleware
func (siw *ServerInterfaceWrapper) FollowAndNotify(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.FollowAndNotify(w, r, userId)
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

	r.HandleFunc(options.BaseURL+"/follows/count/{user_id}", wrapper.GetFollowCounts).Methods("GET")

	r.HandleFunc(options.BaseURL+"/follows/followers/{user_id}", wrapper.GetFollowerInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/follows/following/{user_id}", wrapper.GetFollowingInfos).Methods("GET")

	r.HandleFunc(options.BaseURL+"/follows/requests/received/count", wrapper.GetFollowRequestCount).Methods("GET")

	r.HandleFunc(options.BaseURL+"/follows/requests/received/{user_id}", wrapper.RejectFollowRequest).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/follows/requests/received/{user_id}/accept", wrapper.AcceptFollowRequestAndNotify).Methods("PATCH")

	r.HandleFunc(options.BaseURL+"/follows/requests/{user_id}", wrapper.RequestFollowAndNotify).Methods("POST")

	r.HandleFunc(options.BaseURL+"/follows/{user_id}", wrapper.Unfollow).Methods("DELETE")

	r.HandleFunc(options.BaseURL+"/follows/{user_id}", wrapper.FollowAndNotify).Methods("POST")

	return r
}
