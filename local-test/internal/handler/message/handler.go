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

func NewMessageHandler(svc *service.Service) ServerInterface {
	return &MessageHandler{
		svc: svc,
	}
}

// Get Messages
// (GET /conversations/{user_id}/messages)
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
		utils.RespondError(w, apperrors.NewHandlerError("get messages", err))
		return
	}

	utils.Respond(w, convertToMessageResponse(messages))
}

// Send Message
// (POST /conversations/{user_id}/messages)
func (h *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request, _ string) {
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

	// Validate request
	if err := utils.ValidateRequiredFields(req); err != nil {
		utils.RespondError(w, apperrors.NewRequiredParamError("request body", err))
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
			SenderAccountId: m.SenderAccountID,
			Content:         m.Content,
			IsRead:          m.IsRead,
			CreatedAt:       m.CreatedAt,
		})
	}
	return res
}