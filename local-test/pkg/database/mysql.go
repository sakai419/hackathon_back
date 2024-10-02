package database

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(c config.DBConfig) (db *sql.DB, err error) {
    connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s", c.User, c.Pwd, c.Host, c.Database)

    db, err = sql.Open(c.Type, connStr)
	if err != nil {
		err = fmt.Errorf("fail: sql.Open, %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("fail: db.Ping, %v", err)
	}

	return
}