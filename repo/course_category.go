package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type CourseCategoryRepository interface {
	FindAll() []*models.CourseCategory
	FindByID(ID uint32) (*models.CourseCategory, error)
	Create(result *models.CourseCategory) error
	Delete(result *models.CourseCategory) error
	Update(category models.CourseCategory) error
}

type CourseCategoryRepositoryImpl struct {
	connection *gorm.DB
}

func NewCourseCategoryRepository() CourseCategoryRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &CourseCategoryRepositoryImpl{
		connection: DB,
	}
}

func (s *CourseCategoryRepositoryImpl) FindAll() []*models.CourseCategory {
	var result []*models.CourseCategory
	if err := s.connection.Model(&models.CourseCategory{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *CourseCategoryRepositoryImpl) FindByID(ID uint32) (*models.CourseCategory, error) {
	var result models.CourseCategory
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *CourseCategoryRepositoryImpl) Create(result *models.CourseCategory) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *CourseCategoryRepositoryImpl) Delete(courseCategory *models.CourseCategory) error {
	if err := s.connection.Delete(&courseCategory).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *CourseCategoryRepositoryImpl) Update(courseCategory models.CourseCategory) error {
	if err := s.connection.Save(&courseCategory).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
