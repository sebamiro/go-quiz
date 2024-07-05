package dto

import "fmt"

type QuizSubmit struct {
	Name   string
	Aswers []uint
}

type QuizEnd struct {
	Name         string
	Points       uint
	Position     uint
	TotalQuizers uint
}

func (q QuizEnd) String() string {
	return fmt.Sprintf(
		"Congratulations %s, you scored %d. You are rank %d of %d quizers.",
		q.Name,
		q.Points,
		q.Position,
		q.TotalQuizers,
	)
}
