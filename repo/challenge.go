package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type ChallengeRepository interface {
	FindAll() []*models.Challenge
	FindByID(ID uint32) (*models.Challenge, error)
	Create(result *models.Challenge) error
	Delete(result *models.Challenge) error
	Update(challenge models.Challenge) error
}

type ChallengeRepositoryImpl struct {
	connection *gorm.DB
}

func NewChallengeRepository() ChallengeRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &ChallengeRepositoryImpl{
		connection: DB,
	}
}

func (s *ChallengeRepositoryImpl) FindAll() []*models.Challenge {
	var result []*models.Challenge
	if err := s.connection.Model(&models.Challenge{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}
func (s *ChallengeRepositoryImpl) FindByID(ID uint32) (*models.Challenge, error) {
	var result models.Challenge
	if err := s.connection.First(&result, ID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *ChallengeRepositoryImpl) Create(result *models.Challenge) error {
	if err := s.connection.Create(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *ChallengeRepositoryImpl) Delete(challenge *models.Challenge) error {
	if err := s.connection.Delete(&challenge).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *ChallengeRepositoryImpl) Update(challenge models.Challenge) error {
	if err := s.connection.Save(&challenge).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}
