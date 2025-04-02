package models

import "time"

type CourseCategory struct {
	CourseCategoriesID uint32 `gorm:"primaryKey" json:"course_categories_id"`
	Name               string `json:"name"`
	Description        string `json:"description"`
	CreatedBy          uint32 `json:"created_by"`
	UpdatedBy          uint32 `json:"updated_by"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
