package report

import (
	"database/sql"
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
        utils.RespondError(w, &apperrors.AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Failed to decode request",
            Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "decode request",
					Err: err,
				},
			),
        })
        return
    }

    // Validate request
    if err := req.validate(); err != nil {
        utils.RespondError(w, &apperrors.AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Invalid request",
            Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "validate request",
					Err: err,
				},
			),
        })
    }

    // Create report
    if err := h.svc.CreateReport(r.Context(), &model.CreateReportParams{
		ReporterAccountID: clientAccountID,
		ReportedAccountID: targetAccountID,
		Reason:  model.ReportReason(req.Reason),
		Content: sql.NullString{String: *req.Content, Valid: req.Content != nil && *req.Content != ""},
	}); err != nil {
        utils.RespondError(w, apperrors.WrapHandlerError(
			&apperrors.ErrOperationFailed{
				Operation: "create report",
				Err: err,
			},
		))
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

func (r *CreateReportJSONRequestBody) validate() error {
	if r.Reason == "" {
		return &InvalidParamFormatError{
			ParamName: "reason",
			Err:   errors.New("reason is required"),
		}
	}
	return nil
}