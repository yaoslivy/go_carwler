package engine

import (
	"go_carwler/types"
	"log"
)

//并发版执行引擎，包含调度器和执行者数
type ConcurrentEngine struct {
	Scheduler   Scheduler //任务调度器（队列、管道）
	WorkerCount int       //工作线程数
}

type Scheduler interface {
	Submit(request types.Request) //提交任务
	InitChan(chan types.Request)  //初始化管道
	WorkerReady(chan types.Request)
	Run()
}

func (c *ConcurrentEngine) Run(seeds ...types.Request) {
	//建立输入输出管道
	//in := make(chan engineTypes.Request)
	out := make(chan types.ParseResult)
	//c.Scheduler.InitChan(in)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(out, c.Scheduler)
	}

	//初始化任务队列
	for _, val := range seeds {
		c.Scheduler.Submit(val)
	}
	// 获得解析后的更多任务结果，加入队列中
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("item No.%d：%s", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

// 输入输出管道对接逻辑，使用协程go func，获取页面解析结构输出
func createWorker(out chan types.ParseResult, s Scheduler) {
	in := make(chan types.Request)
	go func() {
		for {
			//需要通知scheduler准备好了
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
