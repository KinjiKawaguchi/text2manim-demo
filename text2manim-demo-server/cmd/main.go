package main

import (
	"text2manim-demo-server/internal/api"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/infrastructure"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/internal/usecase"
	"text2manim-demo-server/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	cfg := config.Load(log)

	entClient := infrastructure.NewEntClient(cfg, log)
	repo := repository.NewGenerationRepository(entClient, log)

	text2ManimClient, err := infrastructure.NewText2ManimClient(cfg.Text2manimApiEndpoint, cfg.Text2manimApiKey, cfg.Environment)
	if err != nil {
		log.Error("Failed to create Text2Manim client", "error", err)
		return
	}

	useCase := usecase.NewVideoGenerationUseCase(repo, cfg.RateLimitRequests, cfg.RateLimitInterval, text2ManimClient, log)
	handler := api.NewHandler(useCase, log)

	router := api.SetupRouter(cfg, handler)

	log.Info("Starting server on " + cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
