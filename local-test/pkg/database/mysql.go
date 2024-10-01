package database

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	User string
	Pwd string
	Host string
	Database string
}

func ConnectToMysql(c DBConfig) (db *sql.DB, err error) {
    connStr := fmt.Sprintf("%s:%s@%s/%s", c.User, c.Pwd, c.Host, c.Database)

    db, err = sql.Open("mysql", connStr)
	if err != nil {
		err = fmt.Errorf("fail: sql.Open, %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("fail: db.Ping, %v", err)
	}

	return
}