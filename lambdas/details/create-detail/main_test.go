package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	dtServiceMock "github.com/mercadofarma/services/services/details/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"io"
	"os"
	"testing"
)

func TestHappyPath(t *testing.T) {
	c := assert.New(t)

	eventJson, _ := os.Open("./samples/event.json")
	defer func(eventJson *os.File) {
		err := eventJson.Close()
		if err != nil {
			fmt.Println("error closing json file")
			os.Exit(1)
		}
	}(eventJson)

	byteValue, _ := io.ReadAll(eventJson)
	var event events.SNSEvent
	err := json.Unmarshal(byteValue, &event)
	c.Nil(err)

	ctrl := gomock.NewController(t)
	ctx := context.Background()
	dtMock := dtServiceMock.NewMockDetailService(ctrl)

	// TODO: create detail object according sample file
	detail := gomock.Any()
	dtMock.EXPECT().InsertDetail(ctx, detail).Return(nil)
	detailService = dtMock

	err = handler(ctx, &event)
	c.Nil(err)
}

func TestDetailServiceError(t *testing.T) {
	c := assert.New(t)

	eventJson, _ := os.Open("./samples/event.json")
	defer func(eventJson *os.File) {
		err := eventJson.Close()
		if err != nil {
			fmt.Println("error closing json file")
			os.Exit(1)
		}
	}(eventJson)

	byteValue, _ := io.ReadAll(eventJson)
	var event events.SNSEvent
	err := json.Unmarshal(byteValue, &event)
	c.Nil(err)

	ctrl := gomock.NewController(t)
	ctx := context.Background()
	dtMock := dtServiceMock.NewMockDetailService(ctrl)

	// TODO: create detail object according sample file
	detail := gomock.Any()
	dtMock.EXPECT().InsertDetail(ctx, detail).Return(assert.AnError)
	detailService = dtMock

	err = handler(ctx, &event)
	c.Error(err, assert.AnError)
}
