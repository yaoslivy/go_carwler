package parser

import (
	"go_carwler/zhenai/types"
	"regexp"
)

// <th><a href="http://album.zhenai.com/u/1633617220" target="_blank">游天</a></th>
//const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a></th>`
// <h3><a class="name" href="http://www.7799520.com/user/6055438.html" target="_blank">红豆大哥</a></h3>
const cityRe = `<a class="name" href="(http://www.7799520.com/user/[0-9]+.html)"[^>]*>([^<]+)</a>`

// ParseCity 解析城市返回用户列表结果
func ParseCity(bytes []byte) types.ParseResult {
	result := types.ParseResult{}
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(bytes, -1)
	for _, matchStrs := range matches {
		url := string(matchStrs[1])
		Name := string(matchStrs[2])
		// 返回用户任务（url和对应解析函数）
		result.Requests = append(result.Requests, types.Request{
			Url: string(url),
			ParseFunc: func(bytes []byte) types.ParseResult {
				//函数式编程，传入名字
				return ParseUser(bytes, Name)
			},
		})
		result.Items = append(result.Items, "User "+Name)
	}
	return result
}
