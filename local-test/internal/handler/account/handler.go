package account

import (
	"local-test/internal/key"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type AccountHandler struct {
	svc *service.Service
}

func NewAccountHandler(svc *service.Service) ServerInterface {
	return &AccountHandler{
		svc: svc,
	}
}

// Create a new account
// (POST /accounts)
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
    // Get client account ID
    clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request
	var req CreateAccountJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Failed to decode request",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "decode request",
					Err: err,
				},
			),
		})
		return
	}

	// Create account
	if err := h.svc.CreateAccount(r.Context(), &model.CreateAccountParams{
		ID:              clientAccountID,
		UserID: 		 req.UserId,
		UserName: 		 req.UserName,
	}); err != nil {
		utils.RespondError(w, apperrors.WrapHandlerError(
			&apperrors.ErrOperationFailed{
				Operation: "create account",
				Err: err,
			},
		))
		return
	}

	resp := CreateAccountResponse{Id: clientAccountID}
	utils.Respond(w, resp)
}

// Delete my account
// (DELETE /accounts/me)
func (h *AccountHandler) DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	// Get user ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	if err := h.svc.DeleteMyAccount(r.Context(), clientAccountID); err != nil {
		utils.RespondError(w, apperrors.WrapHandlerError(
			&apperrors.ErrOperationFailed{
				Operation: "delete account",
				Err: err,
			},
		))
		return
	}

	utils.Respond(w, nil)
}

// Get client account ID from context
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