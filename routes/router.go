package routes

import (
	"github.com/sebamiro/go-quiz/services"
)

func BuildRoutes(c *services.Container) {
	quiz := c.Web.Group("quiz/")

	quiz.GET("/", nil)
	quiz.GET("/:quizid", nil)
	quiz.POST("/:quizid", nil)
}
