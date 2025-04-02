package models

import "time"

type Lesson struct {
	LessonID         uint32         `gorm:"primaryKey;autoIncrement"`
	CourseID         uint32         `json:"course_id"`
	Course           Course         `gorm:"foreignKey:CourseID;references:CourseID"`
	Title            string         `json:"title"`
	LevelID          uint32         `json:"level_id"`
	Level            Level          `gorm:"foreignKey:LevelID;references:LevelID"`
	LessonCategoryID uint32         `json:"lesson_category_id"`
	LessonCategory   LessonCategory `gorm:"foreignKey:LessonCategoryID;references:LessonCategoryID"`
	CreatedBy        uint32         `json:"created_by"`
	UpdatedBy        uint32         `json:"updated_by"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
