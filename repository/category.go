package repository

import (
	"github.com/olajhidey/guess-admin/model"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (categoryRepo *CategoryRepository) Create(category *model.Category) error {
	result := categoryRepo.DB.Create(category)
	return result.Error
}

func (categoryRepo *CategoryRepository) List() ([]model.Category, error) {
	var categories []model.Category
	result := categoryRepo.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (categoryRepo *CategoryRepository) Delete(id string) error {
	result := categoryRepo.DB.Delete(&model.Category{}, id)
	return result.Error
}

func (categoryRepo *CategoryRepository) Get(id string) (*model.Category, error) {
	var category model.Category
	result := categoryRepo.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (categoryRep *CategoryRepository) Update(id string, category *model.Category) error {
	result := categoryRep.DB.Model(&model.Category{}).Where("id = ?", id).Updates(category)
	return result.Error
}