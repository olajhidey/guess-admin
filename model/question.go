package model

import (
	"gorm.io/gorm"
)


type Question struct {
	gorm.Model

	ImageUrl string `json:"image_url"`
	Answer   string `json:"answer"`
	Option1  string `json:"option1"`
	Option2  string `json:"option2"`
	Option3  string `json:"option3,omitempty"`
	Option4  string `json:"option4,omitempty"`
	TopicId string `json:"topic_id"`
}