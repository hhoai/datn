package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type TypeUserRepository interface {
	FindAll() []*models.TypeUser
	FindByID(ID uint32) (*models.TypeUser, error)
	Create(result *models.TypeUser) error
	Delete(result *models.TypeUser) error
	Update(user models.TypeUser) error
}

type TypeUserRepositoryImpl struct {
	connection *gorm.DB
}

func NewTypeUserRepository() TypeUserRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &TypeUserRepositoryImpl{
		connection: DB,
	}
}

func (s *TypeUserRepositoryImpl) FindAll() []*models.TypeUser {
	var result []*models.TypeUser
	if err := s.connection.Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *TypeUserRepositoryImpl) FindByID(ID uint32) (*models.TypeUser, error) {
	var result models.TypeUser
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *TypeUserRepositoryImpl) Create(result *models.TypeUser) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *TypeUserRepositoryImpl) Delete(typeUser *models.TypeUser) error {
	if err := s.connection.Delete(&typeUser).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *TypeUserRepositoryImpl) Update(typeUser models.TypeUser) error {
	if err := s.connection.Save(&typeUser).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
