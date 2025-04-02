package models

import "time"

type Option struct {
	OptionID   uint32    `gorm:"primaryKey;autoIncrement" json:"option_id"`
	QuestionID uint32    `json:"question_id"`
	Question   Question  `json:"foreignKey:QuestionID;references:QuestionID"`
	Content    string    `json:"content"`
	IsCorrect  bool      `json:"is_correct"`
	CreatedBy  uint32    `json:"created_by"`
	UpdatedBy  uint32    `json:"updated_by"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OptionWithoutIsCorrect struct {
	OptionID  uint32 `json:"option_id"`
	Content   string `json:"content"`
	IsCorrect bool   `json:"is_correct"`
}
