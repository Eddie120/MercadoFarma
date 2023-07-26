package main

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/mercadofarma/services/commons"
	"github.com/mercadofarma/services/core"
	"github.com/mercadofarma/services/scrapers/colombia/lopido.com/scraper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCollector_Success(t *testing.T) {
	c := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var urlMock = scraper.BaseUrl
	urlMock = fmt.Sprintf(urlMock, "azitromicina", "azitromicina")
	finalUrl := urlMock + scraper.Remaining

	file := httpmock.File("../samples/azitromicina.json")
	httpmock.RegisterResponder(http.MethodGet, finalUrl, httpmock.NewJsonResponderOrPanic(http.StatusOK, file))

	const query = "azitromicina"
	const country core.Country = "colombia"
	const city core.City = "cali"

	collector := NewCollector()
	ctx := context.Background()
	inputs := &core.SearchInput{
		Query:   query,
		Country: country,
		City:    city,
	}

	var notifier core.HandlerFunc = func(ctx context.Context, d *core.Detail) error {
		c.Equal(core.Found, d.Status)
		c.Len(d.Table.Rows, 11)
		c.Equal("2891", d.Table.Rows[0].Cells[0].Value)
		c.Equal("AZITROMICINA 500 MG (MK)", d.Table.Rows[0].Cells[1].Value)
		c.Equal("12250", d.Table.Rows[0].Cells[3].Value)
		c.Equal("CAJA X 3 TAB", d.Table.Rows[0].Cells[4].Value)

		return nil
	}

	err := collector.Crawl(ctx, inputs, notifier)
	c.Nil(err)
}

func TestCollector_Unexpected_StatusCode(t *testing.T) {
	c := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	var urlMock = scraper.BaseUrl
	urlMock = fmt.Sprintf(urlMock, "asdf", "asdf")
	finalUrl := urlMock + scraper.Remaining

	file := httpmock.File("../samples/not_found.json")
	httpmock.RegisterResponder(http.MethodGet, finalUrl, httpmock.NewJsonResponderOrPanic(http.StatusNotFound, file))

	const query = "asdf"
	const country core.Country = "colombia"
	const city core.City = "cali"

	collector := NewCollector()
	ctx := context.Background()
	inputs := &core.SearchInput{
		Query:   query,
		Country: country,
		City:    city,
	}

	var notifier core.HandlerFunc = func(ctx context.Context, d *core.Detail) error {
		c.Equal(core.Error, d.Status)
		c.Nil(d.Table)

		return nil
	}

	err := collector.Crawl(ctx, inputs, notifier)
	c.Equal(commons.UnexpectedStatusCode, err)
}
