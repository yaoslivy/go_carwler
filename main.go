package main

import (
	"go_carwler/zhenai/engine"
	"go_carwler/zhenai/parser"
	"go_carwler/zhenai/scheduler"
	"go_carwler/zhenai/types"
)

func main() {
	//engine.SingleEngine{}.Run(types.Request{
	//	Url:       "https://www.7799520.com/jiaou", //从征婚网城市列表页面出发
	//	ParseFunc: parser.ParseCityList,
	//})
	//http://www.7799520.com/user/6056626.html

	concurrentEngine := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
	}
	concurrentEngine.Run(types.Request{
		Url:       "https://www.7799520.com/jiaou", //从征婚网城市列表页面出发
		ParseFunc: parser.ParseCityList,
	})

}
