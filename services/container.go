package services

import (
	"github.com/labstack/echo/v4"
	"github.com/sebamiro/go-quiz/database"
)

type Container struct {
	Web      *echo.Echo
	Database database.Database
}

func NewContainer() *Container {
	var c Container
	c.Web = echo.New()
	c.Database = database.Database{}
	return &c
}
