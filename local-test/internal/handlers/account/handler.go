package account

import (
	contextKey "local-test/internal/context"
	"local-test/internal/models"
	"local-test/internal/services"
	"local-test/pkg/utils"
	"net/http"
)

type AccountHandler struct {
	svc *services.Service
}

func NewAccountHandler(svc *services.Service) ServerInterface {
	return &AccountHandler{
		svc: svc,
	}
}

// Create a new account
// (POST /accounts)
func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Decode request
	var req models.CreateAccountRequest
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Failed to decode request",
			Err:     utils.WrapHandlerError(err, "failed to decode request"),
		})
		return
	}

	// validate request
	if err := req.Validate(); err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapHandlerError(err, "invalid request"),
		})
		return
	}

	// Create account
	arg := req.ToParams()
	if err := h.svc.CreateAccount(r.Context(), arg); err != nil {
		utils.RespondError(w, utils.WrapHandlerError(err, "failed to create account"))
		return
	}

	resp := models.CreateAccountResponse{ID: arg.ID}
	utils.Respond(w, resp)
}

// Delete my account
// (DELETE /accounts/me)
func (h *AccountHandler) DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	// Get user ID
	userID, err := contextKey.GetUserID(r.Context())
	if err != nil {
		utils.RespondError(w, &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "User ID not found in context",
            Err:     err,
        })
		return
	}

    // Delete account
	if err := h.svc.DeleteMyAccount(r.Context(), userID); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}