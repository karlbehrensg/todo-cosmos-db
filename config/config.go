package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/karlbehrensg/todo-cosmos-db/utils"
)

type Config struct {
	Azure struct {
		Endpoint     string
		Key          string
		AzClient     *utils.AzureClient
		Database     string
		Container    string
		PartitionKey string
	}
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	cfg := &Config{}

	cfg.Azure.Endpoint = os.Getenv("AZURE_COSMOS_ENDPOINT")
	if cfg.Azure.Endpoint == "" {
		log.Fatal("AZURE_COSMOS_ENDPOINT could not be found")
	}

	cfg.Azure.Key = os.Getenv("AZURE_COSMOS_KEY")
	if cfg.Azure.Key == "" {
		log.Fatal("AZURE_COSMOS_KEY could not be found")
	}

	cfg.Azure.Database = os.Getenv("AZURE_COSMOS_DATABASE_NAME")
	if cfg.Azure.Database == "" {
		log.Fatal("AZURE_COSMOS_DATABASE_NAME could not be found")
	}

	cfg.Azure.Container = os.Getenv("AZURE_COSMOS_TODO_CONTAINER_NAME")
	if cfg.Azure.Container == "" {
		log.Fatal("AZURE_COSMOS_TODO_CONTAINER_NAME could not be found")
	}

	cfg.Azure.PartitionKey = os.Getenv("AZURE_COSMOS_TODO_PARTITION_KEY")
	if cfg.Azure.PartitionKey == "" {
		log.Fatal("AZURE_COSMOS_TODO_PARTITION_KEY could not be found")
	}

	return cfg
}
