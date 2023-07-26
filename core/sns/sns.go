package sns

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Service interface {
	Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

func PublishMessage(ctx context.Context, svc Service, input *sns.PublishInput) (*sns.PublishOutput, error) {
	if input == nil {
		return nil, errors.New("input cannot be nil")
	}

	return svc.Publish(ctx, input)
}
