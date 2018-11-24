package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	file, err := ioutil.ReadFile(
		"city_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCity(file)

	exceptedUrls := []string{
		"http://album.zhenai.com/u/1349973057",
		"http://album.zhenai.com/u/1130447736",
		"http://album.zhenai.com/u/1486293757",
	}
	exceptedUsers := []string{
		"User 虐心砝码", "User 小草屋女人", "User 流浪雪",
	}

	const resultSize = 20
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requestsk but had %d", resultSize, len(result.Requests))
	}
	for i, url := range exceptedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d: %s; but "+
				"was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"requestsk but had %d", resultSize, len(result.Items))
	}

	for i, city := range exceptedUsers {
		if result.Items[i].(string) != city {
			t.Errorf("excepted city #%d: %s; but was %s", i, city, result.Items[i].(string))
		}
	}

}
