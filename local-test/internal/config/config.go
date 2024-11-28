package config

import (
	"local-test/pkg/apperrors"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBConfig *DBConfig
	FirebaseConfig *FirebaseConfig
	ServerConfig *ServerConfig
	GeminiConfig *GeminiConfig
}

type DBConfig struct {
	Driver string
	User string
	Pwd string
	Host string
	Port int
	SSLMode string
	Database string
	Charset string
	MaxOpenConns int
	MaxIdleConns int
	ConnMaxLifetime int
	ConnMaxIdleTime int
	Timeout int
	ReadTimeout int
	WriteTimeout int
	RequiredTables []string
}

type FirebaseConfig struct {
	Type string
	ProjectID string
	PrivateKeyID string
	PrivateKey string
	ClientEmail string
	ClientID string
	AuthURI string
	TokenURI string
	AuthProviderX509CertURL string
	ClientX509CertURL string
}

type GeminiConfig struct {
	APIKey string
}

type ServerConfig struct {
	Port int
}

func generateDBConfig(v *viper.Viper) (*DBConfig, error) {
	return &DBConfig{
		Driver:   v.GetString("db.driver"),
		User:     v.GetString("db.user"),
		Pwd:      v.GetString("db.password"),
		Host:     v.GetString("db.host"),
		Port:     v.GetInt("db.port"),
		SSLMode:  v.GetString("db.ssl_mode"),
		Database: v.GetString("db.database"),
		Charset:  v.GetString("db.charset"),
		MaxOpenConns: v.GetInt("db.max_open_conns"),
		MaxIdleConns: v.GetInt("db.max_idle_conns"),
		ConnMaxLifetime: v.GetInt("db.conn_max_lifetime"),
		ConnMaxIdleTime: v.GetInt("db.conn_max_idle_time"),
		Timeout: v.GetInt("db.timeout"),
		ReadTimeout: v.GetInt("db.read_timeout"),
		WriteTimeout: v.GetInt("db.write_timeout"),
		RequiredTables: v.GetStringSlice("db.required_tables"),
	}, nil
}

func generateFirebaseConfig(v *viper.Viper) (*FirebaseConfig, error) {
	return &FirebaseConfig{
		Type: v.GetString("firebase.type"),
		ProjectID: v.GetString("firebase.project_id"),
		PrivateKeyID: v.GetString("firebase.private_key_id"),
		PrivateKey: v.GetString("firebase.private_key"),
		ClientEmail: v.GetString("firebase.client_email"),
		ClientID: v.GetString("firebase.client_id"),
		AuthURI: v.GetString("firebase.auth_uri"),
		TokenURI: v.GetString("firebase.token_uri"),
		AuthProviderX509CertURL: v.GetString("firebase.auth_provider_x509_cert_url"),
		ClientX509CertURL: v.GetString("firebase.client_x509_cert_url"),
	}, nil
}

func generateServerConfig(v *viper.Viper) (*ServerConfig, error) {
	return &ServerConfig{
		Port: v.GetInt("server.port"),
	}, nil
}

func generateGeminiConfig(v *viper.Viper) (*GeminiConfig, error) {
	return &GeminiConfig{
		APIKey: v.GetString("gemini.api_key"),
	}, nil
}

func LoadConfig() (*Config, error) {
    // Load environment variables from .env file
    if err := godotenv.Load(".env"); err != nil {
        return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "load .env file",
				Err: err,
			},
		)
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
		return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "load config file",
				Err: err,
			},
		)
	}

	// Generate DB config
	DBConfig, err := generateDBConfig(v)
	if err != nil {
		return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "generate db config",
				Err: err,
			},
		)
	}

	// Generate Firebase config
	FirebaseConfig, err := generateFirebaseConfig(v)
	if err != nil {
		return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "generate firebase config",
				Err: err,
			},
		)
	}

	// Generate Server config
	ServerConfig, err := generateServerConfig(v)
	if err != nil {
		return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "generate server config",
				Err: err,
			},
		)
	}

	// Generate Gemini config
	GeminiConfig, err := generateGeminiConfig(v)
	if err != nil {
		return nil, apperrors.WrapConfigError(
			&apperrors.ErrOperationFailed{
				Operation: "generate gemini config",
				Err: err,
			},
		)
	}

	return &Config{
		FirebaseConfig: FirebaseConfig,
		DBConfig:       DBConfig,
		ServerConfig:   ServerConfig,
		GeminiConfig:   GeminiConfig,
	}, nil
}