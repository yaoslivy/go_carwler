package main

import (
	"go_carwler/zhenai/engine"
	"go_carwler/zhenai/parser"
	"go_carwler/zhenai/types"
)

func main() {
	engine.Run(types.Request{
		Url:       "https://www.7799520.com/jiaou", //从征婚网城市列表页面出发
		ParseFunc: parser.ParseCityList,
	})
	//http://www.7799520.com/user/6056626.html

}
