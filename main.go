package main

import (
	"go_carwler/engine"
	"go_carwler/parser"
	"go_carwler/scheduler"
	"go_carwler/types"
)

func main() {
	//engine.SingleEngine{}.Run(engineTypes.Request{
	//	Url:       "https://www.7799520.com/jiaou", //从征婚网城市列表页面出发
	//	ParseFunc: parser.ParseCityList,
	//})
	//http://www.7799520.com/user/6056626.html

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	concurrentEngine.Run(types.Request{
		Url:       "https://www.7799520.com/jiaou", //从征婚网城市列表页面出发
		ParseFunc: parser.ParseCityList,
	})
}
