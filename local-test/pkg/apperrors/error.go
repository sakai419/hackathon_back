package apperrors

import (
	"fmt"
)

// ErrOperationFailed is a custom error type that holds the operation name
type ErrOperationFailed struct {
    Operation string
    Err       error
}

func (e *ErrOperationFailed) Error() string {
    return fmt.Sprintf("operation %s failed: %v", e.Operation, e.Err)
}

func (e *ErrOperationFailed) Unwrap() error {
    return e.Err
}

// ErrRecordNotFound is a custom error type that indicates a record was not found
type ErrRecordNotFound struct {
    Condition string
}

func (e *ErrRecordNotFound) Error() string {
    return fmt.Sprintf("record not found: there is no record matching your request(%s)", e.Condition)
}

// ErrInvalidInput is a custom error type that indicates an invalid input
type ErrInvalidInput struct {
	Message string
}

func (e *ErrInvalidInput) Error() string {
	return fmt.Sprintf("invalid input: %s", e.Message)
}

// ErrInvalidRequest is a custom error type that indicates an invalid request
type ErrInvalidRequest struct {
	Entity string
	Err error
}

func (e *ErrInvalidRequest) Error() string {
	return fmt.Sprintf("invalid request: %s is required: %v", e.Entity, e.Err)
}

// ErrDuplicateEntry is a custom error type that indicates a duplicate entry
type ErrDuplicateEntry struct {
	Entity string
	Err error
}

func (e *ErrDuplicateEntry) Error() string {
	return fmt.Sprintf("duplicate entry: %s: %v", e.Entity, e.Err)
}

func (e *ErrDuplicateEntry) Unwrap() error {
	return e.Err
}

// ErrForbidden is a custom error type that indicates a forbidden operation
type ErrForbidden struct {
	Message string
}

func (e *ErrForbidden) Error() string {
	return fmt.Sprintf("forbidden: %s", e.Message)
}

// ErrEmptyRequest is a custom error type that indicates an empty result
type ErrEmptyRequest struct {
	Message string
}

func (e *ErrEmptyRequest) Error() string {
	return fmt.Sprintf("empty request: %s", e.Message)
}

// ErrJSONMarshal is a custom error type that indicates a JSON marshal error
type ErrJSONMarshal struct {
}


type AppError struct {
    Status  int
    Code    string
    Message string
    Err     error
}

func (e *AppError) Error() string {
    return e.Err.Error()
}

func WrapRepositoryError(err error) error {
	return fmt.Errorf("repository: %w", err)
}

func WrapServiceError(err error) error {
	return fmt.Errorf("service: %w", err)
}

func WrapHandlerError(err error) error {
	return fmt.Errorf("handler: %w", err)
}

func WrapValidationError(err error) error {
	return fmt.Errorf("validation: %w", err)
}

func WrapConfigError(err error) error {
	return fmt.Errorf("config: %w", err)
}

func WrapInitError(err error) error {
	return fmt.Errorf("init: %w", err)
}