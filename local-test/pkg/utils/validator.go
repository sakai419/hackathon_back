package utils

import (
	"database/sql"
	"fmt"
	"local-test/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

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

func generateQueries(db_type string) (create_table_query, drop_table_query, insert_query, select_query, delete_query, update_query string) {
    switch db_type {
    case "mysql":
        create_table_query = "CREATE TABLE test_table (id INT)"
        drop_table_query = "DROP TABLE test_table"
        insert_query = "INSERT INTO test_table (id) VALUES (1)"
        select_query = "SELECT * FROM test_table"
        delete_query = "DELETE FROM test_table"
        update_query = "UPDATE test_table SET id = 2"
    default:
        create_table_query = ""
        drop_table_query = ""
        insert_query = ""
        select_query = ""
        delete_query = ""
        update_query = ""
    }

    return
}

func checkUserPermissions(db *sql.DB, db_type string) error {
    // generate queries for testing user permissions
    create_table_query, drop_table_query, insert_query, select_query, delete_query, update_query := generateQueries(db_type)

    // check user permissions
    _, err := db.Exec(create_table_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to create tables: %v", err)
    }
    _, err = db.Exec(insert_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to insert data: %v", err)
    }
    _, err = db.Exec(select_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to select data: %v", err)
    }
    _, err = db.Exec(update_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to update data: %v", err)
    }
    _, err = db.Exec(delete_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to delete data: %v", err)
    }
    _, err = db.Exec(drop_table_query)
    if err != nil {
        return fmt.Errorf("user does not have permission to drop tables: %v", err)
    }

    return nil
}

func ValidateDBConfig(c *config.DBConfig) error {
	if c.Driver == "" {
		return fmt.Errorf("database driver is required")
	} else if c.User == "" {
		return fmt.Errorf("database user is required")
	} else if c.Pwd == "" {
		return fmt.Errorf("database password is required")
	} else if c.Host == "" {
		return fmt.Errorf("database host is required")
	} else if c.Database == "" {
		return fmt.Errorf("database name is required")
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
