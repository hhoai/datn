package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type AssignmentRepository interface {
	FindAll() []*models.Assignment
	Create(assignment models.Assignment) error
	Delete(models.Assignment) error
	Update(models.Assignment) error
	DeleteByID(assignmentID uint32) error
	UpdateByID(assignmentID uint32) error
	FindByID(assignmentID uint32) (*models.Assignment, error)
	FindInLessonID(lessonID []uint32) ([]*models.Assignment, error)
	DeleteMultiple(assignmentID []uint32) error
	FindByLessonID(lessonID uint32) ([]*models.Assignment, error)
	FindAssignmentIDByLessonID(lessonID uint32) ([]uint32, error)
	GetLastAssignment() (*models.Assignment, error)
	CountAssignmentInCourse(lessonID uint32) (uint32, error)
	CountAssignmentInLesson(lessonID uint32) (uint32, error)
}

type AssignmentRepositoryImpl struct {
	connection *gorm.DB
}

func NewAssignmentRepository() AssignmentRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &AssignmentRepositoryImpl{
		connection: DB,
	}
}

func (s *AssignmentRepositoryImpl) FindAll() []*models.Assignment {
	var result []*models.Assignment
	if err := s.connection.Model(&models.Assignment{}).Preload("TypeAssignment").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *AssignmentRepositoryImpl) Create(assignment models.Assignment) error {
	var typeAssignment models.TypeAssignment
	if err := s.connection.First(&typeAssignment, assignment.TypeAssignmentID).Error; err != nil {
		OutPutDebugError("type assignment not exit: " + err.Error())
		return err
	}
	if err := s.connection.Create(&assignment).Error; err != nil {
		OutPutDebugError("Fail to create: " + err.Error())
		return err
	}
	return nil
}

func (s *AssignmentRepositoryImpl) Update(assignment models.Assignment) error {
	if err := s.connection.Save(&assignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *AssignmentRepositoryImpl) Delete(assignment models.Assignment) error {
	if err := s.connection.Delete(&assignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *AssignmentRepositoryImpl) CountAssignmentInLesson(lessonID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Assignment{}).Where("lesson_id = ?", lessonID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}

func (s *AssignmentRepositoryImpl) DeleteByID(assignmentID uint32) error {
	assignment := models.Assignment{}
	assignment.AssignmentID = assignmentID
	if err := s.connection.Delete(&assignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *AssignmentRepositoryImpl) UpdateByID(assignmentID uint32) error {
	assignment := models.Assignment{}
	assignment.AssignmentID = assignmentID
	if err := s.connection.Save(&assignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *AssignmentRepositoryImpl) FindByID(assignmentID uint32) (*models.Assignment, error) {
	var assignment models.Assignment
	if err := s.connection.Model(&models.Assignment{}).First(&assignment, assignmentID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	if err := s.connection.Preload("TypeAssignment").Preload("User").Preload("Lesson").Find(&assignment).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &assignment, nil
}

func (s *AssignmentRepositoryImpl) FindInLessonID(lessonID []uint32) ([]*models.Assignment, error) {
	var assignments []*models.Assignment
	if err := s.connection.Model(&models.Assignment{}).
		Where("lesson_id IN (?)", lessonID).
		Find(&assignments).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return assignments, nil
}

func (s *AssignmentRepositoryImpl) DeleteMultiple(assignmentID []uint32) error {

	if err := s.connection.Where("assignment_id IN ?", assignmentID).Delete(&models.Assignment{}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *AssignmentRepositoryImpl) FindByLessonID(lessonID uint32) ([]*models.Assignment, error) {
	var result []*models.Assignment
	if err := s.connection.Model(&models.Assignment{}).Preload("TypeAssignment").Where("lesson_id = ?", lessonID).Order("due_date desc").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *AssignmentRepositoryImpl) FindAssignmentIDByLessonID(lessonID uint32) ([]uint32, error) {
	var result []uint32
	if err := s.connection.Table("assignments").
		Where("lesson_id = ?", lessonID).
		Order("due_date desc").Pluck("assignment_id", &result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *AssignmentRepositoryImpl) GetLastAssignment() (*models.Assignment, error) {
	var assignment models.Assignment
	if err := s.connection.Model(&models.Assignment{}).Preload("TypeAssignment").Order("assignment_id desc").First(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (s *AssignmentRepositoryImpl) CountAssignmentInCourse(lessonID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Assignment{}).Where("lesson_id = ?", lessonID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
