package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type LevelRepository interface {
	FindAll() []*models.Level
	FindByID(ID uint32) (*models.Level, error)
	Create(result *models.Level) error
	Delete(result *models.Level) error
	Update(level models.Level) error
}

type LevelRepositoryImpl struct {
	connection *gorm.DB
}

func NewLevelRepository() LevelRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &LevelRepositoryImpl{
		connection: DB,
	}
}

func (s *LevelRepositoryImpl) FindAll() []*models.Level {
	var result []*models.Level
	if err := s.connection.Model(&models.Level{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *LevelRepositoryImpl) FindByID(ID uint32) (*models.Level, error) {
	var result models.Level
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *LevelRepositoryImpl) Create(result *models.Level) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *LevelRepositoryImpl) Delete(level *models.Level) error {
	if err := s.connection.Delete(&level).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *LevelRepositoryImpl) Update(level models.Level) error {
	if err := s.connection.Save(&level).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
