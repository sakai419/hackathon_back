package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"local-test/pkg/apperrors"
	"log"
	"net/http"
	"reflect"
)

type errorResponse struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

func Decode(r *http.Request, v interface{}) error {
    // Check if the request body is empty
    if r.Body == nil {
        return &apperrors.ErrEmptyRequest{Message: "request body is empty"}
    }

    // Decode the request body
    if err := json.NewDecoder(r.Body).Decode(v); err != nil {
        return &apperrors.ErrInvalidInput{Message: "failed to decode request body"}
    }

	// Validate the request body
	if err := validateRequiredFields(v); err != nil {
		return &apperrors.ErrInvalidInput{Message: err.Error()}
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
    var appErr *apperrors.AppError
    if errors.As(err, &appErr) {
        w.WriteHeader(appErr.Status)
        if encodeErr := json.NewEncoder(w).Encode(errorResponse{
            Code:    appErr.Code,
            Message: appErr.Message,
        }); encodeErr != nil {
            log.Printf("Failed to encode error response: %v", encodeErr)
        }
        log.Printf("Error: %v", err)
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

func validateRequiredFields(req interface{}) error {
    v := reflect.ValueOf(req)

	// if the request is a pointer, get the value it points to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

    if v.Kind() != reflect.Struct {
        return &apperrors.ErrInvalidRequest{
            Entity: "All fields",
            Err:    errors.New("provided value is not a struct"),
        }
    }

    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
        fieldType := v.Type().Field(i)

        // Check if the field is a zero value
        if fieldType.Type.Kind() != reflect.Ptr && field.IsZero() {
            return &apperrors.ErrInvalidRequest{
                Entity: fieldType.Name,
                Err:    fmt.Errorf("%s is a zero value", fieldType.Name),
            }
        }
    }
    return nil
}