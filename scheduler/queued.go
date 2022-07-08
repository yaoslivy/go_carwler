package scheduler

import "go_carwler/types"

type QueueScheduler struct {
	requestChan chan types.Request
	workerChan  chan chan types.Request //每一个worker一个channel，这样就可以指定给哪个worker发任务，类似负载均衡
}

func (q *QueueScheduler) Submit(request types.Request) {
	q.requestChan <- request
}

func (q *QueueScheduler) InitChan(requests chan types.Request) {

}

// 可以接收request，哪一个worker空闲，可以分配任务
func (q *QueueScheduler) WorkerReady(w chan types.Request) {
	q.workerChan <- w
}

func (s *QueueScheduler) Run() {
	s.workerChan = make(chan chan types.Request)
	s.requestChan = make(chan types.Request)
	go func() {
		var requestQ []types.Request
		var workerQ []chan types.Request

		for {
			var activateRequest types.Request
			var activateWorker chan types.Request
			// request 和 worker都在排队
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activateRequest = requestQ[0]
				activateWorker = workerQ[0]
			}
			select { //独立的事件，requestChan，workerChan可能有，也可能没有
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activateWorker <- activateRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
