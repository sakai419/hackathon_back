package account

import (
	"local-test/internal/database/sqlc"
	"local-test/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	svc *AccountService
}

func NewAccountHandler(svc *AccountService) *AccountHandler {
	return &AccountHandler{
		svc: svc,
	}
}

func (h *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Decode request
	var req CreateAccountRequest
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Create account
	arg := sqlc.CreateAccountParams{
		ID:       req.ID,
		UserID:   req.UserID,
		UserName: req.UserName,
	}
	if err := h.svc.CreateAccount(r.Context(), &arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

func (h *AccountHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	// Get account ID from query
	vars := mux.Vars(r)
	accountID := vars["account_id"]

	// Delete account
	if err := h.svc.DeleteAccount(r.Context(), accountID); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Encode response
	utils.Respond(w, nil)
}

func (h *AccountHandler) SuspendAccount(w http.ResponseWriter, r *http.Request) {
	// Get account ID from query
	vars := mux.Vars(r)
	accountID := vars["account_id"]

	// Suspend account
	if err := h.svc.SuspendAccount(r.Context(), accountID); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Encode response
	utils.Respond(w, nil)
}

func (h *AccountHandler) UnsuspendAccount(w http.ResponseWriter, r *http.Request) {
	// Get account ID from query
	vars := mux.Vars(r)
	accountID := vars["account_id"]

	// Unsuspend account
	if err := h.svc.UnsuspendAccount(r.Context(), accountID); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Encode response
	utils.Respond(w, nil)
}