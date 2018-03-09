package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test.html")

	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("request should have %d requets;but had %d ", resultSize, len(result.Requests))
	}

}
