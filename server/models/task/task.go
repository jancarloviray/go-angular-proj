package task

import (
	"fmt"
)

type Task struct {
	Title string
	Done  bool
}

var tasks []Task = []Task{
	Task{Title: "Get Milk", Done: true},
	Task{Title: "Drink Milk", Done: false},
}

func List() ([]Task, error) {
	return tasks, nil
}

func Get(id string) (Task, error) {
	fmt.Println("GetTask")
	return Task{}, nil
}

func Create(t Task) (int, error) {
	fmt.Println("CreateTask")
	return 0, nil
}

func Update(id string, t Task) error {
	fmt.Println("PutTask")
	return nil
}

func Delete(id string) error {
	fmt.Println("DeleteTask")
	return nil
}
