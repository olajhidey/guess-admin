package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/olajhidey/guess-admin/config"
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/repository"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (user *UserController) Login(ctx *gin.Context){
	var loginRequest model.LoginForm
	userRepository := repository.UserRepository{DB: user.DB}
	checkLoginFormRequest(&loginRequest, ctx)
	result, err := userRepository.GetUser(loginRequest.Username)

	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	validateUserDetails(loginRequest, result, ctx)
	claims := utils.CustomClaims{
		UUID: result.UUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.SecretKey))

	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token":tokenString})
}
func (user *UserController) Register(ctx *gin.Context){
	var registerRequest model.RegisterForm
	userRepository := repository.UserRepository{DB: user.DB}

	validateRegisterFormRequest(&registerRequest, ctx)
	hashedPassword, _ := utils.HashPassword(registerRequest.Password)
	_ = userRepository.CreateUser(&model.User{
		UUID: utils.GenerateUuid(),
		Username: registerRequest.Username,
		Password: hashedPassword,
		Email: registerRequest.Email,})

		utils.LogMessage(registerRequest.Username)

	// checkIfUserExists(result, ctx)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
	

}
func (user *UserController) GetUser(ctx *gin.Context){}

func (user *UserController) ListUsers(ctx *gin.Context){
	userRepository := repository.UserRepository{DB: user.DB}
	users, err := userRepository.ListUsers()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (user *UserController) DeleteAllUsers(ctx *gin.Context){
	userRepository := repository.UserRepository{DB: user.DB}
	err := userRepository.DeleteAllUsers()
	if utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Users deleted successfully"})
}

func checkLoginFormRequest(loginRequest *model.LoginForm, ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(loginRequest); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fill in the required information"})
		return
	}
	if loginRequest.Username == "" || loginRequest.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fill in the required information"})
		return
	}
}

func validateUserDetails(loginRequest model.LoginForm, result *model.User, ctx *gin.Context) {
	if result.Username == "" || !utils.CheckPassword(result.Password,loginRequest.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
}

func validateRegisterFormRequest(registerRequest *model.RegisterForm, ctx *gin.Context) {
	if err := ctx.ShouldBindBodyWithJSON(registerRequest); utils.ErrorNotNil(err) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Fill in the required information"})
		return
	}
}

func checkIfUserExists(result error, ctx *gin.Context) {
	if strings.Contains(result.Error(), "Error 1062") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	} else {
		utils.LogMessage(result.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
}