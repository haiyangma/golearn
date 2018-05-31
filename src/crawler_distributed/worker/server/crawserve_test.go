package main

import (
	"crawler_distributed/config"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/worker"
	"fmt"
	"testing"
	"time"
)

func TestCrawlServe(t *testing.T) {
	go rpcsupport.ServeRpc(config.WorkerPort0, CrawlService{})
	time.Sleep(1 * time.Second)
	client, err := rpcsupport.NewClient(config.WorkerPort0)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/107444151",
		Parser: worker.SerializedParser{
			FunctionName: config.ParseProfile,
			Args:         "向幸福前进",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("error get the request %v:", req)
	} else {
		fmt.Println(result)
	}
}
