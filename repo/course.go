package repo

import (
	"gorm.io/gorm"
	"lms/models"
	"lms/structs"
)

type CourseRepository interface {
	FindAll() []*models.Course
	Create(course models.Course) error
	Delete(models.Course) error
	Update(models.Course) error
	DeleteByID(courseID uint32) error
	UpdateByID(courseID uint32) error
	FindByID(courseID uint32) (*models.Course, error)
	FindByCourseCode(courseCode string) (*models.Course, error)
	DeleteMultiple(courseIDs []uint32) error
	CountCourse() (uint32, error)
	FindCourseWithLessonByUserID(userID uint32) ([]*structs.CourseDetails, error)
	FindByProgramID(programID uint32) ([]*models.Course, error)
	FindInProgramID(courseID []uint32, existCourseID []uint32) ([]*models.Course, error)
	FindCourseIDByProgramID(programID uint32) ([]uint32, error)
	CountCourseInProgram(programID uint32) (uint32, error)
}

type CourseRepositoryImpl struct {
	connection *gorm.DB
}

func NewCourseRepository() CourseRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &CourseRepositoryImpl{
		connection: DB,
	}
}

func (s *CourseRepositoryImpl) FindAll() []*models.Course {
	var result []*models.Course
	if err := s.connection.Model(&models.Course{}).Preload("Level").Preload("Program").Preload("CourseCategory").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *CourseRepositoryImpl) Create(course models.Course) error {
	if err := s.connection.Create(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseRepositoryImpl) Update(course models.Course) error {
	if err := s.connection.Save(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *CourseRepositoryImpl) Delete(course models.Course) error {
	if err := s.connection.Delete(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseRepositoryImpl) DeleteByID(courseID uint32) error {
	course := models.Course{}
	course.CourseID = courseID
	if err := s.connection.Delete(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseRepositoryImpl) UpdateByID(courseID uint32) error {
	course := models.Course{}
	course.CourseID = courseID
	if err := s.connection.Save(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseRepositoryImpl) FindByID(courseID uint32) (*models.Course, error) {
	var course models.Course
	if err := s.connection.Model(&models.Course{}).First(&course, courseID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	if err := s.connection.Preload("CourseCategory").Find(&course).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &course, nil
}
func (s *CourseRepositoryImpl) FindByCourseCode(courseCode string) (*models.Course, error) {
	var course models.Course
	if err := s.connection.Model(&models.Course{}).Where("course_code = ?", courseCode).First(&course).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	if err := s.connection.Preload("CourseCategory").Find(&course).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &course, nil
}

func (s *CourseRepositoryImpl) DeleteMultiple(courseID []uint32) error {
	if err := s.connection.Where("course_id IN ?", courseID).Delete(&models.Course{}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *CourseRepositoryImpl) CountCourse() (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Course{}).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}

func (s *CourseRepositoryImpl) FindCourseWithLessonByUserID(userID uint32) ([]*structs.CourseDetails, error) {
	var course []*structs.CourseDetails
	if err := s.connection.Table("user_lessons AS ul").
		Select(
			"c.title as title",
			"c.course_id as course_id",
			"c.course_code as course_code",
			"c.description as description",
			"c.status as status",
			"COUNT(ul.lesson_id) AS lesson",
			"COUNT(CASE WHEN ul.status = true THEN 1 END) AS lesson_completed").
		Joins("JOIN lessons AS l ON l.lesson_id = ul.lesson_id").
		Joins("JOIN courses AS c ON c.course_id = l.course_id").
		Where("ul.user_id = ?", userID).
		Group("c.title").
		Find(&course).Error; err != nil {
		return nil, err
	}
	return course, nil
}
func (s *CourseRepositoryImpl) FindByProgramID(programID uint32) ([]*models.Course, error) {
	var result []*models.Course
	if err := s.connection.Model(&models.Course{}).Preload("Level").Preload("Program").Preload("CourseCategory").Where("program_id", programID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *CourseRepositoryImpl) FindInProgramID(courseID []uint32, existCourseID []uint32) ([]*models.Course, error) {
	var courses []*models.Course
	if len(existCourseID) > 0 {
		if err := s.connection.Debug().
			Where("course_id IN (?) and course_id NOT IN (?)", courseID, existCourseID).
			//Where("", existCourseID).
			Find(&courses).Error; err != nil {
			OutPutDebugError(err.Error())
			return nil, err
		}
	} else {
		if err := s.connection.
			Where("course_id IN (?)", courseID).
			//Where("", existCourseID).
			Find(&courses).Error; err != nil {
			OutPutDebugError(err.Error())
			return nil, err
		}
	}

	return courses, nil
}

func (s *CourseRepositoryImpl) FindCourseIDByProgramID(programID uint32) ([]uint32, error) {
	var result []uint32
	if err := s.connection.Model(&models.Course{}).Where("program_id", programID).Pluck("course_id", &result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *CourseRepositoryImpl) CountCourseInProgram(programID uint32) (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Course{}).Where("program_id", programID).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
