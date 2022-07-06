package scheduler

import (
	"go_carwler/types"
)

type SimpleScheduler struct {
	WorkerChan chan types.Request
}

func (s *SimpleScheduler) InitChan(c chan types.Request) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(request types.Request) {
	s.WorkerChan <- request
}
