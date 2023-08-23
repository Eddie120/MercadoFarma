package search_input

import (
	"context"
	"github.com/mercadofarma/services/core"
	searchInputStore "github.com/mercadofarma/services/db/search-input"
)

type SearchInputService interface {
	InsertSearchInput(ctx context.Context, searchInput *core.SearchInput) error
}

type serviceImplementation struct {
	dataStore searchInputStore.SearchInputStore
}

func NewSearchInputService(dataStore searchInputStore.SearchInputStore) SearchInputService {
	return &serviceImplementation{
		dataStore: dataStore,
	}
}

func (sv *serviceImplementation) InsertSearchInput(ctx context.Context, searchInput *core.SearchInput) error {
	return nil
}
