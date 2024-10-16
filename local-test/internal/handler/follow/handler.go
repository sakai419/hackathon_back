package follow

import (
	"errors"
	"local-test/internal/key"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type FollowHandler struct {
	svc *service.Service
}

func NewFollowHandler(svc *service.Service) ServerInterface {
	return &FollowHandler{
		svc: svc,
	}
}

// Follow a user
// (POST /users/{user_id}/follow)
func (h *FollowHandler) FollowAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client ID
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Follow user
	arg := &model.FollowAndNotifyParams{
		FollowerAccountID:  clientID,
		FollowingAccountID: accountIDFromPath,
	}
	if err := h.svc.FollowAndNotify(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Unfollow a user
// (DELETE /users/{user_id}/follow)
func (h *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request, _ string) {
	// Get user ID
	clientAccountID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusUnauthorized,
			Code:    "UNAUTHORIZED",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Unfollow user
	arg := &model.UnfollowParams{
		FollowerAccountID:  clientAccountID,
		FollowingAccountID: accountIDFromPath,
	}
	if err := h.svc.Unfollow(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Get followers
// (GET /users/{user_id}/followers)
func (h *FollowHandler) GetFollowerInfos(w http.ResponseWriter, r *http.Request, _ string, params GetFollowerInfosParams) {
	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Validate request
	if err := params.validate(); err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		})
		return
	}

	// Get followers
	arg := &model.GetFollowerInfosParams{
		FollowingAccountID: accountIDFromPath,
		Limit:              int32(*params.Limit),
		Offset:             int32(*params.Offset),
	}
	followerInfos, err := h.svc.GetFollowerInfos(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}


	// Convert to response
	resp := convertToUserAndProfileInfos(followerInfos)

	utils.Respond(w, resp)
}

// Get followings
// (GET /users/{user_id}/followings)
func (h *FollowHandler) GetFollowingInfos(w http.ResponseWriter, r *http.Request, _ string, params GetFollowingInfosParams) {
	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Validate request
	if err := params.validate(); err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Invalid request",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
		})
		return
	}

	// Get followings
	arg := &model.GetFollowingInfosParams{
		FollowerAccountID: accountIDFromPath,
		Limit:             int32(*params.Limit),
		Offset:            int32(*params.Offset),
	}
	followingInfos, err := h.svc.GetFollowingInfos(r.Context(), arg)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Convert to response
	resp := convertToUserAndProfileInfos(followingInfos)

	utils.Respond(w, resp)
}

// Send follow request
// (POST /users/{user_id}/follow-request)
func (h *FollowHandler) RequestFollowAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client ID
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return

	}

	// Send follow request
	arg := &model.RequestFollowAndNotifyParams{
		RequesterAccountID: clientID,
		RequestedAccountID: accountIDFromPath,
	}
	if err := h.svc.RequestFollowAndNotify(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Accept follow request
// (PUT /users/me/follow-request/{user_id}/accept)
func (h *FollowHandler) AcceptFollowRequestAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Ger client ID
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Accept follow request
	arg := &model.AcceptFollowRequestAndNotifyParams{
		RequestedAccountID: clientID,
		RequesterAccountID: accountIDFromPath,
	}
	if err := h.svc.AcceptFollowRequestAndNotify(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Reject follow request
// (DELETE /users/me/follow-request/{user_id}/reject)
func (h *FollowHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request, _ string) {
	// Get client id
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Get account id from path
	accountIDFromPath, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusBadRequest,
			Code:    "BAD_REQUEST",
			Message: "Account ID not found in path",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

	// Reject follow request
	arg := &model.RejectFollowRequestParams{
		RequestedAccountID: clientID,
		RequesterAccountID: accountIDFromPath,
	}
	if err := h.svc.RejectFollowRequest(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

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

func convertToUserAndProfileInfos(followerInfos []*model.UserAndProfileInfo) []UserAndProfileInfo {
	var resp []UserAndProfileInfo
	for _, followerInfo := range followerInfos {
		resp = append(resp, UserAndProfileInfo{
			Bio: 	followerInfo.Bio,
			ProfileImageUrl: followerInfo.ProfileImageURL,
			UserId: followerInfo.UserID,
			UserName: followerInfo.UserName,
		})
	}
	return resp
}

func (r *GetFollowerInfosParams) validate() error {
	if r.Limit == nil {
		return errors.New("limit is required")
	}
	if r.Offset == nil {
		return errors.New("offset is required")
	}
	return nil
}

func (r *GetFollowingInfosParams) validate() error {
	if r.Limit == nil {
		return errors.New("limit is required")
	}
	if r.Offset == nil {
		return errors.New("offset is required")
	}
	return nil
}