package business

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/mercadofarma/services/repos/models"
	"time"
)

type BusinessRepo interface {
	CreateBusiness(ctx context.Context, businessRecord *models.Business) (*dynamodb.PutItemOutput, error)
}

type BusinessRepoImpl struct {
	client    *dynamodb.Client
	tableName string
}

const businessPrefix = "BNS"

func NewBusinessRepo(client *dynamodb.Client) *BusinessRepoImpl {
	return &BusinessRepoImpl{
		client:    client,
		tableName: "business",
	}
}

func (svc *BusinessRepoImpl) CreateBusiness(ctx context.Context, record *models.Business) (*dynamodb.PutItemOutput, error) {
	current := time.Now()
	record.BusinessId = fmt.Sprintf("%s%s", businessPrefix, uuid.New().String())
	record.CreationDate = &current
	record.UpdateDate = &current

	item, _ := attributevalue.MarshalMap(record)
	output, err := svc.client.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(svc.tableName),
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
