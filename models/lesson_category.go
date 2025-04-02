package models

import "time"

type LessonCategory struct {
	LessonCategoryID uint32    `json:"lesson_category_id" gorm:"primaryKey"`
	Name             string    `json:"name"`
	CreatedBy        uint32    `json:"created_by"`
	UpdatedBy        uint32    `json:"updated_by"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
