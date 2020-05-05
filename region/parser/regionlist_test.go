package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("regionlist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)
	t.Errorf("总数量：%d", len(result.Requests))

}
