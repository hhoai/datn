package models

import "time"

type Level struct {
	LevelID   uint32 `json:"level_id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedBy uint32
	UpdatedBy uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
