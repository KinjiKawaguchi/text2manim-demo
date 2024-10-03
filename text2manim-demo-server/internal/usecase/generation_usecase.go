package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"text2manim-demo-server/internal/domain"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/pkg/ratelimiter"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GenerationUseCase interface {
	CreateGeneration(email, prompt string) (string, error)
	GetGenerationStatus(requestID string) (domain.GenerationStatus, error)
}

type generationUseCase struct {
	repo             repository.GenerationRepository
	rateLimiter      *ratelimiter.RateLimiter
	videoAPIEndpoint string
	log              *slog.Logger
}

func NewGenerationUseCase(repo repository.GenerationRepository, limit int, interval time.Duration, videoAPIEndpoint string, log *slog.Logger) GenerationUseCase {
	return &generationUseCase{
		repo:             repo,
		rateLimiter:      ratelimiter.NewRateLimiter(limit, interval),
		videoAPIEndpoint: videoAPIEndpoint,
		log:              log,
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

	// データベースに保存
	err := uc.repo.Create(generation)
	if err != nil {
		uc.log.Error("Failed to create generation in database", "error", err, "requestID", requestID)
		return "", fmt.Errorf("failed to create generation: %w", err)
	}

	// APIを呼び出し
	err = uc.callVideoGenerationAPI(requestID, prompt)
	if err != nil {
		uc.log.Error("Failed to call video generation API", "error", err, "requestID", requestID)
		// APIの呼び出しに失敗した場合、ステータスを更新
		updateErr := uc.repo.UpdateStatus(requestID, domain.StatusFailed)
		if updateErr != nil {
			uc.log.Error("Failed to update generation status", "error", updateErr, "requestID", requestID)
		}
		return requestID, fmt.Errorf("failed to start video generation: %w", err)
	}

	return requestID, nil
}

func (uc *generationUseCase) callVideoGenerationAPI(requestID, prompt string) error {
	payload := map[string]string{
		"request_id": requestID,
		"prompt":     prompt,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(uc.videoAPIEndpoint, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to send request to video generation API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("video generation API returned non-OK status: %d", resp.StatusCode)
	}

	return nil
}

func (uc *generationUseCase) GetGenerationStatus(requestID string) (domain.GenerationStatus, error) {
	generation, err := uc.repo.FindByRequestID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.log.Info("Generation not found", "requestID", requestID)
			return "", fmt.Errorf("generation not found")
		}
		uc.log.Error("Failed to get generation status", "error", err, "requestID", requestID)
		return "", fmt.Errorf("failed to get generation status: %w", err)
	}
	uc.log.Info("Generation status retrieved", "requestID", requestID, "status", generation.Status)
	return domain.GenerationStatus(generation.Status), nil
}
