package models

import "time"

type Program struct {
	ProgramID   uint32    `json:"program_id" gorm:"primaryKey"`
	ProgramCode string    `json:"program_code"`
	Name        string    `json:"name"`
	CreatedBy   uint32    `json:"created_by"`
	UpdatedBy   uint32    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
