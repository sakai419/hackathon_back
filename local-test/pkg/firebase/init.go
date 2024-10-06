package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"local-test/configs"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitFirebaseClient(c config.FirebaseConfig) (*auth.Client, error) {
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
        return nil, fmt.Errorf("error marshaling firebase credentials: %v", err)
    }

	// Initialize Firebase app
	opt := option.WithCredentialsJSON(credentialsJSON)
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	// Initialize Firebase Auth client
	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return authClient, nil
}