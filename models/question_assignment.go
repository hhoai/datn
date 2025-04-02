package models

import "time"

type QuestionAssignment struct {
	QuestionAssignmentID string     `json:"question_assignment_id"`
	QuestionID           uint32     `json:"question_id"`
	Question             Question   `gorm:"foreignKey:QuestionID;references:QuestionID"`
	AssignmentID         uint32     `json:"assignment_id"`
	Assignment           Assignment `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	CreatedBy            uint32     `json:"created_by"`
	UpdatedBy            uint32     `json:"updated_by"`
	CreatedAt            time.Time  `json:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"`
}
