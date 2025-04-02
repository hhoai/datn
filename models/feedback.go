package models

import "time"

type Feedback struct {
	FeedbackID uint32 `json:"feedback_id" gorm:"primary_key;AUTO_INCREMENT"`
	Feedback   string `json:"feedback"`
	CreatedBy  uint32 `json:"user_id"`
	User       User   `gorm:"foreignkey:CreatedBy; references:UserID"`
	CreatedAt  time.Time
}
