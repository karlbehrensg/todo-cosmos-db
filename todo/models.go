package todo

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required"`
}
