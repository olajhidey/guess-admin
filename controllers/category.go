package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/repository"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func (categoryCtrl *CategoryController) Create(ctx *gin.Context){
	var categoryRequest model.CategoryForm
	validateCategoryRequest(ctx, &categoryRequest)
	categoryRepo := repository.CategoryRepository{DB: categoryCtrl.DB}

	category := &model.Category{
		Name:        categoryRequest.Name,
		Description: categoryRequest.Description,
	}

	if err := categoryRepo.Create(category); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}
	ctx.JSON(http.StatusOK, category)

}

func (categoryCtrl *CategoryController) List(ctx *gin.Context){
	categoryRepo := repository.CategoryRepository{DB: categoryCtrl.DB}
	categories, err := categoryRepo.List()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (categoryCtrl *CategoryController) Delete(ctx *gin.Context){
	id := ctx.Param("id")
	categoryRepo := repository.CategoryRepository{DB: categoryCtrl.DB}
	if err := categoryRepo.Delete(id); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (categoryCtrl *CategoryController) Get(ctx *gin.Context){
	id := ctx.Param("id")
	categoryRepo := repository.CategoryRepository{DB: categoryCtrl.DB}
	category, err := categoryRepo.Get(id)
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve category"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (categoryCtrl *CategoryController) Update(ctx *gin.Context){
	id := ctx.Param("id")
	var categoryRequest model.CategoryForm
	validateCategoryRequest(ctx, &categoryRequest)
	categoryRepo := repository.CategoryRepository{DB: categoryCtrl.DB}

	category := &model.Category{
		Name:        categoryRequest.Name,
		Description: categoryRequest.Description,
	}

	if err := categoryRepo.Update(id, category); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}


func validateCategoryRequest(ctx *gin.Context, categoryRequest *model.CategoryForm) {
	if err := ctx.ShouldBindJSON(categoryRequest); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if categoryRequest.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}
	if categoryRequest.Description == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}
}