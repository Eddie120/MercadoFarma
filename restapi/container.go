package restapi

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/aws"
	"github.com/mercadofarma/services/controllers"
	"github.com/mercadofarma/services/db"
	"github.com/mercadofarma/services/db/details"
	searchInput "github.com/mercadofarma/services/db/search-input"
	search_input "github.com/mercadofarma/services/services/search-input"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()

	provider := func(constructor interface{}, opt ...dig.ProvideOption) {
		err := container.Provide(constructor, opt...)
		panic(err)
	}

	provider(func() db.DynamoDbAPI {
		cfg := aws.NewConfig(context.Background())
		return dynamodb.NewFromConfig(cfg)
	})

	provider(func(client db.DynamoDbAPI) details.DetailStore {
		return details.NewDetailStore(client)
	})

	provider(func(client db.DynamoDbAPI) searchInput.SearchInputStore {
		return searchInput.NewSearchInputStore(client)
	})

	provider(func(searchStore searchInput.SearchInputStore) search_input.SearchInputService {
		return search_input.NewSearchInputService(searchStore)
	})

	// controllers
	provider(controllers.NewSearchInputController)

	return container
}
