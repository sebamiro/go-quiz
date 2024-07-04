package main

import (
	"github.com/sebamiro/go-quiz/routes"
	"github.com/sebamiro/go-quiz/services"
)

func main() {
	c := services.NewContainer()

	routes.BuildRoutes(c)

	c.Web.Start(":3000")
}
