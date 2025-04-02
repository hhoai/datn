package models

import "time"

type FileNews struct {
	FileNewsID uint32 `gorm:"primaryKey;autoIncrement" json:"file_news_id"`
	FileName   string `json:"file_name"`
	NewsID     uint32 `json:"news_id"`
	News       News   `gorm:"foreignKey:NewsID;references:NewsID"`
	Default    bool   `json:"default"`
	CreatedBy  uint32 `json:"created_by"`
	UpdatedBy  uint32 `json:"updated_by"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
