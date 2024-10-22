package profile

import (
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type ProfileHandler struct {
	svc *service.Service
}

func NewProfileHandler(svc *service.Service) ServerInterface {
	return &ProfileHandler{
		svc: svc,
	}
}

// Update profile
// (PATCH /profiles)
func (h *ProfileHandler) UpdateProfiles(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request
	var req UpdateProfilesJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Update profile
	if err := h.svc.UpdateProfiles(r.Context(), &model.UpdateProfilesParams{
		AccountID:       clientAccountID,
		Bio:             req.Bio,
		ProfileImageURL: req.ProfileImageUrl,
		BannerImageURL:  req.BannerImageUrl,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("update profiles", err))
	}

	utils.Respond(w, nil)
}