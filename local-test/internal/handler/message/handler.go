package message

import (
	"errors"
	"local-test/internal/key"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type MessageHandler struct {
	svc *service.Service
}

func NewMessageHandler(svc *service.Service) *MessageHandler {
	return &MessageHandler{
		svc: svc,
	}
}

// Get Messages
// (GET /messages/{user_id})
func (h *MessageHandler) GetMessages(w http.ResponseWriter, r *http.Request, _ string, params GetMessagesParams) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Get messages
	arg := &model.GetMessagesParams{
		ClientAccountID: clientID,
		TargetAccountID: accountIDFromPath,
		Limit:           params.Limit,
		Offset:          params.Offset,
	}
	messages, err := h.svc.GetMessages(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Convert to response
	resp := convertToMessageResponse(messages)

	utils.Respond(w, resp)
}

// Send Message
// (POST /messages/{user_id})
func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Decode request
	var params SendMessageJSONRequestBody
	if err := utils.Decode(r, &params); err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusBadRequest,
				Code:    "BAD_REQUEST",
				Message: "Failed to decode request",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "decode request",
						Err: err,
					},
				),
			},
		)
		return
	}

	// Send message
	arg := &model.SendMessageParams{
		ClientAccountID: clientID,
		TargetAccountID: accountIDFromPath,
		Content:         params.Content,
	}
	err := h.svc.SendMessage(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Mark message as read
// (PATCH /messages/{user_id}/read)
func (h *MessageHandler) MarkMessagesAsRead(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Mark message as read
	arg := &model.MarkMessagesAsReadParams{
		ClientAccountID: clientID,
		TargetAccountID: accountIDFromPath,
	}
	if err := h.svc.MarkMessagesAsRead(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// ErrHandleFunc handles errors
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
	clientID, err := key.GetClientAccountID(r.Context())
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
	return clientID, true
}

func getAccountIDFromPath(w http.ResponseWriter, r *http.Request) (string, bool) {
	accountID, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusBadRequest,
				Code:    "BAD_REQUEST",
				Message: "Account ID not found in path",
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
	return accountID, true
}

func convertToMessageResponse(messages []*model.MessageResponse) []*MessageResponse {
	var res []*MessageResponse
	for _, m := range messages {
		res = append(res, &MessageResponse{
			SenderAccountId: m.SenderAccountID,
			Content:         m.Content,
			IsRead:          m.IsRead,
			CreatedAt:       m.CreatedAt,
		})
	}
	return res
}