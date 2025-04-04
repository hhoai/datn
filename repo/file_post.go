package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type FilePostRepository interface {
	FindAll() []*models.FilePost
	Create(filePost models.FilePost) error
	Delete(filePost models.FilePost) error
	Update(filePost *models.FilePost) error
	UpdateDefaultByFileIDs(fileIDs []string, status bool) error
	DeleteDefault() error
	DeleteByID(filePostID uint32) error
	UpdateByID(filePostID uint32) error
	FindByID(filePostID uint32) (*models.FilePost, error)
	FindByPostID(filePostID uint32) ([]*models.FilePost, error)
}

type FilePostRepositoryImpl struct {
	connection *gorm.DB
}

func NewFilePostRepository() FilePostRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &FilePostRepositoryImpl{
		connection: DB,
	}
}

func (s *FilePostRepositoryImpl) FindAll() []*models.FilePost {
	var result []*models.FilePost
	if err := s.connection.Model(&models.FilePost{}).
		Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *FilePostRepositoryImpl) Create(filePost models.FilePost) error {
	if err := s.connection.Create(&filePost).Error; err != nil {
		OutPutDebugError("Fail to create: " + err.Error())
		return err
	}
	return nil
}

func (s *FilePostRepositoryImpl) Update(filePost *models.FilePost) error {
	if err := s.connection.Save(&filePost).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *FilePostRepositoryImpl) UpdateDefaultByFileIDs(fileIDs []string, status bool) error {
	if err := s.connection.Model(&models.FilePost{}).
		Where("file_post_id in (?)", fileIDs).
		Updates(map[string]interface{}{
			"default": status,
		}).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *FilePostRepositoryImpl) DeleteDefault() error {
	if err := s.connection.Model(&models.FilePost{}).Where("default", true).Update("default", false).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FilePostRepositoryImpl) Delete(filePost models.FilePost) error {
	if err := s.connection.Delete(&filePost).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FilePostRepositoryImpl) DeleteByID(filePostID uint32) error {
	filePost := models.FilePost{}
	filePost.FilePostID = filePostID
	if err := s.connection.Delete(&filePost).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FilePostRepositoryImpl) UpdateByID(filePostID uint32) error {
	filePost := models.FilePost{}
	filePost.FilePostID = filePostID
	if err := s.connection.Save(&filePost).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *FilePostRepositoryImpl) FindByID(filePostID uint32) (*models.FilePost, error) {
	var filePost models.FilePost

	if err := s.connection.Where("file_post_id", filePostID).Find(&filePost).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &filePost, nil
}

func (s *FilePostRepositoryImpl) FindByPostID(postID uint32) ([]*models.FilePost, error) {
	var result []*models.FilePost
	if err := s.connection.Model(&models.FilePost{}).Where("post_id", postID).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}
