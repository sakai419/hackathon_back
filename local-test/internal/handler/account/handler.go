package account

import (
	"errors"
	"local-test/internal/key"
	"local-test/internal/model"
	"local-test/internal/service"
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
    // Get user ID
    accountID, err := key.GetAccountID(r.Context())
    if err != nil {
        utils.RespondError(w, &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "User ID not found in context",
            Err:     utils.WrapHandlerError(
				&utils.ErrOperationFailed{
					Operation: "get user ID",
					Err: err,
				},
			),
        })
        return
    }

	// Decode request
	var req CreateAccountJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Failed to decode request",
			Err:     utils.WrapHandlerError(
				&utils.ErrOperationFailed{
					Operation: "decode request",
					Err: err,
				},
			),
		})
		return
	}

	// Validate request
	if err := req.validate(); err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapHandlerError(
				&utils.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		})
		return
	}

	// Create account
	arg := req.toParams()
    arg.ID = accountID
	if err := h.svc.CreateAccount(r.Context(), arg); err != nil {
		utils.RespondError(w, utils.WrapHandlerError(
			&utils.ErrOperationFailed{
				Operation: "create account",
				Err: err,
			},
		))
		return
	}

	resp := CreateAccountResponse{Id: arg.ID}
	utils.Respond(w, resp)
}

// Delete my account
// (DELETE /accounts/me)
func (h *AccountHandler) DeleteMyAccount(w http.ResponseWriter, r *http.Request) {
	// Get user ID
	accountID, err := key.GetAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "User ID not found in context",
            Err:     utils.WrapHandlerError(
				&utils.ErrOperationFailed{
					Operation: "get user ID",
					Err: err,
				},
			),
        })
		return
	}

    // Delete account
	if err := h.svc.DeleteMyAccount(r.Context(), accountID); err != nil {
		utils.RespondError(w, utils.WrapHandlerError(
			&utils.ErrOperationFailed{
				Operation: "delete account",
				Err: err,
			},
		))
		return
	}

	utils.Respond(w, nil)
}

// validate request
func (r *CreateAccountJSONRequestBody) validate() error {
	if r.UserId == "" {
		return errors.New("UserID is required")
	}
	if r.UserName == "" {
		return errors.New("UserName is required")
	}
	return nil
}

// convert request to params
func (r *CreateAccountJSONRequestBody) toParams() *model.CreateAccountParams {
	return &model.CreateAccountParams{
		UserID:   r.UserId,
		UserName: r.UserName,
	}
}