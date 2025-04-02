package models

import "time"

type Faqs struct {
	FaqsID    uint32    `json:"faqs_id" gorm:"primary_key;AUTO_INCREMENT"`
	Question  string    `json:"question"`
	Answer    string    `json:"answer"`
	IsDisplay bool      `json:"is_display"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy uint32    `json:"created_by"`
	UpdatedBy uint32    `json:"updated_by"`
}
