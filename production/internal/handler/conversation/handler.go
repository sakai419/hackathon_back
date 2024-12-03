package conversation

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type ConversationHandler struct {
	svc *service.Service
}

func NewConversationHandler(svc *service.Service) ServerInterface {
	return &ConversationHandler{
		svc: svc,
	}
}

// Get Conversations
// (GET /conversations)
func (h *ConversationHandler) GetConversations(w http.ResponseWriter, r *http.Request, params GetConversationsParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get conversations
	conversations, err := h.svc.GetConversations(r.Context(), &model.GetConversationsParams{
		ClientAccountID: clientAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get conversations", err))
		return
	}

	utils.Respond(w, convertToConversationResponse(conversations))
}

// Get Unread Conversation Count
// (GET /conversations/unread/count)
func (h *ConversationHandler) GetUnreadConversationCount(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread conversation count
	count, err := h.svc.GetUnreadConversationCount(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get unread conversation count", err))
		return
	}

	utils.Respond(w, UnreadConversationCountResponse{Count: count})
}

// Get Messages
// (GET /conversations/{user_id}/messages)
func (h *ConversationHandler) GetMessages(w http.ResponseWriter, r *http.Request, _ string, params GetMessagesParams) {
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

	// Get messages
	messages, err := h.svc.GetMessages(r.Context(), &model.GetMessagesParams{
		ClientAccountID: clientAccountID,
		TargetAccountID: targetAccountID,
		Limit:           params.Limit,
		Offset:          params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get messages", err))
		return
	}

	utils.Respond(w, convertToMessageResponse(messages))
}

// Send Message
// (POST /conversations/{user_id}/messages)
func (h *ConversationHandler) SendMessage(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) || utils.IsTargetSuspended(w, r) {
		return
	}

	// Get client account ID
	clidentAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Decode request
	var req SendMessageJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Send message
	err := h.svc.SendMessage(r.Context(), &model.SendMessageParams{
		ClientAccountID: clidentAccountID,
		TargetAccountID: targetAccountID,
		Content:         req.Content,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("send message", err))
		return
	}

	utils.Respond(w, nil)
}

// Mark message as read
// (PATCH /conversations/{user_id}/messages)
func (h *ConversationHandler) MarkMessagesAsRead(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client account ID
	clidentAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Mark message as read
	if err := h.svc.MarkMessagesAsRead(r.Context(), &model.MarkMessagesAsReadParams{
		ClientAccountID: clidentAccountID,
		TargetAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("mark messages as read", err))
		return
	}

	utils.Respond(w, nil)
}

// ErrorHandlerFunc is the error handler for the message handler
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

func convertToMessageResponse(messages []*model.MessageResponse) []*MessageResponse {
	var res []*MessageResponse
	for _, m := range messages {
		res = append(res, &MessageResponse{
			Id:              m.ID,
			SenderUserId:    m.SenderUserID,
			Content:         m.Content,
			IsRead:          m.IsRead,
			CreatedAt:       m.CreatedAt,
		})
	}
	return res
}

func convertToConversationResponse(conversations []*model.ConversationResponse) []*Conversation {
	var response []*Conversation
	for _, c := range conversations {
		response = append(response, &Conversation{
			Id: 			c.ID,
			OpponentInfo: 	UserInfoWithoutBio{
				UserId: 		 c.OpponentInfo.UserID,
				UserName: 		 c.OpponentInfo.UserName,
				ProfileImageUrl: c.OpponentInfo.ProfileImageURL,
				IsPrivate: 		 c.OpponentInfo.IsPrivate,
				IsAdmin: 		 c.OpponentInfo.IsAdmin,
				IsFollowing:     c.OpponentInfo.IsFollowing,
				IsFollowed:      c.OpponentInfo.IsFollowed,
			},
			LastMessageTime: c.LastMessageTime,
			Content: 		 c.Content,
			SenderUserId: 	 c.SenderUserID,
			IsRead: 		 c.IsRead,
		})
	}
	return response
}