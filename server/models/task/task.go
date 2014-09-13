package task

import (
	"fmt"
)

type Task struct{}

func GetTasks() ([]Task, error) {
	fmt.Println("GetTasks")
	return []Task{}, nil
}

func GetTask(id int) (Task, error) {
	fmt.Println("GetTask")
	return Task{}, nil
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "task created")
}

func PutTask(t *Task) error {
	fmt.Println("PutTask")
	return nil
}
