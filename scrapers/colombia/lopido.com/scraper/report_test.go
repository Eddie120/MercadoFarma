package scraper

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

func TestSetReport(t *testing.T) {
	c := require.New(t)

	file, err := os.Open("../samples/azitromicina.json")
	c.Nil(err)

	body, err := ioutil.ReadAll(file)
	c.Nil(err)

	response := DefaultResponse{}
	err = json.Unmarshal(body, &response)
	c.Nil(err)

	fmt.Println(response.QueryData[0].Data)
}