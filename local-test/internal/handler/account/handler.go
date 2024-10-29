package account

import (
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
    clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request
	var req CreateAccountJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Create account
	if err := h.svc.CreateAccount(r.Context(), &model.CreateAccountParams{
		ID:              clientAccountID,
		UserID: 		 req.UserId,
		UserName: 		 req.UserName,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("create account", err))
		return
	}

	utils.Respond(w, CreateAccountResponse{Id: clientAccountID})
}

// Delete my account
// (DELETE /accounts/)
func (h *AccountHandler) DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	// Get user ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	if err := h.svc.DeleteMyAccount(r.Context(), clientAccountID); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("delete account", err))
		return
	}

	utils.Respond(w, nil)
}