package scraper

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestSetReportSuccess(t *testing.T) {
	c := require.New(t)

	file, err := os.Open("../samples/azitromicina.json")
	c.Nil(err)

	body, err := ioutil.ReadAll(file)
	c.Nil(err)

	response := DefaultResponse{}
	err = json.Unmarshal(body, &response)
	c.Nil(err)

	report := Report{}
	err = report.setReport(response)
	c.Nil(err)

	c.Equal("2891", report.Table.Rows[0].Cells[0].Value)
	c.Equal("AZITROMICINA 500 MG (MK)", report.Table.Rows[0].Cells[1].Value)
	c.Equal("12250", report.Table.Rows[0].Cells[3].Value)
	c.Equal("CAJA X 3 TAB", report.Table.Rows[0].Cells[4].Value)
}

func TestSetReportNotFound(t *testing.T) {
	c := require.New(t)

	file, err := os.Open("../samples/not_found.json")
	c.Nil(err)

	body, err := ioutil.ReadAll(file)
	c.Nil(err)

	response := DefaultResponse{}
	err = json.Unmarshal(body, &response)
	c.Nil(err)

	report := Report{}
	err = report.setReport(response)
	c.Nil(err)

	c.Len(report.Table.Rows, 0)
}
