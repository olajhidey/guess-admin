package config

import (
	"github.com/joho/godotenv"
	"github.com/olajhidey/guess-admin/utils" 
)

var Port string
var SecretKey string

func LoadConfig() {
	err := godotenv.Load()

	if utils.ErrorNotNil(err) {
		utils.LogMessage("Error loading .env file")
	}
	Port = utils.GetEnv("PORT", "3000")
	SecretKey = utils.GetEnv("SECRET_KEY", "secret")
}