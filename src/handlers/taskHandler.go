package handlers

import (
	"net/http"

	"../entities"
	"../repository"
	"github.com/labstack/echo"
)

type Task entities.Task

var taskRepository repository.TaskRepository

//GetTask get all task
func GetTask(c echo.Context) error {
	id := c.QueryParam("id")
	task := taskRepository.FindById(id)
	return c.JSON(http.StatusOK, task)
}

//CreateTask sad
func CreateTask(c echo.Context) error {
	var task entities.Task
	c.Bind(&task)
	taskRepository.CreateTask(task)
	return c.JSON(http.StatusOK, task)
}
