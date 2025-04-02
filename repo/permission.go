package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type PermissionRepository interface {
	FindAll() []*models.Permission
	FindByID(ID uint32) (*models.Permission, error)
	Create(result *models.Permission) error
}

type PermissionRepositoryImpl struct {
	connection *gorm.DB
}

func NewPermissionRepository() PermissionRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &PermissionRepositoryImpl{
		connection: DB,
	}
}

func (s *PermissionRepositoryImpl) FindAll() []*models.Permission {
	var result []*models.Permission
	if err := s.connection.Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *PermissionRepositoryImpl) FindByID(ID uint32) (*models.Permission, error) {
	var result models.Permission
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *PermissionRepositoryImpl) Create(result *models.Permission) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
