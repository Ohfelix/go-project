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
		return fmt.Errorf("error initializing SQLite: %w", err)
	}
	return nil
}

// GetLogger retorna a instância do logger, inicializando se necessário
func GetLogger(p string) *Logger {
	if logger == nil {
		logger = NewLogger(p)
	}
	return logger
}
