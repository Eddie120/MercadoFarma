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
	Products        []Product `json:"products"`
	RecordsFiltered int       `json:"recordsFiltered"`
	TypeName        string    `json:"__typename"`
}

type Product struct {
	CacheId             string                    `json:"cacheId"`
	ProductId           string                    `json:"productId"`
	Description         string                    `json:"description"`
	ProductName         string                    `json:"productName"`
	ProductReference    string                    `json:"productReference"`
	LinkText            string                    `json:"linkText"`
	Brand               string                    `json:"brand"`
	BrandId             int                       `json:"brandId"`
	Link                string                    `json:"link"`
	Categories          []string                  `json:"categories"`
	CategoryId          string                    `json:"categoryId"`
	PriceRange          PriceRange                `json:"priceRange"`
	SpecificationGroups []SpecificationGroupsItem `json:"specificationGroups"`
}

type SpecificationGroupsProperty struct {
	Name         string   `json:"name"`
	OriginalName string   `json:"originalName"`
	Values       []string `json:"values"`
}

type SpecificationGroupsItem struct {
	Name           string                        `json:"name"`
	OriginalName   string                        `json:"originalName"`
	Specifications []SpecificationGroupsProperty `json:"specifications"`
}

type PriceRange struct {
	SellingPrice SellingPrice `json:"sellingPrice"`
	ListPrice    ListPrice    `json:"listPrice"`
}

type SellingPrice struct {
	HighPrice int64 `json:"highPrice"`
	LowPrice  int64 `json:"lowPrice"`
}

type ListPrice struct {
	HighPrice int64 `json:"highPrice"`
	LowPrice  int64 `json:"lowPrice"`
}
