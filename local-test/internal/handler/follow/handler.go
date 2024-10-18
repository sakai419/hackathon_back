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
	// Check if the user is suspended
	if isClientSuspended(w, r) || isTargetSuspended(w, r) {
		return
	}

	// Get client ID
	clientID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
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
	// Check if the user is suspended
	if isClientSuspended(w, r) {
		return
	}

	// Get user ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
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
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Get followers
	arg := &model.GetFollowerInfosParams{
		FollowingAccountID: accountIDFromPath,
		Limit:              int32(params.Limit),
		Offset:             int32(params.Offset),
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
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Get followings
	arg := &model.GetFollowingInfosParams{
		FollowerAccountID: accountIDFromPath,
		Limit:             int32(params.Limit),
		Offset:            int32(params.Offset),
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
	// Check if the user is suspended
	if isClientSuspended(w, r) || isTargetSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Send follow request
	arg := &model.RequestFollowAndNotifyParams{
		RequesterAccountID: clientAccountID,
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
	// Check if the user is suspended
	if isClientSuspended(w, r) || isTargetSuspended(w, r) {
		return
	}

	// Ger client account ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Accept follow request
	arg := &model.AcceptFollowRequestAndNotifyParams{
		RequestedAccountID: clientAccountID,
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
	// Check if the user is suspended
	if isClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account id from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
	if !ok {
		return
	}

	// Reject follow request
	arg := &model.RejectFollowRequestParams{
		RequestedAccountID: clientAccountID,
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

func isClientSuspended(w http.ResponseWriter, r *http.Request) bool {
	isClientSuspended, err := key.GetIsClientSuspended(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_suspended",
					Err: err,
				},
			),
		})
		return true
	}

	if isClientSuspended {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is suspended",
				},
			),
		})
		return true
	}

	return false
}

func isTargetSuspended(w http.ResponseWriter, r *http.Request) bool {
	isTargetSuspended, err := key.GetIsTargetSuspended(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_suspended",
					Err: err,
				},
			),
		})
		return true
	}

	if isTargetSuspended {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is suspended",
				},
			),
		})
		return true
	}

	return false
}

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

func getAccountIDFromPath(w http.ResponseWriter, r *http.Request) (string, bool) {
	accountID, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusBadRequest,
				Code:    "BAD_REQUEST",
				Message: "Account ID not found in path",
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
	return accountID, true
}


func convertToUserAndProfileInfos(followerInfos []*model.UserInfo) []UserInfo {
	var resp []UserInfo
	for _, followerInfo := range followerInfos {
		resp = append(resp, UserInfo{
			Bio: 	followerInfo.Bio,
			ProfileImageUrl: followerInfo.ProfileImageURL,
			UserId: followerInfo.UserID,
			UserName: followerInfo.UserName,
		})
	}
	return resp
}