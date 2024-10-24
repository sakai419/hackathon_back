package vertex

import (
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
)

func validateVertexConfig(c *config.VertexConfig) error {
	if c == nil {
		return apperrors.WrapConfigError(
			&apperrors.ErrInvalidInput{
				Message: "vertex config is nil",
			},
		)
	}

	fields := map[string]interface{}{
		"project_id": c.ProjectID,
		"location":   c.Location,
		"engine_id":  c.EngineID,
		"scope":      c.Scope,
	}

	// Validate each field
	for fieldName, fieldValue := range fields {
		if err := utils.ValidateField("vertex", fieldName, fieldValue); err != nil {
			return apperrors.WrapValidationError(
				&apperrors.ErrOperationFailed{
					Operation: fmt.Sprintf("validate %s", fieldName),
					Err:       err,
				},
			)
		}
	}

	return nil

}