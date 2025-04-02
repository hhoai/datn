package models

import "time"

type Question struct {
	QuestionID     uint32       `gorm:"primaryKey;autoIncrement" json:"question_id"`
	TypeQuestionID uint32       `json:"type_question_id"`
	TypeQuestion   TypeQuestion `gorm:"foreignKey:TypeQuestionID ;references:TypeQuestionID"`
	Content        string       `json:"content"`
	Score          uint32       `json:"score"`
	CreatedBy      uint32
	UpdatedBy      uint32
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ProgramID      uint32       `json:"program_id"`
	Program        Program      `gorm:"foreignKey:ProgramID;references:ProgramID"`
	LevelID        uint32       `json:"level_id"`
	Level          Level        `gorm:"foreignKey:LevelID;references:LevelID"`
	ChallengeID    uint32       `json:"challenge_id"`
	Challenge      Challenge    `gorm:"foreignKey:ChallengeID;references:ChallengeID"`
	SkillID        uint32       `json:"skill_id"`
	Skill          Skill        `gorm:"foreignKey:SkillID;references:SkillID"`
	Options        []Option     `gorm:"foreignKey:QuestionID;references:QuestionID"`
	UserOptions    []UserOption `gorm:"foreignKey:QuestionID;references:QuestionID"`
}
