package routes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/go-quiz/database"
)

type quiz struct {
	db *database.Database
}

func (q quiz) Get(ctx echo.Context) error {
	quizes := q.db.GetQuizes()
	return ctx.JSON(http.StatusOK, quizes)
}

func (q quiz) GetOne(ctx echo.Context) error {
	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		return err
	}
	quizQuestions, err := q.db.GetQuizQuestions(uint(quizId))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, quizQuestions)
}

type quizSubmit struct {
	name   string
	aswers []uint
}

var unamedCount uint = 0

func (q quiz) Post(ctx echo.Context) error {
	var quizSubmit quizSubmit

	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		return err
	}
	quizQuestions, err := q.db.GetQuizQuestions(uint(quizId))
	if err != nil {
		return err
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&quizSubmit); err != nil {
		return err
	}
	if len(quizQuestions) != len(quizSubmit.aswers) {
		return errors.New("Invalid answer len")
	}

	var correctAnswers uint = 0
	for i, q := range quizQuestions {
		if q.CorrectAnswer == quizSubmit.aswers[i] {
			correctAnswers++
		}
	}
	err = q.db.AddQuizResponse(uint(quizId), database.QuizResponse{
		QuizID:   uint(quizId),
		Username: quizSubmit.name,
		Points:   correctAnswers,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, correctAnswers)
}
