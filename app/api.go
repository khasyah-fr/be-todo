package app

import (
	"be-todo/routes"

	"github.com/labstack/echo/v4"
)

func Run() {
	e := echo.New()
	routes.ActivityRoute(e)
	routes.TodoRoute(e)

	port := "3030"
	go e.Logger.Fatal(e.Start(":" + port))
}
