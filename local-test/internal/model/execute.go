package model

import "local-test/pkg/apperrors"

type ExecuteResult struct {
	Status  string
	Output  *string
	Message *string
}

type ExecuteCodeParams struct {
	Code Code
}

func (p *ExecuteCodeParams) Validate() error {
	if p.Code.Content == "" {
		return &apperrors.ErrInvalidInput{
			Message: "Code is required",
		}
	}

	switch p.Code.Language {
	case "c", "cpp", "java", "python":
		break
	default:
		return &apperrors.ErrInvalidInput{
			Message: "Invalid language",
		}
	}

	return nil
}