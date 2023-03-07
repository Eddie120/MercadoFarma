package scraper

type DefaultResponse struct {
	Query     map[string]string   `json:"query"`
	QueryData []QueryDataResponse `json:"queryData"`
}

type QueryDataResponse struct {
	Query string `json:"query"`
	Data  string `json:"data"`
}

type SearchResult struct {
	ProductSearch ProductSearch `json:"productSearch"`
}

type ProductSearch struct {
	Products        []string `json:"products"`
	RecordsFiltered int      `json:"recordsFiltered"`
	TypeName        string   `json:"__typename"`
}
