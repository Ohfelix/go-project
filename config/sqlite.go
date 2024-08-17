package config

import (
	"example/hello/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Verifique se o diretório existe, caso contrário, crie-o
	if _, err := os.Stat("./db"); os.IsNotExist(err) {
		err := os.Mkdir("./db", 0755)
		if err != nil {
			logger.Error("Failed to create database directory:", err)
			return nil, err
		}
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to SQLite:", err)
		return nil, err
	}

	logger.Info("Connected to SQLite successfully")

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error("Failed to run AutoMigrate:", err)
		return nil, err
	}

	return db, nil
}
