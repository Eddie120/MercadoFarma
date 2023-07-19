package core

import (
	"encoding/json"
	"errors"
)

var (
	ErrCityMissing         = errors.New("city is missing")
	ErrCityNotSupported    = errors.New("city not supported")
	ErrCountryMissing      = errors.New("country missing")
	ErrCountryNotSupported = errors.New("country not supported")
)

type SearchInput struct {
	Query          string  `json:"query"`
	Country        Country `json:"country"`
	City           City    `json:"city"`
	CanonicalQuery string  `json:"canonical_query"`
}

func NewSearchInput(data []byte) (*SearchInput, error) {
	input := &SearchInput{}
	if err := json.Unmarshal(data, input); err != nil {
		return nil, err
	}

	if err := input.Validate(); err != nil {
		return nil, err
	}

	return input, nil
}

func (input *SearchInput) Validate() error {
	if input.Country == "" {
		return ErrCountryMissing
	}

	if !CountriesAllowed[input.Country] {
		return ErrCountryNotSupported
	}

	if input.City == "" {
		return ErrCityMissing
	}

	if !CitiesAllowed[input.City] {
		return ErrCityNotSupported
	}

	return nil
}

func (input *SearchInput) LogName() string {
	return "search-input"
}

func (input *SearchInput) LogProperties() map[string]interface{} {
	return map[string]interface{}{
		"s_query":   input.Query,
		"s_country": input.Country,
		"s_city":    input.City,
	}
}
