package todo

import (
	"github.com/gofiber/fiber/v2"
)

var srv Service

func bootstrap() {
	repo := NewRepository()
	cases := NewCases(repo)
	srv = NewService(cases)
}

func ApplyRoutes(app *fiber.App) {
	bootstrap()

	todo := app.Group("/todo")

	todo.Post("", srv.CreateTodo)
}
