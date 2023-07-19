package details

import (
	"context"
	"errors"
	"github.com/mercadofarma/services/core"
	"github.com/mercadofarma/services/db"
	"log"
)

var logger = log.Default()

type Service interface {
	InsertDetail(ctx context.Context, detail *core.Detail) error
}

type ServiceImplementation struct {
	dbAccess db.DataAccess
}

func NewDetailService(dbAccess db.DataAccess) Service {
	return &ServiceImplementation{
		dbAccess: dbAccess,
	}
}

func (service *ServiceImplementation) InsertDetail(ctx context.Context, detail *core.Detail) error {
	if detail == nil {
		return errors.New("detail cannot be nil")
	}

	if detail.CanonicalQuery == "" {
		return errors.New("canonical query is required")
	}

	if detail.Table == nil {
		return errors.New("table cannot be nil")
	}

	const query string = "INSERT INTO mercadofarma.details (canonical_query,status,message_error,table) VALUES(?,?,?,?);"
	args := []interface{}{detail.CanonicalQuery, detail.Status, detail.MessageError, detail.Table}

	_, err := service.dbAccess.ExecWithContext(ctx, query, args...)
	if err != nil {
		logger.Println("Call service.InsertDetail failed: ", err.Error())

		return err
	}

	return nil
}
