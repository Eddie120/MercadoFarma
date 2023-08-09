package details

import (
	"context"
	"errors"
	"fmt"
	"github.com/mercadofarma/services/core"
	mock_details "github.com/mercadofarma/services/db/details/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServiceImplementation_InsertDetail(t *testing.T) {
	c := assert.New(t)

	ctr := gomock.NewController(t)
	mockDetailStore := mock_details.NewMockDetailFactory(ctr)

	ctx := context.Background()
	detail := &core.Detail{
		Status:         core.Found,
		CanonicalQuery: fmt.Sprintf("%s::%s::%s", "1.2.3.4", "acetaminofen", "20230718T094837"),
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

	mockDetailStore.EXPECT().InsertDetail(ctx, detail).Return(nil, nil)
	svc := &serviceImplementation{
		dataStore: mockDetailStore,
	}

	err := svc.InsertDetail(ctx, detail)
	c.Nil(err)
}

func TestServiceImplementation_InsertDetail_DbError(t *testing.T) {
	c := assert.New(t)

	ctr := gomock.NewController(t)
	mockDetailStore := mock_details.NewMockDetailFactory(ctr)

	ctx := context.Background()
	detail := &core.Detail{
		Status:         core.Found,
		CanonicalQuery: fmt.Sprintf("%s::%s::%s", "acetaminofen", "123", "xyz"),
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

	mockDetailStore.EXPECT().InsertDetail(ctx, detail).Return(nil, errors.New("db error"))
	svc := &serviceImplementation{
		dataStore: mockDetailStore,
	}

	err := svc.InsertDetail(ctx, detail)
	c.Equal(err.Error(), "db error")
}
