package app

import "fmt"

type AppConfig struct {
	Database Database
}

type Database struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

type AppError struct {
	ErrorCode    string
	ErrorMessage string
}

func (e AppError) Error() string {
	return fmt.Sprintf("Error (#%v): %v", e.ErrorCode, e.ErrorMessage)
}
