package dto

import (
	"fmt"

	"github.com/sebamiro/go-quiz/database"
)

type QuizSubmit struct {
	Name    string `json:"name"`
	Answers []uint `json:"answers"`
}

type ResponseEnd struct {
	Title        string `json:"title"`
	Name         string `json:"name"`
	Points       uint   `json:"points"`
	Position     uint   `json:"position"`
	TotalQuizers uint   `json:"totalquizers"`
	Error        string `json:"error"`
}

func (q ResponseEnd) String() string {
	return fmt.Sprintf(
		"Congratulations %s, you scored %d on the %s. You are rank %d of %d quizers.",
		q.Name,
		q.Points,
		q.Title,
		q.Position,
		q.TotalQuizers,
	)
}

type ResopnseQuestions struct {
	Questions []database.QuizQuestion `json:"questions"`
	Error     string                  `json:"error"`
}

type ResopnseLeaderboard struct {
	Title       string                  `json:"title"`
	Leaderboard []database.QuizResponse `json:"leaderboard"`
	Error       string                  `json:"error"`
}
