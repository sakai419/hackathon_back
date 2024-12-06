package execute

import (
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/utils"
	"net/http"
)

type ExecuteHandler struct {
	svc *service.Service
}

func NewExecuteHandler(svc *service.Service) ServerInterface {
	return &ExecuteHandler{
		svc: svc,
	}
}

// Execute source code
// (POST /execute)
func (h *ExecuteHandler) ExecuteCode(w http.ResponseWriter, r *http.Request) {
	// Check if the client is suspended
	if utils.IsClientSuspended(w, r) {
		return
	}

	// Decode request body
	var req ExecuteRequest
	if err := utils.Decode(r, &req); err != nil {
		utils.RespondError(w, err)
		return
	}

	// Execute code
	result, err := h.svc.ExecuteCode(r.Context(), &model.ExecuteCodeParams{
		Code: model.Code{
			Content:  req.Content,
			Language: req.Language,
		},
	})
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// Respond
	ret := &ExecuteResponse{
		Status: result.Status,
	}
	if result.Message != nil {
		ret.Message = result.Message
	}
	if result.Output != nil {
		ret.Output = result.Output
	}

	utils.Respond(w, ret)
}