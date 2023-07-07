package scraper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mercadofarma/services/core"
	"io"
	"log"
	"net/http"
	"strings"
)

var (
	MissingHttpClient    = errors.New("missing http client")
	MissingQueryParam    = errors.New("query can not be empty")
	UnexpectedStatusCode = errors.New("unexpected status code")
	InvalidCountry       = errors.New("invalid country")
	InvalidCity          = errors.New("invalid city")
)

var baseUrl = "https://www.lopido.com/%s?_q=%s"
var remaining = "&map=ft&__pickRuntime=appsEtag%2Cblocks%2CblocksTree%2Ccomponents%2CcontentMap%2Cextensions%2Cmessages%2Cpage%2Cpages%2Cquery%2CqueryData%2Croute%2CruntimeMeta%2Csettings&__device=desktop"
var isValidCountry = core.CountriesAllowed
var isValidCity = core.CitiesAllowed

type Scraper struct {
	Client  *http.Client
	Query   string
	Report  *Report
	Log     *log.Logger
	Country core.Country
	City    core.City
}

func NewScraper(client *http.Client, query string, country core.Country, city core.City, logger *log.Logger) (*Scraper, error) {
	if client == nil {
		return nil, MissingHttpClient
	}

	if len(strings.Trim(query, " ")) == 0 {
		return nil, MissingQueryParam
	}

	if country == "" || !isValidCountry[country] {
		return nil, InvalidCountry
	}

	if city == "" || !isValidCity[city] {
		return nil, InvalidCity
	}

	return &Scraper{
		Client: client,
		Query:  query,
		Report: &Report{},
		Log:    logger,
	}, nil
}

func (s *Scraper) Start(ctx context.Context) error {
	finalUrl := fmt.Sprintf(baseUrl, s.Query, s.Query)
	finalUrl = finalUrl + remaining

	response, err := s.Client.Get(finalUrl)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			s.Log.Println("unable to close response body: ", err.Error())
		}
	}(response.Body)

	var defaultResponse DefaultResponse

	if response.StatusCode != http.StatusOK {
		s.Log.Println("unexpected status code: ", response.Status)
		return UnexpectedStatusCode
	}

	err = json.NewDecoder(response.Body).Decode(&defaultResponse)
	if err != nil {
		s.Log.Println("decode response body in DefaultResponse struct failed: ", err.Error())
		return err
	}

	return s.Report.setReport(defaultResponse)
}
