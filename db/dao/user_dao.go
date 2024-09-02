package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectToDB(db *sql.DB) (_db *sql.DB) {
	// ①-1
	EnvLoad()
	mysqlUser, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		log.Fatal("MYSQL_USER env variable not set")
	}
	mysqlUserPwd, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		log.Fatal("MYSQL_PASSWORD env variable not set")
	}
	mysqlDatabase, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		log.Fatal("MYSQL_DATABASE env variable not set")
	}

	// ①-2
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/%s", mysqlUser, mysqlUserPwd, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
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