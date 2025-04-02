package models

import "time"

type Skill struct {
	SkillID   uint32 `json:"skill_id" gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedBy uint32
	UpdatedBy uint32
	CreatedAt time.Time
	UpdatedAt time.Time
}
