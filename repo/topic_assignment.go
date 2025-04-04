package repo

import (
	"gorm.io/gorm"
	"lms/models"
)

type TopicAssignmentRepository interface {
	FindAll() []*models.TopicAssignment
	Create(topicAssignments *models.TopicAssignment) error
	FindTopicsByAssignmentID(assignmentID uint32) (*models.Topic, error)
	FindTopicsByAssignmentIDTopicID(assignmentID, topicID uint32) (*models.TopicAssignment, error)
	FindByTopicID(topicID uint32) (*models.TopicAssignment, error)
	FindByTopicIDAndAssignmentID(topicID uint32, assignmentID uint32) (*models.TopicAssignment, error)
	Delete(ta *models.TopicAssignment) error
}

type TopicAssignmentRepositoryImpl struct {
	connection *gorm.DB
}

func NewTopicAssignmentRepository() TopicAssignmentRepository {
	if DB == nil {
		_, err := ConnectDB()
		if err != nil {
			OutPutDebugError(err.Error())
		}
	}
	return &TopicAssignmentRepositoryImpl{
		connection: DB,
	}
}

func (s *TopicAssignmentRepositoryImpl) FindAll() []*models.TopicAssignment {
	var result []*models.TopicAssignment
	if err := s.connection.Model(&models.TopicAssignment{}).Preload("Topic").Preload("Assignment").Find(&result).Error; err != nil {
		OutPutDebugError(err.Error())
	}
	return result
}

func (s *TopicAssignmentRepositoryImpl) Create(topicAssignment *models.TopicAssignment) error {
	if err := s.connection.Create(topicAssignment).Error; err != nil {
		OutPutDebugError(err.Error())
		return err
	}
	return nil
}
func (s *TopicAssignmentRepositoryImpl) FindTopicsByAssignmentID(assignmentID uint32) (*models.Topic, error) {
	var topics *models.Topic
	if err := s.connection.Model(&models.Topic{}).
		Joins("join topic_assignments on topics.topic_id = topic_assignments.topic_id").
		Where("topic_assignments.assignment_id = ?", assignmentID).
		First(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (s *TopicAssignmentRepositoryImpl) FindTopicsByAssignmentIDTopicID(assignmentID, topicID uint32) (*models.TopicAssignment, error) {
	var result *models.TopicAssignment
	if err := s.connection.
		Preload("Topic").
		Where("topic_id", topicID).
		Where("assignment_id", assignmentID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TopicAssignmentRepositoryImpl) FindByTopicID(topicID uint32) (*models.TopicAssignment, error) {
	var result *models.TopicAssignment
	if err := s.connection.
		Preload("Topic").
		Where("topic_id", topicID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TopicAssignmentRepositoryImpl) FindByTopicIDAndAssignmentID(topicID uint32, assignmentID uint32) (*models.TopicAssignment, error) {
	var result *models.TopicAssignment
	if err := s.connection.
		Preload("Topic").
		Where("topic_id = ? and assignment_id = ?", topicID, assignmentID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (s *TopicAssignmentRepositoryImpl) Delete(ta *models.TopicAssignment) error {
	if err := s.connection.Delete(ta).Error; err != nil {
		return err
	}
	return nil
}
