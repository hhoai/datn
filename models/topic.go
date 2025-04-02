package models

import "time"

type Topic struct {
	TopicID      uint32 `json:"topic_id" gorm:"primaryKey" `
	TotalScore   uint32 `json:"total_score" `
	MinimumScore uint32 `json:"minimum_score" `
	Name         string `json:"name"`
	Description  string `json:"description"`
	Status       bool
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedBy    uint32    `json:"created_by"`
	UpdatedBy    uint32    `json:"updated_by"`
}
