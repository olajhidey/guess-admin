package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/repository"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type TopicController struct {
	DB *gorm.DB
}

func (topicCtrl *TopicController) Create(ctx *gin.Context) {
	var topicRequest *model.TopicForm
	topicRepository := repository.TopicRepository{DB: topicCtrl.DB}

	if err := ctx.ShouldBindBodyWithJSON(&topicRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	insertTopic := model.Topic{
		Name: topicRequest.Name,
		Description: topicRequest.Description,
		CategoryID: uint(topicRequest.CategoryID),
	}
	err := topicRepository.Create(&insertTopic)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create topic"})
		return 
	}

	ctx.JSON(http.StatusOK, insertTopic)

}

func (topicCtrl *TopicController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	topicRepo := repository.TopicRepository{DB: topicCtrl.DB}

	topic, err:= topicRepo.Get(id)
	if utils.ErrorNotNil(err){
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get topic"})
		return
	}

	ctx.JSON(http.StatusOK, topic)
	
}

func (topicCtrl *TopicController) List(ctx *gin.Context) {
	topicRepo := repository.TopicRepository{DB: topicCtrl.DB}
	topics, err := topicRepo.List()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get list of topics"})
		return 
	}
	ctx.JSON(http.StatusOK, topics)
}

func (topicCtrl *TopicController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	topicRepo := repository.TopicRepository{DB: topicCtrl.DB}
	var topicRequest *model.TopicForm

	if err:= ctx.ShouldBindBodyWithJSON(&topicRequest); utils.ErrorNotNil(err){
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return 
	}

	updatedTopic := model.Topic{
		Name: topicRequest.Name,
		Description: topicRequest.Description,
		CategoryID: uint(topicRequest.CategoryID),
	}
	err := topicRepo.Update(id, &updatedTopic)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating topic"})
		return 
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Topic updated successfully"})
}

func (topicCtrl *TopicController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	topicRepo := repository.TopicRepository{DB: topicCtrl.DB}
	topics, err := topicRepo.GetByCategoryID(id)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get topics"})
		return 
	}
	ctx.JSON(http.StatusOK, topics)
}

func (topicCtrl *TopicController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	topicRepo := repository.TopicRepository{DB: topicCtrl.DB}
	err := topicRepo.Delete(id)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Unable to delete Topic"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Topic deleted successfully"})
}
