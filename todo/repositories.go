package todo

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/uuid"
	"github.com/karlbehrensg/todo-cosmos-db/config"
)

type repository struct {
	client *azcosmos.ContainerClient
}

func NewRepository(cfg *config.Config) Repository {
	client, err := cfg.Azure.AzClient.Client.NewContainer(cfg.Azure.Database, cfg.Azure.Container)
	if err != nil {
		panic(err)
	}
	return &repository{
		client: client,
	}
}

func (r *repository) CreateTodo(todo *Todo) error {
	// Generate id with uuid and set it to string
	id := uuid.New()
	todo.ID = id.String()

	// specifies the value of the partiton key
	pk := azcosmos.NewPartitionKeyString(todo.ID)

	b, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	// setting the item options upon creating ie. consistency level
	itemOptions := azcosmos.ItemOptions{
		ConsistencyLevel: azcosmos.ConsistencyLevelSession.ToPtr(),
	}

	// this is a helper function that swallows 409 errors
	errorIs409 := func(err error) bool {
		var responseErr *azcore.ResponseError
		return err != nil && errors.As(err, &responseErr) && responseErr.StatusCode == 409
	}

	ctx := context.TODO()
	itemResponse, err := r.client.CreateItem(ctx, pk, b, &itemOptions)

	switch {
	case errorIs409(err):
		log.Printf("Item with partitionkey value %s already exists\n", pk)
	case err != nil:
		return err
	default:
		log.Printf("Status %d. Item %v created. ActivityId %s. Consuming %v Request Units.\n", itemResponse.RawResponse.StatusCode, pk, itemResponse.ActivityID, itemResponse.RequestCharge)
	}

	return nil
}
