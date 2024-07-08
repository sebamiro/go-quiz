package main

import (
	"github.com/sebamiro/go-quiz/pkg/routes"
	"github.com/sebamiro/go-quiz/pkg/services"
)

func main() {
	c := services.NewContainer()

	routes.BuildRoutes(c)
	c.Web.Start(":3000")
}
