package block

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type BlockHandler struct {
	svc *service.Service
}

func NewBlockHandler(svc *service.Service) ServerInterface {
	return &BlockHandler{
		svc: svc,
	}
}


// Block a user
// (POST /blocks/{user_id})
func (h *BlockHandler) BlockUser(w http.ResponseWriter, r *http.Request, _ string) {
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

	// Block user
	if err := h.svc.BlockUser(r.Context(), &model.BlockUserParams{
		BlockerAccountID: clientAccountID,
		BlockedAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("block user", err))
		return
	}

	utils.Respond(w, nil)
}

// Unblock a user
// (DELETE /blocks/{user_id})
func (h *BlockHandler) UnblockUser(w http.ResponseWriter, r *http.Request, _ string) {
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

	// Unblock user
	if err := h.svc.UnblockUser(r.Context(), &model.UnblockUserParams{
		BlockerAccountID: clientAccountID,
		BlockedAccountID: targetAccountID,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("unblock user", err))
		return
	}

	utils.Respond(w, nil)
}

// Get blocked users
// (GET /blocks)
func (h *BlockHandler) GetBlockedInfos(w http.ResponseWriter, r *http.Request, params GetBlockedInfosParams) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get blocked users
	blockedInfos, err := h.svc.GetBlockedInfos(r.Context(), &model.GetBlockedInfosParams{
		BlockerAccountID: clientAccountID,
		Limit:            params.Limit,
		Offset:           params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get blocked users", err))
		return
	}

	resp := convertToUserInfos(blockedInfos)

	utils.Respond(w, resp)
}

// Get block count
// (GET /blocks/count)
func (h *BlockHandler) GetBlockCount(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get block count
	count, err := h.svc.GetBlockCount(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get block count", err))
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

func convertToUserInfos(blockedInfos []*model.UserInfo) []UserInfo {
	var resp []UserInfo
	for _, blockedInfo := range blockedInfos {
		resp = append(resp, UserInfo{
			UserId:          blockedInfo.UserID,
			UserName:        blockedInfo.UserName,
			Bio:	         blockedInfo.Bio,
			ProfileImageUrl: blockedInfo.ProfileImageURL,
			IsPrivate:       blockedInfo.IsPrivate,
			IsAdmin:         blockedInfo.IsAdmin,
			IsFollowing:     blockedInfo.IsFollowing,
			IsFollowed:      blockedInfo.IsFollowed,
		})
	}
	return resp
}