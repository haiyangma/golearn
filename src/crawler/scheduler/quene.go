package scheduler

import "crawler/engine"

type QueneScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueneScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueneScheduler) ConfigureWorkerChan(chan engine.Request) {
}

func (q *QueneScheduler) Submit(request engine.Request) {
	q.requestChan <- request
}

func (q *QueneScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueneScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerChanQ []chan engine.Request
		var request engine.Request
		var workerChan chan engine.Request
		for {
			if len(requestQ) > 0 && len(workerChanQ) > 0 {
				request = requestQ[0]
				workerChan = workerChanQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				workerChanQ = append(workerChanQ, w)
			case workerChan <- request:
				if len(requestQ) > 0 {
					requestQ = requestQ[1:]
				}
				if len(workerChanQ) > 0 {
					workerChanQ = workerChanQ[1:]
				}
			}
		}
	}()
}
