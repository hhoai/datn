package models

import "time"

type UserAssignment struct {
	UserAssignmentID uint32     `json:"user_assignment_id" gorm:"primaryKey"`
	UserID           uint32     `json:"user_id"`
	User             User       `gorm:"foreignKey:UserID;references:UserID"`
	AssignmentID     uint32     `json:"assignment_id"`
	Assignment       Assignment `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	Score            uint32     `json:"score"`
	Status           bool       `json:"status"`
	Comment          string     `json:"comment"`
	StartTime        time.Time  `json:"start_time"`
	EndTime          time.Time  `json:"end_time"`
	CreatedBy        uint32     `json:"created_by"`
	UpdatedBy        uint32     `json:"updated_by"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	CompletedAt      time.Time  `json:"completed_at"`
}
