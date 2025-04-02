package models

type TypeUser struct {
	TypeUserID uint32 `json:"type_user_id" gorm:"primaryKey"`
	Name       string `json:"name"`
}
