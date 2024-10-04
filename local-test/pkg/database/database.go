package database

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

func generateConnStr(c config.DBConfig) string {
	switch c.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Pwd, c.Host, c.Database)
	default:
		return ""
	}
}

func ConnectToDB(c config.DBConfig) (db *sql.DB, err error) {
	// Validate database configuration
	if err := utils.ValidateDBConfig(&c); err != nil {
		return nil, fmt.Errorf("fail: config.ValidateDBConfig, %v", err)
	}

	// Generate connection string
    connStr := generateConnStr(c)

	// Open database connection
    db, err = sql.Open(c.Type, connStr)
	if err != nil {
		err = fmt.Errorf("fail: sql.Open, %v", err)
	}

	// Check if the database is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("fail: db.Ping, %v", err)
	}

	// Check user permissions
	if err := utils.ValidateDB(db, c); err != nil {
		return nil, fmt.Errorf("fail: checkUserPermissions, %v", err)
	}

	return
}