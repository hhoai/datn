package models

import "time"

type News struct {
	NewsID    uint32    `gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"news_title"`
	Body      string    `json:"news_body" gorm:"type:text"`
	CreatedBy uint32    `json:"user_id"`
	User      User      `gorm:"foreignKey:CreatedBy;references:UserID"`
	UpdatedBy uint32    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
