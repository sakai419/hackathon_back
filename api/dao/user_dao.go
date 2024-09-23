package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB(db *sql.DB) (_db *sql.DB) {
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PWD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")

    connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)

    _db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}

	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	return _db
}

func CloseDB(db *sql.DB) error {
	return db.Close()
}

func GetTX(db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}

func RollbackTX(tx *sql.Tx) error {
	return tx.Rollback()
}

func CommitTX(tx *sql.Tx) error {
	return tx.Commit()
}

func InsertUser(tx *sql.Tx, id, name string, age int) error {
	if _, err := tx.Exec("INSERT INTO user(id, name, age) VALUES(?, ?, ?)", id, name, age); err != nil {
		log.Printf("fail: tx.Exec, %v\n", err)
		return err
	}
	return nil
}

func SelectUserByName(db *sql.DB, name string) (*sql.Rows, error) {
	return db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
}