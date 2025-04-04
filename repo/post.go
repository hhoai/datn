package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type PostRepository interface {
	FindAll() []*models.Post
	Create(post models.Post) error
	Delete(post models.Post) error
	Update(post models.Post) error
	DeleteByID(postID uint32) error
	UpdateByID(postID uint32) error
	FindByID(postID uint32) (*models.Post, error)
	FindByLessonID(lessonID uint32) ([]*models.Post, error)
}

type PostRepositoryImpl struct {
	connection *gorm.DB
}

func NewPostRepository() PostRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &PostRepositoryImpl{
		connection: DB,
	}
}

func (s *PostRepositoryImpl) FindAll() []*models.Post {
	var result []*models.Post
	if err := s.connection.Model(&models.Post{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *PostRepositoryImpl) Create(post models.Post) error {
	if err := s.connection.Create(&post).Error; err != nil {
		OutPutDebugError("Fail to create: " + err.Error())
		return err
	}
	return nil
}

func (s *PostRepositoryImpl) Update(post models.Post) error {
	if err := s.connection.Save(&post).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *PostRepositoryImpl) Delete(post models.Post) error {
	if err := s.connection.Delete(&post).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *PostRepositoryImpl) DeleteByID(postID uint32) error {
	post := models.Post{}
	post.PostID = postID
	if err := s.connection.Delete(&post).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *PostRepositoryImpl) UpdateByID(postID uint32) error {
	post := models.Post{}
	post.PostID = postID
	if err := s.connection.Save(&post).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}

func (s *PostRepositoryImpl) FindByID(postID uint32) (*models.Post, error) {
	var post models.Post
	if err := s.connection.Preload("User").Where("post_id = ?", postID).First(&post).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return &post, nil
}

func (s *PostRepositoryImpl) FindByLessonID(lessonID uint32) ([]*models.Post, error) {
	var result []*models.Post
	if err := s.connection.Model(&models.Post{}).Preload("User").Where("lesson_id = ?", lessonID).Order("created_at desc").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return result, nil
}
