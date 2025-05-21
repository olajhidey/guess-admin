package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olajhidey/guess-admin/config"
	"github.com/olajhidey/guess-admin/database"
	"github.com/olajhidey/guess-admin/routes"
)

func main(){
	config.LoadConfig()
	db := database.ConnectDB()
	router := gin.Default()
	routes.LoadRoutes(router, db)
	router.Run(":" + config.Port)
}