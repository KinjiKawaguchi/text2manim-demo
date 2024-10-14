package main

import (
	"text2manim-demo-server/internal/api"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/infrastructure"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/internal/usecase"
	"text2manim-demo-server/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()
	cfg := config.Load(log)
	db := infrastructure.NewDatabase(cfg, log)
	repo := repository.NewGenerationRepository(db, log)
	useCase := usecase.NewVideoGenerationUseCase(repo, cfg.RateLimitRequests, cfg.RateLimitInterval, cfg.Text2manimApiEndpoint, cfg.Text2manimApiKey, log)
	handler := api.NewHandler(useCase, log)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Api-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	r.POST("/v1/generations", handler.CreateGeneration)
	r.GET("/v1/generations/:request_id", handler.GetGeneration)
	r.GET("/health", handler.HealthCheck)

	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
