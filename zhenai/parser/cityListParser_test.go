package parser

import (
	"GoStudy_mk/src/go13/g04/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	// 使用保存的文件，避免直接抓取数据失败
	//bytes, err := ioutil.ReadFile("src/go12/g04/zhenai/parser/test_data_city.html")
	bytes, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(bytes)
	const resultSize = 470
	exprctedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCitys := []string{
		"City 阿坝",
		"City 阿克苏",
		"City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("Requests长度应为%d，实际为%d", resultSize, len(result.Requests))
	}
	for i, url := range exprctedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("第%d个URL应为%s，实际为%s", i, result.Requests[i].Url, url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("Items长度应为%d，实际为%d", resultSize, len(result.Requests))
	}
	for i, city := range expectedCitys {
		if result.Items[i].(string) != city {
			t.Errorf("第%d个city应为%s，实际为%s", i, result.Items[i].(string), city)
		}
	}
}
