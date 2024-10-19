package conversation

import (
	"errors"
	"local-test/internal/key"
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
	clientAccountID, ok := getClientAccountID(w, r)
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
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, convertToConversationResponse(conversations))
}

// Get Unread Conversation Count
// (GET /conversations/unread/count)
func (h *ConversationHandler) GetUnreadConversationCount(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get unread conversation count
	count, err := h.svc.GetUnreadConversationCount(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, UnreadConversationCountResponse{Count: count})
}

func ErrHandleFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid parameter format",
			Err:     err,
		})
		return
	} else {
		utils.RespondError(w, err)
	}
}

func getClientAccountID(w http.ResponseWriter, r *http.Request) (string, bool) {
	clidentAccountID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusInternalServerError,
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "Account ID not found in context",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return clidentAccountID, true
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