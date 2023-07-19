package details

import (
	"context"
	"errors"
	"fmt"
	"github.com/mercadofarma/services/core"
	mockdb "github.com/mercadofarma/services/db/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServiceImplementation_InsertDetail(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dataAccess := mockdb.NewMockDataAccess(ctrl)
	detailService := NewDetailService(dataAccess)

	ctx := context.Background()
	detail := &core.Detail{
		Id:             int64(1),
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

	dataAccess.EXPECT().ExecWithContext(ctx, gomock.Any(), gomock.Any()).Return(nil, nil)
	err := detailService.InsertDetail(ctx, detail)
	c.Nil(err)
}

func TestServiceImplementation_InsertDetail_NilValue(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dataAccess := mockdb.NewMockDataAccess(ctrl)
	detailService := NewDetailService(dataAccess)

	ctx := context.Background()

	err := detailService.InsertDetail(ctx, nil)
	c.NotNil(err)
	c.Equal("detail cannot be nil", err.Error())
}

func TestServiceImplementation_InsertDetail_Error(t *testing.T) {
	c := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dataAccess := mockdb.NewMockDataAccess(ctrl)
	detailService := NewDetailService(dataAccess)

	ctx := context.Background()
	detail := &core.Detail{
		Id:             int64(1),
		CanonicalQuery: fmt.Sprintf("%s::%s::%s", "1.2.3.4", "acetaminofen", "20230718T094837"),
		Status:         core.Found,
		Table:          &core.Table{},
	}

	dataAccess.EXPECT().ExecWithContext(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("db error"))
	err := detailService.InsertDetail(ctx, detail)
	c.NotNil(err)
	c.Equal("db error", err.Error())
}
