package models

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(RolePermission{})
	gob.Register([]RolePermission{})
}

type RolePermission struct {
	ID           uint32     `json:"id" gorm:"primaryKey"`
	RoleID       uint32     `json:"role_id"`
	Role         Role       `gorm:"foreignKey:RoleID;references:RoleID"`
	PermissionID uint32     `json:"permission_id"`
	Permission   Permission `gorm:"foreignKey:PermissionID;references:PermissionID"`
	CreatedBy    uint32     `json:"created_by"`
	UpdatedBy    uint32     `json:"updated_by"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
