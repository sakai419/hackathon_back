package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type Config struct {
    FirebaseAuthClient *auth.Client
}

func InitFirebaseClient() (*auth.Client, error) {
	opt := option.WithCredentialsFile("firebase.json")
	firebaseApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}

	return authClient, nil
}

func LoadConfig() (*Config, error) {
	authClient, err := InitFirebaseClient()
	if err != nil {
		return nil, err
	}

	return &Config{
		FirebaseAuthClient: authClient,
	}, nil
}