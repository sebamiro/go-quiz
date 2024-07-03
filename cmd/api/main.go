package main

import (
	"github.com/sebamiro/go-quiz/services"
)

func main() {
	c := services.NewContainer()
	c.Web.Start(":3000")
}
