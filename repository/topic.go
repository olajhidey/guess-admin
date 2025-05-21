package repository

import (
	"github.com/olajhidey/guess-admin/model"
	"gorm.io/gorm"
)

type TopicRepository struct {
	DB *gorm.DB
}

func (repo *TopicRepository) Create(topic *model.Topic) error {
	if err := repo.DB.Create(topic).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TopicRepository) Get(id string) (*model.Topic, error) {
	var topic model.Topic
	if err := repo.DB.First(&topic, id).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (repo *TopicRepository) List() ([]model.Topic, error) {
	var topics []model.Topic
	if err := repo.DB.Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (repo *TopicRepository) Update(id string, topic *model.Topic) error {
	if err := repo.DB.Model(&model.Topic{}).Where("id = ?", id).Updates(topic).Error; err != nil {
		return err
	}
	return nil
}

func (repo *TopicRepository) GetByCategoryID(categoryID string) ([]model.Topic, error) {
	var topics []model.Topic
	if err := repo.DB.Where("category_id = ?", categoryID).Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (repo *TopicRepository) Delete(id string) error {
	if err := repo.DB.Delete(&model.Topic{}, id).Error; err != nil {
		return err
	}
	return nil
}