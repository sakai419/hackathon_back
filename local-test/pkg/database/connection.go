package database

import (
	"database/sql"
	"fmt"
	"local-test/configs"
	"local-test/pkg/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func generateConnStr(c config.DBConfig) string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&timeout=%ds&readTimeout=%ds&writeTimeout=%ds", c.User, c.Pwd, c.Host, c.Database, c.Charset, c.Timeout, c.ReadTimeout, c.WriteTimeout)
	default:
		return ""
	}
}

func ConnectToDB(c config.DBConfig) (*sql.DB, error) {
	// Validate database configuration
	if err := utils.ValidateDBConfig(&c); err != nil {
		return nil, fmt.Errorf("database: invalid config: %w", err)
	}

	// Generate connection string
    connStr := generateConnStr(c)

	// Open database connection
    db, err := sql.Open(c.Driver, connStr)
	if err != nil {
		return nil, fmt.Errorf("database: failed to open db: %w", err)
	}

	// Set database connection pool settings
	db.SetMaxOpenConns(c.MaxOpenConns)
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(c.ConnMaxIdleTime) * time.Minute)

	// Check if the database is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database: failed to ping db: %w", err)
	}

	// Check user permissions
	if err := utils.ValidateDB(db, c); err != nil {
		return nil, fmt.Errorf("database: invalid user permissions: %w", err)
	}

	return db, nil
}