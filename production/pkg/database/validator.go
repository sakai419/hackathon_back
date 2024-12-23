package database

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"
	"local-test/pkg/apperrors"
	"local-test/pkg/utils"

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

func validateDBConfig(c *config.DBConfig) error {
	if c == nil {
		return apperrors.WrapConfigError(
			&apperrors.ErrInvalidInput{
				Message: "database config is nil",
			},
		)
	}

	fields := map[string]interface{}{
		"driver":            c.Driver,
		"user":              c.User,
		"password":          c.Pwd,
		"host":              c.Host,
		"port":              c.Port,
		"sslmode":           c.SSLMode,
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

	// Validate each field
	for fieldName, fieldValue := range fields {
		if err := utils.ValidateField("database", fieldName, fieldValue); err != nil {
			return apperrors.WrapValidationError(
				&apperrors.ErrOperationFailed{
					Operation: fmt.Sprintf("validate %s", fieldName),
					Err: err,
				},
			)
		}
	}

	return nil
}

func validateDB(db *sql.DB, c *config.DBConfig) error {
	// Validate database configuration
	if err := checkRequiredTables(db, c); err != nil {
		return apperrors.WrapValidationError(
			&apperrors.ErrOperationFailed{
				Operation: "check required tables",
				Err: err,
			},
		)
	}

	// Validate user permissions
	if err := checkUserPermissions(db, c.Driver); err != nil {
		return apperrors.WrapValidationError(
			&apperrors.ErrOperationFailed{
				Operation: "check user permissions",
				Err: err,
			},
		)
	}

	return nil
}

func checkRequiredTables(db *sql.DB, c *config.DBConfig) error {
	for _, table := range c.RequiredTables {
		if _, err := db.Exec(fmt.Sprintf("SELECT 1 FROM %s LIMIT 1", table)); err != nil {
			return fmt.Errorf("table %s does not exist or is not accessible: %v", table, err)
		}
	}
	return nil
}

func generateQueries(driver string) (*Queries) {
    switch driver {
	case "mysql":
		return &Queries{
			createTableQuery: "CREATE TABLE test_table (id INT)",
			dropTableQuery:   "DROP TABLE test_table",
			insertQuery:      "INSERT INTO test_table (id) VALUES (1)",
			selectQuery:      "SELECT * FROM test_table",
			deleteQuery:      "DELETE FROM test_table",
			updateQuery:      "UPDATE test_table SET id = 2",
		}
	case "postgres":
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

func checkUserPermissions(db *sql.DB, driver string) error {
	// Generate queries based on database type
    queries := generateQueries(driver)
    permissions := map[string]string{
        "create tables": queries.createTableQuery,
        "insert data":   queries.insertQuery,
        "select data":   queries.selectQuery,
        "update data":   queries.updateQuery,
        "delete data":   queries.deleteQuery,
        "drop tables":   queries.dropTableQuery,
    }

	// Begin transaction
    tx, err := db.Begin()
    if err != nil {
        return fmt.Errorf("failed to begin transaction: %v", err)
    }

	// Defer rollback if panic
    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
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
