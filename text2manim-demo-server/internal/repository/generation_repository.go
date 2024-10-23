package repository

import (
	"context"
	"fmt"
	"log/slog"
	"text2manim-demo-server/internal/domain/ent"
	"text2manim-demo-server/internal/domain/ent/generation"
	"time"

	"github.com/google/uuid"
)

type GenerationRepository interface {
	Create(ctx context.Context, generation *ent.Generation) (*ent.Generation, error)
	FindByRequestID(ctx context.Context, requestID string) (*ent.Generation, error)
	Update(ctx context.Context, ID uuid.UUID, generation *ent.Generation) error
	Ping(ctx context.Context) error
}

type generationRepository struct {
	entClient *ent.Client
	log       *slog.Logger
}

func NewGenerationRepository(entClient *ent.Client, log *slog.Logger) GenerationRepository {
	return &generationRepository{entClient: entClient, log: log}
}

func (r *generationRepository) Create(ctx context.Context, generation *ent.Generation) (*ent.Generation, error) {
	start := time.Now()

	// Createのビルダーを開始
	creator := r.entClient.Generation.Create()

	if generation.Prompt != "" {
		creator.SetPrompt(generation.Prompt)
	}

	if generation.Email != "" {
		creator.SetEmail(generation.Email)
	}

	if generation.Status != "" {
		creator.SetStatus(generation.Status)
	}

	if generation.RequestID != "" {
		creator.SetRequestID(generation.RequestID)
	}

	if generation.VideoURL != "" {
		creator.SetVideoURL(generation.VideoURL)
	}

	if generation.ScriptURL != "" {
		creator.SetScriptURL(generation.ScriptURL)
	}

	if generation.ErrorMessage != "" {
		creator.SetErrorMessage(generation.ErrorMessage)
	}

	// 保存を実行
	createdGeneration, err := creator.Save(ctx)

	duration := time.Since(start)
	if err != nil {
		r.log.Error("Failed to create generation in database",
			"error", err,
			"email", generation.Email,
			"duration", duration)
		return nil, fmt.Errorf("failed to create generation: %w", err)
	}

	r.log.Info("Generation created in database",
		"id", createdGeneration.ID,
		"email", generation.Email,
		"duration", duration)

	return createdGeneration, nil
}

func (r *generationRepository) FindByRequestID(ctx context.Context, requestID string) (*ent.Generation, error) {
	start := time.Now()
	generation, err := r.entClient.Generation.Query().
		Where(generation.RequestID(requestID)).
		Only(ctx)
	duration := time.Since(start)

	if err != nil {
		if ent.IsNotFound(err) {
			r.log.Warn("Generation not found",
				"requestID", requestID,
				"duration", duration)
			return nil, fmt.Errorf("generation not found: %w", err)
		}
		r.log.Error("Failed to find generation",
			"error", err,
			"requestID", requestID,
			"duration", duration)
		return nil, fmt.Errorf("failed to find generation: %w", err)
	}

	r.log.Info("Generation found",
		"requestID", requestID,
		"status", generation.Status,
		"duration", duration)
	return generation, nil
}

func (r *generationRepository) Update(ctx context.Context, ID uuid.UUID, generation *ent.Generation) error {
	start := time.Now()

	updater := r.entClient.Generation.UpdateOneID(ID)

	if generation.RequestID != "" {
		updater.SetRequestID(generation.RequestID)
	}

	if generation.Prompt != "" {
		updater.SetPrompt(generation.Prompt)
	}

	if generation.Status != "" {
		updater.SetStatus(generation.Status)
	}

	if generation.VideoURL != "" {
		updater.SetVideoURL(generation.VideoURL)
	}

	if generation.ScriptURL != "" {
		updater.SetScriptURL(generation.ScriptURL)
	}

	if generation.ErrorMessage != "" {
		updater.SetErrorMessage(generation.ErrorMessage)
	}

	if generation.Email != "" {
		updater.SetEmail(generation.Email)
	}

	updatedGeneration, err := updater.Save(ctx)

	duration := time.Since(start)

	if err != nil {
		r.log.Error("Failed to update generation status",
			"error", err,
			"ID", ID,
			"status", generation.Status,
			"duration", duration)
		return fmt.Errorf("failed to update generation: %w", err)
	}

	r.log.Info("Generation status updated",
		"ID", ID,
		"status", updatedGeneration.Status,
		"duration", duration)
	return nil
}

func (r *generationRepository) Ping(ctx context.Context) error {
	return r.entClient.Schema.Create(ctx)
}
