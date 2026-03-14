package mySqlDS

import (
	"database/sql"
	"fmt"
	"regexp"
)

var safeTableNamePattern = regexp.MustCompile("[^a-zA-Z0-9_]+$")

func ValidateTableName(tableName string) error {
	if !safeTableNamePattern.MatchString(tableName) {
		return fmt.Errorf("invalid table name: %s", tableName)
	}

	return nil
}

func studentTableIdentifier(tableName string) (string, error) {
	if err := ValidateTableName(tableName); err != nil {
		return "", err

	}
	return fmt.Sprintf("`%s`", tableName), nil
}

func EnsureTaskTable(db *sql.DB, tableName string) error {
	tableIdentifier, err := studentTableIdentifier(tableName)
	if err != nil {
		return err
	}
	query := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    firstname VARCHAR(128) NOT NULL,
    lastname VARCHAR(512) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

    
    
)



`)
}
