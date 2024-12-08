package admin

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type AdminHandler struct {
	svc *service.Service
}

func NewAdminHandler(svc *service.Service) ServerInterface {
	return &AdminHandler{
		svc: svc,
	}
}

// Get reported user infos order by report count
// (GET /admin/reports/users)
func (h *AdminHandler) GetReportedUsers(w http.ResponseWriter, r *http.Request, params GetReportedUsersParams) {
	// Check if client is admin
	if !utils.IsClientAdmin(w, r) {
		return
	}

	// Get reported user infos order by report count
	reportedUserInfos, err := h.svc.GetReportedUserInfosOrderByReportCount(r.Context(), &model.GetReportedUserInfosOrderByReportCountParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get reported user infos order by report count", err))
		return
	}

	// convert to response
	resp := make([]ReportedUserInfo, 0, len(reportedUserInfos))
	for _, reportedUserInfo := range reportedUserInfos {
		resp = append(resp, ReportedUserInfo{
			UserInfo:    convertToUserInfo(reportedUserInfo.UserInfo),
			ReportCount: reportedUserInfo.ReportCount,
		})
	}

	utils.Respond(w, resp)
}

// ErrorHandlerFunc is the error handler for tweet handlers
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

func convertToUserInfo(info model.UserInfo) UserInfo {
	return UserInfo{
		UserId:          info.UserID,
		UserName:        info.UserName,
		ProfileImageUrl: info.ProfileImageURL,
		Bio:             info.Bio,
		IsPrivate:       info.IsPrivate,
		IsAdmin:         info.IsAdmin,
		IsFollowing:     info.IsFollowing,
		IsFollowed:      info.IsFollowed,
		IsPending: 	     info.IsPending,
	}
}