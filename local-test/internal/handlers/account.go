package handlers

import (
	"local-test/pkg/database/generated"
	"local-test/internal/models"
	"local-test/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// Decode request
	var req models.CreateAccountRequest
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Create account
	arg := queries.CreateAccountParams{
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

func (h *Handler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) SuspendAccount(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) UnsuspendAccount(w http.ResponseWriter, r *http.Request) {
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