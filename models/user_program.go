package models

import "time"

type UserProgram struct {
	UserProgramID uint32    `gorm:"primary_key;AUTO_INCREMENT"`
	ProgramID     uint32    `json:"program_id"`
	Program       Program   `gorm:"foreignKey:program_id;references:program_id"`
	UserID        uint32    `json:"user_id"`
	User          User      `gorm:"foreignKey:user_id;references:user_id"`
	CompletedAt   time.Time `json:"completed_at"`
	Status        bool      `json:"status"`
}
