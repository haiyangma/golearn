package main

import (
	"crawler/engine"
	keepConfig "crawler/keep/config"
	"crawler/keep/parser"
	"crawler/scheduler"
	"crawler_distributed/config"
	itemSaver "crawler_distributed/persist/client"
	"crawler_distributed/rpcsupport"
	worker "crawler_distributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")

	workerHosts = flag.String("worker_hosts", "", "worker hosts(comma seperated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(config.ItemSaverHost)
	if err != nil {
		panic(err)
	}
	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor, err := worker.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueneScheduler{},
		WorkerCount:      10,
		SaverChan:        itemChan,
		RequestProcessor: processor,
	}
	e.Run(engine.Request{
		Url:    keepConfig.KEEP_HOST+keepConfig.KEEP_EXPLORE_URI, //http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseExplorePostList,"ParseExplorePostList"),
	})

}
func createClientPool(hosts []string) chan *rpc.Client {
	clientChan := make(chan *rpc.Client, len(hosts))
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clientChan <- client
			log.Printf("Connected to %s", h)
		} else {
			log.Println("Error connecting to %s:%v", h, err)
		}
	}

	return clientChan
}
