package search_input

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/core"
	"github.com/mercadofarma/services/db"
	log2 "log"
)

var log = log2.Default()

const searchInputsTable = "search-inputs"

type SearchInputStore interface {
	InsertSearchInput(ctx context.Context, searchInput *core.SearchInput) (*dynamodb.PutItemOutput, error)
}

type ServiceImplementation struct {
	DynamoDbClient db.DynamoDbAPI
	table          string
}

func NewSearchInputStore() SearchInputStore {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return &ServiceImplementation{
		DynamoDbClient: client,
		table:          searchInputsTable,
	}
}

func (service *ServiceImplementation) InsertSearchInput(ctx context.Context, searchInput *core.SearchInput) (*dynamodb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(searchInput)
	if err != nil {
		log.Println("error calling service.InsertSearchInput: marshalling map error - ", err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(service.table),
	}

	output, err := service.DynamoDbClient.PutItem(ctx, input)
	if err != nil {
		log.Println("error calling service.InsertSearchInput: put item failed - ", err.Error())
		return nil, err
	}

	return output, nil
}
