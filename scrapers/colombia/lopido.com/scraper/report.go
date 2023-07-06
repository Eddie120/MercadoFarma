package scraper

import (
	"encoding/json"
	"github.com/mercadofarma/services/core"
	"strconv"
)

type Report struct {
	Table *core.Table
}

func (r *Report) setReport(response DefaultResponse) error {
	searchResult := SearchResult{}
	data := response.QueryData[0].Data

	err := json.Unmarshal([]byte(data), &searchResult)
	if err != nil {
		return err
	}

	r.Table = &core.Table{
		Rows: []*core.Row{},
	}
	r.setTableName()
	r.setRows(searchResult)

	return nil
}

func (r *Report) setTableName() {
	r.Table.TableName = "Basic Information"
}

func (r *Report) setCells(cells *[]*core.Cell, product Product) {
	if product.ProductId != "" {
		*cells = append(*cells, &core.Cell{
			Name:  string(core.ProductReference),
			Value: product.ProductId,
		})
	}

	if product.ProductName != "" {
		*cells = append(*cells, &core.Cell{
			Name:  string(core.ProductName),
			Value: product.ProductName,
		})
	}

	if product.Description != "" {
		*cells = append(*cells, &core.Cell{
			Name:  string(core.ProductDescription),
			Value: product.Description,
		})
	}

	*cells = append(*cells, &core.Cell{
		Name:  string(core.ProductPrice),
		Value: strconv.Itoa(int(product.PriceRange.SellingPrice.HighPrice)),
	})

	const groupName = "Pum"
	specificationItem := SpecificationGroupsItem{}
	if len(product.SpecificationGroups) > 0 {
		for _, specificationGroup := range product.SpecificationGroups {
			if specificationGroup.Name != groupName {
				continue
			}

			specificationItem = specificationGroup
		}
	}

	const presentationUnitOfMeasure = "Presentacionunidadmedida"
	for _, specificationGroupProperty := range specificationItem.Specifications {
		if specificationGroupProperty.Name == presentationUnitOfMeasure {
			*cells = append(*cells, &core.Cell{
				Name:  string(core.ProductPresentation),
				Value: specificationGroupProperty.Values[0],
			})
		}
	}
}

func (r *Report) setRows(results SearchResult) {
	if len(results.ProductSearch.Products) == 0 {
		r.Table.Rows = []*core.Row{}
		return
	}

	r.Table.Rows = make([]*core.Row, 0)

	products := results.ProductSearch.Products
	for _, product := range products {
		var cells []*core.Cell
		r.setCells(&cells, product)
		r.Table.Rows = append(r.Table.Rows, &core.Row{
			Cells: cells,
		})
	}
}
