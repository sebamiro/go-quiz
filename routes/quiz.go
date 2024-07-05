package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/go-quiz/database"
	"github.com/sebamiro/go-quiz/dto"
)

type quiz struct {
	db *database.Database
}

func (q quiz) Get(ctx echo.Context) error {
	quizes := q.db.GetQuizes()
	return ctx.JSON(http.StatusOK, quizes)
}

func (q quiz) GetLeaderboard(ctx echo.Context) error {
	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		return err
	}
	leaderboard, err := q.db.GetQuizResponsesOrderdByPoints(uint(quizId))
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, leaderboard)
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

var userCount uint = 0

func (q quiz) Post(ctx echo.Context) error {
	var QuizSubmit dto.QuizSubmit

	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		log.Println("[ERROR] Atoi:", err)
		return err
	}
	quizQuestions, err := q.db.GetQuizQuestions(uint(quizId))
	if err != nil {
		log.Println("[ERROR] DB:", err)
		return err
	}
	if err := json.NewDecoder(ctx.Request().Body).Decode(&QuizSubmit); err != nil {
		log.Println("[ERROR] json:", err)
		return err
	}
	if len(quizQuestions) != len(QuizSubmit.Aswers) {
		log.Println("[ERROR] len:", err)
		return errors.New("Invalid answer len")
	}

	var correctAnswers uint = 0
	for i, q := range quizQuestions {
		if q.CorrectAnswer == QuizSubmit.Aswers[i] {
			correctAnswers++
		}
	}

	if QuizSubmit.Name == "" {
		QuizSubmit.Name = "user"
	}
	username := fmt.Sprintf("%s#%.4d", QuizSubmit.Name, userCount)
	err = q.db.AddQuizResponse(uint(quizId), database.QuizResponse{
		QuizID:   uint(quizId),
		Username: username,
		Points:   correctAnswers,
	})
	userCount++
	if err != nil {
		return err
	}

	leaderboard, _ := q.db.GetQuizResponsesOrderdByPoints(uint(quizId))
	position := 0
	for i, l := range leaderboard {
		if l.Username == username {
			position = i
			break
		}
	}

	return ctx.JSON(http.StatusOK, dto.QuizEnd{
		Name:         username,
		Points:       correctAnswers,
		TotalQuizers: uint(len(leaderboard)),
		Position:     uint(position),
	})
}
