package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type SkillRepository interface {
	FindAll() []*models.Skill
	FindByID(ID uint32) (*models.Skill, error)
	Create(result *models.Skill) error
	Delete(result *models.Skill) error
	Update(models.Skill) error
}

type SkillRepositoryImpl struct {
	connection *gorm.DB
}

func NewSkillRepository() SkillRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &SkillRepositoryImpl{
		connection: DB,
	}
}

func (s *SkillRepositoryImpl) FindAll() []*models.Skill {
	var result []*models.Skill
	if err := s.connection.Model(&models.Skill{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *SkillRepositoryImpl) FindByID(ID uint32) (*models.Skill, error) {
	var result models.Skill
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *SkillRepositoryImpl) Create(result *models.Skill) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *SkillRepositoryImpl) Delete(skill *models.Skill) error {
	if err := s.connection.Delete(&skill).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *SkillRepositoryImpl) Update(skill models.Skill) error {
	if err := s.connection.Save(&skill).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
