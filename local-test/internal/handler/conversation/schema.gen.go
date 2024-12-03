// Package conversation provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package conversation

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

// Conversation defines model for Conversation.
type Conversation struct {
	Content         string             `json:"content"`
	Id              int64              `json:"id"`
	IsRead          bool               `json:"is_read"`
	LastMessageTime time.Time          `json:"last_message_time"`
	OpponentInfo    UserInfoWithoutBio `json:"opponent_info"`
	SenderUserId    string             `json:"sender_user_id"`
}

// ConversationResponses defines model for ConversationResponses.
type ConversationResponses struct {
	Conversations *[]Conversation `json:"conversations,omitempty"`
}

// MessageRequest defines model for MessageRequest.
type MessageRequest struct {
	Content string `json:"content"`
}

// MessageResponse defines model for MessageResponse.
type MessageResponse struct {
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	Id           int64     `json:"id"`
	IsRead       bool      `json:"is_read"`
	SenderUserId string    `json:"sender_user_id"`
}

// MessageResponses defines model for MessageResponses.
type MessageResponses = []MessageResponse

// UnreadConversationCountResponse defines model for UnreadConversationCountResponse.
type UnreadConversationCountResponse struct {
	Count int64 `json:"count"`
}

// UserInfoWithoutBio defines model for UserInfoWithoutBio.
type UserInfoWithoutBio struct {
	IsAdmin         bool   `json:"is_admin"`
	IsFollowed      bool   `json:"is_followed"`
	IsFollowing     bool   `json:"is_following"`
	IsPending       bool   `json:"is_pending"`
	IsPrivate       bool   `json:"is_private"`
	ProfileImageUrl string `json:"profile_image_url"`
	UserId          string `json:"user_id"`
	UserName        string `json:"user_name"`
}

// GetConversationsParams defines parameters for GetConversations.
type GetConversationsParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// GetMessagesParams defines parameters for GetMessages.
type GetMessagesParams struct {
	Limit  int32 `form:"limit" json:"limit"`
	Offset int32 `form:"offset" json:"offset"`
}

// SendMessageJSONRequestBody defines body for SendMessage for application/json ContentType.
type SendMessageJSONRequestBody = MessageRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get conversations of a user
	// (GET /conversations)
	GetConversations(w http.ResponseWriter, r *http.Request, params GetConversationsParams)
	// Get unread conversation count of a user
	// (GET /conversations/unread/count)
	GetUnreadConversationCount(w http.ResponseWriter, r *http.Request)
	// Get messages of a user
	// (GET /conversations/{user_id}/messages)
	GetMessages(w http.ResponseWriter, r *http.Request, userId string, params GetMessagesParams)
	// Send a message to a user
	// (POST /conversations/{user_id}/messages)
	SendMessage(w http.ResponseWriter, r *http.Request, userId string)
	// Mark messages as read
	// (PATCH /conversations/{user_id}/messages/read)
	MarkMessagesAsRead(w http.ResponseWriter, r *http.Request, userId string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetConversations operation middleware
func (siw *ServerInterfaceWrapper) GetConversations(w http.ResponseWriter, r *http.Request) {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetConversationsParams

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
		siw.Handler.GetConversations(w, r, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetUnreadConversationCount operation middleware
func (siw *ServerInterfaceWrapper) GetUnreadConversationCount(w http.ResponseWriter, r *http.Request) {

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUnreadConversationCount(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetMessages operation middleware
func (siw *ServerInterfaceWrapper) GetMessages(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetMessagesParams

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
		siw.Handler.GetMessages(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// SendMessage operation middleware
func (siw *ServerInterfaceWrapper) SendMessage(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.SendMessage(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// MarkMessagesAsRead operation middleware
func (siw *ServerInterfaceWrapper) MarkMessagesAsRead(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", mux.Vars(r)["user_id"], &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.MarkMessagesAsRead(w, r, userId)
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

	r.HandleFunc(options.BaseURL+"/conversations", wrapper.GetConversations).Methods("GET")

	r.HandleFunc(options.BaseURL+"/conversations/unread/count", wrapper.GetUnreadConversationCount).Methods("GET")

	r.HandleFunc(options.BaseURL+"/conversations/{user_id}/messages", wrapper.GetMessages).Methods("GET")

	r.HandleFunc(options.BaseURL+"/conversations/{user_id}/messages", wrapper.SendMessage).Methods("POST")

	r.HandleFunc(options.BaseURL+"/conversations/{user_id}/messages/read", wrapper.MarkMessagesAsRead).Methods("PATCH")

	return r
}
