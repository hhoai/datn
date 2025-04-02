package models

import "time"

type TopicQuestion struct {
	TopicQuestionID uint32    `json:"topic_question_id" gorm:"primaryKey"`
	TopicID         uint32    `json:"topic_id"`
	Topic           Topic     `gorm:"foreignKey:TopicID;references:TopicID"`
	QuestionID      uint32    `json:"question_id"`
	Question        Question  `gorm:"foreignKey:QuestionID;references:QuestionID"`
	CreatedBy       uint32    `json:"created_by"`
	UpdatedBy       uint32    `json:"updated_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TopicQuestionResponse struct {
	TopicQuestionID uint32                    `json:"topic_question_id"`
	TopicID         uint32                    `json:"topic_id"`
	QuestionID      uint32                    `json:"question_id"`
	TypeQuestionID  uint32                    `json:"type_question_id"`
	Content         string                    `json:"content"`
	Score           uint32                    `json:"score"`
	Options         []*OptionWithoutIsCorrect `json:"options"`
	IsCorrect       bool                      `json:"is_correct"`
}
