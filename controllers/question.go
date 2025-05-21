package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/repository"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type QuestionController struct {
	DB *gorm.DB
}

func (qc *QuestionController) Create(ctx *gin.Context) {
	var questionRequest model.QuestionForm
	if err := ctx.ShouldBindBodyWithJSON(&questionRequest); utils.ErrorNotNil(err) {
		utils.LogMessage(questionRequest)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	questionRepository := repository.QuestionRepository{DB: qc.DB}
	question := model.Question{
		ImageUrl: questionRequest.ImageUrl,
		Answer:   questionRequest.Answer,
		Option1: questionRequest.Option1,
		Option2: questionRequest.Option2,
		Option3: questionRequest.Option3,
		Option4: questionRequest.Option4,
		TopicId: questionRequest.TopicId,
	}

	if err := questionRepository.Create(&question); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create question"})
		return
	}

	ctx.JSON(http.StatusOK, question)
}

func (qc *QuestionController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	questionRepository := repository.QuestionRepository{DB: qc.DB}
	question, err := questionRepository.Get(id)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	ctx.JSON(http.StatusOK, question)
}

func (qc *QuestionController) List(ctx *gin.Context) {
	questionRepository := repository.QuestionRepository{DB: qc.DB}
	questions, err := questionRepository.List()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list questions"})
		return
	}

	ctx.JSON(http.StatusOK, questions)
}

func (qc *QuestionController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var questionRequest model.QuestionForm

	if err := ctx.ShouldBindJSON(&questionRequest); utils.ErrorNotNil(err) {
		utils.LogMessage(questionRequest)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	questionRepository := repository.QuestionRepository{DB: qc.DB}
	question := model.Question{
		ImageUrl: questionRequest.ImageUrl,
		Answer:   questionRequest.Answer,
		Option1: questionRequest.Option1,
		Option2: questionRequest.Option2,
		Option3: questionRequest.Option3,
		Option4: questionRequest.Option4,
		TopicId: questionRequest.TopicId,
	}

	if err := questionRepository.Update(id, &question); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update question"})
		return
	}

	ctx.JSON(http.StatusOK, question)
}

func (qc *QuestionController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	questionRepository := repository.QuestionRepository{DB: qc.DB}
	if err := questionRepository.Delete(id); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete question"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Question deleted successfully"})
}

func (qc *QuestionController) GetById(ctx *gin.Context){
	id := ctx.Param("id")
	repo := repository.QuestionRepository{DB: qc.DB}
	questions, err := repo.GetByTopic(id)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve questions"})
		return
	}
	ctx.JSON(http.StatusOK, questions)
}

func (qc *QuestionController) DeleteAll(ctx *gin.Context) {
	questionRepository := repository.QuestionRepository{DB: qc.DB}
	if err := questionRepository.DeleteAll(); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all questions"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "All questions deleted successfully"})
}