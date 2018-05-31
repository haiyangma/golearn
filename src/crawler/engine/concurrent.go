package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	SaverChan        chan Item
	RequestProcessor Processor
}
type Processor func(r Request) (ParseResult, error)

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
	var visited = make(map[string]int)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), e.Scheduler, out)
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
			if count, ok := visited[request.Url]; !ok {
				visited[request.Url] = 1
				e.Scheduler.Submit(request)
			} else {
				visited[request.Url] = count + 1
			}
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, ready ReadyNotifier, out chan ParseResult) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
