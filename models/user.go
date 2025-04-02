package models

import (
	"encoding/gob"
	"time"
)

func init() {
	gob.Register(User{})
	gob.Register(UserWithoutPass{})
}

type User struct {
	UserID      uint32   `json:"user_id" gorm:"primaryKey;autoIncrement" `
	TypeUserID  uint32   `json:"type_user_id"`
	TypeUser    TypeUser `gorm:"foreignKey:TypeUserID;references:TypeUserID"`
	RoleID      uint32   `json:"role_id" `
	Role        Role     `gorm:"foreignKey:RoleID;references:RoleID"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Vault       string   `json:"vault"`
	IsActivated bool
	Token       string
	TokenExpiry time.Time
	TypeAccount string
	Password    string `json:"password"`
	SessionID   string
	Status      bool
	CreatedBy   uint32    `json:"created_by"`
	UpdatedBy   uint32    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserWithoutPass struct {
	UserID      uint32   `json:"user_id" gorm:"primaryKey" `
	TypeUserID  uint32   `json:"type_user_id"`
	TypeUser    TypeUser `gorm:"foreignKey:TypeUserID;references:TypeUserID"`
	RoleID      uint32   `json:"role_id"`
	Role        Role     `gorm:"foreignKey:RoleID;references:RoleID"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	IsActivated bool
	TypeAccount string
	Token       string
	SessionID   string
	Status      bool
	TokenExpiry time.Time
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   uint32    `json:"created_by"`
	UpdatedBy   uint32    `json:"updated_by"`
}

type LearningProcess struct {
	CourseID       uint32    `json:"course_id"`
	CourseName     string    `json:"course_name"`
	LessonID       uint32    `json:"lesson_id"`
	LessonName     string    `json:"lesson_name"`
	AssignmentID   uint32    `json:"assignment_id"`
	AssignmentName string    `json:"assignment_name"`
	TimelineDate   time.Time `json:"timeline_date"`
	SubmitAt       time.Time `json:"submit_at"`
}
