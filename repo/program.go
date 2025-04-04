package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type ProgramRepository interface {
	FindAll() []*models.Program
	Create(result *models.Program) error
	FindByID(ID uint32) (*models.Program, error)
	Delete(result *models.Program) error
	Update(models.Program) error
	FindByUserID(userID uint32) ([]*models.UserProgram, error)
	FindByProgramCode(programCode string) (*models.Program, error)
	SearchProgram(query string) ([]*models.Program, error)
}

type ProgramRepositoryImpl struct {
	connection *gorm.DB
}

func NewProgramRepository() ProgramRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &ProgramRepositoryImpl{
		connection: DB,
	}
}

func (s *ProgramRepositoryImpl) FindAll() []*models.Program {
	var result []*models.Program
	if err := s.connection.Model(&models.Program{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *ProgramRepositoryImpl) Create(result *models.Program) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *ProgramRepositoryImpl) FindByID(ID uint32) (*models.Program, error) {
	var result models.Program
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *ProgramRepositoryImpl) Delete(program *models.Program) error {
	if err := s.connection.Delete(&program).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *ProgramRepositoryImpl) Update(program models.Program) error {
	if err := s.connection.Save(&program).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *ProgramRepositoryImpl) FindByUserID(userID uint32) ([]*models.UserProgram, error) {
	var result []*models.UserProgram
	if err := s.connection.Preload("User").Preload("Program").Where("user_id", userID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *ProgramRepositoryImpl) FindByProgramCode(programCode string) (*models.Program, error) {
	var result *models.Program
	if err := s.connection.Where("program_code = ?", programCode).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *ProgramRepositoryImpl) SearchProgram(query string) ([]*models.Program, error) {
	var result []*models.Program
	if err := s.connection.Where("name LIKE ?", "%"+query+"%").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}
