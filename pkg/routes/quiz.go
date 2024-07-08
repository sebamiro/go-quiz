package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sebamiro/go-quiz/database"
	"github.com/sebamiro/go-quiz/pkg/dto"
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
		return ctx.JSON(http.StatusBadRequest, dto.ResopnseLeaderboard{Error: err.Error()})
	}
	quiz, err := q.db.GetQuizesById(uint(quizId))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dto.ResopnseLeaderboard{Error: err.Error()})
	}
	leaderboard, _ := q.db.GetQuizResponsesOrderdByPoints(uint(quizId))
	return ctx.JSON(http.StatusOK, dto.ResopnseLeaderboard{
		Title:       quiz.Title,
		Leaderboard: leaderboard,
	})
}

func (q quiz) GetOne(ctx echo.Context) error {
	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResopnseQuestions{Error: err.Error()})
	}
	quizQuestions, err := q.db.GetQuizQuestions(uint(quizId))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dto.ResopnseQuestions{Error: err.Error()})
	}
	return ctx.JSON(http.StatusOK, dto.ResopnseQuestions{Questions: quizQuestions})
}

var userCount uint = 0

func (q quiz) Post(ctx echo.Context) error {
	var QuizSubmit dto.QuizSubmit

	quizId, err := strconv.Atoi(ctx.Param("quizid"))
	if err != nil {
		log.Println("[ERROR] Atoi:", err)
		return ctx.JSON(http.StatusBadRequest, dto.ResponseEnd{Error: err.Error()})
	}
	quiz, err := q.db.GetQuizesById(uint(quizId))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, dto.ResponseEnd{Error: err.Error()})
	}
	quizQuestions, _ := q.db.GetQuizQuestions(uint(quizId))

	if err := json.NewDecoder(ctx.Request().Body).Decode(&QuizSubmit); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseEnd{Error: err.Error()})
	}
	if len(quizQuestions) != len(QuizSubmit.Answers) {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseEnd{Error: "Invalid answer len"})
	}

	var correctAnswers uint = 0
	for i, q := range quizQuestions {
		if q.CorrectAnswer == QuizSubmit.Answers[i] {
			correctAnswers++
		}
	}

	if QuizSubmit.Name == "" {
		QuizSubmit.Name = "user"
	}
	username := fmt.Sprintf("%s#%.4d", QuizSubmit.Name, userCount)
	_ = q.db.AddQuizResponse(uint(quizId), database.QuizResponse{
		QuizID:   uint(quizId),
		Username: username,
		Points:   correctAnswers,
	})
	userCount++

	leaderboard, _ := q.db.GetQuizResponsesOrderdByPoints(uint(quizId))
	position := 0
	for i, l := range leaderboard {
		if l.Username == username {
			position = i
			break
		}
	}

	return ctx.JSON(http.StatusOK, dto.ResponseEnd{
		Name:         username,
		Points:       correctAnswers,
		Title:        quiz.Title,
		TotalQuizers: uint(len(leaderboard)),
		Position:     uint(position + 1),
	})
}
