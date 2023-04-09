package todo

import "fmt"

var validStatus = [2]string{"pending", "completed"}

type cases struct {
	repo Repository
}

func NewCases(repo Repository) Cases {
	return &cases{
		repo: repo,
	}
}

func (r *cases) validateStatus(status string) error {
	for _, s := range validStatus {
		if s == status {
			return nil
		}
	}
	return fmt.Errorf("invalid status: %s", status)
}

func (r *cases) CreateTodo(todo *Todo) error {
	if err := r.validateStatus(todo.Status); err != nil {
		return err
	}
	if err := r.repo.CreateTodo(todo); err != nil {
		return err
	}
	return nil
}
