package usecase

import (
	"errors"
	"log/slog"
	"text2manim-demo-server/internal/domain"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/pkg/ratelimiter"
	"time"

	"github.com/google/uuid"
)

type GenerationUseCase interface {
	CreateGeneration(email, prompt string) (string, error)
	GetGenerationStatus(requestID string) (domain.GenerationStatus, error)
}

type generationUseCase struct {
	repo        repository.GenerationRepository
	rateLimiter *ratelimiter.RateLimiter
	log         *slog.Logger
}

func NewGenerationUseCase(repo repository.GenerationRepository, limit int, interval time.Duration, log *slog.Logger) GenerationUseCase {
	return &generationUseCase{
		repo:        repo,
		rateLimiter: ratelimiter.NewRateLimiter(limit, interval),
		log:         log,
	}
}

func (uc *generationUseCase) CreateGeneration(email, prompt string) (string, error) {
	if !uc.rateLimiter.Allow() {
		uc.log.Warn("Rate limit exceeded", "email", email)
		return "", errors.New("rate limit exceeded")
	}

	requestID := uuid.New().String()
	generation := &domain.Generation{
		RequestID: requestID,
		Email:     email,
		Prompt:    prompt,
		Status:    string(domain.StatusPending),
		CreatedAt: time.Now(),
	}

	err := uc.repo.Create(generation)
	if err != nil {
		uc.log.Error("Failed to create generation", "error", err, "email", email)
		return "", err
	}

	uc.log.Info("Generation created", "requestID", requestID, "email", email)
	return requestID, nil
}

func (uc *generationUseCase) GetGenerationStatus(requestID string) (domain.GenerationStatus, error) {
	generation, err := uc.repo.FindByRequestID(requestID)
	if err != nil {
		uc.log.Error("Failed to get generation status", "error", err, "requestID", requestID)
		return "", err
	}
	uc.log.Info("Generation status retrieved", "requestID", requestID, "status", generation.Status)
	return domain.GenerationStatus(generation.Status), nil
}
