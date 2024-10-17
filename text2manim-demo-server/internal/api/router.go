package api

import (
	"text2manim-demo-server/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg *config.Config, handler *Handler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Api-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60,
	}))

	setupRoutes(r, handler)

	return r
}

func setupRoutes(r *gin.Engine, handler *Handler) {
	r.POST("/v1/generations", handler.CreateGeneration)
	r.GET("/v1/generations/:request_id", handler.GetGeneration)
	r.GET("/v1/health", handler.HealthCheck)
}
