package repository

import (
	"time"

	"../entities"
	uuid "github.com/satori/go.uuid"
)

type Task entities.Task

//TaskRepository write and read the database
type TaskRepository struct{}

func (t TaskRepository) FindById(id string) Task {
	var task Task
	newId, _ := uuid.FromString(id)
	instance.Where("id = ?", newId).First(&task)
	return task
}

func (t TaskRepository) CreateTask(task entities.Task) {
	task.Start = time.Now()
	task.ID = uuid.Must(uuid.NewV4())
	instance.Create(&task)
}
