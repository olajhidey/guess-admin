package model

import "gorm.io/gorm"

type Topic struct{
	gorm.Model

	Name string `json:"name"`
	Description string `json:"description"`
	CategoryID uint `json:"category_id"`
}