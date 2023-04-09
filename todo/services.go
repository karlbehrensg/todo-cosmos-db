package todo

import "github.com/gofiber/fiber/v2"

type services struct {
	cases Cases
}

func NewService(cases Cases) Service {
	return &services{
		cases: cases,
	}
}

func (s *services) CreateTodo(ctx *fiber.Ctx) error {
	var todo Todo
	if err := ctx.BodyParser(&todo); err != nil {
		return err
	}

	if err := s.cases.CreateTodo(&todo); err != nil {
		return err
	}

	return ctx.JSON(todo)
}
