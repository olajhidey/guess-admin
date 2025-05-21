package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/repository"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type GameController struct{
	DB *gorm.DB
}

func (gameController *GameController) CreateGame(ctx *gin.Context) {
	var gameRequest model.GameForm
	gameRepository := repository.GameRepository{DB: gameController.DB}
	if err:= ctx.ShouldBindBodyWithJSON(&gameRequest); utils.ErrorNotNil(err){
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	insertGame := model.Game{
		Code: gameRequest.Code,
		TopicId: gameRequest.TopicId,
		PlayerOneName: gameRequest.PlayerOneName,
		PlayerTwoName: gameRequest.PlayerTwoName,
		PlayerOneScore: gameRequest.PlayerOneScore,
		PlayerTwoScore: gameRequest.PlayerTwoScore,
	}
	err := gameRepository.Create(&insertGame)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create game"})
		return
	}
	ctx.JSON(http.StatusOK, insertGame)
}

func (gameController *GameController) ListGames(ctx *gin.Context) {
	gameRepository := repository.GameRepository{DB: gameController.DB}
	games, err := gameRepository.List()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get list of games"})
		return
	}
	ctx.JSON(http.StatusOK, games)
}