package domain

import "time"

type Generation struct {
	ID           uint      `gorm:"primaryKey"`
	RequestID    string    `gorm:"unique;not null"`
	Email        string    `gorm:"not null"`
	Prompt       string    `gorm:"not null"`
	Status       string    `gorm:"not null"`
	VideoURL     string    `gorm:"default:null"`
	ScriptURL    string    `gorm:"default:null"`
	ErrorMessage string    `gorm:"default:null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

type GenerationStatus string

const (
	StatusPending   GenerationStatus = "PENDING"
	StatusCompleted GenerationStatus = "COMPLETED"
	StatusFailed    GenerationStatus = "FAILED"
)

type GenerationResponse struct {
	Status       string `json:"status"`
	VideoURL     string `json:"video_url"`
	ScriptURL    string `json:"script_url"`
	Prompt       string `json:"prompt"`
	ErrorMessage string `json:"error_message"`
	UpdatedAt    int64  `json:"updated_at"`
}
