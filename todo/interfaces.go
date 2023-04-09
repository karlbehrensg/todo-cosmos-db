package todo

import "github.com/gofiber/fiber/v2"

type Repository interface {
	CreateTodo(todo *Todo) error
}

type Cases interface {
	CreateTodo(todo *Todo) error
}

type Service interface {
	CreateTodo(ctx *fiber.Ctx) error
}
