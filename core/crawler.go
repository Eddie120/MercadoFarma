package core

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
)

type factory func() Crawler

type CrawlerConfig struct {
	CrawlerName string `json:"crawler_name"`
}

type CrawlerHandlerProcessor struct {
	instance Crawler
	config   CrawlerConfig
	inputs   *SearchInput
}

func NewCrawlerHandlerProcessor(instance Crawler, data []byte, config CrawlerConfig) (*CrawlerHandlerProcessor, error) {
	inputs, err := NewSearchInput(data)
	if err != nil {
		return nil, err
	}

	return &CrawlerHandlerProcessor{
		instance: instance,
		config:   config,
		inputs:   inputs,
	}, nil
}

func (handler *CrawlerHandlerProcessor) Run(ctx context.Context) error {
	var notifier HandlerFunc
	return handler.instance.Crawl(ctx, handler.inputs, notifier)
}

type Notifier interface {
	Notify(details *Detail) error
}

type Crawler interface {
	Crawl(ctx context.Context, input *SearchInput, notifier Notifier) error
}

func CrawlerProcessor(method factory, config CrawlerConfig) func(ctx context.Context, event *events.SNSEvent) error {
	return func(ctx context.Context, event *events.SNSEvent) error {
		crawler, err := NewCrawlerHandlerProcessor(method(), []byte(event.Records[0].SNS.Message), config)
		if err != nil {
			return err
		}

		return crawler.Run(ctx)
	}
}
