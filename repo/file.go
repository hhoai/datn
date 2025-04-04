package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type FileAssignmentRepository interface {
	FindAll() []*models.FileAssignment
	Create(fileAssignment models.FileAssignment) error
	Delete(fileAssignment models.FileAssignment) error
	Update(fileAssignment models.FileAssignment) error
	DeleteByID(fileAssignmentID uint32) error
	UpdateByID(fileAssignmentID uint32) error
	FindByID(fileAssignmentID uint32) (*models.FileAssignment, error)
	FindByAssignmentID(fileAssignmentID uint32) ([]*models.FileAssignment, error)
	DeleteMultiple(fileAssignmentID []uint32) error
	FindByAssignmentIDAndUserID(assignmentID, userID uint32) ([]*models.FileAssignment, error)
}

type FileAssignmentRepositoryImpl struct {
	connection *gorm.DB
}

func NewFileAssignmentRepository() FileAssignmentRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &FileAssignmentRepositoryImpl{
		connection: DB,
	}
}

func (s *FileAssignmentRepositoryImpl) FindAll() []*models.FileAssignment {
	var result []*models.FileAssignment
	if err := s.connection.Model(&models.FileAssignment{}).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *FileAssignmentRepositoryImpl) Create(fileAssignment models.FileAssignment) error {
	if err := s.connection.Create(&fileAssignment).Error; err != nil {
		OutPutDebugError("Fail to create: " + err.Error())
		return err
	}
	return nil
}

func (s *FileAssignmentRepositoryImpl) Update(fileAssignment models.FileAssignment) error {
	if err := s.connection.Save(&fileAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *FileAssignmentRepositoryImpl) Delete(fileAssignment models.FileAssignment) error {
	if err := s.connection.Delete(&fileAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FileAssignmentRepositoryImpl) DeleteByID(fileAssignmentID uint32) error {
	fileAssignment := models.FileAssignment{}
	fileAssignment.FileAssignmentID = fileAssignmentID
	if err := s.connection.Delete(&fileAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FileAssignmentRepositoryImpl) UpdateByID(fileAssignmentID uint32) error {
	fileAssignment := models.FileAssignment{}
	fileAssignment.FileAssignmentID = fileAssignmentID
	if err := s.connection.Save(&fileAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FileAssignmentRepositoryImpl) FindByID(fileAssignmentID uint32) (*models.FileAssignment, error) {
	var fileAssignment models.FileAssignment

	if err := s.connection.Where("file_assignment_id", fileAssignmentID).Find(&fileAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &fileAssignment, nil
}

func (s *FileAssignmentRepositoryImpl) FindByAssignmentID(assignmentID uint32) ([]*models.FileAssignment, error) {
	var result []*models.FileAssignment
	if err := s.connection.Model(&models.FileAssignment{}).Where("assignment_id", assignmentID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *FileAssignmentRepositoryImpl) DeleteMultiple(fileAssignmentID []uint32) error {

	if err := s.connection.Where("file_assignment_id IN ?", fileAssignmentID).Delete(&models.FileAssignment{}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FileAssignmentRepositoryImpl) FindByAssignmentIDAndUserID(assignmentID, userID uint32) ([]*models.FileAssignment, error) {
	var result []*models.FileAssignment
	if err := s.connection.Model(&models.FileAssignment{}).Preload("Assignment").Preload("User").Where("assignment_id = ? and created_by = ?", assignmentID, userID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}
