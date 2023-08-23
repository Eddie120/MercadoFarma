package details

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/core"
	"github.com/mercadofarma/services/db"
	log2 "log"
)

var log = log2.Default()

const detailTable = "details"

type DetailStore interface {
	InsertDetail(ctx context.Context, detail *core.Detail) (*dynamodb.PutItemOutput, error)
}

type ServiceImplementation struct {
	DynamoDbClient db.DynamoDbAPI
	table          string
}

func NewDetailStore(client db.DynamoDbAPI) DetailStore {
	return &ServiceImplementation{
		DynamoDbClient: client,
		table:          detailTable,
	}
}

func (service *ServiceImplementation) InsertDetail(ctx context.Context, detail *core.Detail) (*dynamodb.PutItemOutput, error) {
	item, err := attributevalue.MarshalMap(detail)
	if err != nil {
		log.Println("error calling service.InsertDetail: marshalling map error - ", err.Error())
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(service.table),
	}

	output, err := service.DynamoDbClient.PutItem(ctx, input)
	if err != nil {
		log.Println("error calling service.InsertDetail: put item failed - ", err.Error())
		return nil, err
	}

	return output, nil
}
