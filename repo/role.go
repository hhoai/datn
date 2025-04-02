package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type RoleRepository interface {
	FindAll() []*models.Role
	FindByID(ID uint32) (*models.Role, error)
	FindByRoleID(roleID uint32) ([]*models.Role, error)
	Create(result *models.Role) error
	ExistsByID(ID uint32) bool
	Delete(role *models.Role) error
	Update(models.Role) error
}

type RoleRepositoryImpl struct {
	connection *gorm.DB
}

func NewRoleRepository() RoleRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &RoleRepositoryImpl{
		connection: DB,
	}
}

func (s *RoleRepositoryImpl) FindAll() []*models.Role {
	var result []*models.Role
	if err := s.connection.Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *RoleRepositoryImpl) ExistsByID(ID uint32) bool {
	var count int64

	if err := s.connection.Model(&models.Role{}).Where("id = ?", ID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
		return false
	}
	return count > 0
}

func (s *RoleRepositoryImpl) FindByID(ID uint32) (*models.Role, error) {
	var result models.Role
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *RoleRepositoryImpl) FindByRoleID(roleID uint32) ([]*models.Role, error) {
	var result []*models.Role
	if err := s.connection.Model(&models.Role{}).
		Joins("Permission").
		Where("role_permissions.role_id", roleID).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *RoleRepositoryImpl) Create(result *models.Role) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *RoleRepositoryImpl) Delete(role *models.Role) error {
	if err := s.connection.Delete(&role).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *RoleRepositoryImpl) DeleteByID(roleID uint32) error {
	role := models.Role{}
	role.RoleID = roleID
	if err := s.connection.Delete(&role).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *RoleRepositoryImpl) Update(role models.Role) error {
	if err := s.connection.Save(&role).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
