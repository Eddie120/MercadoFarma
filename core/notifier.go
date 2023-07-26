package core

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	sns_ "github.com/mercadofarma/services/core/sns"
)

// TODO:setup topic arn for detail service
const topicARN = ""

var client *sns.Client

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client = sns.NewFromConfig(cfg)
}

type HandlerFunc func(ctx context.Context, detail *Detail) error

func (f HandlerFunc) Notify(ctx context.Context, detail *Detail) error {
	return f(ctx, detail)
}

var notifier HandlerFunc = func(ctx context.Context, detail *Detail) error {
	data, err := json.Marshal(detail)
	if err != nil {
		return err
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(data)),
		TopicArn: aws.String(topicARN),
	}

	_, err = sns_.PublishMessage(ctx, client, input)
	return err
}
