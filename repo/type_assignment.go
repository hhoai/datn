package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type TypeAssignmentRepository interface {
	FindAll() []*models.TypeAssignment
}

type TypeAssignmentRepositoryImpl struct {
	connection *gorm.DB
}

func NewTypeAssignmentRepository() TypeAssignmentRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &TypeAssignmentRepositoryImpl{
		connection: DB,
	}
}

func (s *TypeAssignmentRepositoryImpl) FindAll() []*models.TypeAssignment {
	var result []*models.TypeAssignment
	if err := s.connection.Model(&models.TypeAssignment{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
