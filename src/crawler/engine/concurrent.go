package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	SaverChan   chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), e.Scheduler, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			e.SaverChan <- item

		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, ready ReadyNotifier, out chan ParseResult) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
