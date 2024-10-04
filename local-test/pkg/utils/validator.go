package utils

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

type Queries struct {
    createTableQuery string
    dropTableQuery   string
    insertQuery      string
    selectQuery      string
    deleteQuery      string
    updateQuery      string
}

func validateField(fieldName string, fieldValue interface{}) error {
	switch v := fieldValue.(type) {
	case string:
		if v == "" {
			return fmt.Errorf("database %s is required", fieldName)
		}
	case int:
		if v == 0 {
			return fmt.Errorf("database %s is required", fieldName)
		}
	default:
		if fieldValue == nil {
			return fmt.Errorf("database %s is required", fieldName)
		}
	}
	return nil
}

func checkRequiredTables(db *sql.DB) error {
	requiredTables := []string{
        "accounts",
        "blocks",
        "follow_requests",
        "follows",
        "hashtags",
        "interests",
        "likes",
        "messages",
        "notifications",
        "profiles",
        "replies",
        "retweets_and_quotes",
        "tweet_hashtags",
        "tweets",
    }
	for _, table := range requiredTables {
		if _, err := db.Exec(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", table)); err != nil {
			return fmt.Errorf("table %s does not exist or is not accessible: %v", table, err)
		}
	}
	return nil
}

func generateQueries(db_type string) (*Queries) {
    switch db_type {
    case "mysql":
		return &Queries{
			createTableQuery: "CREATE TABLE test_table (id INT)",
			dropTableQuery:   "DROP TABLE test_table",
			insertQuery:      "INSERT INTO test_table (id) VALUES (1)",
			selectQuery:      "SELECT * FROM test_table",
			deleteQuery:      "DELETE FROM test_table",
			updateQuery:      "UPDATE test_table SET id = 2",
		}
    default:
		return nil
    }
}

func checkUserPermissions(db *sql.DB, dbType string) error {
    queries := generateQueries(dbType)
    permissions := map[string]string{
        "create tables": queries.createTableQuery,
        "insert data":   queries.insertQuery,
        "select data":   queries.selectQuery,
        "update data":   queries.updateQuery,
        "delete data":   queries.deleteQuery,
        "drop tables":   queries.dropTableQuery,
    }

    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %v", err)
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p) // re-throw panic after Rollback
        } else if err != nil {
            tx.Rollback() // err is non-nil; don't change it
        } else {
            err = tx.Commit() // err is nil; if Commit returns error update err
        }
    }()

    // create tables
    if err = executeQuery(tx, queries.createTableQuery, "create tables"); err != nil {
        return err
    }

    // other operations
    for action, query := range permissions {
        if action == "create tables" || action == "drop tables" {
            continue
        }
        if err = executeQuery(tx, query, action); err != nil {
            return err
        }
    }

    // drop tables
    if err = executeQuery(tx, queries.dropTableQuery, "drop tables"); err != nil {
        return err
    }

    return nil
}

func executeQuery(tx *sql.Tx, query, action string) error {
    _, err := tx.Exec(query)
    if err != nil {
        return fmt.Errorf("user does not have permission to %s: %v", action, err)
    }
    return nil
}

func ValidateDBConfig(c *config.DBConfig) error {
    fields := map[string]interface{}{
        "driver":            c.Driver,
        "user":              c.User,
        "password":          c.Pwd,
        "host":              c.Host,
        "database":          c.Database,
        "charset":           c.Charset,
        "max open conns":    c.MaxOpenConns,
        "max idle conns":    c.MaxIdleConns,
        "conn max lifetime": c.ConnMaxLifetime,
        "conn max idle time": c.ConnMaxIdleTime,
        "timeout":           c.Timeout,
        "read timeout":      c.ReadTimeout,
        "write timeout":     c.WriteTimeout,
    }

    for fieldName, fieldValue := range fields {
        if err := validateField(fieldName, fieldValue); err != nil {
            return err
        }
    }

    return nil
}

func ValidateDB(db *sql.DB, c config.DBConfig) error {
    // Validate database configuration
	if err := checkRequiredTables(db); err != nil {
		return fmt.Errorf("fail: checkRequiredTables, %v", err)
	}
	if err := checkUserPermissions(db, c.Driver); err != nil {
		return fmt.Errorf("fail: checkUserPermissions, %v", err)
	}
	return nil
}
