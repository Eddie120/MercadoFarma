package restapi

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	_dynamodb "github.com/mercadofarma/services/db/dynamodb"
	"github.com/mercadofarma/services/repos/business"
	businessService "github.com/mercadofarma/services/services/business"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	container := dig.New()

	provider := func(constructor interface{}, opt ...dig.ProvideOption) {
		err := container.Provide(constructor, opt...)
		panic(err)
	}

	provider(func() *dynamodb.Client {
		return _dynamodb.NewDynamoDBClient(context.Background())
	})

	// business storage
	provider(func(client *dynamodb.Client) business.BusinessRepo {
		return business.NewBusinessRepo(client)
	})

	// business service
	provider(func(storage business.BusinessRepo) businessService.BusinessService {
		return businessService.NewBusinessService(storage)
	})

	return container
}
