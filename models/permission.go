package models

type Permission struct {
	PermissionID uint32 `json:"permission_id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Permission   string `json:"permission"`
}
