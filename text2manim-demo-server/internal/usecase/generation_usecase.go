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

type VideoGenerationUseCase interface {
	RequestVideoGeneration(email, prompt string) (string, error)
	GetVideoGenerationStatus(requestID string) (domain.Generation, error)
	CheckDatabaseConnection() error
	CheckText2ManimAPIConnection() error
}

type videoGenerationUseCase struct {
	repo             repository.GenerationRepository
	rateLimiter      *ratelimiter.RateLimiter
	text2ManimAPIURL string
	text2ManimAPIKey string
	logger           *slog.Logger
}

func NewVideoGenerationUseCase(repo repository.GenerationRepository, limit int, interval time.Duration, text2ManimAPIURL string, text2ManimAPIKey string, logger *slog.Logger) VideoGenerationUseCase {
	return &videoGenerationUseCase{
		repo:             repo,
		rateLimiter:      ratelimiter.NewRateLimiter(limit, interval),
		text2ManimAPIURL: text2ManimAPIURL,
		text2ManimAPIKey: text2ManimAPIKey,
		logger:           logger,
	}
}

func (uc *videoGenerationUseCase) RequestVideoGeneration(email, prompt string) (string, error) {
	if !uc.rateLimiter.Allow() {
		uc.logger.Warn("Rate limit exceeded", "email", email)
		return "", errors.New("rate limit exceeded")
	}

	requestID := uuid.New().String()
	generation := &domain.Generation{
		RequestID: requestID,
		Email:     email,
		Prompt:    prompt,
		Status:    string(domain.StatusPending),
	}

	if err := uc.repo.Create(generation); err != nil {
		uc.logger.Error("Failed to create generation record", "error", err, "requestID", requestID)
		return "", fmt.Errorf("failed to create generation record: %w", err)
	}

	if err := uc.initiateVideoGeneration(requestID, prompt); err != nil {
		uc.logger.Error("Failed to initiate video generation", "error", err, "requestID", requestID)
		generation.Status = string(domain.StatusFailed)
		if updateErr := uc.repo.Update(requestID, generation); updateErr != nil {
			uc.logger.Error("Failed to update generation status", "error", updateErr, "requestID", requestID)
		}
		return requestID, fmt.Errorf("failed to initiate video generation: %w", err)
	}

	return requestID, nil
}

func (uc *videoGenerationUseCase) initiateVideoGeneration(requestID, prompt string) error {
	payload := map[string]string{
		"request_id": requestID,
		"prompt":     prompt,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/v1/generations", uc.text2ManimAPIURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", uc.text2ManimAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request to video generation API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("video generation API returned non-OK status: %d", resp.StatusCode)
	}

	return nil
}

func (uc *videoGenerationUseCase) GetVideoGenerationStatus(requestID string) (domain.Generation, error) {
	generation, err := uc.repo.FindByRequestID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.logger.Info("Generation not found", "requestID", requestID)
			return domain.Generation{}, fmt.Errorf("generation not found")
		}
		uc.logger.Error("Failed to get generation status", "error", err, "requestID", requestID)
		return domain.Generation{}, fmt.Errorf("failed to get generation status: %w", err)
	}

	apiStatus, err := uc.fetchVideoGenerationStatus(requestID)
	if err != nil {
		uc.logger.Error("Failed to fetch video generation status", "error", err, "requestID", requestID)
		return *generation, nil
	}

	generation.Status = string(apiStatus.Status)
	generation.VideoURL = apiStatus.VideoURL
	generation.ScriptURL = apiStatus.ScriptURL
	generation.ErrorMessage = apiStatus.ErrorMessage
	generation.UpdatedAt = time.Unix(apiStatus.UpdatedAt, 0)

	if err := uc.repo.Update(requestID, generation); err != nil {
		uc.logger.Error("Failed to update generation record", "error", err, "requestID", requestID)
	}

	uc.logger.Info("Generation status updated", "requestID", requestID, "generation", generation)
	return *generation, nil
}

func (uc *videoGenerationUseCase) fetchVideoGenerationStatus(requestID string) (*domain.GenerationResponse, error) {
	url := fmt.Sprintf("%s/v1/generations/%s", uc.text2ManimAPIURL, requestID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Api-Key", uc.text2ManimAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to video generation API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("video generation API returned non-OK status: %d", resp.StatusCode)
	}

	var result struct {
		GenerationStatus domain.GenerationResponse `json:"generation_status"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	return &result.GenerationStatus, nil
}

func (uc *videoGenerationUseCase) CheckDatabaseConnection() error {
	return uc.repo.Ping()
}

func (uc *videoGenerationUseCase) CheckText2ManimAPIConnection() error {
	url := fmt.Sprintf("%s/v1/health", uc.text2ManimAPIURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Api-Key", uc.text2ManimAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to connect to video API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("video API returned non-OK status: %d", resp.StatusCode)
	}

	return nil
}
