package report

import (
	"errors"
	"local-test/internal/model"
	"local-test/internal/service"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
	"net/http"
)

type ReportHandler struct {
	svc *service.Service
}

func NewReportHandler(svc *service.Service) ServerInterface {
	return &ReportHandler{
		svc: svc,
	}
}

// Create a new report
// (POST /reports/{user_id})
func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request, _ string) {
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

    // Decode request
    var req CreateReportJSONRequestBody
    if err := utils.Decode(r, &req); err != nil {
        utils.RespondError(w, apperrors.NewDecodeError(err))
        return
    }

    // Create report
    if err := h.svc.CreateReport(r.Context(), &model.CreateReportParams{
		ReporterAccountID: clientAccountID,
		ReportedAccountID: targetAccountID,
		Reason:  model.ReportReason(req.Reason),
		Content: req.Content,
	}); err != nil {
        utils.RespondError(w, apperrors.NewHandlerError("create report", err))
        return
    }

	utils.Respond(w, nil)
}

// ErrorHandlerFunc is the error handler for the report handler
func ErrorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	var invalidParamFormatError *InvalidParamFormatError
	if errors.As(err, &invalidParamFormatError) {
		utils.RespondError(w, apperrors.NewInvalidParamFormatError(
			invalidParamFormatError.ParamName,
			invalidParamFormatError.Err,
		))
		return
	}

	utils.RespondError(w, apperrors.NewUnexpectedError(err))
}