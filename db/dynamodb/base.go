package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"os"
)

const defaultRegion = "us-west-2"

func NewDynamoDBClient(ctx context.Context) *dynamodb.Client {
	defaultConfig, err := config.LoadDefaultConfig(ctx, func(opts *config.LoadOptions) error {
		envRegion := os.Getenv("AWS_REGION")
		if envRegion != "" {
			opts.Region = envRegion
		} else {
			opts.Region = defaultRegion
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(defaultConfig)
}
