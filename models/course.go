package models

import "time"

type Course struct {
	CourseID             uint32         `gorm:"primaryKey;autoIncrement" json:"course_id"`
	CourseCode           string         `json:"code"`
	Title                string         `json:"title"`
	Description          string         `json:"description"`
	LevelID              uint32         `json:"level_id"`
	Level                Level          `gorm:"foreignKey:LevelID;references:LevelID"`
	Image                string         `json:"image"`
	Amount               uint32         `json:"amount"`
	CourseCategory       CourseCategory `gorm:"foreignKey:CourseCategoriesID;references:CourseCategoriesID"`
	CourseCategoriesID   uint32         `json:"course_categories_id"`
	CreatedBy            uint32         `gorm:"foreignKey:CreatedBy;references:UserID"`
	UpdatedBy            uint32         `gorm:"foreignKey:UpdatedBy;references:UserID"`
	StartTime            time.Time      `json:"start_time"`
	EndTime              time.Time      `json:"end_time"`
	Status               bool           `json:"status"`
	CreatedAt            time.Time      `json:"created_at"`
	UpdatedAt            time.Time      `json:"updated_at"`
	Program              Program        `gorm:"foreignKey:ProgramID;references:ProgramID"`
	ProgramID            uint32         `json:"program_id"`
	PrerequisiteCourseID uint32         `json:"prerequisite_course_id"`
}
