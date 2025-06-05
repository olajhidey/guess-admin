package controllers

import (
	"net/http"
	"strconv"

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
	num, _ := strconv.Atoi(gameRequest.PlayerScore)
	insertGame := model.Game{
		Code: gameRequest.Code,
		TopicId: gameRequest.TopicId,
		PlayerName: gameRequest.PlayerName,
		PlayerScore: num,
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
func (gameController *GameController) ListGamesByCode(ctx *gin.Context) {
	code := ctx.Param("code")
	gameRepository := repository.GameRepository{DB: gameController.DB}
	games, err := gameRepository.ListByCode(code)
	if utils.ErrorNotNil(err){
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get list of games"})
		return
	}
	ctx.JSON(http.StatusOK, games)
}

// https://www.thesportsdb.com/images/media/player/thumb/kgojrb1711448509.jpg
// https://www.thesportsdb.com/images/media/player/thumb/75kyl01734192183.jpg
// https://www.thesportsdb.com/images/media/player/thumb/w561fs1732109007.jpg
// https://www.thesportsdb.com/images/media/player/cutout/29024814.png