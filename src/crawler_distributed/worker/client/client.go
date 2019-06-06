package client

import (
	"crawler/engine"
	"crawler_distributed/config"
	"crawler_distributed/worker"
	"log"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(r engine.Request) (engine.ParseResult, error) {
		client := <-clientChan
		log.Printf("Fetching %s \n", r.Url)
		var result worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, worker.SerializeRequest(r), &result)
		clientChan <- client
		if err != nil {
			log.Printf("Fetcher :error fetching url %s: %v", r.Url, err)
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(result), nil
	}, nil
}
