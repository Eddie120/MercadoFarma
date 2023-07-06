package scraper

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/mercadofarma/services/core"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewScraperSuccess(t *testing.T) {
	c := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var urlMock = baseUrl
	urlMock = fmt.Sprintf(urlMock, "dolex", "dolex")
	finalUrl := urlMock + remaining

	file := httpmock.File("../samples/dolex.json")
	httpmock.RegisterResponder(http.MethodGet, finalUrl, httpmock.NewJsonResponderOrPanic(http.StatusOK, file))

	const query = "dolex"
	const country core.Country = "colombia"
	const city core.City = "cali"
	scraper, _ := NewScraper(&http.Client{}, query, country, city)

	ctx := context.Background()
	err := scraper.Start(ctx)
	c.Nil(err)

	rows := scraper.Report.Table.Rows
	c.Len(rows, 16)
}