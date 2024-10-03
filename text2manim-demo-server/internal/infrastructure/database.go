package infrastructure

import (
	"fmt"
	"log/slog"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config, log *slog.Logger) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=require",
		cfg.SupabaseHost,
		cfg.SupabaseUser,
		cfg.SupabasePassword,
		cfg.SupabaseDBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
