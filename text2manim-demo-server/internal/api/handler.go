package api

import (
	"log/slog"
	"net/http"
	"text2manim-demo-server/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase usecase.VideoGenerationUseCase
	log     *slog.Logger
}

func NewHandler(useCase usecase.VideoGenerationUseCase, log *slog.Logger) *Handler {
	return &Handler{useCase: useCase, log: log}
}

func (h *Handler) CreateGeneration(c *gin.Context) {
	var request struct {
		Email  string `json:"email" binding:"required,email"`
		Prompt string `json:"prompt" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.Warn("Invalid request", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestID, err := h.useCase.CreateGeneration(c, request.Email, request.Prompt)
	if err != nil {
		h.log.Error("Failed to create generation", "error", err, "email", request.Email)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("Generation created successfully", "requestID", requestID, "email", request.Email)
	c.JSON(http.StatusOK, requestID)
}

func (h *Handler) GetGeneration(c *gin.Context) {
	requestID := c.Param("request_id")

	generation, err := h.useCase.GetGenerationStatus(c, requestID)
	if err != nil {
		h.log.Error("Failed to get generation", "error", err, "requestID", requestID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.log.Info("Generation retrieved", "requestID", requestID, "generation", generation)
	c.JSON(http.StatusOK, generation)
}

func (h *Handler) HealthCheck(c *gin.Context) {
	// データベース接続を確認
	if err := h.useCase.CheckDatabaseConnection(c); err != nil {
		h.log.Error("Database health check failed", "error", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "message": "Database connection failed"})
		return
	}

	// 動画生成APIの状態を確認
	if err := h.useCase.HealthCheck(c); err != nil {
		h.log.Error("Video API health check failed", "error", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy", "message": "Video API connection failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}
