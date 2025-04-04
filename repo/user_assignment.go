package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type UserAssignmentRepository interface {
	FindAll() []*models.UserAssignment
	FindByAssignmentID(courseID uint32) ([]*models.UserAssignment, error)
	FindByUserAssignmentID(userAssignmentID uint32) (*models.UserAssignment, error)
	FindByUserID(userID uint32) ([]*models.UserAssignment, error)
	Create(requestCourse models.UserAssignment) error
	CreateAny(requestCourse []*models.UserAssignment) error
	Update(userAssignment models.UserAssignment) error
	Delete(requestCourse models.UserAssignment) error
	FindByUserIDAndAssignmentID(assignmentID, userID uint32) (*models.UserAssignment, error)
	FindByUserIDAndAssignmentIDs(assignmentIDs []uint32, userID uint32) ([]*models.UserAssignment, error)
	FindCompletedAssignment(userid uint32) ([]*models.LearningProcess, error)
	CountAssignmentCompletedInCourse(lessonID, userID uint32) (uint32, error)
	FindByUserIDAndLessonID(lessonID, userID uint32) ([]*models.UserAssignment, error)
	FindFileAssignmentsByUserIDAndAssignmentID(userID, assignmentID uint32) ([]*models.FileAssignment, error)
	CountAssignmentCompletedInLesson(lessonID, userID uint32) (uint32, error)
}

type UserAssignmentRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserAssignmentRepository() UserAssignmentRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &UserAssignmentRepositoryImpl{
		connection: DB,
	}
}

func (s *UserAssignmentRepositoryImpl) FindAll() []*models.UserAssignment {
	var result []*models.UserAssignment
	if err := s.connection.Model(&models.UserAssignment{}).Preload("User").Preload("Assignment").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *UserAssignmentRepositoryImpl) FindByAssignmentID(ID uint32) ([]*models.UserAssignment, error) {
	var result []*models.UserAssignment
	if err := s.connection.Model(&models.UserAssignment{}).Preload("User").Preload("Assignment").Where("assignment_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserAssignmentRepositoryImpl) FindByUserID(ID uint32) ([]*models.UserAssignment, error) {
	var result []*models.UserAssignment
	if err := s.connection.Model(&models.UserAssignment{}).Preload("User").Preload("Assignment").Where("user_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserAssignmentRepositoryImpl) FindByUserAssignmentID(userAssignmentID uint32) (*models.UserAssignment, error) {
	var result *models.UserAssignment
	if err := s.connection.Where("user_assignment_id = ?", userAssignmentID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserAssignmentRepositoryImpl) Create(req models.UserAssignment) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserAssignmentRepositoryImpl) CreateAny(req []*models.UserAssignment) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserAssignmentRepositoryImpl) Update(userAssignment models.UserAssignment) error {
	if err := s.connection.Save(&userAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserAssignmentRepositoryImpl) Delete(req models.UserAssignment) error {
	if err := s.connection.Delete(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserAssignmentRepositoryImpl) FindByUserIDAndAssignmentID(assignmentID, userID uint32) (*models.UserAssignment, error) {
	var result *models.UserAssignment
	if err := s.connection.Preload("User").Preload("Assignment").Where("assignment_id = ? AND user_id = ?", assignmentID, userID).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *UserAssignmentRepositoryImpl) FindByUserIDAndAssignmentIDs(assignmentIDs []uint32, userID uint32) ([]*models.UserAssignment, error) {
	var result []*models.UserAssignment
	if err := s.connection.Preload("User").
		Preload("Assignment").
		Where("assignment_id IN (?) AND user_id = ?", assignmentIDs, userID).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *UserAssignmentRepositoryImpl) FindByUserIDAndLessonID(lessonID, userID uint32) ([]*models.UserAssignment, error) {
	var result []*models.UserAssignment
	if err := s.connection.Table("user_assignments").Preload("Assignment").
		Joins("JOIN assignments a ON a.assignment_id = user_assignments.assignment_id").
		Where("a.lesson_id = ? AND user_assignments.user_id = ?", lessonID, userID).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *UserAssignmentRepositoryImpl) FindCompletedAssignment(userid uint32) ([]*models.LearningProcess, error) {
	var result []*models.LearningProcess

	if err := s.connection.Table("user_assignments").
		Select("courses.course_id AS course_id, courses.title AS course_name, lessons.lesson_id AS lesson_id, lessons.title AS lesson_name, "+
			"user_assignments.completed_at AS timeline_date, user_assignments.completed_at AS submit_at, user_assignments.assignment_id as assignment_id, assignments.title as assignment_name").
		Joins("JOIN assignments ON assignments.assignment_id = user_assignments.assignment_id").
		Joins("JOIN users ON users.user_id = user_assignments.user_id").
		Joins("JOIN lessons ON lessons.lesson_id = assignments.lesson_id").
		Joins("JOIN courses ON courses.course_id = lessons.course_id").
		Where("user_assignments.status = ? AND user_assignments.user_id = ?", 1, userid).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserAssignmentRepositoryImpl) CountAssignmentCompletedInCourse(lessonID, userID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Table("user_assignments").
		Joins("JOIN assignments on assignments.assignment_id = user_assignments.assignment_id").
		Where("assignments.lesson_id = ? AND user_assignments.status = ? AND user_assignments.user_id = ?", lessonID, 1, userID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
func (s *UserAssignmentRepositoryImpl) FindFileAssignmentsByUserIDAndAssignmentID(userID, assignmentID uint32) ([]*models.FileAssignment, error) {
	var results []*models.FileAssignment

	if err := s.connection.Table("file_assignments").
		Joins("JOIN user_assignments ON file_assignments.assignment_id = user_assignments.assignment_id").
		Where("user_assignments.user_id = ? AND user_assignments.assignment_id = ?", userID, assignmentID).
		Preload("Assignment").
		Preload("User").
		Find(&results).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return results, nil
}

func (s *UserAssignmentRepositoryImpl) CountAssignmentCompletedInLesson(lessonID, userID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Table("user_assignments").
		Joins("JOIN assignments on assignments.assignment_id = user_assignments.assignment_id").
		Where("assignments.lesson_id = ? AND user_assignments.status = ? AND user_assignments.user_id = ?", lessonID, true, userID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
