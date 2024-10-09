package utils

import "fmt"

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

// ErrInvalidRequest is a custom error type that indicates an invalid request
type ErrInvalidRequest struct {
	Message string
	Err     error
}

func (e *ErrInvalidRequest) Error() string {
	return fmt.Sprintf("invalid request: %s: %v", e.Message, e.Err)
}

func (e *ErrInvalidRequest) Unwrap() error {
	return e.Err
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

type AppError struct {
    Status  int
    Code    string
    Message string
    Err     error
}

func WrapRepositoryError(err error) error {
	return fmt.Errorf("repository: %w", err)
}

func WrapSerivceError(err error) error {
	return fmt.Errorf("service: %w", err)
}

func WrapHandlerError(err error) error {
	return fmt.Errorf("handler: %w", err)
}