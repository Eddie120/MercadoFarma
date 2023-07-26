package sns

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	mock_sns "github.com/mercadofarma/services/core/sns/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestPublishMessage(t *testing.T) {
	c := assert.New(t)

	input := &sns.PublishInput{
		Message:  aws.String("Hello"),
		TopicArn: aws.String("dummytopicarn"),
	}

	ctrl := gomock.NewController(t)
	mockService := mock_sns.NewMockService(ctrl)

	ctx := context.Background()
	mockService.EXPECT().Publish(ctx, gomock.Any(), gomock.Any()).Return(&sns.PublishOutput{
		MessageId: aws.String("123"),
	}, nil)

	response, err := PublishMessage(ctx, mockService, input)
	c.Nil(err)
	c.Equal(*response.MessageId, "123")
}

func TestPublishMessage_Error(t *testing.T) {
	c := assert.New(t)

	input := &sns.PublishInput{
		Message:  aws.String("Hello"),
		TopicArn: aws.String("dummytopicarn"),
	}

	ctrl := gomock.NewController(t)
	mockService := mock_sns.NewMockService(ctrl)

	ctx := context.Background()
	mockService.EXPECT().Publish(ctx, gomock.Any(), gomock.Any()).Return(nil, errors.New("sns failed"))

	response, err := PublishMessage(ctx, mockService, input)
	c.Nil(response)
	c.NotNil(err)
	c.Equal(err.Error(), "sns failed")

	response, err = PublishMessage(ctx, mockService, nil)
	c.Nil(response)
	c.NotNil(err)
	c.Equal(err.Error(), "input cannot be nil")
}
