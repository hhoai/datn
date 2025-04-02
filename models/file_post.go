package models

import "time"

type FilePost struct {
	FilePostID uint32 `gorm:"primaryKey;autoIncrement"`
	Filename   string `json:"file_name"`
	PostID     uint32 `json:"post_id"`
	Post       Post   `gorm:"foreignKey:PostID;references:PostID"`
	Default    bool   `json:"default"`
	CreatedBy  uint32 `json:"created_by"`
	UpdatedBy  uint32 `json:"updated_by"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
