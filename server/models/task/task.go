package task

import (
	"fmt"
)

type Task struct{}

func List() ([]Task, error) {
	fmt.Println("List")
	return []Task{}, nil
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
