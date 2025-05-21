package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
