package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type CourseUserRepository interface {
	FindAll() []*models.CourseUser
	FindUserByCourseID(courseID, typeUserID uint32) []*models.CourseUser
	FindCourseByUserID(userID, typeUserID uint32) []*models.CourseUser
	AddUserToCourse(courseID uint32, userID uint32) error
	AddStudentToCourse(courseID uint32, userID []uint32) error
	RemoveUserFromCourse(courseID, userID uint32) error
	FindUsersNotInCourse(courseID, typeUser uint32) ([]models.User, error)
	Create(courseUser models.CourseUser) error
	CreateBatch(courseUsers []models.CourseUser) error
	FindCourseIDByProgramID(programID uint32, userID uint32) ([]uint32, error)
	Update(courseUser models.CourseUser) error
	FindByUserIDAndCourseID(courseID, userID uint32) (*models.CourseUser, error)
	FindCompletedCourse(userid uint32) ([]*models.CourseUser, error)
}

type CourseUserRepositoryImpl struct {
	connection *gorm.DB
}

func NewCourseUserRepository() CourseUserRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &CourseUserRepositoryImpl{
		connection: DB,
	}
}

func (s *CourseUserRepositoryImpl) FindAll() []*models.CourseUser {
	var result []*models.CourseUser
	if err := s.connection.Model(&models.CourseUser{}).Preload("User").Preload("Course").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *CourseUserRepositoryImpl) FindUserByCourseID(courseID, typeUser uint32) []*models.CourseUser {
	var result []*models.CourseUser

	if err := s.connection.
		Joins("JOIN users b ON course_users.user_id = b.user_id").
		Preload("User").
		Joins("JOIN type_users c ON c.type_user_id = b.type_user_id").
		Where("c.type_user_id = ? AND course_users.course_id = ?", typeUser, courseID).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil
	}

	return result
}
func (s *CourseUserRepositoryImpl) FindCourseByUserID(userID, typeUserID uint32) []*models.CourseUser {
	var result []*models.CourseUser

	err := s.connection.
		Joins("JOIN users b ON course_users.user_id = b.user_id").
		Preload("User").Preload("Course").
		Joins("JOIN type_users c ON c.type_user_id = b.type_user_id").
		Where("c.type_user_id = ? AND course_users.user_id = ?", typeUserID, userID).
		Find(&result).Error

	if err != nil {
		OutPutDebugError(err.Error())
		return nil
	}

	return result

}

func (s *CourseUserRepositoryImpl) AddUserToCourse(courseID, userID uint32) error {
	courseUser := models.CourseUser{
		CourseID: courseID,
		UserID:   userID,
	}
	if err := s.connection.Create(&courseUser).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseUserRepositoryImpl) AddStudentToCourse(courseID uint32, userIDs []uint32) error {
	var courseUsers []models.CourseUser
	for _, userID := range userIDs {
		courseUsers = append(courseUsers, models.CourseUser{
			CourseID: courseID,
			UserID:   userID,
		})
	}

	if err := s.connection.Create(&courseUsers).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseUserRepositoryImpl) FindUsersNotInCourse(courseID uint32, typeUser uint32) ([]models.User, error) {
	var users []models.User
	err := s.connection.
		Table("users").
		Joins("LEFT JOIN course_users ON users.user_id = course_users.user_id AND course_users.course_id = ?", courseID).
		Where("course_users.user_id IS NULL AND users.type_user_id = ?", typeUser).
		Find(&users).Error

	if err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return users, nil
}

func (s *CourseUserRepositoryImpl) RemoveUserFromCourse(courseID, userID uint32) error {
	err := s.connection.Where("course_id = ? AND user_id = ?", courseID, userID).Delete(&models.CourseUser{}).Error
	if err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *CourseUserRepositoryImpl) Create(courseUser models.CourseUser) error {
	if err := s.connection.Create(&courseUser).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseUserRepositoryImpl) CreateBatch(courseUsers []models.CourseUser) error {
	if err := s.connection.Create(&courseUsers).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseUserRepositoryImpl) FindCourseIDByProgramID(programID uint32, userID uint32) ([]uint32, error) {
	var result []uint32
	//if err := s.connection.Model(&models.CourseUser{}).Preload("Course").Where("program_id", programID).Pluck("course_id", &result).Error; err != nil {
	//	OutPutDebugError(err.Error())
	//	return nil, err
	//}

	if err := s.connection.Table("course_users").
		Joins("JOIN courses c ON course_users.course_id = c.course_id").
		Where("c.program_id = ? and course_users.user_id = ?", programID, userID).Pluck("course_users.course_id", &result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *CourseUserRepositoryImpl) Update(courseUser models.CourseUser) error {
	if err := s.connection.Save(&courseUser).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseUserRepositoryImpl) FindByUserIDAndCourseID(courseID, userID uint32) (*models.CourseUser, error) {
	var result *models.CourseUser

	if err := s.connection.Preload("User").Preload("Course").Where("course_id = ? AND user_id = ?", courseID, userID).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *CourseUserRepositoryImpl) FindCompletedCourse(userid uint32) ([]*models.CourseUser, error) {
	var result []*models.CourseUser
	if err := s.connection.Preload("User").Preload("Course").Where("status = ? AND user_id = ?", 1, userid).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}
