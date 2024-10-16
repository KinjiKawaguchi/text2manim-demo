package main

import (
	"text2manim-demo-server/internal/api"
	"text2manim-demo-server/internal/config"
	"text2manim-demo-server/internal/infrastructure"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/internal/usecase"
	"text2manim-demo-server/pkg/logger"

	pb "github.com/KinjiKawaguchi/text2manim/api/pkg/pb/text2manim/v1"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	log := logger.NewLogger()
	cfg := config.Load(log)
	db := infrastructure.NewDatabase(cfg, log)
	repo := repository.NewGenerationRepository(db, log)
	log.Info("Connecting to Text2Manim API", "endpoint", cfg.Text2manimApiEndpoint)
	creds := credentials.NewClientTLSFromCert(nil, "")

	// セキュアなチャネルを通じてサーバーに接続
	conn, err := grpc.NewClient(cfg.Text2manimApiEndpoint, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Error("Failed to connect to Text2Manim API", "error", err)
	}
	defer conn.Close()

	text2ManimClient := pb.NewText2ManimServiceClient(conn)

	useCase := usecase.NewVideoGenerationUseCase(repo, cfg.RateLimitRequests, cfg.RateLimitInterval, text2ManimClient, log)
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
	r.GET("/v1/health", handler.HealthCheck)

	log.Info("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Error("Failed to start server", "error", err)
	}
}
