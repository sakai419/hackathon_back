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

func NewConversationHandler(svc *service.Service) *ConversationHandler {
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

// ErrorHandlerFunc is the error handler for the conversation handler
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

func convertToConversationResponse(conversations []*model.ConversationResponse) []*Conversation {
	var response []*Conversation
	for _, c := range conversations {
		response = append(response, &Conversation{
			Id: 			c.ID,
			OpponentInfo: 	UserInfoWithoutBio{
				UserId: 		 c.OpponentInfo.UserID,
				Username: 		 c.OpponentInfo.UserName,
				ProfileImageUrl: c.OpponentInfo.ProfileImageURL,
			},
			LastMessageTime: c.LastMessageTime,
			Content: 		 c.Content,
			SenderUserId: 	 c.SenderUserID,
			IsRead: 		 c.IsRead,
		})
	}
	return response
}