package sidebar

import (
	"local-test/internal/model"
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

// Get left sidebar info
// (GET /sidebar/left)
func (h *SidebarHandler) GetLeftSidebarInfo(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get sidebar info
	sidebarInfo, err := h.svc.GetLeftSidebarInfo(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get sidebar info", err))
		return
	}

	utils.Respond(w, LeftSidebarInfo{
		UnreadConversationCount: sidebarInfo.UnreadConversationCount,
		UnreadNotificationCount: sidebarInfo.UnreadNotificationCount,
	})
}

// Get right sidebar info
// (GET /sidebar/right)
func (h *SidebarHandler) GetRightSidebarInfo(w http.ResponseWriter, r *http.Request) {
	// Get client account ID
	clientAccountID, ok := utils.GetClientAccountID(w, r)
	if !ok {
		return
	}

	// Get sidebar info
	sidebarInfo, err := h.svc.GetRightSidebarInfo(r.Context(), clientAccountID)
	if err != nil {
		utils.RespondError(w, apperrors.NewHandlerError("get sidebar info", err))
		return
	}

	utils.Respond(w, RightSidebarInfo{
		RecentLabels: 	   convertToLabelCounts(sidebarInfo.RecentLabels),
		FollowSuggestions: convertToUserInfoWithoutBios(sidebarInfo.FollowSuggestions),
	})
}

func convertToLabelCounts(labelCounts []*model.LabelCount) []LabelCount {
	ret := make([]LabelCount, 0, len(labelCounts))
	for _, labelCount := range labelCounts {
		ret = append(ret, LabelCount{
			Label: string(labelCount.Label),
			Count: labelCount.Count,
		})
	}
	return ret
}

func convertToUserInfoWithoutBios(userInfos []*model.UserInfoWithoutBio) []UserInfoWithoutBio {
	ret := make([]UserInfoWithoutBio, 0, len(userInfos))
	for _, userInfo := range userInfos {
		ret = append(ret, UserInfoWithoutBio{
			UserId:          userInfo.UserID,
			UserName:        userInfo.UserName,
			ProfileImageUrl: userInfo.ProfileImageURL,
			IsPrivate:       userInfo.IsPrivate,
			IsAdmin:         userInfo.IsAdmin,
			IsFollowing:     userInfo.IsFollowing,
			IsFollowed:      userInfo.IsFollowed,
			IsPending:       userInfo.IsPending,
		})
	}
	return ret
}