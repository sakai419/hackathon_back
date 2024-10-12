package account

import (
	"errors"
	contextKey "local-test/internal/context"
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
    userID, err := contextKey.GetUserID(r.Context())
    if err != nil {
        utils.RespondError(w, &utils.AppError{
            Status:  http.StatusInternalServerError,
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "User ID not found in context",
            Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "get user ID", Err: err}),
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
			Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "decode request", Err: err}),
		})
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "validate request", Err: err}),
		})
		return
	}

	// Create account
	arg := req.ToParams()
    arg.ID = userID
	if err := h.svc.CreateAccount(r.Context(), arg); err != nil {
		utils.RespondError(w, utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "create account", Err: err}))
		return
	}

	resp := CreateAccountResponse{Id: arg.ID}
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
            Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "get user ID", Err: err}),
        })
		return
	}

    // Delete account
	if err := h.svc.DeleteMyAccount(r.Context(), userID); err != nil {
		utils.RespondError(w, utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "delete account", Err: err}))
		return
	}

	utils.Respond(w, nil)
}

// validate request
func (r *CreateAccountJSONRequestBody) Validate() error {
	if r.UserId == "" {
		return errors.New("UserID is required")
	}
	if r.UserName == "" {
		return errors.New("UserName is required")
	}
	return nil
}

// convert request to params
func (r *CreateAccountJSONRequestBody) ToParams() *model.CreateAccountParams {
	return &model.CreateAccountParams{
		UserID:   r.UserId,
		UserName: r.UserName,
	}
}