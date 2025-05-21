package database

import (
	"github.com/olajhidey/guess-admin/model"
	"github.com/olajhidey/guess-admin/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if utils.ErrorNotNil(err) {
		utils.LogMessage("Failed to connect to database")
	}
	utils.LogMessage("Connected to database")

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Topic{}, &model.Question{}, &model.Game{})
	return db
}