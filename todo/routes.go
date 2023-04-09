package todo

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karlbehrensg/todo-cosmos-db/config"
)

var srv Service

func bootstrap(cfg *config.Config) {
	repo := NewRepository(cfg)
	cases := NewCases(repo)
	srv = NewService(cases)
}

func ApplyRoutes(app *fiber.App, cfg *config.Config) {
	bootstrap(cfg)

	todo := app.Group("/todo")

	todo.Post("", srv.CreateTodo)
}
