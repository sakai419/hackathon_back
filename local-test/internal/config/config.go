package config

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

type Config struct {
	FirebaseAuthClient *auth.Client
	DBConfig DBConfig
}

type DBConfig struct {
	Type string
	User string
	Pwd string
	Host string
	Database string
}

func initFirebaseClient(v *viper.Viper) (*auth.Client, error) {
	// Create a map of Firebase credentials
	firebaseCredentials := map[string]string{
        "type":                        v.GetString("firebase.type"),
        "project_id":                  v.GetString("firebase.project_id"),
        "private_key_id":              v.GetString("firebase.private_key_id"),
        "private_key":                 v.GetString("firebase.private_key"),
        "client_email":                v.GetString("firebase.client_email"),
        "client_id":                   v.GetString("firebase.client_id"),
        "auth_uri":                    v.GetString("firebase.auth_uri"),
        "token_uri":                   v.GetString("firebase.token_uri"),
        "auth_provider_x509_cert_url": v.GetString("firebase.auth_provider_x509_cert_url"),
        "client_x509_cert_url":        v.GetString("firebase.client_x509_cert_url"),
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

func LoadConfig() (*Config, error) {
    // Load environment variables from .env file
    if err := godotenv.Load(".env"); err != nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }

	// Initialize Viper
	v := viper.New()
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Load config file
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs/")
	if err := v.MergeInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	// Initialize Firebase Auth client
	authClient, err := initFirebaseClient(v)
	if err != nil {
		return nil, err
	}

	return &Config{
		FirebaseAuthClient: authClient,
		DBConfig: DBConfig{
			Type:     v.GetString("db.type"),
			User:     v.GetString("db.user"),
			Pwd:      v.GetString("db.password"),
			Host:     v.GetString("db.host"),
			Database: v.GetString("db.database"),
		},
	}, nil
}