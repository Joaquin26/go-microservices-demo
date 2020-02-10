package handlers

import (
	"fmt"
	"net/http"

	"../entities"
	"github.com/labstack/echo"
)

//GetTask get all task
func GetTask(c echo.Context) error {
	newTask := entities.CreateTask("titulo", "content")
	return c.String(http.StatusOK, fmt.Sprintf("your task title is : %s\n", newTask.Start))
}
