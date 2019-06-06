package main

import (
	"crawler/engine"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"flag"
	"fmt"
)

type CrawlService struct {
}

var port = flag.Int("port", 0, "the port for worker to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), CrawlService{})
}

func (CrawlService) Process(req worker.Request, result *worker.ParseResult) error {
	engineRequest, err := worker.DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}
	*result = worker.SerializeResult(engineResult)
	return nil
}
