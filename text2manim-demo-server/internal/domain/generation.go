package domain

import "time"

type Generation struct {
	ID        uint      `gorm:"primaryKey"`
	RequestID string    `gorm:"unique;not null"`
	Email     string    `gorm:"not null"`
	Prompt    string    `gorm:"not null"`
	Status    string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

type GenerationStatus string

const (
	StatusPending   GenerationStatus = "PENDING"
	StatusCompleted GenerationStatus = "COMPLETED"
	StatusFailed    GenerationStatus = "FAILED"
)
