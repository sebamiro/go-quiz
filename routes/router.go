package routes

import (
	"github.com/sebamiro/go-quiz/services"
)

func BuildRoutes(c *services.Container) {
	quizGroup := c.Web.Group("quiz/")

	quiz := quiz{db: &c.Database}
	quizGroup.GET("", quiz.Get)
	quizGroup.GET(":quizid", quiz.GetOne)
	quizGroup.POST(":quizid", quiz.Post)
}
