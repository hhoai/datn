package models

import "time"

type Assignment struct {
	AssignmentID     uint32         `gorm:"primaryKey;autoIncrement"`
	LessonID         uint32         `json:"lesson_id"`
	Lesson           Lesson         `gorm:"foreignKey:LessonID;references:LessonID"`
	TypeAssignmentID uint32         `json:"type_assignment_id"`
	TypeAssignment   TypeAssignment `gorm:"foreignKey:TypeAssignmentID;references:TypeAssignmentID"`
	CreatedBy        uint32         `json:"user_id"`
	User             User           `gorm:"foreignKey:CreatedBy;references:UserID"`
	UpdatedBy        uint32         `gorm:"foreignKey:UpdatedBy;references:UserID"`
	Title            string         `json:"title"`
	Body             string         `json:"assignment_body"`
	Score            uint32         `json:"score"`
	Status           bool           `json:"status"`
	DueDate          time.Time      `json:"due_date"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
