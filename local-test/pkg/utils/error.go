package utils

import "fmt"

func WrapRepositoryError(err error, message string) error {
	return fmt.Errorf("repository: %s: %w", message, err)
}

func WrapSerivceError(err error, message string) error {
	return fmt.Errorf("service: %s: %w", message, err)
}

func WrapHandlerError(err error, message string) error {
	return fmt.Errorf("handler: %s: %w", message, err)
}