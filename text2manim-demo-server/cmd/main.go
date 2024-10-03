package main

import (
	"text2manim-demo-server/internal/api"
	"text2manim-demo-server/internal/infrastructure"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/internal/usecase"
	"text2manim-demo-server/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.NewLogger()

	db := infrastructure.NewDatabase(log)
	repo := repository.NewGenerationRepository(db, log)
	useCase := usecase.NewGenerationUseCase(repo, 100, time.Hour, log)
	handler := api.NewHandler(useCase, log)

	r := gin.Default()
	r.POST("/v1/generations", handler.CreateGeneration)
	r.GET("/v1/generations/:request_id", handler.GetGenerationStatus)

	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
