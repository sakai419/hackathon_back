package setting

import (
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type SettingHandler struct {
	svc *service.Service
}

func NewSettingHandler(svc *service.Service) ServerInterface {
	return &SettingHandler{
		svc: svc,
	}
}

// Update settings
// (PATCH /settings)
func (h *SettingHandler) UpdateSettings(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Decode request
	var req UpdateSettingsJSONRequestBody
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, apperrors.NewDecodeError(err))
		return
	}

	// Update settings
	if err := h.svc.UpdateSettings(r.Context(), &model.UpdateSettingsParams{
		AccountID:       clientAccountID,
		IsPrivate:       req.IsPrivate,
	}); err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("update settings", err))
	}

	utils.Respond(w, nil)
}