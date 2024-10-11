package report

import (
	"database/sql"
	"errors"
	contextKey "local-test/internal/context"
	"local-test/internal/models"
	"local-test/internal/services"
	"local-test/pkg/utils"
	"net/http"
)

type ReportHandler struct {
	svc *services.Service
}

func NewReportHandler(svc *services.Service) ServerInterface {
	return &ReportHandler{
		svc: svc,
	}
}

// Create a new report
// (POST /reports)
func (h *ReportHandler) CreateReportByUserId(w http.ResponseWriter, r *http.Request, reportedUserID string) {
	// Get Reporter Account ID
	reporterAccountID, err := contextKey.GetUserID(r.Context())
	if err != nil {
		utils.RespondError(w, &utils.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Reporter ID not found in context",
			Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "get reporter ID", Err: err}),
		})
		return
	}

    // Decode request
    var req CreateReportByUserIdJSONRequestBody
    if err := utils.Decode(r, &req); err != nil {
        utils.RespondError(w, &utils.AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Failed to decode request",
            Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "decode request", Err: err}),
        })
        return
    }

    // Validate request
    if err := req.Validate(); err != nil {
        utils.RespondError(w, &utils.AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Invalid request",
            Err:     utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "validate request", Err: err}),
        })
    }

    // Create report
    arg := req.ToParams()
    arg.ReporterAccountID = reporterAccountID
    arg.ReportedUserID = reportedUserID
    if err := h.svc.CreateReportByUserID(r.Context(), arg); err != nil {
        utils.RespondError(w, utils.WrapHandlerError(&utils.ErrOperationFailed{Operation: "create report", Err: err}))
        return
    }

	utils.Respond(w, nil)
}

func ErrHandleFunc(w http.ResponseWriter, r *http.Request, err error) {
    var invalidParamFormatError *InvalidParamFormatError
    if errors.As(err, &invalidParamFormatError) {
        utils.RespondError(w, &utils.AppError{
            Status:  http.StatusBadRequest,
            Code:    "BAD_REQUEST",
            Message: "Invalid parameter format",
            Err:     err,
        })
        return
    } else {
        utils.RespondError(w, err)
    }
}

func (r *CreateReportByUserIdJSONRequestBody) Validate() error {
    if r.Reason == "" {
        return errors.New("reason is required")
    }
    return nil
}

func (r *CreateReportByUserIdJSONRequestBody) ToParams() *models.CreateReportByUserIDParams {
    if r.Content == nil {
        return &models.CreateReportByUserIDParams{
            Reason: models.ReportReason(r.Reason),
        }
    }

    return &models.CreateReportByUserIDParams{
        Reason: models.ReportReason(r.Reason),
        Content: sql.NullString{String: *r.Content, Valid: r.Content != nil},
    }
}