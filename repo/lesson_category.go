package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type LessonCategoryRepository interface {
	FindAll() []*models.LessonCategory
	FindByID(ID uint32) (*models.LessonCategory, error)
	Create(result *models.LessonCategory) error
	Delete(result *models.LessonCategory) error
	Update(category models.LessonCategory) error
}

type LessonCategoryRepositoryImpl struct {
	connection *gorm.DB
}

func NewLessonCategoryRepository() LessonCategoryRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &LessonCategoryRepositoryImpl{
		connection: DB,
	}
}

func (s *LessonCategoryRepositoryImpl) FindAll() []*models.LessonCategory {
	var result []*models.LessonCategory
	if err := s.connection.Model(&models.LessonCategory{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *LessonCategoryRepositoryImpl) FindByID(ID uint32) (*models.LessonCategory, error) {
	var result models.LessonCategory
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *LessonCategoryRepositoryImpl) Create(result *models.LessonCategory) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *LessonCategoryRepositoryImpl) Delete(lessonCategory *models.LessonCategory) error {
	if err := s.connection.Delete(&lessonCategory).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *LessonCategoryRepositoryImpl) Update(lessonCategory models.LessonCategory) error {
	if err := s.connection.Save(&lessonCategory).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
