package search

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type SearchHandler struct {
	svc *service.Service
}

func NewSearchHandler(svc *service.Service) ServerInterface {
	return &SearchHandler{
		svc: svc,
	}
}

// Search for users
// (GET /search/users)
func (h *SearchHandler) SearchUsers(w http.ResponseWriter, r *http.Request, params SearchUsersParams) {
	// get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// search users
	userInfos, err := h.svc.SearchUsers(r.Context(), &model.SearchUsersParams{
		ClientAccountID: clientAccountID,
		SortType:        model.SortType(params.SortType),
		Keyword:         params.Keyword,
		Offset:          params.Offset,
		Limit:           params.Limit,
	})
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("Failed to search users", err))
		return
	}

	// convert to response
	resp := make(UserInfos, len(userInfos))
	for i, u := range userInfos {
		resp[i] = UserInfo{
			UserId:          u.UserID,
			UserName:        u.UserName,
			Bio:             u.Bio,
			ProfileImageUrl: u.ProfileImageURL,
			IsPrivate:       u.IsPrivate,
			IsAdmin:         u.IsAdmin,
		}
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