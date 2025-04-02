package models

import "time"

type UserAnswer struct {
	UserAnswerID      uint32          `json:"user_answer_id" gorm:"primaryKey"`
	TopicAssignmentID uint32          `json:"topic_assignment_id"`
	TopicAssignment   TopicAssignment `gorm:"foreignKey:TopicAssignmentID;references:TopicAssignmentID"`
	QuestionID        uint32          `json:"question_id"`
	Question          Question        `gorm:"foreignKey:QuestionID ;references:QuestionID "`
	TopicQuestionID   uint32          `json:"topic_question_id"`
	UserID            uint32          `json:"user_id"`
	User              User            `gorm:"foreignKey:UserID ;references:UserID"`
	IsCorrect         bool            `json:"is_correct"`
	CreatedBy         uint32          `json:"created_by"`
	UpdatedBy         uint32          `json:"updated_by"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}
