package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model

	Code string `json:"code"`
	PlayerName string `json:"player_name"`
	PlayerScore int `json:"player_score"`
	TopicId string `json:"topic_id"`
}