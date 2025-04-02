package models

import "time"

type Comment struct {
	CommentID      uint32 `gorm:"primaryKey;autoIncrement"`
	CreatedBy      uint32 `json:"user_id"`
	User           User   `gorm:"foreignKey:CreatedBy;references:UserID"`
	CommentContent string `json:"comment_content"`
	PostID         uint32 `json:"post_id"`
	Post           Post   `gorm:"foreignKey:PostID;references:PostID"`
	NewsID         uint32 `json:"news_id"`
	News           News   `gorm:"foreignKey:NewsID;references:NewsID"`
	UpdatedBy      uint32 `json:"updated_by"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
