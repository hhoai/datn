package repo

import (
	"gorm.io/gorm"
	"lms/models"
	"lms/structs"
	"time"
)

type UserRepository interface {
	FindAll() []*models.UserWithoutPass
	FindByID(userID uint32) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByIDWithoutPass(userID uint32) (*models.UserWithoutPass, error)
	FindByEmailWithoutPass(email string) (*models.UserWithoutPass, error)
	FindByTypeUserID(typeUserID uint32) ([]*models.UserWithoutPass, error)
	Create(user models.User) error
	Delete(models.User) error
	Update(models.User) error
	DeleteByID(userID uint32) error
	UpdateByID(userID uint32) error
	UpdatePassword(userID uint32, newPassword string) error
	GetLearningProcess(userID uint32) (*models.LearningProcess, error)
	CountUser() (uint32, error)
	FindNewUsersLast7Days() []structs.HighchartsData
	CreateAny(req []*models.User) error
}

type UserRepositoryImpl struct {
	connection *gorm.DB
}

func NewUserRepository() UserRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &UserRepositoryImpl{
		connection: DB,
	}
}

func (s *UserRepositoryImpl) FindAll() []*models.UserWithoutPass {
	var result []*models.UserWithoutPass
	if err := s.connection.Model(&models.User{}).Preload("TypeUser").Preload("Role").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *UserRepositoryImpl) FindByID(userID uint32) (*models.User, error) {
	var result models.User

	if err := s.connection.Model(&models.User{}).Preload("TypeUser").Preload("Role").First(&result, userID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}

	return &result, nil
}

func (s *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var result models.User
	if err := s.connection.Model(&models.User{}).Preload("TypeUser").Preload("Role").Where("email", email).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *UserRepositoryImpl) FindByIDWithoutPass(userID uint32) (*models.UserWithoutPass, error) {
	var result models.UserWithoutPass
	if err := s.connection.Model(&models.User{}).Preload("TypeUser").Preload("Role").First(&result, userID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *UserRepositoryImpl) FindByEmailWithoutPass(email string) (*models.UserWithoutPass, error) {
	var result models.UserWithoutPass
	if err := s.connection.Model(&models.User{}).Preload("TypeUser").Preload("Role").Where("email = ?", email).First(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}

func (s *UserRepositoryImpl) FindByTypeUserID(typeUserID uint32) ([]*models.UserWithoutPass, error) {
	var result []*models.UserWithoutPass
	if err := s.connection.Model(&models.User{}).Where("type_user_id = ?", typeUserID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}

func (s *UserRepositoryImpl) Create(user models.User) error {
	if err := s.connection.Create(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) Update(user models.User) error {
	if err := s.connection.Save(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) Delete(user models.User) error {
	if err := s.connection.Delete(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) DeleteByID(userID uint32) error {
	user := models.User{}
	user.UserID = userID
	if err := s.connection.Delete(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) UpdateByID(userID uint32) error {
	user := models.User{}
	user.UserID = userID
	if err := s.connection.Save(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) UpdatePassword(userID uint32, newPassword string) error {
	var user models.User
	if err := s.connection.Model(&models.User{}).First(&user, userID).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	user.Password = newPassword
	if err := s.connection.Save(&user).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *UserRepositoryImpl) GetLearningProcess(userID uint32) (*models.LearningProcess, error) {
	var process models.LearningProcess
	if err := s.connection.Model(&models.LearningProcess{}).First(&process, userID).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &process, nil
}

func (s *UserRepositoryImpl) CountUser() (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.User{}).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}

func (s *UserRepositoryImpl) FindNewUsersLast7Days() []structs.HighchartsData {
	var result []structs.HighchartsData

	var rs []struct {
		Date  time.Time `json:"date"`
		Count uint32    `json:"count"`
	}

	s.connection.Model(&models.User{}).
		Select("DATE(created_at) as date, COUNT(*) as count").
		Where("DATE(created_at) >= DATE(NOW()) - INTERVAL 7 DAY").
		Group("DATE(created_at)").
		Order("date").
		Scan(&rs)

	for _, row := range rs {
		result = append(result, structs.HighchartsData{
			Date:  row.Date.Format("2006-01-02"),
			Count: row.Count,
		})
	}

	return result
}

func (s *UserRepositoryImpl) CreateAny(req []*models.User) error {
	if err := s.connection.Create(&req).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
