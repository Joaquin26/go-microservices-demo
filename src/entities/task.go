package entities

import (
	"time"

	pgis "github.com/aodin/aspect/postgis"
	uuid "github.com/satori/go.uuid"
)

//Task is a public struct
type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string
	Content     string
	Start       time.Time
	coordinates pgis.Polygon
}

//CreateTask return a new tarea
func CreateTask(title string, content string) Task {
	start := time.Now()
	return Task{
		Title:   title,
		Content: content,
		Start:   start,
	}
}
