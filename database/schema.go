package database

import (
	"fmt"
	"strings"
)

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

func (q QuizQuestion) String() string {
	lines := make([]string, 0)
	lines = append(lines, fmt.Sprintf("- %s\n", q.Question))
	for i, a := range q.Answers {
		lines = append(lines, fmt.Sprintf("\t%d - %s\n", i, a))
	}
	return strings.Join(lines, "")
}

type QuizResponse struct {
	ID       uint   `json:"id"`
	QuizID   uint   `json:"quiz_id"`
	Username string `json:"username"`
	Points   uint   `json:"points"`
}
