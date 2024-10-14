package follow

import (
	"errors"
	"local-test/internal/key"
	"local-test/internal/model"
	"local-test/internal/service"
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
func (h *FollowHandler) FollowUser(w http.ResponseWriter, r *http.Request, userID string) {
	// Get user ID
	followerAccountID, err := key.GetAccountID(r.Context())
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

	// Follow user
	arg := &model.FollowUserParams{
		FollowerAccountID: followerAccountID,
		FollowingUserID:    userID,
	}
	if err := h.svc.FollowUser(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Unfollow a user
// (DELETE /users/{user_id}/follow)
func (h *FollowHandler) UnfollowUser(w http.ResponseWriter, r *http.Request, userID string) {
	// Get user ID
	followerAccountID, err := key.GetAccountID(r.Context())
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

	// Unfollow user
	arg := &model.UnfollowUserParams{
		FollowerAccountID: followerAccountID,
		FollowingUserID:    userID,
	}
	if err := h.svc.UnfollowUser(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Get followers
// (GET /users/{user_id}/followers)
func (h *FollowHandler) GetFollowers(w http.ResponseWriter, r *http.Request, userID string, params GetFollowersParams) {
	// Get followers
	arg := &model.GetFollowerInfosParams{
		FollowingUserID: userID,
		Limit:  int32(*params.Limit),
		Offset: int32(*params.Offset),
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
func (h *FollowHandler) GetFollowings(w http.ResponseWriter, r *http.Request, userID string, params GetFollowingsParams) {
	// Get followings
	arg := &model.GetFollowingInfosParams{
		FollowerUserID: userID,
		Limit:  int32(*params.Limit),
		Offset: int32(*params.Offset),
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
func (h *FollowHandler) SendFollowRequest(w http.ResponseWriter, r *http.Request, userID string) {
	// Get user ID
	requesterAccountID, err := key.GetAccountID(r.Context())
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

	// Send follow request
	arg := &model.RequestFollowParams{
		RequesterAccountID: requesterAccountID,
		RequestedUserID:    userID,
	}
	if err := h.svc.RequestFollow(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Accept follow request
// (POST /users/{user_id}/follow-request/accept)
func (h *FollowHandler) AcceptFollowRequest(w http.ResponseWriter, r *http.Request, userID string) {
	// Get user ID
	requestedAccountID, err := key.GetAccountID(r.Context())
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

	// Accept follow request
	arg := &model.AcceptFollowRequestParams{
		RequestedAccountID: requestedAccountID,
		RequesterUserID:    userID,
	}
	if err := h.svc.AcceptFollowRequest(r.Context(), arg); err != nil {
		utils.RespondError(w, err)
		return
	}

	utils.Respond(w, nil)
}

// Reject follow request
// (POST /users/{user_id}/follow-request/reject)
func (h *FollowHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request, userID string) {
	// Get user ID
	requestedAccountID, err := key.GetAccountID(r.Context())
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

	// Reject follow request
	arg := &model.RejectFollowRequestParams{
		RequestedAccountID: requestedAccountID,
		RequesterUserID:    userID,
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
		utils.RespondError(w, &utils.AppError{
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