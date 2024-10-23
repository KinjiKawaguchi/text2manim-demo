package infrastructure

import (
	"context"
	"fmt"
	"log/slog"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/domain/ent"

	_ "github.com/lib/pq"
)

func NewEntClient(cfg *config.Config, log *slog.Logger) *ent.Client {
	DSN := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", cfg.SupabaseHost, cfg.SupabaseUser, cfg.SupabasePassword, cfg.SupabaseDBName, cfg.SupabasePort)
	entClient, err := ent.Open("postgres", DSN)
	if err != nil {
		log.Error("Failed to create Ent client", "error", err)
		panic(err)
	}
	log.Info("Successfully created Ent client")
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Error("Failed to create schema", "error", err)
		panic(err)
	}
	return entClient
}
