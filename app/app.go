package app

import (
	"fmt"

	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

func NewApp(config AppConfig) (*App, error) {
	app := &App{}
	if conn, err := app.initDBConn(config.Database); err != nil {
		return nil, err
	} else {
		app.db = conn
	}

	if err := app.restoreState(config.Database.Type); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() {
	fmt.Println("run app")
}
