package database

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func createConnStr(c config.DBConfig) string{
	if (c.Type == "mysql") {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Pwd, c.Host, c.Database)
	} else {
		return ""
	}
}

func ConnectToDB(c config.DBConfig) (db *sql.DB, err error) {
	// Create connection string
    connStr := createConnStr(c)

	// Open database connection
    db, err = sql.Open(c.Type, connStr)
	if err != nil {
		err = fmt.Errorf("fail: sql.Open, %v", err)
	}

	// Check if the database is alive
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("fail: db.Ping, %v", err)
	}

	return
}