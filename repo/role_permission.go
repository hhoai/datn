package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type RolePerRepository interface {
	FindAll() []*models.RolePermission
	FindByID(ID uint32) (*models.RolePermission, error)
	FindByRoleID(roleID uint32) ([]models.RolePermission, error)
	Create(result *models.RolePermission) error
	CreateAny(result []*models.RolePermission) error
	DeleteByRoleID(roleID uint32) error
}

type RolePerRepositoryImpl struct {
	connection *gorm.DB
}

func NewRolePerRepository() RolePerRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &RolePerRepositoryImpl{
		connection: DB,
	}
}

func (s *RolePerRepositoryImpl) FindAll() []*models.RolePermission {
	var result []*models.RolePermission
	if err := s.connection.Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *RolePerRepositoryImpl) FindByID(ID uint32) (*models.RolePermission, error) {
	var result models.RolePermission
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *RolePerRepositoryImpl) DeleteByRoleID(roleID uint32) error {

	if err := s.connection.Where("role_id = ?", roleID).Delete(&models.RolePermission{}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *RolePerRepositoryImpl) FindByRoleID(roleID uint32) ([]models.RolePermission, error) {
	var result []models.RolePermission
	if err := s.connection.Model(&models.RolePermission{}).
		Joins("Permission").
		Where("role_permissions.role_id", roleID).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *RolePerRepositoryImpl) Create(result *models.RolePermission) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *RolePerRepositoryImpl) CreateAny(result []*models.RolePermission) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *RolePerRepositoryImpl) Delete(result *models.RolePermission) error {
	if err := s.connection.Delete(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
