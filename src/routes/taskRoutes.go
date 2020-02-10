package routes

import (
	"../handlers"
	"github.com/labstack/echo"
)

//SetTaskRoutes define task routes
func SetTaskRoutes(e *echo.Echo) {
	taskRoutes := e.Group("/task")
	taskRoutes.GET("/", handlers.GetTask)

}
