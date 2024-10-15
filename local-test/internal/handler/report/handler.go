package report

import (
	"database/sql"
	"errors"
	"local-test/internal/key"
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
// (POST /reports)
func (h *ReportHandler) CreateReportByUserId(w http.ResponseWriter, r *http.Request, reportedUserID string) {
	// Get Reporter Account ID
	reporterAccountID, err := key.GetAccountID(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusUnauthorized,
			Code:    "UNAUTHORIZED",
			Message: "Account ID not found in context",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get account ID",
					Err: err,
				},
			),
		})
		return
	}

    // Decode request
    var req CreateReportByUserIdJSONRequestBody
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
    arg := req.toParams()
    arg.ReporterAccountID = reporterAccountID
    arg.ReportedUserID = reportedUserID
    if err := h.svc.CreateReportByUserID(r.Context(), arg); err != nil {
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

func ErrHandleFunc(w http.ResponseWriter, r *http.Request, err error) {
    var invalidParamFormatError *InvalidParamFormatError
    if errors.As(err, &invalidParamFormatError) {
        utils.RespondError(w, &apperrors.AppError{
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

func (r *CreateReportByUserIdJSONRequestBody) validate() error {
    if r.Reason == "" {
        return errors.New("reason is required")
    }
    return nil
}

func (r *CreateReportByUserIdJSONRequestBody) toParams() *model.CreateReportByUserIDParams {
    if r.Content == nil {
        return &model.CreateReportByUserIDParams{
            Reason: model.ReportReason(r.Reason),
        }
    }

    return &model.CreateReportByUserIDParams{
        Reason: model.ReportReason(r.Reason),
        Content: sql.NullString{String: *r.Content, Valid: r.Content != nil},
    }
}