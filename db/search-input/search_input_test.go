package search_input

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/core"
	dynamodb_mock "github.com/mercadofarma/services/db/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServiceImplementation_InsertSearchInput(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	searchInputStore := new(ServiceImplementation)
	mockDynamoDbClient := dynamodb_mock.NewMockDynamoDbAPI(ctrl)
	mockDynamoDbClient.EXPECT().PutItem(context.Background(), gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)

	searchInputStore.DynamoDbClient = mockDynamoDbClient

	ctx := context.Background()
	input := &core.SearchInput{
		Query:   "dolex",
		Country: "CO",
		City:    "Cali",
	}

	_, err := searchInputStore.InsertSearchInput(ctx, input)
	c.Nil(err)
}

func TestServiceImplementation_InsertSearchInput_Error(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	searchInputStore := new(ServiceImplementation)
	mockDynamoDbClient := dynamodb_mock.NewMockDynamoDbAPI(ctrl)
	mockDynamoDbClient.EXPECT().PutItem(context.Background(), gomock.Any(), gomock.Any()).Return(nil, errors.New("db error"))

	searchInputStore.DynamoDbClient = mockDynamoDbClient

	ctx := context.Background()
	input := &core.SearchInput{
		Query:   "dolex",
		Country: "CO",
		City:    "Cali",
	}

	_, err := searchInputStore.InsertSearchInput(ctx, input)
	c.NotNil(err)
	c.Equal("db error", err.Error())
}
