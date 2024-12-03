package database

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/apperrors"
	"time"

	_ "github.com/lib/pq"
)

func generateConnStr(c *config.DBConfig) string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", c.User, c.Pwd, c.Host, c.Port, c.Database, c.Charset, c.Timeout, c.ReadTimeout, c.WriteTimeout)
	case "postgres":
		return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Pwd, c.Database, c.SSLMode)
	default:
		return ""
	}
}

func ConnectToDB(c *config.DBConfig) (*sql.DB, error) {
	// Validate database configuration
	if err := validateDBConfig(c); err != nil {
		return nil, apperrors.WrapInitError(
			&apperrors.ErrOperationFailed{
				Operation: "validate db config",
				Err:       err,
			},
		)
	}

	// Generate connection string
    connStr := generateConnStr(c)

	// Open database connection
    db, err := sql.Open(c.Driver, connStr)
	if err != nil {
		return nil, apperrors.WrapInitError(
			&apperrors.ErrOperationFailed{
				Operation: "open db connection",
				Err:       err,
			},
		)
	}

	// Set database connection pool settings
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Minute)


	// Check if the database is alive
	if err := db.Ping(); err != nil {
		return nil, apperrors.WrapInitError(
			&apperrors.ErrOperationFailed{
				Operation: "ping db",
				Err:       err,
			},
		)
	}


	// Check user permissions
	if err := validateDB(db, c); err != nil {
		return nil, apperrors.WrapInitError(
			&apperrors.ErrOperationFailed{
				Operation: "validate db",
				Err:       err,
			},
		)
	}

	return db, nil
}