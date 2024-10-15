package domain

import (
	"gorm.io/gorm"
)

type Generation struct {
	gorm.Model
	RequestID    string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Prompt       string `gorm:"not null"`
	Status       string `gorm:"not null"`
	VideoURL     string `gorm:"default:null"`
	ScriptURL    string `gorm:"default:null"`
	ErrorMessage string `gorm:"default:null"`
}

type GenerationStatus int

const (
	StatusUnspecified GenerationStatus = iota
	StatusPending
	StatusProcessing
	StatusCompleted
	StatusFailed
)

// jsonの定義を確認
type GenerationResponse struct {
	RequestID    string `json:"requestId"`
	Status       string `json:"status"`
	VideoURL     string `json:"videoUrl"`
	ScriptURL    string `json:"scriptUrl"`
	Prompt       string `json:"prompt"`
	ErrorMessage string `json:"errorMessage"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CreateGenerationResponse struct {
	RequestID string `json:"requestId"`
}
