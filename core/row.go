package core

type code string

// Inputs fields defined for all collectors
const (
	ProductName         code = "PRODUCT_NAME"
	ProductDescription  code = "PRODUCT_DESCRIPTION"
	ProductId           code = "PRODUCT_ID"
	ProductReference    code = "PRODUCT_REFERENCE"
	ProductPrice        code = "PRODUCT_PRICE"
	ProductPresentation code = "PRODUCT_PRESENTATION"
)

type Row struct {
	Cells []*Cell `json:"cells"`
}
