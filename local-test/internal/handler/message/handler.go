package message

import (
	"errors"
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
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, convertToMessageResponse(messages))
}

// Send Message
// (POST /messages/{user_id})
func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request, _ string) {
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
	err := h.svc.SendMessage(r.Context(), &model.SendMessageParams{
		ClientAccountID: clidentAccountID,
		TargetAccountID: targetAccountID,
		Content:         params.Content,
	})
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Mark message as read
// (PATCH /messages/{user_id}/read)
func (h *MessageHandler) MarkMessagesAsRead(w http.ResponseWriter, r *http.Request, _ string) {
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