package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model

	Code string `json:"code"`
	PlayerOneName string `json:"player_one_name"`
	PlayerTwoName string `json:"player_two_name"`
	PlayerOneScore int `json:"player_one_score"`
	PlayerTwoScore int `json:"player_two_score"`
	TopicId string `json:"topic_id"`
}