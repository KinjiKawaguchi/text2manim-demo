package domain

import (
	"gorm.io/gorm"
)

type Generation struct {
	gorm.Model
	RequestID    string `gorm:"unique;not null"`
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

type GenerationResponse struct {
	RequestID    string `json:"request_id"`
	Status       string `json:"status"`
	VideoURL     string `json:"video_url"`
	ScriptURL    string `json:"script_url"`
	Prompt       string `json:"prompt"`
	ErrorMessage string `json:"error_message"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
