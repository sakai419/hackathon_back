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
// (POST /reports/{user_id})
func (h *ReportHandler) CreateReport(w http.ResponseWriter, r *http.Request, _ string) {
	// Check if the user is suspended
	if isClientSuspended(w, r) {
		return
	}

	// Get client account ID
	clientAccountID, ok := getClientAccountID(w, r)
	if !ok {
		return
	}

	// Get account ID from path
	accountIDFromPath, ok := getAccountIDFromPath(w, r)
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
    arg := req.toParams(clientAccountID, accountIDFromPath)
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

func isClientSuspended(w http.ResponseWriter, r *http.Request) bool {
	isClientSuspended, err := key.GetIsClientSuspended(r.Context())
	if err != nil {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusInternalServerError,
			Code:    "INTERNAL_SERVER_ERROR",
			Message: "Failed to get is_suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrOperationFailed{
					Operation: "get is_suspended",
					Err: err,
				},
			),
		})
		return true
	}

	if isClientSuspended {
		utils.RespondError(w, &apperrors.AppError{
			Status:  http.StatusForbidden,
			Code:    "FORBIDDEN",
			Message: "User is suspended",
			Err:     apperrors.WrapHandlerError(
				&apperrors.ErrForbidden{
					Message: "User is suspended",
				},
			),
		})
		return true
	}

	return false
}

func getClientAccountID(w http.ResponseWriter, r *http.Request) (string, bool) {
	clientID, err := key.GetClientAccountID(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusInternalServerError,
				Code:    "INTERNAL_SERVER_ERROR",
				Message: "Account ID not found in context",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return clientID, true
}

func getAccountIDFromPath(w http.ResponseWriter, r *http.Request) (string, bool) {
	accountID, err := key.GetAccountIDFromPath(r.Context())
	if err != nil {
		utils.RespondError(w,
			&apperrors.AppError{
				Status:  http.StatusBadRequest,
				Code:    "BAD_REQUEST",
				Message: "Account ID not found in path",
				Err:     apperrors.WrapHandlerError(
					&apperrors.ErrOperationFailed{
						Operation: "get account ID",
						Err: err,
					},
				),
			},
		)
		return "", false
	}
	return accountID, true
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

func (r *CreateReportJSONRequestBody) toParams(reporterAccountID, reportedAccountID string) *model.CreateReportParams {
	return &model.CreateReportParams{
		ReporterAccountID: reporterAccountID,
		ReportedAccountID: reportedAccountID,
		Reason:  model.ReportReason(r.Reason),
		Content: sql.NullString{String: *r.Content, Valid: r.Content != nil && *r.Content != ""},
	}
}