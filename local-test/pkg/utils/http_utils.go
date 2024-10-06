package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type AppError struct {
    Status  int
    Code    string
    Message string
    Err     error
}

type errorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func (e *AppError) Error() string {
    return e.Err.Error()
}

func Decode(r *http.Request, v interface{}) error {
    // Check if the request body is empty
    if r.Body == nil {
        return &AppError{
            Status:  http.StatusBadRequest,
            Code:    "EMPTY_BODY",
            Message: "Request body is empty",
            Err:     errors.New("empty request body"),
        }
    }

    // Decode the request body
    if err := json.NewDecoder(r.Body).Decode(v); err != nil {
        return &AppError{
            Status:  http.StatusBadRequest,
            Code:    "INVALID_REQUEST",
            Message: "Failed to decode request",
            Err:     err,
        }
    }
    return nil
}

func Respond(w http.ResponseWriter, data interface{}, statusCode ...int) {
    w.Header().Set("Content-Type", "application/json")

    code := http.StatusOK
    if len(statusCode) > 0 {
        code = statusCode[0]
    } else if data == nil {
        code = http.StatusNoContent
    }

    w.WriteHeader(code)

    if data != nil {
        if err := json.NewEncoder(w).Encode(data); err != nil {
            log.Printf("Failed to encode response: %v", err)
            w.Write([]byte(`{"error":"Internal Server Error"}`))
        }
    }
}

func RespondError(w http.ResponseWriter, err error) {
    var appErr *AppError
    if errors.As(err, &appErr) {
        w.WriteHeader(appErr.Status)
        if encodeErr := json.NewEncoder(w).Encode(errorResponse{
            Code:    appErr.Code,
            Message: appErr.Message,
        }); encodeErr != nil {
            log.Printf("Failed to encode error response: %v", encodeErr)
        }
        log.Printf("Error: %v", appErr.Err)
    } else {
        w.WriteHeader(http.StatusInternalServerError)
        if encodeErr := json.NewEncoder(w).Encode(errorResponse{
            Code:    "INTERNAL_SERVER_ERROR",
            Message: "An unexpected error occurred",
        }); encodeErr != nil {
            log.Printf("Failed to encode error response: %v", encodeErr)
        }
        log.Printf("Unexpected error: %v", err)
    }
}