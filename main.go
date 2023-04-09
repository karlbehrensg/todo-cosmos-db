package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/karlbehrensg/todo-cosmos-db/config"
	"github.com/karlbehrensg/todo-cosmos-db/todo"
	"github.com/karlbehrensg/todo-cosmos-db/utils"
)

var cfg *config.Config
var err error

func init() {
	cfg = config.NewConfig()
}

func main() {
	cfg.Azure.AzClient, err = utils.NewAzureClient(cfg.Azure.Endpoint, cfg.Azure.Key)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	if err := cfg.Azure.AzClient.CreateDatabase(cfg.Azure.Database); err != nil {
		log.Fatal("Failed to create database: ", err)
	}

	if err := cfg.Azure.AzClient.CreateContainer(cfg.Azure.Database, cfg.Azure.Container, cfg.Azure.PartitionKey); err != nil {
		log.Fatal("Failed to create container: ", err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	todo.ApplyRoutes(app, cfg)

	app.Listen(":3000")
}
