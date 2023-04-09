package todo

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title" validate:"required"`
	Status string `json:"status" validate:"required"`
}
