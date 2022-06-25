package scheduler

import "go_carwler/zhenai/types"

type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}

func (s *SimpleScheduler) Submit(request types.Request) {
	//使用go rountine避免死锁
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) WorkerReady(chan types.Request) {
}
