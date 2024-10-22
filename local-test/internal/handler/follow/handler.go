package follow

import (
	"errors"
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
// (POST /follows/{user_id})
func (h *FollowHandler) FollowAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) || utils.IsTargetSuspended(w, r) || utils.IsTargetPrivate(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Follow user
	if err := h.svc.FollowAndNotify(r.Context(), &model.FollowAndNotifyParams{
		FollowerAccountID:  clientAccountID,
		FollowingAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("follow user", err))
		return
	}

	utils.Respond(w, nil)
}

// Unfollow a user
// (DELETE /follows/{user_id})
func (h *FollowHandler) Unfollow(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the clident is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Unfollow user
	if err := h.svc.Unfollow(r.Context(), &model.UnfollowParams{
		FollowerAccountID:  clientAccountID,
		FollowingAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("unfollow user", err))
		return
	}

	utils.Respond(w, nil)
}

// Send follow request
// (POST /follows/requests/{user_id})
func (h *FollowHandler) RequestFollowAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) || utils.IsTargetSuspended(w, r) || utils.IsNotTargetPrivate(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Send follow request
	if err := h.svc.RequestFollowAndNotify(r.Context(), &model.RequestFollowAndNotifyParams{
		RequesterAccountID: clientAccountID,
		RequestedAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("send follow request", err))
		return
	}

	utils.Respond(w, nil)
}

// Accept follow request
// (PUT /follows/requests/received/{user_id}/accept)
func (h *FollowHandler) AcceptFollowRequestAndNotify(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) || utils.IsTargetSuspended(w, r) {
		return
	}

	// Ger client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Accept follow request
	if err := h.svc.AcceptFollowRequestAndNotify(r.Context(), &model.AcceptFollowRequestAndNotifyParams{
		RequestedAccountID: clientAccountID,
		RequesterAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("accept follow request", err))
		return
	}

	utils.Respond(w, nil)
}

// Reject follow request
// (DELETE /follows/requests/received/{user_id})
func (h *FollowHandler) RejectFollowRequest(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Reject follow request
	if err := h.svc.RejectFollowRequest(r.Context(), &model.RejectFollowRequestParams{
		RequestedAccountID: clientAccountID,
		RequesterAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("reject follow request", err))
		return
	}

	utils.Respond(w, nil)
}

// Get followers
// (GET /follows/followers/{user_id})
func (h *FollowHandler) GetFollowerInfos(w http.ResponseWriter, r *http.Request, _ string, params GetFollowerInfosParams) {
	// Get tager account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Get followers
	followerInfos, err := h.svc.GetFollowerInfos(r.Context(), &model.GetFollowerInfosParams{
		FollowingAccountID: targetAccountID,
		Limit:              params.Limit,
		Offset:             params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get followers", err))
		return
	}


	// Convert to response
	resp := convertToUserInfos(followerInfos)

	utils.Respond(w, resp)
}

// Get followings
// (GET /follows/following/{user_id})
func (h *FollowHandler) GetFollowingInfos(w http.ResponseWriter, r *http.Request, _ string, params GetFollowingInfosParams) {
	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Get followings
	followingInfos, err := h.svc.GetFollowingInfos(r.Context(), &model.GetFollowingInfosParams{
		FollowerAccountID: targetAccountID,
		Limit:             params.Limit,
		Offset: 		   params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get followings", err))
		return
	}

	// Convert to response
	resp := convertToUserInfos(followingInfos)

	utils.Respond(w, resp)
}

// Get follow counts
// (GET /follows/count/{user_id})
func (h *FollowHandler) GetFollowCounts(w http.ResponseWriter, r *http.Request, _ string) {
	// Get target account ID
	targetAccountID, ok := utils.GetTargetAccountID(w, r)
	if !ok {
		return
	}

	// Get follow counts
	counts, err := h.svc.GetFollowCounts(r.Context(), targetAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get follow counts", err))
		return
	}

	utils.Respond(w, &FollowCounts{
		FollowersCount: counts.FollowersCount,
		FollowingCount: counts.FollowingCount,
	})
}

// Get follow requests count
// (GET /follows/requests/received/count)
func (h *FollowHandler) GetFollowRequestCount(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get follow requests count
	count, err := h.svc.GetFollowRequestsCount(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get follow requests count", err))
		return
	}

	utils.Respond(w, Count{Count: count})
}


// ErrorHandlerFunc is the error handler for the follow handler
func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	var requiredParamError *RequiredParamError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, apperrors.NewInvalidParamFormatError(
			invalidParamFormatError.ParamName,
			invalidParamFormatError.Err,
		))
		return
	} else if errors.As(err, &requiredParamError) {
		utils.RespondError(w, apperrors.NewRequiredParamError(
			requiredParamError.ParamName,
			requiredParamError,
		))
		return
	}

	utils.RespondError(w, apperrors.NewUnexpectedError(err))
}

func convertToUserInfos(followerInfos []*model.UserInfo) []UserInfo {
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