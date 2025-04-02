package models

import "time"

type FileAssignment struct {
	FileAssignmentID uint32     `gorm:"primaryKey;autoIncrement"`
	FileName         string     `json:"file_name"`
	AssignmentID     uint32     `json:"assignment_id"`
	Assignment       Assignment `gorm:"foreignKey:AssignmentID;references:AssignmentID"`
	CreatedBy        uint32     `json:"created_by"`
	User             User       `gorm:"foreignKey:CreatedBy;references:UserID"`
	UpdatedBy        uint32     `json:"updated_by"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
