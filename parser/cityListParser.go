package parser

import (
	"go_carwler/types"
	"regexp"
)

// 解析城市URL和名称的正则表达式
//const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)" data-v-[0-9a-zA-Z]+[^>]*>([^<]+)</a>`
//<a href="http://www.7799520.com/jiaou/anhui">安徽</a>
const cityListRe = `<a href="(http://www.7799520.com/jiaou/[0-9a-zA-Z]+)">([^<]+)</a>`

func ParseCityList(bytes []byte) types.ParseResult {
	result := types.ParseResult{}
	// 使用[^X]匹配到非X的表达式，提取出url和城市名
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(bytes, -1)
	//limit := 10
	for _, matchStrs := range matches {
		url := string(matchStrs[1])
		cityName := string(matchStrs[2])
		// 返回城市任务(url和对应解析器)
		result.Requests = append(result.Requests, types.Request{
			Url:       string(url),
			ParseFunc: ParseCity,
		})
		result.Items = append(result.Items, "City "+cityName)
		//limit--
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
