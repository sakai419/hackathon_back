package firebase

import (
	"context"
	"encoding/json"
	"local-test/internal/config"
	"local-test/pkg/apperrors"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitFirebaseClient(c *config.FirebaseConfig) (*auth.Client, error) {
    // Validate Firebase configuration
    if err := validateFirebaseConfig(c); err != nil {
        return nil, apperrors.WrapInitError(
            &apperrors.ErrOperationFailed{
                Operation: "validate firebase config",
                Err:       err,
            },
        )
    }

	// Create a map of Firebase credentials
	firebaseCredentials := map[string]string{
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

	// Marshal the map into JSON
    credentialsJSON, err := json.Marshal(firebaseCredentials)
    if err != nil {
        return nil, apperrors.WrapInitError(
            &apperrors.ErrOperationFailed{
                Operation: "marshal firebase credentials",
                Err:       err,
            },
        )
    }

	// Initialize Firebase app
	opt := option.WithCredentialsJSON(credentialsJSON)
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, apperrors.WrapInitError(
            &apperrors.ErrOperationFailed{
                Operation: "initialize Firebase app",
                Err:       err,
            },
        )
	}

	// Initialize Firebase Auth client
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, apperrors.WrapInitError(
            &apperrors.ErrOperationFailed{
                Operation: "initialize Firebase Auth client",
                Err:       err,
            },
        )
	}

	return authClient, nil
}