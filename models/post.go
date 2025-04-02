package models

import "time"

type Post struct {
	PostID    uint32    `gorm:"primaryKey;autoIncrement" json:"post_id"`
	Title     string    `json:"post_title"`
	Body      string    `json:"post_body"`
	LessonID  uint32    `json:"lesson_id"`
	CreatedBy uint32    `json:"user_id"`
	User      User      `gorm:"foreignKey:CreatedBy;references:UserID"`
	UpdatedBy uint32    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
