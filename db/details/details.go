package details

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/core"
	log2 "log"
)

var log = log2.Default()

const detailTable = "details"

type DynamoDbAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type DetailStore interface {
	InsertDetail(ctx context.Context, detail *core.Detail) (*dynamodb.PutItemOutput, error)
}

type ServiceImplementation struct {
	DynamoDbClient DynamoDbAPI
	table          string
}

func NewDetailStore() DetailStore {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := dynamodb.NewFromConfig(cfg)

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
