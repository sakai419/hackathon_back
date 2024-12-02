package firebase

import (
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"
)

func validateFirebaseConfig(c *config.FirebaseConfig) error {
	if c == nil {
		return apperrors.WrapConfigError(
			&apperrors.ErrInvalidInput{
				Message: "firebase config is nil",
			},
		)
	}

	fields := map[string]interface{}{
		"type":                        c.Type,
		"project_id":                  c.ProjectID,
		"private_key_id":              c.PrivateKeyID,
		"private_key":                 c.PrivateKey,
		"client_email":                c.ClientEmail,
		"client_id":                   c.ClientID,
		"auth_uri":                    c.AuthURI,
		"token_uri":                   c.TokenURI,
		"auth_provider_x509_cert_url": c.AuthProviderX509CertURL,
		"client_x509_cert_url":        c.ClientX509CertURL,
	}

	// Validate each field
	for fieldName, fieldValue := range fields {
		if err := utils.ValidateField("firebase", fieldName, fieldValue); err != nil {
			return apperrors.WrapValidationError(
				&apperrors.ErrOperationFailed{
					Operation: fmt.Sprintf("validate %s", fieldName),
					Err: err,
				},
			)
		}
	}

	return nil
}