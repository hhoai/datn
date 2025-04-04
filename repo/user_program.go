package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type UserProgramRepository interface {
	FindAll() []*models.UserProgram
	FindByProgramID(ProgramID uint32) ([]*models.UserProgram, error)
	FindByUserID(userID uint32) ([]*models.UserProgram, error)
	Create(requestCourse models.UserProgram) error
	CreateAny(requestCourse []*models.UserProgram) error
	Update(userProgram models.UserProgram) error
	Delete(requestCourse models.UserProgram) error
	FindByUserIDAndProgramID(ProgramID, userID uint32) (*models.UserProgram, error)
	//FindCompletedProgram(userid uint32) ([]*models.LearningProcess, error)
	//CountProgramCompletedInCourse(courseID, userID uint32) (uint32, error)
}
type UserProgramRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserProgramRepository() UserProgramRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &UserProgramRepositoryImpl{
		connection: DB,
	}
}

func (s *UserProgramRepositoryImpl) FindAll() []*models.UserProgram {
	var result []*models.UserProgram
	if err := s.connection.Model(&models.UserProgram{}).Preload("User").Preload("Program").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *UserProgramRepositoryImpl) FindByProgramID(ID uint32) ([]*models.UserProgram, error) {
	var result []*models.UserProgram
	if err := s.connection.Model(&models.UserProgram{}).Preload("User").Preload("Program").Where("program_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserProgramRepositoryImpl) FindByUserID(ID uint32) ([]*models.UserProgram, error) {
	var result []*models.UserProgram
	if err := s.connection.Model(&models.UserProgram{}).Preload("User").Preload("Program").Where("user_id = ?", ID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result, nil
}

func (s *UserProgramRepositoryImpl) Create(req models.UserProgram) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserProgramRepositoryImpl) CreateAny(req []*models.UserProgram) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserProgramRepositoryImpl) Update(userProgram models.UserProgram) error {
	if err := s.connection.Save(&userProgram).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserProgramRepositoryImpl) Delete(req models.UserProgram) error {
	if err := s.connection.Delete(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *UserProgramRepositoryImpl) FindByUserIDAndProgramID(ProgramID, userID uint32) (*models.UserProgram, error) {
	var result *models.UserProgram

	if err := s.connection.Preload("User").Preload("Program").Where("program_id = ? AND user_id = ?", ProgramID, userID).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return result, nil
}

//func (s *UserProgramRepositoryImpl) FindCompletedProgram(userid uint32) ([]*models.LearningProcess, error) {
//	var result []*models.LearningProcess
//
//	if err := s.connection.Table("user_programs").
//		Select("courses.course_id as course_id, courses.title as course_name, Programs.Program_id as Program_id, Programs.title as Program_name,"+
//			"user_Programs.completed_at as timeline_date, user_Programs.completed_at as submit_at").
//		Joins("JOIN users ON users.user_id = user_Programs.user_id").
//		Joins("JOIN Programs ON Programs.Program_id = user_Programs.Program_id").
//		Joins("JOIN courses on courses.course_id = Programs.course_id").
//		Where("user_Programs.status = ? AND user_programs.user_id = ?", 1, userid).Find(&result).Error; err != nil {
//		return nil, err
//	}
//	return result, nil
//}
//
//func (s *UserProgramRepositoryImpl) CountProgramCompletedInCourse(courseID, userID uint32) (uint32, error) {
//	var count int64
//	if err := s.connection.Table("user_programs").
//		Joins("JOIN programs on programs.program_id = user_programs.program_id").
//		Where("programs.course_id = ? AND user_Programs.status = ? AND user_Programs.user_id = ?", courseID, 1, userID).Count(&count).Error; err != nil {
//		OutPutDebugError(err.Error())
//	}
//	return uint32(count), nil
//}
