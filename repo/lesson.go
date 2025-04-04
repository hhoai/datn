package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type LessonRepository interface {
	FindAll() []*models.Lesson
	Create(lesson models.Lesson) error
	Delete(models.Lesson) error
	Update(models.Lesson) error
	DeleteByID(lessonID uint32) error
	UpdateByID(lessonID uint32) error
	FindByID(lessonID uint32) (*models.Lesson, error)
	FindByCourseID(courseID uint32) ([]*models.Lesson, error)
	FindInCourseID(courseID []uint32) ([]*models.Lesson, error)
	FindLessonIDByCourseID(courseID uint32) ([]uint32, error)
	FindLessonIDByCourseIDs(courseID []uint32) ([]uint32, error)
	DeleteMultiple(lessonID []uint32) error
	GetLastLesson() (*models.Lesson, error)
	CountLessonInCourse(courseID uint32) (uint32, error)
	FindLessonsByUserAndCourse(userID uint32, courseID uint32) ([]models.Lesson, error)
}

type LessonRepositoryImpl struct {
	connection *gorm.DB
}

func NewLessonRepository() LessonRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &LessonRepositoryImpl{
		connection: DB,
	}
}

func (s *LessonRepositoryImpl) FindAll() []*models.Lesson {
	var result []*models.Lesson
	if err := s.connection.Model(&models.Lesson{}).Preload("LessonCategory").
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *LessonRepositoryImpl) Create(lesson models.Lesson) error {
	var lessonCategory models.LessonCategory
	if err := s.connection.First(&lessonCategory, lesson.LessonCategoryID).Error; err != nil {
		OutPutDebugError("lesson category not exit: " + err.Error())
		return err
	}
	if err := s.connection.Create(&lesson).Error; err != nil {
		OutPutDebugError("Fail to create: " + err.Error())
		return err
	}
	return nil
}

func (s *LessonRepositoryImpl) Update(lesson models.Lesson) error {
	if err := s.connection.Save(&lesson).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *LessonRepositoryImpl) Delete(lesson models.Lesson) error {
	if err := s.connection.Delete(&lesson).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *LessonRepositoryImpl) DeleteByID(lessonID uint32) error {
	lesson := models.Lesson{}
	lesson.LessonID = lessonID
	if err := s.connection.Delete(&lesson).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *LessonRepositoryImpl) UpdateByID(lessonID uint32) error {
	lesson := models.Lesson{}
	lesson.LessonID = lessonID
	if err := s.connection.Save(&lesson).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *LessonRepositoryImpl) FindByID(lessonID uint32) (*models.Lesson, error) {
	var lesson models.Lesson
	if err := s.connection.Model(&models.Lesson{}).Preload("LessonCategory").First(&lesson, lessonID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	if err := s.connection.Find(&lesson).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &lesson, nil
}

func (s *LessonRepositoryImpl) FindByCourseID(courseID uint32) ([]*models.Lesson, error) {
	var result []*models.Lesson
	if err := s.connection.Model(&models.Lesson{}).Preload("LessonCategory").Preload("Level").Where("course_id", courseID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *LessonRepositoryImpl) FindInCourseID(courseID []uint32) ([]*models.Lesson, error) {
	var lessons []*models.Lesson
	if err := s.connection.Model(&models.Lesson{}).
		Where("course_id IN (?)", courseID).
		Find(&lessons).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return lessons, nil
}

func (s *LessonRepositoryImpl) FindLessonIDByCourseID(courseID uint32) ([]uint32, error) {
	var result []uint32
	if err := s.connection.Model(&models.Lesson{}).Where("course_id", courseID).Pluck("lesson_id", &result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *LessonRepositoryImpl) FindLessonIDByCourseIDs(courseID []uint32) ([]uint32, error) {
	var result []uint32
	if err := s.connection.Model(&models.Lesson{}).Where("course_id IN (?)", courseID).Pluck("lesson_id", &result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *LessonRepositoryImpl) DeleteMultiple(lessonID []uint32) error {
	if err := s.connection.Where("lesson_id IN ?", lessonID).Delete(&models.Lesson{}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *LessonRepositoryImpl) GetLastLesson() (*models.Lesson, error) {
	var lesson models.Lesson
	if err := s.connection.Model(&models.Lesson{}).Preload("LessonCategory").Order("lesson_id desc").First(&lesson).Error; err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (s *LessonRepositoryImpl) CountLessonInCourse(courseID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Lesson{}).Where("course_id = ?", courseID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}

func (s *LessonRepositoryImpl) FindLessonsByUserAndCourse(userID uint32, courseID uint32) ([]models.Lesson, error) {
	var lessons []models.Lesson
	err := s.connection.Model(&models.Lesson{}).
		Joins("JOIN courses ON lessons.course_id = courses.course_id").
		Joins("JOIN course_users ON courses.course_id = course_users.course_id").
		Joins("JOIN users ON course_users.user_id = users.user_id").
		Where("users.user_id = ? AND courses.course_id = ?", userID, courseID).
		Preload("Level").
		Preload("LessonCategory").
		Find(&lessons).Error
	if err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return lessons, nil
}
