package details

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/core"
	mock_details "github.com/mercadofarma/services/db/details/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServiceImplementation_InsertDetail(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailStore := new(ServiceImplementation)
	mockDynamoDbClient := mock_details.NewMockDynamoDbAPI(ctrl)
	mockDynamoDbClient.EXPECT().PutItem(context.Background(), gomock.Any(), gomock.Any()).Return(&dynamodb.PutItemOutput{}, nil)

	detailStore.DynamoDbClient = mockDynamoDbClient

	ctx := context.Background()
	detail := &core.Detail{
		Id:             "123",
		CanonicalQuery: fmt.Sprintf("%s::%s::%s", "1.2.3.4", "acetaminofen", "20230718T094837"),
		Status:         core.Found,
		Table: &core.Table{
			TableName: "Basic Information",
			Rows: []*core.Row{
				{
					Cells: []*core.Cell{
						{Name: string(core.ProductReference), Value: "2891"},
						{Name: string(core.ProductName), Value: "AZITROMICINA 500 MG (MK)"},
						{Name: string(core.ProductPresentation), Value: "CAJA X 3 TAB"},
					},
				},
			},
		},
	}

	_, err := detailStore.InsertDetail(ctx, detail)
	c.Nil(err)
}

func TestServiceImplementation_InsertDetail_Error(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	detailStore := new(ServiceImplementation)
	mockDynamoDbClient := mock_details.NewMockDynamoDbAPI(ctrl)
	mockDynamoDbClient.EXPECT().PutItem(context.Background(), gomock.Any(), gomock.Any()).Return(nil, errors.New("db error"))

	detailStore.DynamoDbClient = mockDynamoDbClient

	ctx := context.Background()
	detail := &core.Detail{
		Id:             "123",
		CanonicalQuery: fmt.Sprintf("%s::%s::%s", "1.2.3.4", "acetaminofen", "20230718T094837"),
		Status:         core.Found,
		Table:          &core.Table{},
	}

	_, err := detailStore.InsertDetail(ctx, detail)
	c.NotNil(err)
	c.Equal("db error", err.Error())
}
