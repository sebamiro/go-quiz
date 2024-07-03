package database

type Quiz struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type QuizQuestion struct {
	ID            uint     `json:"id"`
	QuizID        uint     `json:"quiz_id"`
	Question      string   `json:"question"`
	Answers       []string `json:"answers"`
	CorrectAnswer uint
}

type QuizResponse struct {
	ID       uint   `json:"id"`
	QuizID   uint   `json:"quiz_id"`
	Username string `json:"username"`
	Points   uint   `json:"points"`
}
