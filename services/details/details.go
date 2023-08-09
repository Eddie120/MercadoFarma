package details

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/mercadofarma/services/core"
	detailstore "github.com/mercadofarma/services/db/details"
)

var (
	errMissingStatus         = errors.New("status cannot be empty")
	errInvalidDetailStatus   = errors.New("invalid detail status")
	errMissingCanonicalQuery = errors.New("canonical query is required")
	errMissingTable          = errors.New("table cannot be nil")

	validDetailStatuses = map[core.DetailStatus]bool{
		core.None:     true,
		core.Error:    true,
		core.NotFound: true,
		core.Found:    true,
	}
)

type DetailService interface {
	InsertDetail(ctx context.Context, detail *core.Detail) error
}

type serviceImplementation struct {
	dataStore detailstore.DetailStore
}

func NewDetailService() DetailService {
	return &serviceImplementation{
		dataStore: detailstore.NewDetailStore(),
	}
}

func (sv *serviceImplementation) InsertDetail(ctx context.Context, detail *core.Detail) error {
	if detail.Status == "" {
		return errMissingStatus
	}

	if !validDetailStatuses[detail.Status] {
		return errInvalidDetailStatus
	}

	if detail.CanonicalQuery == "" {
		return errMissingCanonicalQuery
	}

	if detail.Table == nil {
		return errMissingTable
	}

	detail.Id = uuid.New().String()
	_, err := sv.dataStore.InsertDetail(ctx, detail)

	return err
}
