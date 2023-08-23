package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/mercadofarma/services/aws"
	"github.com/mercadofarma/services/core"
	detailstore "github.com/mercadofarma/services/db/details"
	dtservice "github.com/mercadofarma/services/services/details"
)

// setting up the services
var globalConfig = aws.NewConfig(context.Background())
var detailStore = detailstore.NewDetailStore(dynamodb.NewFromConfig(globalConfig))
var detailService = dtservice.NewDetailService(detailStore)

const errEmptyMessage = "empty message"

type eventHandler struct {
	service dtservice.DetailService
}

func (handler *eventHandler) processRecord(ctx context.Context, detail *core.Detail) error {
	return handler.service.InsertDetail(ctx, detail)
}

func handler(ctx context.Context, event *events.SNSEvent) error {
	if len(event.Records) == 0 {
		return errors.New(errEmptyMessage)
	}

	svc := &eventHandler{
		service: detailService,
	}

	for _, record := range event.Records {
		detailRecord := &core.Detail{}
		snsRecord := record.SNS

		err := json.Unmarshal([]byte(snsRecord.Message), detailRecord)
		if err != nil {
			return err
		}

		if err = svc.processRecord(ctx, detailRecord); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
