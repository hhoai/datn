package models

import "time"

type UserLesson struct {
	UserLessonID uint32    `gorm:"primary_key;AUTO_INCREMENT"`
	LessonID     uint32    `json:"lesson_id"`
	Lesson       Lesson    `gorm:"foreignKey:lesson_id;references:lesson_id"`
	UserID       uint32    `json:"user_id"`
	User         User      `gorm:"foreignKey:user_id;references:user_id"`
	CompletedAt  time.Time `json:"completed_at"`
	Status       bool      `json:"status"`
}
