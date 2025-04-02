package models

import "time"

type Challenge struct {
	ChallengeID uint32    `json:"challenge_id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   uint32
	UpdatedBy   uint32
}
