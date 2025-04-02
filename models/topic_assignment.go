package models

import "time"

type TopicAssignment struct {
	TopicAssignmentID uint32     `json:"topic_assignment_id" gorm:"primaryKey"`
	AssignmentID      uint32     `json:"assignment_id"`
	Assignment        Assignment `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	TopicID           uint32     `json:"topic_id"`
	Topic             Topic      `gorm:"foreignKey:TopicID;references:TopicID"`
	CreatedBy         uint32     `json:"created_by"`
	UpdatedBy         uint32     `json:"updated_by"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
