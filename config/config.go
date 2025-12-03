package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite database: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetSQLite() *gorm.DB {
	return db
}
