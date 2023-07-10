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
	logger.Println("Starting Crawler for lopido.com")

	detail := core.NewDetail()
	client := http.Client{}

	scraper, err := lopido.NewScraper(&client, input.Query, input.Country, input.City, logger)
	if err != nil {
		logger.Println("Create a new scraper failed: ", err.Error())
		if err := notify(err, detail, nil, notifier, logger); err != nil {
			return err
		}

		return err
	}

	err = scraper.Start(ctx)
	if err != nil {
		logger.Println("Scraper execution failed: ", err.Error())
		if err := notify(err, detail, nil, notifier, logger); err != nil {
			return err
		}

		return err
	}

	return notify(err, detail, scraper.Report.Table, notifier, logger)
}

func notify(err error, detail *core.Detail, table *core.Table, notifier core.Notifier, logger *log.Logger) error {
	if err != nil {
		detail.Status = core.Error
		detail.MessageError = err.Error()
	}

	if err == nil {
		detail.Status = core.Found
		detail.Table = table
	}

	logger.Println("Finishing Crawler for lopido.com ...")

	if err := notifier.Notify(detail); err != nil {
		logger.Println("Notify detail failed: ", err.Error())
		return err
	}

	return nil
}

func main() {
	lambda.Start(core.CrawlerProcessor(NewCollector, config))
}
