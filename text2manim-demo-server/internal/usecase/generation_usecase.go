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
	GetGeneration(requestID string) (domain.Generation, error)
	CheckDatabaseConnection() error
	CheckText2ManimApiConnection() error
}

type generationUseCase struct {
	repo                  repository.GenerationRepository
	rateLimiter           *ratelimiter.RateLimiter
	text2ManimApiEndpoint string
	text2ManimApiKey      string
	log                   *slog.Logger
}

func NewGenerationUseCase(repo repository.GenerationRepository, limit int, interval time.Duration, text2ManimApiEndpoint string, text2ManimApiKey string, log *slog.Logger) GenerationUseCase {
	return &generationUseCase{
		repo:                  repo,
		rateLimiter:           ratelimiter.NewRateLimiter(limit, interval),
		text2ManimApiEndpoint: text2ManimApiEndpoint,
		text2ManimApiKey:      text2ManimApiKey, // APIキーを設定
		log:                   log,
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
	err = uc.callVideoGenerationApi(requestID, prompt)
	if err != nil {
		uc.log.Error("Failed to call video generation API", "error", err, "requestID", requestID)
		// APIの呼び出しに失敗した場合、ステータスを更新
		generation.Status = string(domain.StatusFailed)
		updateErr := uc.repo.Update(requestID, generation)
		if updateErr != nil {
			uc.log.Error("Failed to update generation status", "error", updateErr, "requestID", requestID)
		}
		return requestID, fmt.Errorf("failed to start video generation: %w", err)
	}

	return requestID, nil
}

func (uc *generationUseCase) callVideoGenerationApi(requestID, prompt string) error {
	payload := map[string]string{
		"request_id": requestID,
		"prompt":     prompt,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	url := fmt.Sprintf("%s/v1/generations", uc.text2ManimApiEndpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", uc.text2ManimApiKey) // APIキーをヘッダーに追加

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

func (uc *generationUseCase) GetGeneration(requestID string) (domain.Generation, error) {
	generation, err := uc.repo.FindByRequestID(requestID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			uc.log.Info("Generation not found", "requestID", requestID)
			return domain.Generation{}, fmt.Errorf("generation not found")
		}
		uc.log.Error("Failed to get generation status", "error", err, "requestID", requestID)
		return domain.Generation{}, fmt.Errorf("failed to get generation status: %w", err)
	}

	// APIから最新の状態を取得
	resp, err := uc.checkVideoGenerationStatus(requestID)
	if err != nil {
		uc.log.Error("Failed to check video generation status", "error", err, "requestID", requestID)
		return *generation, nil // エラーが発生しても現在の状態を返す
	}

	// データベースの状態を更新
	generation.Status = string(resp.Status)
	generation.VideoURL = resp.VideoURL
	generation.ScriptURL = resp.ScriptURL
	generation.ErrorMessage = resp.ErrorMessage
	generation.UpdatedAt = time.Unix(resp.UpdatedAt, 0)

	err = uc.repo.Update(requestID, generation)
	if err != nil {
		uc.log.Error("Failed to update generation in database", "error", err, "requestID", requestID)
	}

	uc.log.Info("Generation status retrieved and updated", "requestID", requestID, "generation", generation)
	return *generation, nil
}

func (uc *generationUseCase) checkVideoGenerationStatus(requestID string) (*domain.GenerationResponse, error) {
	url := fmt.Sprintf("%s/v1/generations/%s", uc.text2ManimApiEndpoint, requestID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Api-Key", uc.text2ManimApiKey)

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

func (uc *generationUseCase) CheckDatabaseConnection() error {
	return uc.repo.Ping()
}

func (uc *generationUseCase) CheckText2ManimApiConnection() error {
	url := fmt.Sprintf("%s/v1/health", uc.text2ManimApiEndpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("X-Api-Key", uc.text2ManimApiKey) // APIキーをヘッダーに追加

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
