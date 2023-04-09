package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/karlbehrensg/todo-cosmos-db/todo"
	"github.com/karlbehrensg/todo-cosmos-db/utils"
)

var databaseName = "personal"
var containerName = "todo"
var partitionKey = "/id"

var endpoint string
var key string

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	endpoint = os.Getenv("AZURE_COSMOS_ENDPOINT")
	if endpoint == "" {
		log.Fatal("AZURE_COSMOS_ENDPOINT could not be found")
	}

	key = os.Getenv("AZURE_COSMOS_KEY")
	if key == "" {
		log.Fatal("AZURE_COSMOS_KEY could not be found")
	}
}

func main() {
	azClient, err := utils.NewAzureClient(endpoint, key)
	if err != nil {
		log.Fatal("Failed to create Azure Cosmos DB client: ", err)
	}

	if err := azClient.CreateDatabase(databaseName); err != nil {
		log.Fatal("Failed to create database: ", err)
	}

	if err := azClient.CreateContainer(databaseName, containerName, partitionKey); err != nil {
		log.Fatal("Failed to create container: ", err)
	}

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	todo.ApplyRoutes(app)

	app.Listen(":3000")
}
