package sidebar

import (
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type SidebarHandler struct {
	svc *service.Service
}

func NewSidebarHandler(svc *service.Service) ServerInterface {
	return &SidebarHandler{
		svc: svc,
	}
}

// Get sidebar info
// (GET /sidebar)
func (h *SidebarHandler) GetSidebarInfo(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get sidebar info
	sidebarInfo, err := h.svc.GetSidebarInfo(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get sidebar info", err))
		return
	}

	utils.Respond(w, SidebarInfo{
		UnreadConversationCount: sidebarInfo.UnreadConversationCount,
		UnreadNotificationCount: sidebarInfo.UnreadNotificationCount,
	})
}