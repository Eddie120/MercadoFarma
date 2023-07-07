package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mercadofarma/services/core"
	lopido "github.com/mercadofarma/services/scrapers/colombia/lopido.com/scraper"
	"log"
	"net/http"
)

var config = core.CrawlerConfig{
	CrawlerName: "lopido.com",
}

type Collector struct{}

func NewCollector() core.Crawler {
	return &Collector{}
}

func (c *Collector) Crawl(ctx context.Context, input *core.SearchInput, notifier core.Notifier) error {
	logger := log.Default()
	logger.Println("Starting Crawl for lopido.com ...")

	client := http.Client{}
	scraper, err := lopido.NewScraper(&client, input.Query, input.Country, input.City, logger)
	if err != nil {
		logger.Println("Create a new scraper failed: ", err.Error())
		return err
	}

	err = scraper.Start(ctx)
	if err != nil {
		logger.Println("Scraper execution failed: ", err.Error())
		return err
	}

	return nil
}

func main() {
	lambda.Start(core.CrawlerProcessor(NewCollector, config))
}
