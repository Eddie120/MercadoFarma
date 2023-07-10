package commons

type CollectorError string

const MissingHttpClient CollectorError = "MISSING_HTTP_CLIENT"
const MissingQueryParam CollectorError = "MISSING_QUERY_PARAM"
const UnexpectedStatusCode CollectorError = "UNEXPECTED_STATUS_CODE"
const InvalidCountry CollectorError = "INVALID_COUNTRY"
const InvalidCity CollectorError = "INVALID_CITY"

func (c CollectorError) Error() string {
	return string(c)
}
