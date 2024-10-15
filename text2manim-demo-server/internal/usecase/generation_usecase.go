package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"text2manim-demo-server/internal/domain"
	"text2manim-demo-server/internal/repository"
	"text2manim-demo-server/pkg/ratelimiter"
	"time"

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
	// TODO: Rate limit引っかかったユーザーを保存しておきたい  https://github.com/KinjiKawaguchi/text2manim-demo/issues/15
	if !uc.rateLimiter.Allow() {
		uc.logger.Warn("Rate limit exceeded", "email", email)
		return "", errors.New("rate limit exceeded")
	}

	generation := &domain.Generation{
		Email:  email,
		Prompt: prompt,
		Status: string(domain.StatusPending),
	}

	if err := uc.repo.Create(generation); err != nil {
		uc.logger.Error("Failed to create generation record", "error", err, "id", generation.ID)
		return "", fmt.Errorf("failed to create generation record: %w", err)
	}

	resp, err := uc.initiateVideoGeneration(prompt)
	if err != nil {
		uc.logger.Error("Failed to initiate video generation", "error", err, "id", generation.ID)
		generation.Status = string(domain.StatusFailed)
		if updateErr := uc.repo.Update(generation.ID, generation); updateErr != nil {
			uc.logger.Error("Failed to update generation status", "error", updateErr, "id", generation.ID)
		}
		return "", fmt.Errorf("failed to initiate video generation: %w", err)
	}

	// Update the generation record with the request ID
	generation.RequestID = resp.RequestID
	if updateErr := uc.repo.Update(generation.ID, generation); updateErr != nil {
		uc.logger.Error("Failed to update generation with request ID", "error", updateErr, "id", generation.ID)
		// Note: We don't return an error here as the video generation process has already been initiated
	}

	return resp.RequestID, nil
}

func (uc *videoGenerationUseCase) initiateVideoGeneration(prompt string) (*domain.CreateGenerationResponse, error) {
	payload := map[string]string{
		"prompt": prompt,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/v1/generations", uc.text2ManimAPIURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
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

	// レスポンスボディを読み取り
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// レスポンスの内容をログに出力
	uc.logger.Info("API Response", "body", string(body))

	var apiResponse domain.CreateGenerationResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode API response: %w", err)
	}

	return &apiResponse, nil
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

	// Only fetch from API if the status is Pending or Processing
	if generation.Status == string(domain.StatusPending) || generation.Status == string(domain.StatusProcessing) {
		apiStatus, err := uc.fetchVideoGenerationStatus(requestID)
		if err != nil {
			uc.logger.Error("Failed to fetch video generation status", "error", err, "requestID", requestID)
			return *generation, nil
		}

		// Update generation with API response
		generation.Status = string(apiStatus.Status)
		generation.VideoURL = apiStatus.VideoURL
		generation.ScriptURL = apiStatus.ScriptURL
		generation.ErrorMessage = apiStatus.ErrorMessage
		generation.UpdatedAt = time.Unix(apiStatus.UpdatedAt, 0)

		if err := uc.repo.Update(generation.ID, generation); err != nil {
			uc.logger.Error("Failed to update generation record", "error", err, "requestID", requestID)
		}

		uc.logger.Info("Generation status updated", "requestID", requestID, "generation", generation)
	} else {
		uc.logger.Info("Generation status not updated (already completed or failed)", "requestID", requestID, "status", generation.Status)
	}

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("video generation API returned non-OK status: %d", resp.StatusCode)
	}

	var result struct {
		GenerationStatus domain.GenerationResponse `json:"generationStatus"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
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
