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
	entClient     *ent.Client
	log           *slog.Logger
	notFoundError error
}

func NewGenerationRepository(entClient *ent.Client, log *slog.Logger) GenerationRepository {
	return &generationRepository{entClient: entClient, log: log}
}

func (r *generationRepository) Create(ctx context.Context, generation *ent.Generation) (*ent.Generation, error) {
	start := time.Now()
	createdGeneration, err := r.entClient.Generation.Create().
		SetRequestID(generation.RequestID).
		SetPrompt(generation.Prompt).
		SetStatus(generation.Status).
		SetVideoURL(generation.VideoURL).
		SetScriptURL(generation.ScriptURL).
		SetErrorMessage(generation.ErrorMessage).
		SetEmail(generation.Email).
		Save(ctx)
	duration := time.Since(start)

	if err != nil {
		r.log.Error("Failed to create generation in database",
			"error", err,
			"requestID", generation.RequestID,
			"email", generation.Email,
			"duration", duration)
		return nil, fmt.Errorf("failed to create generation: %w", err)
	}

	r.log.Info("Generation created in database",
		"requestID", generation.RequestID,
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

	_, err := r.entClient.Generation.UpdateOneID(ID).
		SetRequestID(generation.RequestID).
		SetPrompt(generation.Prompt).
		SetStatus(generation.Status).
		SetVideoURL(generation.VideoURL).
		SetScriptURL(generation.ScriptURL).
		SetErrorMessage(generation.ErrorMessage).
		SetEmail(generation.Email).
		Save(ctx)
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
		"status", generation.Status,
		"duration", duration)
	return nil
}

func (r *generationRepository) Ping(ctx context.Context) error {
	return r.entClient.Schema.Create(ctx)
}
