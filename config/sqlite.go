package config

import (
	"os"

	"github.com/DaviRodrigues/Project-go-with-gin/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	const dbPath string = "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		// Create the databse file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return  nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return  nil, err
		}

		file.Close()
	}

	// Create DB and connected
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite openning error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigrate error: %v", err)
		return nil, err
	}

	// Return the db
	return db, nil
}
