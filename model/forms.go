package model

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type CategoryForm struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type TopicForm struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID  int    `json:"category_id" binding:"required"`
}

type QuestionForm struct {
	ImageUrl string `json:"image_url" binding:"required"`
	Answer   string `json:"answer" binding:"required"`
	Option1  string `json:"option1" binding:"required"`
	Option2  string `json:"option2" binding:"required"`
	Option3  string `json:"option3,omitempty"`
	Option4  string `json:"option4,omitempty"`
	TopicId  string `json:"topic_id" binding:"required"`
}

type GameForm struct {
	Code        string `json:"code" binding:"required"`
	PlayerName  string `json:"player_name" binding:"required"`
	PlayerScore int    `json:"player_score" binding:"required"`
	TopicId     string `json:"topic_id" binding:"required"`
}
