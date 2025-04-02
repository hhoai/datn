package models

import "time"

type Banner struct {
	BannerID    uint32 `gorm:"primaryKey;autoIncrement" json:"banner_id"`
	FileName    string `json:"file_name"`
	Description string `json:"description"`
	CreatedBy   uint32 `json:"created_by"`
	UpdatedBy   uint32 `json:"updated_by"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
