package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type TopicRepository interface {
	FindAll() []*models.Topic
	Create(result *models.Topic) error
	FindByTopicID(topicID uint32) (*models.Topic, error)
	Delete(topic models.Topic) error
	FindByID(ID uint32) (*models.Topic, error)
	Update(topic models.Topic) error
	CountTopic() (uint32, error)
}

type TopicRepositoryImpl struct {
	connection *gorm.DB
}

func NewTopicRepository() TopicRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &TopicRepositoryImpl{
		connection: DB,
	}
}

func (s *TopicRepositoryImpl) FindAll() []*models.Topic {
	var result []*models.Topic
	if err := s.connection.Model(&models.Topic{}).Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *TopicRepositoryImpl) Create(result *models.Topic) error {
	if err := s.connection.Create(result).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *TopicRepositoryImpl) FindByTopicID(topicID uint32) (*models.Topic, error) {
	var topic models.Topic
	if err := s.connection.Where("topic_id = ?", topicID).First(&topic).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &topic, nil
}

func (s *TopicRepositoryImpl) Delete(topic models.Topic) error {
	if err := s.connection.Delete(&topic).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *TopicRepositoryImpl) FindByID(topicID uint32) (*models.Topic, error) {
	var result models.Topic
	if err := s.connection.First(&result, topicID).Error; err != nil {
		OutPutDebugError(err.Error())
		return nil, err
	}
	return &result, nil
}
func (s *TopicRepositoryImpl) Update(topic models.Topic) error {
	if err := s.connection.Save(&topic).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil

}

func (s *TopicRepositoryImpl) CountTopic() (uint32, error) {
	var count int64
	if err := s.connection.Model(&models.Topic{}).Count(&count).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return uint32(count), nil
}
