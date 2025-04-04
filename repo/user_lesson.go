package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type UserLessonRepository interface {
	FindAll() []*models.UserLesson
	FindByLessonID(lessonID uint32) ([]*models.UserLesson, error)
	FindByUserID(userID uint32) ([]*models.UserLesson, error)
	Create(requestCourse models.UserLesson) error
	CreateAny(requestCourse []*models.UserLesson) error
	Update(userLesson models.UserLesson) error
	Delete(requestCourse models.UserLesson) error
	FindByUserIDAndLessonID(lessonID, userID uint32) (*models.UserLesson, error)
	FindCompletedLesson(userid uint32) ([]*models.LearningProcess, error)
	CountLessonCompletedInCourse(courseID, userID uint32) (uint32, error)
}
type UserLessonRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserLessonRepository() UserLessonRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &UserLessonRepositoryImpl{
		connection: DB,
	}
}

func (s *UserLessonRepositoryImpl) FindAll() []*models.UserLesson {
	var result []*models.UserLesson
	if err := s.connection.Model(&models.UserLesson{}).Preload("User").Preload("Lesson").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *UserLessonRepositoryImpl) FindByLessonID(ID uint32) ([]*models.UserLesson, error) {
	var result []*models.UserLesson
	if err := s.connection.Model(&models.UserLesson{}).Preload("User").Preload("Lesson").Where("lesson_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserLessonRepositoryImpl) FindByUserID(ID uint32) ([]*models.UserLesson, error) {
	var result []*models.UserLesson
	if err := s.connection.Model(&models.RequestCourse{}).Preload("User").Preload("Lesson").Where("user_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserLessonRepositoryImpl) Create(req models.UserLesson) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserLessonRepositoryImpl) CreateAny(req []*models.UserLesson) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserLessonRepositoryImpl) Update(userLesson models.UserLesson) error {
	if err := s.connection.Save(&userLesson).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserLessonRepositoryImpl) Delete(req models.UserLesson) error {
	if err := s.connection.Delete(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *UserLessonRepositoryImpl) FindByUserIDAndLessonID(lessonID, userID uint32) (*models.UserLesson, error) {
	var result *models.UserLesson

	if err := s.connection.Preload("User").Preload("Lesson").Where("lesson_id = ? AND user_id = ?", lessonID, userID).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

func (s *UserLessonRepositoryImpl) FindCompletedLesson(userid uint32) ([]*models.LearningProcess, error) {
	var result []*models.LearningProcess

	if err := s.connection.Table("user_lessons").
		Select("courses.course_id as course_id, courses.title as course_name, lessons.lesson_id as lesson_id, lessons.title as lesson_name,"+
			"user_lessons.completed_at as timeline_date, user_lessons.completed_at as submit_at").
		Joins("JOIN users ON users.user_id = user_lessons.user_id").
		Joins("JOIN lessons ON lessons.lesson_id = user_lessons.lesson_id").
		Joins("JOIN courses on courses.course_id = lessons.course_id").
		Where("user_lessons.status = ? AND user_lessons.user_id = ?", 1, userid).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *UserLessonRepositoryImpl) CountLessonCompletedInCourse(courseID, userID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Table("user_lessons").
		Joins("JOIN lessons on lessons.lesson_id = user_lessons.lesson_id").
		Where("lessons.course_id = ? AND user_lessons.status = ? AND user_lessons.user_id = ?", courseID, 1, userID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
