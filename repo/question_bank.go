package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type QuestionBankRepository interface {
	FindAll() []*models.Question
	Create(question models.Question) (error, models.Question)
	Delete(models.Question) error
	Update(models.Question) error
	FindByID(questionID uint32) (*models.Question, error)
	UpdateByID(questionID uint32) error
	DeleteByID(questionID uint32) error
	Search(programID, SkillID, ChallengeID, TypeQuestionID, LevelID uint32, questionIDs []uint32) ([]*models.Question, error)
	CountQuestion() (uint32, error)
	CheckQuestionExist(content string) (bool, error)
}

type QuestionBankRepositoryImpl struct {
	connection *gorm.DB
}

func NewQuestionBankRepository() QuestionBankRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &QuestionBankRepositoryImpl{
		connection: DB,
	}
}

func (s *QuestionBankRepositoryImpl) FindAll() []*models.Question {
	var result []*models.Question
	if err := s.connection.Model(&models.Question{}).Preload("Level").Preload("Skill").Preload("Program").Preload("Challenge").Preload("TypeQuestion").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *QuestionBankRepositoryImpl) Create(question models.Question) (error, models.Question) {
	if err := s.connection.Create(&question).Error; err != nil {
		OutPutDebugError(err.Error())
		return err, models.Question{}
	}

	return nil, question
}
func (s *QuestionBankRepositoryImpl) CheckQuestionExist(content string) (bool, error) {
	var count int64
	if err := s.connection.Model(&models.Question{}).Where("content = ?", content).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
		return false, err
	}
	return count > 0, nil
}

func (s *QuestionBankRepositoryImpl) Delete(question models.Question) error {
	if err := s.connection.Delete(&question).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *QuestionBankRepositoryImpl) Update(question models.Question) error {
	if err := s.connection.Save(&question).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *QuestionBankRepositoryImpl) FindByID(questionID uint32) (*models.Question, error) {
	var result models.Question

	if err := s.connection.Model(&models.Question{}).Preload("Level").Preload("Skill").Preload("Program").Preload("Challenge").Preload("TypeQuestion").First(&result, questionID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *QuestionBankRepositoryImpl) DeleteByID(questionID uint32) error {
	question := models.Question{}
	question.QuestionID = questionID
	if err := s.connection.Delete(&question).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *QuestionBankRepositoryImpl) UpdateByID(questionID uint32) error {
	question := models.Question{}
	question.QuestionID = questionID
	if err := s.connection.Save(&question).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *QuestionBankRepositoryImpl) Search(programID, skillID, challengeID, typeQuestionID, levelID uint32, questionIDs []uint32) ([]*models.Question, error) {
	var results []*models.Question

	query := s.connection.Model(&models.Question{}).Preload("Level").Preload("Skill").Preload("Program").Preload("Challenge").Preload("TypeQuestion")

	if programID != 0 {
		query = query.Where("program_id = ?", programID)
	}
	if skillID != 0 {
		query = query.Where("skill_id = ?", skillID)
	}
	if challengeID != 0 {
		query = query.Where("challenge_id = ?", challengeID)
	}
	if typeQuestionID != 0 {
		query = query.Where("type_question_id = ?", typeQuestionID)
	}
	if levelID != 0 {
		query = query.Where("level_id = ?", levelID)
	}

	if len(questionIDs) > 0 {
		query = query.Where("question_id NOT IN (?)", questionIDs)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (s *QuestionBankRepositoryImpl) CountQuestion() (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Question{}).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
