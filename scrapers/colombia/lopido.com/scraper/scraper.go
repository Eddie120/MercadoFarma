package scraper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mercadofarma/services/commons"
	"github.com/mercadofarma/services/core"
	"io"
	"log"
	"net/http"
	"strings"
)

const BaseUrl = "https://www.lopido.com/%s?_q=%s"
const Remaining = "&map=ft&__pickRuntime=appsEtag%2Cblocks%2CblocksTree%2Ccomponents%2CcontentMap%2Cextensions%2Cmessages%2Cpage%2Cpages%2Cquery%2CqueryData%2Croute%2CruntimeMeta%2Csettings&__device=desktop"

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
		return nil, commons.MissingHttpClient
	}

	if len(strings.Trim(query, " ")) == 0 {
		return nil, commons.MissingQueryParam
	}

	if country == "" || !isValidCountry[country] {
		return nil, commons.InvalidCountry
	}

	if city == "" || !isValidCity[city] {
		return nil, commons.InvalidCity
	}

	return &Scraper{
		Client: client,
		Query:  query,
		Report: &Report{},
		Log:    logger,
	}, nil
}

func (s *Scraper) Start(ctx context.Context) error {
	finalUrl := fmt.Sprintf(BaseUrl, s.Query, s.Query)
	finalUrl = finalUrl + Remaining

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
		return commons.UnexpectedStatusCode
	}

	err = json.NewDecoder(response.Body).Decode(&defaultResponse)
	if err != nil {
		s.Log.Println("decode response body in DefaultResponse struct failed: ", err.Error())
		return err
	}

	return s.Report.setReport(defaultResponse)
}
