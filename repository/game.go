package repository

import (
	"github.com/olajhidey/guess-admin/model"
	"gorm.io/gorm"
)

type GameRepository struct {
	DB *gorm.DB
}

func (gameRepo *GameRepository) Create(game *model.Game) error {
	if err := gameRepo.DB.Create(game).Error; err != nil {
		return err
	}
	return nil
}

func (gameRepo *GameRepository) List() ([]model.Game, error) {
	var games []model.Game
	if err := gameRepo.DB.Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func (gameRepo *GameRepository) ListByCode(code string) ([]model.Game, error){
	var games []model.Game

	if err := gameRepo.DB.Where("code = ?", code).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}