package repository

import (
	"log/slog"
	"text2manim-demo-server/internal/domain"
	"time"

	"gorm.io/gorm"
)

type GenerationRepository interface {
	Create(generation *domain.Generation) error
	FindByRequestID(requestID string) (*domain.Generation, error)
	UpdateStatus(requestID string, status domain.GenerationStatus) error
}

type generationRepository struct {
	db  *gorm.DB
	log *slog.Logger
}

func NewGenerationRepository(db *gorm.DB, log *slog.Logger) GenerationRepository {
	return &generationRepository{db: db, log: log}
}

func (r *generationRepository) Create(generation *domain.Generation) error {
	start := time.Now()
	err := r.db.Create(generation).Error
	duration := time.Since(start)

	if err != nil {
		r.log.Error("Failed to create generation in database",
			"error", err,
			"requestID", generation.RequestID,
			"email", generation.Email,
			"duration", duration)
		return err
	}

	r.log.Info("Generation created in database",
		"requestID", generation.RequestID,
		"email", generation.Email,
		"duration", duration)
	return nil
}

func (r *generationRepository) FindByRequestID(requestID string) (*domain.Generation, error) {
	start := time.Now()
	var generation domain.Generation
	err := r.db.Where("request_id = ?", requestID).First(&generation).Error
	duration := time.Since(start)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			r.log.Warn("Generation not found",
				"requestID", requestID,
				"duration", duration)
		} else {
			r.log.Error("Failed to find generation",
				"error", err,
				"requestID", requestID,
				"duration", duration)
		}
		return nil, err
	}

	r.log.Info("Generation found",
		"requestID", requestID,
		"status", generation.Status,
		"duration", duration)
	return &generation, nil
}

func (r *generationRepository) UpdateStatus(requestID string, status domain.GenerationStatus) error {
	start := time.Now()
	err := r.db.Model(&domain.Generation{}).Where("request_id = ?", requestID).Update("status", status).Error
	duration := time.Since(start)

	if err != nil {
		r.log.Error("Failed to update generation status",
			"error", err,
			"requestID", requestID,
			"newStatus", status,
			"duration", duration)
		return err
	}

	r.log.Info("Generation status updated",
		"requestID", requestID,
		"newStatus", status,
		"duration", duration)
	return nil
}
