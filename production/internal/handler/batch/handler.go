package batch

import (
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"log"
	"net/http"
)

type BatchHandler struct {
	svc *service.Service
}

func NewBatchHandler(svc *service.Service) ServerInterface {
	return &BatchHandler{
		svc: svc,
	}
}

// Update user interests scores
// (POST /batch/users/interests/update)
func (h *BatchHandler) UpdateUserInterests(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating user interests scores")
	// Check if user is authenticated
	if ok := utils.IsClientAdmin(w, r); !ok {
		return
	}

	log.Println("Updating user interests scores")

	// Update user interests scores
	if err := h.svc.UpdateUserInterests(r.Context()); err != nil {
		log.Println("Error updating user interests scores")
		utils.RespondError(w, apperrors.NewHandlerError("update user interests scores", err))
		return
	}

	log.Println("User interests scores updated")

	utils.Respond(w, nil)
}
