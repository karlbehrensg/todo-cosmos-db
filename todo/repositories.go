package todo

import "fmt"

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CreateTodo(todo *Todo) error {
	fmt.Printf("todo: %v\n", todo)
	return nil
}
