package repository

import (
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (userRepo *UserRepository) CreateUser(user *model.User) error {
	result := userRepo.DB.Create(user)
	return result.Error
}

func (userRepo *UserRepository) ListUsers() ([]model.User, error) {
	var users []model.User
	result := userRepo.DB.Find(&users)
	if utils.ErrorNotNil(result.Error) {
		return nil, result.Error
	}
	return users, nil
}

func (userRepo *UserRepository) DeleteAllUsers() error {
	result := userRepo.DB.Exec("DELETE FROM users")
	return result.Error
}

func (userRepo *UserRepository) GetUser(username string) (*model.User, error) {
	var user model.User
	result := userRepo.DB.Where("username = ?", username).First(&user)
	if utils.ErrorNotNil(result.Error) {
		return nil, result.Error
	}
	return &user, nil
}