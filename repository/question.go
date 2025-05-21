package repository

import (
	"github.com/olajhidey/guess-admin/model"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	DB *gorm.DB
}

func (repo *QuestionRepository) Create(question *model.Question) error {
	if err := repo.DB.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (repo *QuestionRepository) Get(id string) (*model.Question, error) {
	var question model.Question
	if err := repo.DB.First(&question, id).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

func (repo *QuestionRepository) List() ([]model.Question, error) {
	var questions []model.Question
	if err := repo.DB.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (repo *QuestionRepository) Update(id string, question *model.Question) error {
	if err := repo.DB.Model(&model.Question{}).Where("id = ?", id).Updates(question).Error; err != nil {
		return err
	}
	return nil
}

func (repo *QuestionRepository) Delete(id string) error {
	if err := repo.DB.Delete(&model.Question{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *QuestionRepository) GetByTopic(id string) ([]model.Question, error) {
	var questions []model.Question
	if err := repo.DB.Where("topic_id = ?", id).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (repo *QuestionRepository) DeleteAll() error {
	if err := repo.DB.Exec("DELETE FROM questions").Error; err != nil {
		return err
	}
	return nil
}