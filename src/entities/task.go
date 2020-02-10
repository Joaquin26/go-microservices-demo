package entities

import "time"

//task is a public struct
type task struct {
	Title   string
	Content string
	Start   time.Time
}

//CreateTask return a new tarea
func CreateTask(title string, content string) task {
	start := time.Now()
	return task{
		Title:   title,
		Content: content,
		Start:   start,
	}
}
