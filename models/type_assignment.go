package models

type TypeAssignment struct {
	TypeAssignmentID uint32 `gorm:"primaryKey,autoIncrement" json:"type_assignment_id"`
	Name             string `json:"name"`
}
