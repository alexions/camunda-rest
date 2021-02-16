package app

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	ErrorWrongDBTypeCode = "wrong_db_type"
)

var sqliteFixture = `
CREATE TABLE IF NOT EXISTS topic_links (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    topic TEXT NOT NULL,
    method TEXT NOT NULL,
    endpoint TEXT NOT NULL
);
`

func (a *App) initDBConn(config Database) (*gorm.DB, error) {
	switch config.Type {
	case "sqlite":
		return gorm.Open(sqlite.Open(config.Dbname), &gorm.Config{})

	case "mysql":
		fallthrough
	case "postgres":
		fallthrough
	default:
		return nil, AppError{
			ErrorCode:    ErrorWrongDBTypeCode,
			ErrorMessage: fmt.Sprintf("Selected DB type (%s) is not supported", config.Type),
		}
	}
}

func (a *App) restoreState(dbType string) error {
	switch dbType {
	case "sqlite":
		return a.db.Exec(sqliteFixture).Error
	}

	return nil
}
