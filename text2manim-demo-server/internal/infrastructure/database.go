package infrastructure

import (
	"log/slog"
	"text2manim-demo-server/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase(log *slog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Error("Failed to connect to database", "error", err)
		panic(err)
	}

	err = db.AutoMigrate(&domain.Generation{})
	if err != nil {
		log.Error("Failed to auto migrate database", "error", err)
		panic(err)
	}

	log.Info("Database connected and migrated successfully")
	return db
}
