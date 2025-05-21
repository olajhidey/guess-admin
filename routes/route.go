package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/controllers"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/gorm"
)

func LoadRoutes(router *gin.Engine, db *gorm.DB) {

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	SetupWebRoutes(router)
	SetupUserRoutes(router, db)
	SetupCategoryRoutes(router, db)
	SetupTopicRoutes(router, db)
	SetupQuestionRoutes(router, db)
	SetupGameSessionRoutes(router, db)
	
}

func SetupWebRoutes(router *gin.Engine) {
	router.Static("/static", "../../www")
	router.NoRoute(func(c *gin.Context) {
		c.File("../../www/index.html")
	})
}

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
	user := controllers.UserController{DB: db}
	userGroup := router.Group("/api/auth")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.GET("/get/:id", user.GetUser)
		userGroup.GET("/list", user.ListUsers)
		userGroup.DELETE("/delete", user.DeleteAllUsers)
	}
}

func SetupCategoryRoutes(router *gin.Engine, db *gorm.DB) {
	category := controllers.CategoryController{DB: db}
	categoryGroup := router.Group("/api/category")
	{
		categoryGroup.POST("/create", utils.AuthMiddleWare(), category.Create)
		categoryGroup.GET("/list", utils.AuthMiddleWare() ,category.List)
		categoryGroup.GET("/get/:id", category.Get)
		categoryGroup.PUT("/update/:id", utils.AuthMiddleWare(), category.Update)
		categoryGroup.DELETE("/delete/:id", utils.AuthMiddleWare(), category.Delete)
	}
}

func SetupTopicRoutes(router *gin.Engine, db *gorm.DB) {
	topic := controllers.TopicController{DB: db}
	topicGroup := router.Group("/api/topic")
	{
		topicGroup.POST("/create", utils.AuthMiddleWare(), topic.Create)
		topicGroup.GET("/list", utils.AuthMiddleWare(), topic.List)
		topicGroup.GET("/get/:id", topic.Get)
		topicGroup.GET("/list/:id", topic.GetById)
		topicGroup.PUT("/update/:id", utils.AuthMiddleWare(), topic.Update)
		topicGroup.DELETE("/delete/:id", utils.AuthMiddleWare(), topic.Delete)
	}
}

func SetupQuestionRoutes(router *gin.Engine, db *gorm.DB){
	questionRoute := router.Group("/api/question")
	questionCtrl := controllers.QuestionController{DB: db}

	questionRoute.POST("/create", utils.AuthMiddleWare(), questionCtrl.Create)
	questionRoute.GET("/get/:id", utils.AuthMiddleWare(), questionCtrl.Get)
	questionRoute.GET("/list", utils.AuthMiddleWare(), questionCtrl.List)
	questionRoute.PUT("/update/:id", utils.AuthMiddleWare(), questionCtrl.Update)
	questionRoute.DELETE("/delete/:id", utils.AuthMiddleWare(), questionCtrl.Delete)
	questionRoute.DELETE("/delete/nuke", utils.AuthMiddleWare(), questionCtrl.DeleteAll)
	questionRoute.GET("/list/:id", utils.AuthMiddleWare(), questionCtrl.GetById)
}

func SetupGameSessionRoutes(router *gin.Engine, db *gorm.DB) {
	game := controllers.GameController{DB: db}
	gameGroup := router.Group("/api/game")
	
	gameGroup.POST("/create", utils.AuthMiddleWare(), game.CreateGame)
	gameGroup.GET("/list", utils.AuthMiddleWare(), game.ListGames)
}