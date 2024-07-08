package routes

import (
	"github.com/sebamiro/go-quiz/pkg/services"
)

func BuildRoutes(c *services.Container) {
	quizGroup := c.Web.Group("quiz/")

	quiz := quiz{db: &c.Database}
	quizGroup.GET("", quiz.Get)
	quizGroup.GET(":quizid", quiz.GetOne)
	quizGroup.GET(":quizid/leaderboard", quiz.GetLeaderboard)
	quizGroup.POST(":quizid", quiz.Post)
}
