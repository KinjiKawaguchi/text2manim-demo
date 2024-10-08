package main

import (
	"text2manim-demo-server/internal/api"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/infrastructure"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/internal/usecase"
	"text2manim-demo-server/pkg/logger"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()
	cfg := config.Load(log)

	db := infrastructure.NewDatabase(cfg, log)
	repo := repository.NewGenerationRepository(db, log)
	useCase := usecase.NewGenerationUseCase(repo, cfg.RateLimitRequests, cfg.RateLimitInterval, cfg.Text2manimAPIEndpoint, log)
	handler := api.NewHandler(useCase, log)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:4200", "http://localhost:4200"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/v1/generations", handler.CreateGeneration)
	r.GET("/v1/generations/:request_id", handler.GetGenerationStatus)

	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
