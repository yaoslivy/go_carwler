package main

import (
	"go_carwler/zhenai/engine"
	"go_carwler/zhenai/parser"
	"go_carwler/zhenai/types"
)

func main() {
	engine.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun", //从征婚网城市列表页面出发
		ParseFunc: parser.ParseCityList,
	})

}
