package models

import "time"

type UserOption struct {
	UserOptionID uint32     `gorm:"primaryKey;autoIncrement" json:"user_option_id"`
	OptionID     uint32     `json:"option_id"`
	QuestionID   uint32     `json:"question_id"`
	Question     Question   `json:"foreignKey:QuestionID;references:QuestionID"`
	UserAnswerID uint32     `json:"user_answer_id"`
	UserAnswer   UserAnswer `gorm:"foreignKey:UserAnswerID;references:UserAnswerID"`
	IsCorrect    bool       `json:"is_correct"`
	Content      string     `json:"content"`
	CreatedBy    uint32     `json:"created_by"`
	UpdatedBy    uint32     `json:"updated_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
