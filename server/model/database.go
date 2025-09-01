package model

import (
	"log"
	"log/slog"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	slog.Info("Initializing database connection")
	var err error
	db, err = gorm.Open(sqlite.Open("ero.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database", "error", err)
	}
	slog.Info("Database connection established")

	if err := db.AutoMigrate(&meta{}, &User{}, &Container{}); err != nil {
		log.Fatal("Failed to migrate database", "error", err)
	}
	slog.Info("Database migration completed")
}

var once sync.Once

func GetDB() *gorm.DB {
	once.Do(initDB)
	return db
}
