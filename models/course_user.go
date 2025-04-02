package models

import (
	"time"
)

type CourseUser struct {
	CourseUserID uint32 `gorm:"primaryKey;autoIncrement"`
	UserID       uint32 `json:"user_id"`
	User         User   `gorm:"foreignKey:UserID;references:UserID"`
	CourseID     uint32 `json:"course_id"`
	Course       Course `gorm:"foreignKey:CourseID;references:CourseID"`
	Status       bool   `json:"status"`
	CreatedBy    uint32 `json:"created_by"`
	UpdatedBy    uint32 `json:"updated_by"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CompletedAt  time.Time
}
